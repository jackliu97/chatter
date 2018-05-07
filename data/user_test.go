package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test ./data -run Test_VerifyPassword_Pass -v
func Test_VerifyPassword_Pass(t *testing.T) {
	testPassword := "jackliu!123"
	user, _ := MakeUser("jackliu", testPassword)

	assert.Nil(t, user.VerifyPassword(testPassword))
}

// go test ./data -run Test_VerifyPassword_Fail -v
func Test_VerifyPassword_Fail(t *testing.T) {
	testPassword := "jackliu!123"
	testPassword2 := "jackliu!12345"
	user, _ := MakeUser("jackliu", testPassword)

	assert.Error(t, user.VerifyPassword(testPassword2))
}
