// Code generated by "enumer -type=ClientOpcode,ServerOpcode -trimprefix=Op -output=enumer.go"; DO NOT EDIT.

package realmd

import (
	"fmt"
	"strings"
)

const (
	_ClientOpcodeName_0      = "ClientCharCreateClientCharListClientCharDelete"
	_ClientOpcodeLowerName_0 = "clientcharcreateclientcharlistclientchardelete"
	_ClientOpcodeName_1      = "ClientPlayerLogin"
	_ClientOpcodeLowerName_1 = "clientplayerlogin"
	_ClientOpcodeName_2      = "ClientLogout"
	_ClientOpcodeLowerName_2 = "clientlogout"
	_ClientOpcodeName_3      = "ClientLogoutCancel"
	_ClientOpcodeLowerName_3 = "clientlogoutcancel"
	_ClientOpcodeName_4      = "ClientPing"
	_ClientOpcodeLowerName_4 = "clientping"
	_ClientOpcodeName_5      = "ClientAuthSession"
	_ClientOpcodeLowerName_5 = "clientauthsession"
	_ClientOpcodeName_6      = "ClientGetStorageClientPutStorage"
	_ClientOpcodeLowerName_6 = "clientgetstorageclientputstorage"
	_ClientOpcodeName_7      = "ClientRealmSplit"
	_ClientOpcodeLowerName_7 = "clientrealmsplit"
	_ClientOpcodeName_8      = "ClientGetUnixTime"
	_ClientOpcodeLowerName_8 = "clientgetunixtime"
	_ClientOpcodeName_9      = "ClientReadyForAccountDataTimes"
	_ClientOpcodeLowerName_9 = "clientreadyforaccountdatatimes"
)

var (
	_ClientOpcodeIndex_0 = [...]uint8{0, 16, 30, 46}
	_ClientOpcodeIndex_1 = [...]uint8{0, 17}
	_ClientOpcodeIndex_2 = [...]uint8{0, 12}
	_ClientOpcodeIndex_3 = [...]uint8{0, 18}
	_ClientOpcodeIndex_4 = [...]uint8{0, 10}
	_ClientOpcodeIndex_5 = [...]uint8{0, 17}
	_ClientOpcodeIndex_6 = [...]uint8{0, 16, 32}
	_ClientOpcodeIndex_7 = [...]uint8{0, 16}
	_ClientOpcodeIndex_8 = [...]uint8{0, 17}
	_ClientOpcodeIndex_9 = [...]uint8{0, 30}
)

func (i ClientOpcode) String() string {
	switch {
	case 54 <= i && i <= 56:
		i -= 54
		return _ClientOpcodeName_0[_ClientOpcodeIndex_0[i]:_ClientOpcodeIndex_0[i+1]]
	case i == 61:
		return _ClientOpcodeName_1
	case i == 75:
		return _ClientOpcodeName_2
	case i == 78:
		return _ClientOpcodeName_3
	case i == 476:
		return _ClientOpcodeName_4
	case i == 493:
		return _ClientOpcodeName_5
	case 522 <= i && i <= 523:
		i -= 522
		return _ClientOpcodeName_6[_ClientOpcodeIndex_6[i]:_ClientOpcodeIndex_6[i+1]]
	case i == 908:
		return _ClientOpcodeName_7
	case i == 1270:
		return _ClientOpcodeName_8
	case i == 1279:
		return _ClientOpcodeName_9
	default:
		return fmt.Sprintf("ClientOpcode(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ClientOpcodeNoOp() {
	var x [1]struct{}
	_ = x[OpClientCharCreate-(54)]
	_ = x[OpClientCharList-(55)]
	_ = x[OpClientCharDelete-(56)]
	_ = x[OpClientPlayerLogin-(61)]
	_ = x[OpClientLogout-(75)]
	_ = x[OpClientLogoutCancel-(78)]
	_ = x[OpClientPing-(476)]
	_ = x[OpClientAuthSession-(493)]
	_ = x[OpClientGetStorage-(522)]
	_ = x[OpClientPutStorage-(523)]
	_ = x[OpClientRealmSplit-(908)]
	_ = x[OpClientGetUnixTime-(1270)]
	_ = x[OpClientReadyForAccountDataTimes-(1279)]
}

var _ClientOpcodeValues = []ClientOpcode{OpClientCharCreate, OpClientCharList, OpClientCharDelete, OpClientPlayerLogin, OpClientLogout, OpClientLogoutCancel, OpClientPing, OpClientAuthSession, OpClientGetStorage, OpClientPutStorage, OpClientRealmSplit, OpClientGetUnixTime, OpClientReadyForAccountDataTimes}

var _ClientOpcodeNameToValueMap = map[string]ClientOpcode{
	_ClientOpcodeName_0[0:16]:       OpClientCharCreate,
	_ClientOpcodeLowerName_0[0:16]:  OpClientCharCreate,
	_ClientOpcodeName_0[16:30]:      OpClientCharList,
	_ClientOpcodeLowerName_0[16:30]: OpClientCharList,
	_ClientOpcodeName_0[30:46]:      OpClientCharDelete,
	_ClientOpcodeLowerName_0[30:46]: OpClientCharDelete,
	_ClientOpcodeName_1[0:17]:       OpClientPlayerLogin,
	_ClientOpcodeLowerName_1[0:17]:  OpClientPlayerLogin,
	_ClientOpcodeName_2[0:12]:       OpClientLogout,
	_ClientOpcodeLowerName_2[0:12]:  OpClientLogout,
	_ClientOpcodeName_3[0:18]:       OpClientLogoutCancel,
	_ClientOpcodeLowerName_3[0:18]:  OpClientLogoutCancel,
	_ClientOpcodeName_4[0:10]:       OpClientPing,
	_ClientOpcodeLowerName_4[0:10]:  OpClientPing,
	_ClientOpcodeName_5[0:17]:       OpClientAuthSession,
	_ClientOpcodeLowerName_5[0:17]:  OpClientAuthSession,
	_ClientOpcodeName_6[0:16]:       OpClientGetStorage,
	_ClientOpcodeLowerName_6[0:16]:  OpClientGetStorage,
	_ClientOpcodeName_6[16:32]:      OpClientPutStorage,
	_ClientOpcodeLowerName_6[16:32]: OpClientPutStorage,
	_ClientOpcodeName_7[0:16]:       OpClientRealmSplit,
	_ClientOpcodeLowerName_7[0:16]:  OpClientRealmSplit,
	_ClientOpcodeName_8[0:17]:       OpClientGetUnixTime,
	_ClientOpcodeLowerName_8[0:17]:  OpClientGetUnixTime,
	_ClientOpcodeName_9[0:30]:       OpClientReadyForAccountDataTimes,
	_ClientOpcodeLowerName_9[0:30]:  OpClientReadyForAccountDataTimes,
}

var _ClientOpcodeNames = []string{
	_ClientOpcodeName_0[0:16],
	_ClientOpcodeName_0[16:30],
	_ClientOpcodeName_0[30:46],
	_ClientOpcodeName_1[0:17],
	_ClientOpcodeName_2[0:12],
	_ClientOpcodeName_3[0:18],
	_ClientOpcodeName_4[0:10],
	_ClientOpcodeName_5[0:17],
	_ClientOpcodeName_6[0:16],
	_ClientOpcodeName_6[16:32],
	_ClientOpcodeName_7[0:16],
	_ClientOpcodeName_8[0:17],
	_ClientOpcodeName_9[0:30],
}

// ClientOpcodeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ClientOpcodeString(s string) (ClientOpcode, error) {
	if val, ok := _ClientOpcodeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ClientOpcodeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ClientOpcode values", s)
}

// ClientOpcodeValues returns all values of the enum
func ClientOpcodeValues() []ClientOpcode {
	return _ClientOpcodeValues
}

// ClientOpcodeStrings returns a slice of all String values of the enum
func ClientOpcodeStrings() []string {
	strs := make([]string, len(_ClientOpcodeNames))
	copy(strs, _ClientOpcodeNames)
	return strs
}

// IsAClientOpcode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ClientOpcode) IsAClientOpcode() bool {
	for _, v := range _ClientOpcodeValues {
		if i == v {
			return true
		}
	}
	return false
}

const _ServerOpcodeName = "ServerCharCreateServerCharListServerCharDeleteServerCharLoginFailedServerLogoutServerLogoutCompleteServerLogoutCancelACKServerUpdateObjectServerPlayCinematicServerTutorialFlagsServerHearthLocationServerPongServerAuthChallengeServerAuthResponseServerAccountStorageTimesServerGetStorageServerCharLoginVerifyWorldServerRealmSplitServerSystemFeaturesServerPutStorageOKServerUnixTime"
const _ServerOpcodeLowerName = "servercharcreateservercharlistserverchardeleteservercharloginfailedserverlogoutserverlogoutcompleteserverlogoutcancelackserverupdateobjectserverplaycinematicservertutorialflagsserverhearthlocationserverpongserverauthchallengeserverauthresponseserveraccountstoragetimesservergetstorageservercharloginverifyworldserverrealmsplitserversystemfeaturesserverputstorageokserverunixtime"

var _ServerOpcodeMap = map[ServerOpcode]string{
	58:   _ServerOpcodeName[0:16],
	59:   _ServerOpcodeName[16:30],
	60:   _ServerOpcodeName[30:46],
	65:   _ServerOpcodeName[46:67],
	76:   _ServerOpcodeName[67:79],
	77:   _ServerOpcodeName[79:99],
	79:   _ServerOpcodeName[99:120],
	169:  _ServerOpcodeName[120:138],
	250:  _ServerOpcodeName[138:157],
	253:  _ServerOpcodeName[157:176],
	341:  _ServerOpcodeName[176:196],
	477:  _ServerOpcodeName[196:206],
	492:  _ServerOpcodeName[206:225],
	494:  _ServerOpcodeName[225:243],
	521:  _ServerOpcodeName[243:268],
	524:  _ServerOpcodeName[268:284],
	566:  _ServerOpcodeName[284:310],
	907:  _ServerOpcodeName[310:326],
	969:  _ServerOpcodeName[326:346],
	1123: _ServerOpcodeName[346:364],
	1271: _ServerOpcodeName[364:378],
}

func (i ServerOpcode) String() string {
	if str, ok := _ServerOpcodeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("ServerOpcode(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ServerOpcodeNoOp() {
	var x [1]struct{}
	_ = x[OpServerCharCreate-(58)]
	_ = x[OpServerCharList-(59)]
	_ = x[OpServerCharDelete-(60)]
	_ = x[OpServerCharLoginFailed-(65)]
	_ = x[OpServerLogout-(76)]
	_ = x[OpServerLogoutComplete-(77)]
	_ = x[OpServerLogoutCancelACK-(79)]
	_ = x[OpServerUpdateObject-(169)]
	_ = x[OpServerPlayCinematic-(250)]
	_ = x[OpServerTutorialFlags-(253)]
	_ = x[OpServerHearthLocation-(341)]
	_ = x[OpServerPong-(477)]
	_ = x[OpServerAuthChallenge-(492)]
	_ = x[OpServerAuthResponse-(494)]
	_ = x[OpServerAccountStorageTimes-(521)]
	_ = x[OpServerGetStorage-(524)]
	_ = x[OpServerCharLoginVerifyWorld-(566)]
	_ = x[OpServerRealmSplit-(907)]
	_ = x[OpServerSystemFeatures-(969)]
	_ = x[OpServerPutStorageOK-(1123)]
	_ = x[OpServerUnixTime-(1271)]
}

var _ServerOpcodeValues = []ServerOpcode{OpServerCharCreate, OpServerCharList, OpServerCharDelete, OpServerCharLoginFailed, OpServerLogout, OpServerLogoutComplete, OpServerLogoutCancelACK, OpServerUpdateObject, OpServerPlayCinematic, OpServerTutorialFlags, OpServerHearthLocation, OpServerPong, OpServerAuthChallenge, OpServerAuthResponse, OpServerAccountStorageTimes, OpServerGetStorage, OpServerCharLoginVerifyWorld, OpServerRealmSplit, OpServerSystemFeatures, OpServerPutStorageOK, OpServerUnixTime}

var _ServerOpcodeNameToValueMap = map[string]ServerOpcode{
	_ServerOpcodeName[0:16]:         OpServerCharCreate,
	_ServerOpcodeLowerName[0:16]:    OpServerCharCreate,
	_ServerOpcodeName[16:30]:        OpServerCharList,
	_ServerOpcodeLowerName[16:30]:   OpServerCharList,
	_ServerOpcodeName[30:46]:        OpServerCharDelete,
	_ServerOpcodeLowerName[30:46]:   OpServerCharDelete,
	_ServerOpcodeName[46:67]:        OpServerCharLoginFailed,
	_ServerOpcodeLowerName[46:67]:   OpServerCharLoginFailed,
	_ServerOpcodeName[67:79]:        OpServerLogout,
	_ServerOpcodeLowerName[67:79]:   OpServerLogout,
	_ServerOpcodeName[79:99]:        OpServerLogoutComplete,
	_ServerOpcodeLowerName[79:99]:   OpServerLogoutComplete,
	_ServerOpcodeName[99:120]:       OpServerLogoutCancelACK,
	_ServerOpcodeLowerName[99:120]:  OpServerLogoutCancelACK,
	_ServerOpcodeName[120:138]:      OpServerUpdateObject,
	_ServerOpcodeLowerName[120:138]: OpServerUpdateObject,
	_ServerOpcodeName[138:157]:      OpServerPlayCinematic,
	_ServerOpcodeLowerName[138:157]: OpServerPlayCinematic,
	_ServerOpcodeName[157:176]:      OpServerTutorialFlags,
	_ServerOpcodeLowerName[157:176]: OpServerTutorialFlags,
	_ServerOpcodeName[176:196]:      OpServerHearthLocation,
	_ServerOpcodeLowerName[176:196]: OpServerHearthLocation,
	_ServerOpcodeName[196:206]:      OpServerPong,
	_ServerOpcodeLowerName[196:206]: OpServerPong,
	_ServerOpcodeName[206:225]:      OpServerAuthChallenge,
	_ServerOpcodeLowerName[206:225]: OpServerAuthChallenge,
	_ServerOpcodeName[225:243]:      OpServerAuthResponse,
	_ServerOpcodeLowerName[225:243]: OpServerAuthResponse,
	_ServerOpcodeName[243:268]:      OpServerAccountStorageTimes,
	_ServerOpcodeLowerName[243:268]: OpServerAccountStorageTimes,
	_ServerOpcodeName[268:284]:      OpServerGetStorage,
	_ServerOpcodeLowerName[268:284]: OpServerGetStorage,
	_ServerOpcodeName[284:310]:      OpServerCharLoginVerifyWorld,
	_ServerOpcodeLowerName[284:310]: OpServerCharLoginVerifyWorld,
	_ServerOpcodeName[310:326]:      OpServerRealmSplit,
	_ServerOpcodeLowerName[310:326]: OpServerRealmSplit,
	_ServerOpcodeName[326:346]:      OpServerSystemFeatures,
	_ServerOpcodeLowerName[326:346]: OpServerSystemFeatures,
	_ServerOpcodeName[346:364]:      OpServerPutStorageOK,
	_ServerOpcodeLowerName[346:364]: OpServerPutStorageOK,
	_ServerOpcodeName[364:378]:      OpServerUnixTime,
	_ServerOpcodeLowerName[364:378]: OpServerUnixTime,
}

var _ServerOpcodeNames = []string{
	_ServerOpcodeName[0:16],
	_ServerOpcodeName[16:30],
	_ServerOpcodeName[30:46],
	_ServerOpcodeName[46:67],
	_ServerOpcodeName[67:79],
	_ServerOpcodeName[79:99],
	_ServerOpcodeName[99:120],
	_ServerOpcodeName[120:138],
	_ServerOpcodeName[138:157],
	_ServerOpcodeName[157:176],
	_ServerOpcodeName[176:196],
	_ServerOpcodeName[196:206],
	_ServerOpcodeName[206:225],
	_ServerOpcodeName[225:243],
	_ServerOpcodeName[243:268],
	_ServerOpcodeName[268:284],
	_ServerOpcodeName[284:310],
	_ServerOpcodeName[310:326],
	_ServerOpcodeName[326:346],
	_ServerOpcodeName[346:364],
	_ServerOpcodeName[364:378],
}

// ServerOpcodeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ServerOpcodeString(s string) (ServerOpcode, error) {
	if val, ok := _ServerOpcodeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ServerOpcodeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ServerOpcode values", s)
}

// ServerOpcodeValues returns all values of the enum
func ServerOpcodeValues() []ServerOpcode {
	return _ServerOpcodeValues
}

// ServerOpcodeStrings returns a slice of all String values of the enum
func ServerOpcodeStrings() []string {
	strs := make([]string, len(_ServerOpcodeNames))
	copy(strs, _ServerOpcodeNames)
	return strs
}

// IsAServerOpcode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ServerOpcode) IsAServerOpcode() bool {
	_, ok := _ServerOpcodeMap[i]
	return ok
}