package tools

import (
	"testing"
)

func TestJwt(t *testing.T) {
	//token, err := CreateToken(1)
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	token := "dadadsadasda"
	claims, err := VerifyToken(token)
	if err != nil {
		t.Log("-----", err)
		return
	}
	t.Log(claims)
}
