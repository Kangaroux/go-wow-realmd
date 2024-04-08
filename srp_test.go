package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustDecodeHex(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Test_passVerify(t *testing.T) {
	type testCase struct {
		username string
		password string
		salt     string
		expected string
	}

	// First 10 testCases from:
	// https://gtker.com/implementation-guide-for-the-world-of-warcraft-flavor-of-srp6/verification_values/calculate_v_values.txt
	testCases := []testCase{
		{"LF2BGFQIFQ3HZ1ZF", "MVRVMUJFWRA0IBVK", "AFE5D28E925DBB3DAFED5D91ACA0928940E8FBFEF2D2A3CC154ADA0FE6ABEF6F", "21B4153B0A938D0A69D28F2690CC3F79A99A13C40CACB525B3B79D4201EB33FF"},
		{"TL2XRNJ09DGYK9MQ", "MHMUJMEIX8LGRKZ6", "FCAD0C7D6BD03F3DCE8AEB9C64FF1D1550EDB77BF9F1EF2CE437D83C9150DEC1", "0C12CF1EB9855705B21B3E0EC9AB2807A6F660659C7C62040E43AE3908B9B827"},
		{"1LGT7JVQ29N79FIG", "IL4ZXPC440CFFK9A", "5DCC9CFAD04F1CDFDE58A9D76211D79D7A8EFAC4ECC8EFA43FF4AC52FBB0CE9A", "18AD8E1169E8DDFF4CD1513E35958B8D929AC09CCA3C7D65120DE860A711F1DC"},
		{"8ACOYJOYZNKC7F9M", "YJPND51QB9HQDKJ2", "4ABA2F715C0E92FD73CA154F6820D8A73D268B0D76BFE952F8EA5DFECB152DD4", "13951E42C3C1126E30ED6142A1BC9FD98C1AC67548102137A161B64AD398AAB9"},
		{"FB81TKOIMJ0ZHL5W", "J2ZCUVI42BJC1K7I", "B0E680DD2DBB8EB6EED4624E5CC8E78DAC57F4CCAACB04AEC86AF48B4A45F38E", "4D4AD0E1D19E0527C54C79497FD50602A9A0E30D8EBEA005CAE3DF9E098884CE"},
		{"ASAXNEO9O9O9PC7E", "3R72J2RJAWB7VE9T", "9AAEC1D1EDD9E0FF97B0400CCD744E3E3C15BE4C62CB82A1AB2953E5595C9B31", "77A0AACB0427381D67A832C2D2704D762DB01D558FDE2D79EEBB40697D67D788"},
		{"NW8FFHXOI3D5TDH4", "HP4RKUMSO13L16ZJ", "7E7BFDDAB232AB9328AC2C9BB3B40BDBACAB9CCCFE9BDC6EEC704BA2B65443A6", "6ACFC3A152E23BD367949B93D4E460447562FC8921B74DF28317342720736B72"},
		{"RDL7YCLB6QL3Y40X", "C63MYZK061FQ1TAJ", "AA08C8DCE74FE0DECCCA7EC3FABCA6EA6C015CA9B97C750A88F03EB9AB478ABF", "3C5C6F755D6FEFBD35C7A520C4FD33B623AC27B4039D68F90F14CF6E631235E9"},
		{"8F6TYYE7OXZUMYEM", "8V1QAMPZ51FBO6NR", "09CE0BFECFAE1C2EA79DD8A87B4CCB4FF2F4EA4F8E35A6889580AEBAE1F95925", "62C100DC3AE4BECFB1E58445847F95BF7E81CCCD65B6DB3F1EC2ED8C87BDE344"},
		{"D61MXHXTJTHM2N7Y", "2VL37KANR9NQHND6", "0B8DE644B7ABDBEBDC615E89C4E56A2FD3D171C5F8E9AFC55A5D8FFB1ABFA6AA", "1E908AD2EE91F76AAB18BC6163B2A55996A1E86AD161F33F9EC77D30FDD7F2BF"},
	}

	for _, tc := range testCases {
		salt := ReverseBytes(mustDecodeHex(tc.salt))
		expected := ReverseBytes(mustDecodeHex(tc.expected))
		assert.Equal(t, expected, passVerify(tc.username, tc.password, salt))
	}
}

func Test_calcX(t *testing.T) {
	type testCase struct {
		salt     string
		expected string
	}

	// First 10 testCases from:
	// https://gtker.com/implementation-guide-for-the-world-of-warcraft-flavor-of-srp6/verification_values/calculate_x_salt_values.txt
	testCases := []testCase{
		{"CAC94AF32D817BA64B13F18FDEDEF92AD4ED7EF7AB0E19E9F2AE13C828AEAF57", "D927E98BE3E9AF84FDC99DE9034F8E70ED7E90D6"},
		{"0CE6107EAC9DBBED6C7AD14EB0D4DDFBAAB82BE636F4EFF23B2A5E39EEC16A4B", "D30F36712DBB25FFAD213AE1F8E6A836FC9512DC"},
		{"422E5E9DE1CD44BF81D64FCB8BAE78F23C66C3E323D3DBC0FEEFC991475CECFA", "FC4EC39D5AAD4875F3B706EB3D10DAE7C7742273"},
		{"E0EDCFCA5AF8A23FFE02D505EB4BD34C0BAD4228DC3EC0975EADF18214F9C960", "0BF89E0845FC52040C8227D9F54C6D0E800B3C2B"},
		{"E5E8E756C71E1674CFD385FBB1FCD9A9A19A6454D6CF6C0DEBF9B6DDCD46B2A3", "BFEF246D5F39B46B2CF1D789748EF6885EE1854B"},
		{"B6DDDFBEBB9BBA928A3E76150BF1BFA306DA6BDAB2BD7043B87D0DEA3C45ADE6", "0CA7F727D8DBBCACB51E51E1E848C8EA12A53399"},
		{"EF969E1EE76DCCEFCF05288F5EBB2DB13A429DD8B35BAE2EE2E603B2EB1F1F0F", "9D4D2352AD71BC4256225FF7C18902548F15B13E"},
		{"96E1AA3F3DA9CB15AD05D304DB5E36703E2DEFCFD0FCC1E6452B6306C89F9DC6", "35FFFEEC0C5D4623D693B55B4C39A6648417A575"},
		{"FFCDF9AED1A5EC3D56CB0BFE6026A5CFA6CCCDDCCEBFF6FC7AC0F5A71ACC54E8", "5840003306F80C70FDBD0E692F3866629E6CB92B"},
		{"CADD8EEBE2FD8BCACDCDCEE3BEA0BF4F0928EC463E70E07CCAA05CD91A44DC1F", "73AB2AACD12BB1A5E1AF467CF2F46F0B76F5418C"},
	}

	username := "USERNAME123"
	password := "PASSWORD123"

	for _, tc := range testCases {
		salt := ReverseBytes(mustDecodeHex(tc.salt))
		expected := ReverseBytes(mustDecodeHex(tc.expected))
		assert.Equal(t, expected, calcX(username, password, salt))
	}
}

func Test_calcServerPublicKey(t *testing.T) {
	type testCase struct {
		verifier         string
		serverPrivateKey string
		expected         string
	}

	// First 10 testCases from:
	// https://gtker.com/implementation-guide-for-the-world-of-warcraft-flavor-of-srp6/verification_values/calculate_B_values.txt
	testCases := []testCase{
		{"870A98A3DA8CCAFE6B2F4B0C43A022A0C6CEF4374BA4A50CEBF3FACA60237DC4", "ACDCB7CB1DE67DB1D5E0A37DAE80068BCCE062AE0EDA0CBEADF560BCDAE6D6B9", "85A204C987B68764FA69C523E32B940D1E1822B9E0F134FDC5086B1408A2BB43"},
		{"4FAA98BAEA63EC2CF1DF81A93B3617C5CD3A2FEB33A3B53ABCD9A09F2B11A106", "AEA005D3A5BE7692EC111F4A67B2E23C8F3FBA772291ECBDFABC6195DFC33E09", "5FB8F971FC6080ED82F0E7358F6917BABA99AAAFC783AA863AE0356E3FEF7D81"},
		{"AAB5B6E2CC99A14E1C09344ADD8E9D8303C5BFE2437FA2BDEE263ACA569BFDAB", "9B4A2FC8D11D2A3C4B32D0F1EE9904A208A945A8DF9AB8924A7BF8CA2CDA4E7F", "44F0672ED04E3EDF9EF9806AE8A241DC383A3A81741C4C034DBE43756AC76BF1"},
		{"EDDBF853DFF2A4DFEBB3B33E50156028B21AE6362C263DB28BC090DFD0D5FC11", "C04E0676EC36E85B6ED6EF4A6F835AEBCEBFB3A0AADCD1C15B516BABC9ED4AB6", "4CCCEA5FEA33D64E837C99F547B9EA5A317B6D6B5ECF592ABBD8AE2183D8101D"},
		{"E70FC28C7EE166B557A8CD7FCADBCDE42BFA1FA4761AD10B60BB00EFA182EDB4", "D323EF4E3E2ADFFD5CE6A4BFE139DC3A97A8A4DEB0A97ADD82E576D541EB4BF9", "319ACB101913BEF86EB27FA0FC744BB1C12171D80E612B43DF4C0BDB49549DCE"},
		{"46FC1ADBBF60D0EB1BE5C9CDB82FCF33E3FBB40EB62CC4CCDF32EC322AE7A37F", "A616D9CC4A46BA108BA428CDDC0D47583CD7BAEF59D92FB5CAEEEFC3DD3FBCB1", "150C8D487F9E3C66A239F8CDD81ECE6992E2A9CED7F52EC89074F3BBDC059BF6"},
		{"7AC81FC7C98ADB3DDACBEF8DDEEA6D4C7F539AA70CE6D5BC1DDE39AA2813A3D1", "5A48582BBDCED65F683EC595C9AE5397FA90EF0FD62F863CB69A7CCDDCFFCC1E", "4DC7461B056EB87446441E0AA4539E8A4623F349975F5B9D34F40328BB92F9C4"},
		{"7FC9BE2EB6490FA71B9B30ECBF2F0002EE8DCDBF6A0E133ED3C300EA38C09BD4", "C58E11AA32C0A7923CBA29AE19BC13596A25DDFBE417F4613B77EB579F146EC5", "82D955E0F3BF897BD3723DC8EB2F413456FB75130E2BCE9FF07F57C3584B493C"},
		{"5C00BEF9D8E4ED8FB1E0E48CE3A128BFEACAECBDFB2815CBC2244D24BC73D52A", "EC3FCBE21DA1852D0AD40D06F60F3B940B0A3DD08A37A3BA5A55411CBEAA816E", "25088266E411E5746120C7E4D9C0D99F91A009CC5F262080BA744E3B4317978B"},
		{"1F9147A8C9A5DCE13584A7C26AEC5C0E39917E7D0609CA39EE628990639F9B2B", "59D2AFE6AFE3BAFA05BECD25DBA7B8AFD1BBF9551BE68C1E6D4AFB7DA37A25B0", "54172E50A01EDA99C1286F9FD0E5F74A786090F7BE0EB7DC9F3B45BBE7AB0D82"},
	}

	for _, tc := range testCases {
		verifier := mustDecodeHex(tc.verifier)
		serverPrivateKey := mustDecodeHex(tc.serverPrivateKey)
		expected := mustDecodeHex(tc.expected)
		assert.Equal(t, expected, calcServerPublicKey(verifier, serverPrivateKey))
	}
}

func Test_calcClientSKey(t *testing.T) {
	type testCase struct {
		serverPublicKey  string
		clientPrivateKey string
		x                string
		u                string
		expected         string
	}

	// First 10 testCases from:
	// https://gtker.com/implementation-guide-for-the-world-of-warcraft-flavor-of-srp6/verification_values/calculate_client_S_values.txt
	testCases := []testCase{
		{"E232D2C71AD1BF58DB9F7DBE51FFE271B6BDC61524F2E6B32ABFFFCAB09D09AB", "FC3D610C4E2CEC5ECC7E47344D0ED81D2ACB938AB198EC7E2ED474AEFCC3ABD1", "A4A7CB7DFBE00D26EE06F6B3DACC51E5779D7E8B", "FDAFAEF0E77F0FE1BD2956CF1820D4BC964E5283", "3898DF5193EA6AA8111524A253DB480A51EA6160D1E41BC4B662420299B4A435"},
		{"8CBDF6ADB7AB7C440ADF2A6EF35504A16D0CFC1D6BDB2B9D490A9FE0DBFC2ED4", "32B7EF3E95B0F0B8DDCDAAEDFA8763B50CA388D37CE6DC9EECDB621A5C844B25", "DE80EB1158911D56B8FDC761DC0AEA0C7C9D7EFD", "BAFAEBD30DC342C57F28C0CDB68DFD238AB2CACF", "879E097E9C97C6F5AB5970FBAD87C038405E0CBF401F8864FD0DADD5FAF2CB76"},
		{"96AEAF1CEB90FED799A7419448CBC2BADCF71E9C252BE5DAD29DA4E7DE97ED7F", "B079AA9EBE8B97BBD8B1EBACF60DF6EC116D92FCAEC2CAEB4DDCEADCF9CFCAE1", "82D4EC32D8ACFC69D1DCFAEC6A6BF3EEC24FB004", "E11CD6C189EAEAAC96813EDBDEC99C0EA229BA66", "355035CD9F74FDD39EFFEBECF7F2876B44FB640F122A4E00F57DCC184BF47AB4"},
		{"9DB2DFBADCBFBF1366EEEBECE5D0ACCB1FA21E7BF5BECBD147F2DBFC1FDF8F25", "20F2A5C85E0EB9DF17CE009BAEC094D0DF2F2B50C1BBBFB1A8304E269FF5BC14", "2AAED45224AB1FD6BA3FA44FFD92ADDFEFDED2CE", "316B6A64F2EEFDBFDDC5AEA86C88FD78A3AFBEBF", "08F7FB15DC7FE3142AC04F4E35E195A378ED73A9327A69DF4C0A49C3FB77F3BA"},
		{"AEB4DF97B83CCC5E799B0EFEE07A46FCCAA25EDCEA70D5BE2BD9FF0EAE7E848B", "5917F7199EDFF0DEFB4DAE2B6FA9A63BEF602336ADCC76A28F46ABFCF83DE4FB", "BDCD64DEDDED6D8B27A84BFE51E9A0CBEBB0CACE", "5FED0CCDBD59A2DDDFD76AF1FAEE7EC690A7260C", "6206D9824B99195D371347F3A1A2B994BC325545A40D746CE465E63313283078"},
		{"626FE1E5F6F2B87DF7AF3B9AE7E50FCEB3E771ACAB9A39B99208A48A6B90BC8C", "473F4D6DD13D104DBA7952FFEDAAB73E41D9F39F0EC19AF2C4980B4BFEBAB051", "DCD6D4BAF581E7AA64E42AA3BDC2CD6D440C7A96", "9B88AB34F7A32BAAF9DBDBEF95DA59ACE2FBFD61", "57998A49E9375A3C8FAD71893E212660C444D6CB0EE8922FC77E0C23C6016249"},
		{"520A747330C4EDD2FFD563F84E8CA17E6DB61A9D6EAB78EA71CD01F194FA5693", "FEDBD9D32D5CAD99BDCBE63BBCCACECC6D9A38CDDCD09CE0BAAC6ABFC2437F2B", "6AAC18EC61ED876BED4CFAE7B60EE40E2BACF10E", "A1B4594FDDE0CB9DA954F6FDA0A5F8CBFCBEAEB0", "5CEE5BB2CBB2CEE4FD23C70DCD67C3716CCBB5683070DB1E01BC605775FB24BA"},
		{"2EF4CB8BFA59DCBBE8EFC56CFBDBF9C1B6CDD93DC602A0D75E8A27F4B106E6E3", "F4C885E10B7AA0FFABAFEE608D0C79E20A0ED16E1C0AA302631DD2789DF36EFE", "ACC11325FCB99550E38752FAAF2B401AB389E58F", "79CDFBC14D6D3A0C2AAFE32A06AEB1CD4FCB6E96", "2CD30D2425833FF323F78C8540C0ED06C6642242DD30E1F253B3D187594C10EA"},
		{"39FDCCEFBDD50B2ADE8F9FAC9FBCDCEDB7F5C0F420FAFA5EBC9731CF8EFDE308", "F17818C580D097D71ECC3DCDFEA47DBBDDA4DD24CC58D7123D36F0ED99559518", "BDA1980EC03D1F956FD9D1AFFCEE4016AFF00FBD", "923BCA0E6A58E404A9006F425F09E8A03A9B5A46", "65234100D8A2F780FBB953F996E0AF71269BE5313328A8EF9ED490916ECEBAE7"},
		{"23ABACECC3E4F5BD9C078D0BAD9BCA631996FA7AE3A038EB448FBC4D525D1BBC", "8452AFD7D7E005CB4B7AF6F5CD7ACCFFE0F9AB8DDE19A4B224D8376A70CCA77E", "48DFEDFF6F76DAF8F87A6DBFC02BAC82D72C479B", "BEEEC0790DDFE96CAB0E9FCAFF50DDE76B74D70F", "3F1025C1A869131507BFA7AAB83667B26CBDEFAEAB1DDD088B7D0045470532D8"},
	}

	for _, tc := range testCases {
		serverPublicKey := mustDecodeHex(tc.serverPublicKey)
		clientPrivateKey := mustDecodeHex(tc.clientPrivateKey)
		x := mustDecodeHex(tc.x)
		u := mustDecodeHex(tc.u)
		expected := mustDecodeHex(tc.expected)
		assert.Equal(t, expected, calcClientSKey(clientPrivateKey, serverPublicKey, x, u))
	}
}

func Test_calcServerSKey(t *testing.T) {
	type testCase struct {
		clientPublicKey  string
		verifier         string
		u                string
		serverPrivateKey string
		expected         string
	}

	// First 10 testCases from:
	// https://gtker.com/implementation-guide-for-the-world-of-warcraft-flavor-of-srp6/verification_values/calculate_S_values.txt
	testCases := []testCase{
		{"51CCDDFACF7F960EDF5030F09F0B033C0D08DB1E43FCBA3A92ABB4BE3535D1DB", "6FC7D4ACFCFFFDCF780EE9BBD17AE507FFCDF586F83B2C9AEE2198F195DB3AB5", "F9CEDDD82E776BEDB1A94852A9A7FFA4FCADD5DE", "A5DBBFCB4C7A1B7C3041CAC9DDBD36CD646F9FBABDAD66A019BCBB8FEDF2FAAE", "3503B289A60D6DD59EBD6FD88DF24836833433E39048ECAFF7E887313554F85C"},
		{"B4AFA1EDB6A84ACF163FDC316DEB2D8CCDC5C39C092D7AEEA6C072220CC2C1C9", "E0A62EE6EB0CA1BFBAF3C5ACE3ABE1EDEC9E9780C57AFD6BB5EBDBAB8ED2FFE5", "D4BF6B2CD01F53DAEE447F695F20E0FD65FCDBF2", "E29CB6EDA5181B3ADBF32A2EBE28FE1B2BD5FC5A1E2DF2CBB68F42CBBFAC1B5D", "58C1574090182445A6FFB54D054923B949937D943E9C72C90FD1A5040395F80E"},
		{"BD5645CCE1BC10FDD6A62D503CB069C4BEA53739DAEDBA4BFEE1F716D9DBC1A9", "FBD507EF153C8B6FEC9A9AE71B1C72817E80C8FACFFED8ECFEBCBEE84DDD9E63", "7ECAC7862C45B4EAB1F8AFFAAB24C2FE34516BAA", "772EBD4AC253CE7FDE0145BEAF6AE437F1E63CE4BCCF3B4B21DCCE3F0CAD5DCB", "12BB8C26CABE175DEBD6AF751A99641A58BEDDD403ACDD2FB540232654D84527"},
		{"F703BC2BB45A04D0FDF25FA19671BAA7AFB387CC87E22D1B3CCDDABBF2E90807", "B3BABDDAED8607881BF14151AADD0CC9AB8A44A96EDD7DDBAAE8ABECC8C7EE7E", "BEB07F782CDAAF3C6B50D7BDD9E5ED1E57CCEDC4", "73CBF92ADCABDB427EFAD7BB5FF1152DC0CBB8FE5A4A7CDE68E07932BBDDA3BD", "3AAAD610E538D9DBBCEA0C22355A96917223AACDE0A6BE9366BE1685C1D10E3A"},
		{"6E545CB8DBBEDF8CAD8EA9D9DCDAA56F0E7DDF7575B3A9E2E7AFEEEBD6A0E0B6", "BFAB3DF87E36A3A46C06AC8BBBBD9F0BFB3C63AD6DAEC2F2D74D7A4DB9E8038E", "CD84B91A5BC090A8FED7B0DFF74682E15B53BABC", "F4974A1EE11E02AC5F6988D93D0F0DDFDC5DE5FCAD7A1F97AAE31D0BE34EF279", "81840648ACE795176BDC6210270DB9C88B880D969C270140DD74FB14684CF2AA"},
		{"FABC2F0ADC54FC1382AD87BA9EA7BAF4EF6F9EA1C4BAE8E955DB82D7FFB6D0B1", "B78575C969BCFCFC52FCAECC4BCFB6FD4488ACF87EDAEBDCC46F879F3CDF83CA", "90EFDFE37AEA142AF6361D006CE82AF6CDEC74B2", "A501A55F82DADCEA69277C6B6DCA41D6EBAE24BAA51F18635B8B894D8762A78A", "1E68E8ACC369838FE52D4A64CB7A6F108CBF5725A496870289682DDA6615BAD7"},
		{"3B1EEFCDFE9C0BFFE0EF9AEE7E2D0D0FEB65D8EAAEED56B55BFEDDE2E5FDCC03", "8D815F84ACA1A32EDF4FDDAA7376EE19DE9C54021EDED9C502B08A0FCFFE2FFA", "EE5FAD35AEE31F446DEEFDBA4CBEB3BE3ABBBA5F", "8919168B111BADFB6AFB3B5BCFCACCFE7DCCEDFEC53BBD3ACB138F3DC730AD96", "711FDE0840BA3F789A79A20F7DCB1E16C6333F093C48385884EB15313B252DB3"},
		{"CBBDF1FDA1C1A9EE5F8A6FAE2AF063CDC90E4B69160AA18D479CF72DC0E1B1F4", "E8D66FB95DB077AC85C5C3DF222C9A54C0B90673FACAB6F4E890549D45F28DBC", "AAAF5F1BD3C5EE0BEADAE54BF6EEA67D296BC5AB", "C7EFA31A4632FE1CB02A88302FE9D86E891AC8E85FDFA2B36FA2CB80E34A1A25", "190496CE448901B4C6EC0C9635EC8F8CFCBAE947AC0099AE3716B69612D4E8CF"},
		{"849A2A6E1BC4FDB28E89FBE0DCF2BDC16D0B6A45ABBADADC81AD4CC6F1CCBFCC", "E06B2E9A68C076B1861106E03F1A390A0A72FCC0FF7F841FBC9E8AA545AF9FC5", "1CCA72DCCBDC3145956D2E90CBFC85E3E03B3DCC", "D8E7CBD8D5EBEF29C104E7CF87E3CE54F9A0FAE1E8E5CBE5AFB12F9CB8AF04C0", "788B3608A68B14EDBE81DDF342953A9571A071F1D3A3F5EF6AAC334C2B01910B"},
		{"5DBDC58BE03004E0519133C0F76CBF26AB9BAC3D21FA4ABF9BECC454C523BF60", "B3A31993AD788E13EA6ADBCEEDC8F4D42CDD89358AAC93C6FBE7ECE736ABEAA3", "5E3FE7EEB927F0E1F121AA4DC6BC4CA6A2AEDDDF", "99D7BB4BC18E8D64D71DBE6CA184CFCEECC6CCBCB642CE291F5E98026DB1C0A5", "5AC6454FC58F5A1AAF1EA9DA703DA9B2A907B1DF60497C71EFA09CF8BE2C74D1"},
	}

	for _, tc := range testCases {
		clientPublicKey := mustDecodeHex(tc.clientPublicKey)
		verifier := mustDecodeHex(tc.verifier)
		u := mustDecodeHex(tc.u)
		serverPrivateKey := mustDecodeHex(tc.serverPrivateKey)
		expected := mustDecodeHex(tc.expected)
		assert.Equal(t, expected, calcServerSKey(clientPublicKey, verifier, u, serverPrivateKey))
	}
}

func Test_calcU(t *testing.T) {
	type testCase struct {
		clientPublicKey string
		serverPublicKey string
		expected        string
	}

	// First 10 testCases from:
	// https://gtker.com/implementation-guide-for-the-world-of-warcraft-flavor-of-srp6/verification_values/calculate_u_values.txt
	testCases := []testCase{
		{"6FCEEEE7D40AAF0C7A08DFE1EFD3FCE80A152AA436CECB77FC06DAF9E9E5BDF3", "F8CD769BDE603FC8F48B9BE7C5BEAAA7BD597ABDBDAC1AEFCACF0EE13443A3B9", "1309BD7851A1A505B95D6F60A8D884133458D24E"},
		{"EC1E5C8F333A31944E20E40648CE7A60C7C29EEA8DE185FB7CB7485AEE19EAAD", "A43CC0A6B05DCDA9CB3EA4B267D3773ECCC4086C1BB93FCDE35CB3E41373413C", "CB81A3170E6AA5E605B99F374062C33D1DAD84F2"},
		{"C2EFF57445E0DFB4B8CD1905DD2F75B75B1BF9618203810B4B4CCEFBFC7AFE24", "FA902EB5E315BDE5E17AB08DACB0BCBC58ECA0CBDACEE558FB18BCC1CA04ACFE", "1A28E10049E7156056520157DF9913C4E4A39D71"},
		{"FEBE2AEAFEF1D5DB5BBD7F7189AD33DFE8F4DCCBAAA55E2BAA8D9CB8F5B36C6F", "923C1A70609A1BE21B92FF54E62D10AEFA0FECF9FAF0A02DA411DE5B7BF41DBA", "D451CE4F977656D65329A029596DEB61BCBE4383"},
		{"87CDEFAC8532B1C5FF81AC3B4E0A9DFA2B1AC8BDFFCDDA4E3D3B1F2D54EA474B", "E62CF5B1C296BC93B8FDCDD93CF65DA1539AEB98A0FF73CCC442D5B2C4D9328B", "E4CB8A8A4F20D540A4153D26CF0FA6FC0F9938FF"},
		{"3FE5D456A24D0DC6B5FD652C129E2AC2E826D03B6CB43ACB73A4F7F3D3B2D8EA", "CF63BEFB67B231E553FDD0072FFF615DDDEAEFAA4992FAC81DCB175F47D711A6", "B33E75C7F4151D06350AF51361E98AA1D9B6E208"},
		{"6FF43F21C63EB88EA93BBA014FDDDF7CECB5BBEC975FB80E516ACA2BD15062C5", "11DA67B9CA120B8BB0DCFCA67FDD1C17B2184A6EA7EE40CD0E486F48CC35AC9A", "31C872B7EA603BA6B4B3E7966DA4A66E6C3E3B06"},
		{"1F8ADCFF3CCEEF2D71AACCED3E0405F2D3B26FA8DDFFFEE22170A0D30FBFCCD4", "6EADD50B7AFDB77A267E46AF3C5A97AA1DC5FD0F3D0D8EFA3EBC94A393241034", "FA07C28229516786942D55A9ECBAA553213AC8E5"},
		{"F42A43AB607087C7AEBFB49CFC536DA0F451E2DA9126771A8D1ACB088BF2ABBD", "103DDDC6DB0BAD0F872E87BFEBEDB3EFB1A5F6DAC0DCCA7A3191CA60E7E12625", "462E6586CE76D17DA547E4782CE77B328C1F777D"},
		{"6F8DA5EC2D5D86BED0F1CC04DE9981CEBFAF2579B4777B268D80AAAACFF117F9", "1B3529F8FBCAFEB2379DEF59CBB7EFAB4AEFCCCB40D79D6980B5EE3FE36AD46F", "5097ED6E620C301D9E931D190019545DE4892D03"},
	}

	for _, tc := range testCases {
		clientPublicKey := ReverseBytes(mustDecodeHex(tc.clientPublicKey))
		serverPublicKey := ReverseBytes(mustDecodeHex(tc.serverPublicKey))
		expected := ReverseBytes(mustDecodeHex(tc.expected))
		assert.Equal(t, expected, calcU(clientPublicKey, serverPublicKey))
	}
}
