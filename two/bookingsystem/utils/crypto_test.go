package utils

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	verified := "202cb962ac59075b964b07152d234b70"
	v := GetMD5Hash("123")
	assert.Equal(t, verified, v)

}

func TestAesCryptor(t *testing.T) {
	cs := "This is over 32 bytes gooddddddddddddddddddddddddddgoodddddddddddddddddddddddddd"
	etext, err := Encrypt("test", cs)
	assert.NoError(t, err)

	dtext, err := Decrypt(etext, cs)
	assert.NoError(t, err)
	assert.Equal(t, "test", dtext)

	// 考慮如果沒帶加密Key的情況，使用設定檔(預設的)AesKey
	etext1, err := Encrypt("test1", "")
	assert.NoError(t, err)

	dtext1, err := Decrypt(etext1, "")
	assert.NoError(t, err)
	assert.Equal(t, "test1", dtext1)
}

func TestRsaCryptor(t *testing.T) {
	err := InitPubPrvKeyAndAesSert("../conf/rsa_pub.key", "../conf/rsa_prv.key", "tester")
	assert.NoError(t, err)

	testString := "10483a9caa990714c520a2f7c162b554"
	encryptString, err := RsaEncrypt([]byte(testString))
	assert.NoError(t, err)
	log.Printf("%v", encryptString)

	decryptBytes, err := RsaDecrypt(encryptString)
	assert.NoError(t, err)
	assert.Equal(t, testString, string(decryptBytes))

}

func TestMD5Base64Encrypt(t *testing.T) {
	tmp := "Account=test&Agent=65"
	tmprsae, _ := RsaEncrypt([]byte(tmp))
	v := GetMD5Hash(base64.StdEncoding.EncodeToString([]byte(tmprsae)))
	// assert.Equal(t, "", v)

	reqM := map[string]interface{}{
		"Param": map[string]interface{}{
			"Account": "test",
			"Agent":   6,
		},
	}

	bodyBytes, _ := json.Marshal(reqM)
	tmp2rsae, _ := RsaEncrypt(bodyBytes)
	v2 := GetMD5Hash(base64.StdEncoding.EncodeToString([]byte(tmp2rsae)))
	// assert.Equal(t, "", v2)

	assert.Equal(t, v, v2)
}

type requestMap struct {
	Param map[string]Params
	Sign  string
}

type Params struct {
	Variables map[string]interface{}
}
