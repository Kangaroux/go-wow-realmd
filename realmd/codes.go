package realmd

type ServerOpcode uint16

const (
	OpServerAuthChallenge        ServerOpcode = 0x1EC
	OpServerAuthResponse         ServerOpcode = 0x1EE
	OpServerPong                 ServerOpcode = 0x1DD
	OpServerAccountStorageTimes  ServerOpcode = 0x209
	OpServerCharList             ServerOpcode = 0x3B
	OpServerRealmSplit           ServerOpcode = 0x38B
	OpServerCharCreate           ServerOpcode = 0x3A
	OpServerCharDelete           ServerOpcode = 0x3C
	OpServerCharLoginFailed      ServerOpcode = 0x41
	OpServerCharLoginVerifyWorld ServerOpcode = 0x236
	OpServerUpdateObject         ServerOpcode = 0xA9
	OpServerTutorialFlags        ServerOpcode = 0xFD
	OpServerSystemFeatures       ServerOpcode = 0x3C9
	OpServerHearthLocation       ServerOpcode = 0x155 // SMSG_BINDPOINTUPDATE
	OpServerPlayCinematic        ServerOpcode = 0xFA  // SMSG_TRIGGER_CINEMATIC
	OpServerLogout               ServerOpcode = 0x4C
	OpServerLogoutComplete       ServerOpcode = 0x4D
	OpServerLogoutCancelACK      ServerOpcode = 0x4F
	OpServerPutStorageOK         ServerOpcode = 0x463 // SMSG_UPDATE_ACCOUNT_DATA_COMPLETE
	OpServerGetStorage           ServerOpcode = 0x20C // SMSG_UPDATE_ACCOUNT_DATA
)

type ClientOpcode uint32

const (
	OpClientAuthSession              ClientOpcode = 0x1ED
	OpClientRealmSplit               ClientOpcode = 0x38C
	OpClientPing                     ClientOpcode = 0x1DC
	OpClientCharList                 ClientOpcode = 0x37
	OpClientCharCreate               ClientOpcode = 0x36
	OpClientCharDelete               ClientOpcode = 0x38
	OpClientPlayerLogin              ClientOpcode = 0x3D
	OpClientReadyForAccountDataTimes ClientOpcode = 0x4FF
	OpClientLogout                   ClientOpcode = 0x4B
	OpClientLogoutCancel             ClientOpcode = 0x4E
	OpClientPutStorage               ClientOpcode = 0x20B // CMSG_UPDATE_ACCOUNT_DATA
	OpClientGetStorage               ClientOpcode = 0x20A // CMSG_REQUEST_ACCOUNT_DATA
)

type ResponseCode byte

const (
	RespCodeResponseSuccess                                ResponseCode = 0x00
	RespCodeResponseFailure                                ResponseCode = 0x01
	RespCodeResponseCancelled                              ResponseCode = 0x02
	RespCodeResponseDisconnected                           ResponseCode = 0x03
	RespCodeResponseFailedToConnect                        ResponseCode = 0x04
	RespCodeResponseConnected                              ResponseCode = 0x05
	RespCodeResponseVersionMismatch                        ResponseCode = 0x06
	RespCodeCStatusConnecting                              ResponseCode = 0x07
	RespCodeCStatusNegotiatingSecurity                     ResponseCode = 0x08
	RespCodeCStatusNegotiationComplete                     ResponseCode = 0x09
	RespCodeCStatusNegotiationFailed                       ResponseCode = 0x0A
	RespCodeCStatusAuthenticating                          ResponseCode = 0x0B
	RespCodeAuthOk                                         ResponseCode = 0x0C
	RespCodeAuthFailed                                     ResponseCode = 0x0D
	RespCodeAuthReject                                     ResponseCode = 0x0E
	RespCodeAuthBadServerProof                             ResponseCode = 0x0F
	RespCodeAuthUnavailable                                ResponseCode = 0x10
	RespCodeAuthSystemError                                ResponseCode = 0x11
	RespCodeAuthBillingError                               ResponseCode = 0x12
	RespCodeAuthBillingExpired                             ResponseCode = 0x13
	RespCodeAuthVersionMismatch                            ResponseCode = 0x14
	RespCodeAuthUnknownAccount                             ResponseCode = 0x15
	RespCodeAuthIncorrectPassword                          ResponseCode = 0x16
	RespCodeAuthSessionExpired                             ResponseCode = 0x17
	RespCodeAuthServerShuttingDown                         ResponseCode = 0x18
	RespCodeAuthAlreadyLoggingIn                           ResponseCode = 0x19
	RespCodeAuthLoginServerNotFound                        ResponseCode = 0x1A
	RespCodeAuthWaitQueue                                  ResponseCode = 0x1B
	RespCodeAuthBanned                                     ResponseCode = 0x1C
	RespCodeAuthAlreadyOnline                              ResponseCode = 0x1D
	RespCodeAuthNoTime                                     ResponseCode = 0x1E
	RespCodeAuthDbBusy                                     ResponseCode = 0x1F
	RespCodeAuthSuspended                                  ResponseCode = 0x20
	RespCodeAuthParentalControl                            ResponseCode = 0x21
	RespCodeAuthLockedEnforced                             ResponseCode = 0x22
	RespCodeRealmListInProgress                            ResponseCode = 0x23
	RespCodeRealmListSuccess                               ResponseCode = 0x24
	RespCodeRealmListFailed                                ResponseCode = 0x25
	RespCodeRealmListInvalid                               ResponseCode = 0x26
	RespCodeRealmListRealmNotFound                         ResponseCode = 0x27
	RespCodeAccountCreateInProgress                        ResponseCode = 0x28
	RespCodeAccountCreateSuccess                           ResponseCode = 0x29
	RespCodeAccountCreateFailed                            ResponseCode = 0x2A
	RespCodeCharListRetrieving                             ResponseCode = 0x2B
	RespCodeCharListRetrieved                              ResponseCode = 0x2C
	RespCodeCharListFailed                                 ResponseCode = 0x2D
	RespCodeCharCreateInProgress                           ResponseCode = 0x2E
	RespCodeCharCreateSuccess                              ResponseCode = 0x2F
	RespCodeCharCreateError                                ResponseCode = 0x30
	RespCodeCharCreateFailed                               ResponseCode = 0x31
	RespCodeCharCreateNameInUse                            ResponseCode = 0x32
	RespCodeCharCreateDisabled                             ResponseCode = 0x33
	RespCodeCharCreatePvpTeamsViolation                    ResponseCode = 0x34
	RespCodeCharCreateServerLimit                          ResponseCode = 0x35
	RespCodeCharCreateAccountLimit                         ResponseCode = 0x36
	RespCodeCharCreateServerQueue                          ResponseCode = 0x37
	RespCodeCharCreateOnlyExisting                         ResponseCode = 0x38
	RespCodeCharCreateExpansion                            ResponseCode = 0x39
	RespCodeCharCreateExpansionClass                       ResponseCode = 0x3A
	RespCodeCharCreateLevelRequirement                     ResponseCode = 0x3B
	RespCodeCharCreateUniqueClassLimit                     ResponseCode = 0x3C
	RespCodeCharCreateCharacterInGuild                     ResponseCode = 0x3D
	RespCodeCharCreateRestrictedRaceclass                  ResponseCode = 0x3E
	RespCodeCharCreateCharacterChooseRace                  ResponseCode = 0x3F
	RespCodeCharCreateCharacterArenaLeader                 ResponseCode = 0x40
	RespCodeCharCreateCharacterDeleteMail                  ResponseCode = 0x41
	RespCodeCharCreateCharacterSwapFaction                 ResponseCode = 0x42
	RespCodeCharCreateCharacterRaceOnly                    ResponseCode = 0x43
	RespCodeCharCreateCharacterGoldLimit                   ResponseCode = 0x44
	RespCodeCharCreateForceLogin                           ResponseCode = 0x45
	RespCodeCharDeleteInProgress                           ResponseCode = 0x46
	RespCodeCharDeleteSuccess                              ResponseCode = 0x47
	RespCodeCharDeleteFailed                               ResponseCode = 0x48
	RespCodeCharDeleteFailedLockedForTransfer              ResponseCode = 0x49
	RespCodeCharDeleteFailedGuildLeader                    ResponseCode = 0x4A
	RespCodeCharDeleteFailedArenaCaptain                   ResponseCode = 0x4B
	RespCodeCharLoginInProgress                            ResponseCode = 0x4C
	RespCodeCharLoginSuccess                               ResponseCode = 0x4D
	RespCodeCharLoginNoWorld                               ResponseCode = 0x4E
	RespCodeCharLoginDuplicateCharacter                    ResponseCode = 0x4F
	RespCodeCharLoginNoInstances                           ResponseCode = 0x50
	RespCodeCharLoginFailed                                ResponseCode = 0x51
	RespCodeCharLoginDisabled                              ResponseCode = 0x52
	RespCodeCharLoginNoCharacter                           ResponseCode = 0x53
	RespCodeCharLoginLockedForTransfer                     ResponseCode = 0x54
	RespCodeCharLoginLockedByBilling                       ResponseCode = 0x55
	RespCodeCharLoginLockedByMobileAh                      ResponseCode = 0x56
	RespCodeCharNameSuccess                                ResponseCode = 0x57
	RespCodeCharNameFailure                                ResponseCode = 0x58
	RespCodeCharNameNoName                                 ResponseCode = 0x59
	RespCodeCharNameTooShort                               ResponseCode = 0x5A
	RespCodeCharNameTooLong                                ResponseCode = 0x5B
	RespCodeCharNameInvalidCharacter                       ResponseCode = 0x5C
	RespCodeCharNameMixedLanguages                         ResponseCode = 0x5D
	RespCodeCharNameProfane                                ResponseCode = 0x5E
	RespCodeCharNameReserved                               ResponseCode = 0x5F
	RespCodeCharNameInvalidApostrophe                      ResponseCode = 0x60
	RespCodeCharNameMultipleApostrophes                    ResponseCode = 0x61
	RespCodeCharNameThreeConsecutive                       ResponseCode = 0x62
	RespCodeCharNameInvalidSpace                           ResponseCode = 0x63
	RespCodeCharNameConsecutiveSpaces                      ResponseCode = 0x64
	RespCodeCharNameRussianConsecutiveSilentCharacters     ResponseCode = 0x65
	RespCodeCharNameRussianSilentCharacterAtBeginningOrEnd ResponseCode = 0x66
	RespCodeCharNameDeclensionDoesntMatchBaseName          ResponseCode = 0x67
)
