package authd

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"log"
	mrand "math/rand"
	"time"

	"github.com/kangaroux/gomaggus/internal"
	"github.com/kangaroux/gomaggus/internal/authd/packets"
	"github.com/kangaroux/gomaggus/internal/models"
	"github.com/kangaroux/gomaggus/internal/srp"
)

func handleLoginChallenge(services *Services, c *Client, data []byte) error {
	log.Println("Starting login challenge")

	var err error

	p := packets.ClientLoginChallenge{}
	if err = p.Read(data); err != nil {
		return err
	}
	c.username = p.Username

	log.Printf("client trying to login as '%s'\n", c.username)

	c.account, err = services.accounts.Get(&models.AccountGetParams{Username: c.username})
	if err != nil {
		return err
	}

	// https://gtker.com/wow_messages/docs/cmd_auth_logon_challenge_server.html#protocol-version-8
	resp := &bytes.Buffer{}
	resp.WriteByte(OP_LOGIN_CHALLENGE)
	resp.WriteByte(0) // protocol version

	// Always return success to prevent a bad actor from mining usernames
	resp.WriteByte(WOW_SUCCESS)

	var publicKey []byte
	var salt []byte

	if c.account == nil {
		publicKey = make([]byte, srp.KeySize)
		if _, err := rand.Read(publicKey); err != nil {
			return err
		}

		// A real account will always return the same salt, so our fake account needs to do that, too.
		// Using the username as a seed for the fake salt guarantees we always generate the same data.
		// Ironically, using crypto/rand here is actually less secure.
		//
		// If we didn't do this, a bad actor could send two challenges with the same username and compare
		// the salts. The salts would be the same for real accounts and different for fake accounts.
		// This would allow someone to mine usernames and start building an attack vector.
		seededRand := mrand.New(mrand.NewSource(internal.FastHash(c.username)))
		salt = make([]byte, srp.SaltSize)
		if _, err := seededRand.Read(salt); err != nil {
			return err
		}
	} else {
		if err = c.account.DecodeSrp(); err != nil {
			return err
		}
		publicKey = srp.CalculateServerPublicKey(c.account.Verifier(), c.privateKey)
		c.serverPublicKey = publicKey
		salt = c.account.Salt()
	}

	resp.Write(publicKey)
	resp.WriteByte(1)  // generator size (1 byte)
	resp.WriteByte(7)  // generator
	resp.WriteByte(32) // large prime size (32 bytes)
	resp.Write(srp.LargeSafePrime())
	resp.Write(salt)
	resp.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) // crc hash
	resp.WriteByte(0)

	if _, err := c.conn.Write(resp.Bytes()); err != nil {
		return err
	}

	log.Println("Replied to login challenge")
	c.state = StateAuthProof

	return nil
}

func handleLoginProof(services *Services, c *Client, data []byte) error {
	log.Println("Starting login proof")

	var serverProof []byte
	authenticated := false

	if c.account != nil {
		p := packets.ClientLoginProof{}
		if err := p.Read(data); err != nil {
			return err
		}

		c.clientPublicKey = p.ClientPublicKey[:]
		c.sessionKey = srp.CalculateServerSessionKey(
			c.clientPublicKey,
			c.serverPublicKey,
			c.privateKey,
			c.account.Verifier(),
		)
		calculatedClientProof := srp.CalculateClientProof(
			c.account.Username,
			c.account.Salt(),
			c.clientPublicKey,
			c.serverPublicKey,
			c.sessionKey,
		)
		authenticated = bytes.Equal(calculatedClientProof, p.ClientProof[:])

		if authenticated {
			serverProof = srp.CalculateServerProof(c.clientPublicKey, p.ClientProof[:], c.sessionKey)
		}
	}

	// https://gtker.com/wow_messages/docs/cmd_auth_logon_proof_server.html#protocol-version-8
	resp := &bytes.Buffer{}
	resp.WriteByte(OP_LOGIN_PROOF)

	if !authenticated {
		resp.WriteByte(WOW_FAIL_UNKNOWN_ACCOUNT)
		resp.Write([]byte{0, 0}) // padding
	} else {
		resp.WriteByte(WOW_SUCCESS)
		resp.Write(serverProof)
		resp.Write([]byte{0, 0, 0, 0}) // Account flag
		resp.Write([]byte{0, 0, 0, 0}) // Hardware survey ID
		resp.Write([]byte{0, 0})       // Unknown
	}

	if _, err := c.conn.Write(resp.Bytes()); err != nil {
		return err
	}

	log.Println("Replied to login proof")

	if authenticated {
		err := services.sessions.UpdateOrCreate(&models.Session{
			AccountId:     c.account.Id,
			SessionKeyHex: hex.EncodeToString(c.sessionKey),
			Connected:     1,
			ConnectedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		})
		if err != nil {
			return err
		}

		c.state = StateAuthenticated
	} else {
		c.state = StateInvalid
	}

	return nil
}

func handleReconnectChallenge(services *Services, c *Client, data []byte) error {
	log.Println("Starting reconnect challenge")

	var err error
	p := packets.ClientLoginChallenge{}
	if err = p.Read(data); err != nil {
		return err
	}
	c.username = p.Username

	log.Printf("client trying to reconnect as '%s'\n", c.username)

	c.account, err = services.accounts.Get(&models.AccountGetParams{Username: c.username})
	if err != nil {
		return err
	}

	// Generate random data that will be used for the reconnect proof
	if _, err := rand.Read(c.reconnectData); err != nil {
		return err
	}

	// https://gtker.com/wow_messages/docs/cmd_auth_reconnect_challenge_server.html#protocol-version-8
	resp := &bytes.Buffer{}
	resp.WriteByte(OP_RECONNECT_CHALLENGE)

	// Always return success to prevent a bad actor from mining usernames
	resp.WriteByte(WOW_SUCCESS)
	resp.Write(c.reconnectData)
	resp.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) // checksum salt

	if _, err := c.conn.Write(resp.Bytes()); err != nil {
		return err
	}

	log.Println("Replied to reconnect challenge")

	c.state = StateReconnectProof

	return nil
}

func handleReconnectProof(services *Services, c *Client, data []byte) error {
	log.Println("Starting reconnect proof")

	authenticated := false

	if c.account != nil {
		session, err := services.sessions.Get(c.account.Id)
		if err != nil {
			return err
		}

		// We can only try to reconnect the client if we have a previous session key
		if session != nil {
			if err := session.Decode(); err != nil {
				return err
			}
			c.sessionKey = session.SessionKey()

			p := packets.ClientReconnectProof{}
			if err := p.Read(data); err != nil {
				return err
			}

			serverProof := srp.CalculateReconnectProof(c.username, p.ProofData[:], c.reconnectData, c.sessionKey)
			authenticated = bytes.Equal(serverProof, p.ClientProof[:])
		}
	}

	// https://gtker.com/wow_messages/docs/cmd_auth_reconnect_proof_server.html#protocol-version-8
	resp := &bytes.Buffer{}
	resp.WriteByte(OP_RECONNECT_PROOF)

	if !authenticated {
		resp.WriteByte(WOW_FAIL_UNKNOWN_ACCOUNT)
	} else {
		resp.WriteByte(WOW_SUCCESS)
	}

	resp.Write([]byte{0, 0}) // padding

	if _, err := c.conn.Write(resp.Bytes()); err != nil {
		return err
	}

	log.Println("Replied to reconnect proof")

	if authenticated {
		err := services.sessions.UpdateOrCreate(&models.Session{
			AccountId:     c.account.Id,
			SessionKeyHex: hex.EncodeToString(c.sessionKey),
			Connected:     1,
			ConnectedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		})
		if err != nil {
			return err
		}
		c.state = StateAuthenticated
	} else {
		c.state = StateInvalid
	}

	return nil
}

func handleRealmList(services *Services, c *Client) error {
	realmList, err := services.realms.List()
	if err != nil {
		return err
	}

	// https://gtker.com/wow_messages/docs/cmd_realm_list_server.html#protocol-version-8
	resp := &bytes.Buffer{}
	resp.WriteByte(OP_REALM_LIST)

	inner := &bytes.Buffer{}
	inner.Write([]byte{0, 0, 0, 0}) // header padding
	binary.Write(inner, binary.LittleEndian, uint16(len(realmList)))
	for _, r := range realmList {
		inner.WriteByte(byte(r.Type))
		inner.WriteByte(0)                    // locked
		inner.WriteByte(byte(REALMFLAG_NONE)) // TODO?
		inner.WriteString(r.Name)
		inner.WriteByte(0) // name is NUL-terminated
		inner.WriteString(r.Host)
		inner.WriteByte(0)                                   // host is NUL-terminated
		binary.Write(inner, binary.LittleEndian, float32(0)) // TODO: population
		inner.WriteByte(byte(0))                             // TODO: number of chars on realm
		inner.WriteByte(byte(r.Region))
		inner.WriteByte(byte(r.Id))
	}
	inner.Write([]byte{0, 0}) // footer padding

	// Write size of realm list payload
	binary.Write(resp, binary.LittleEndian, uint16(inner.Len()))
	// Concat to main payload
	inner.WriteTo(resp)

	if _, err := c.conn.Write(resp.Bytes()); err != nil {
		return err
	}

	log.Println("Replied to realm list")

	return nil
}
