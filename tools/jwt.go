package tools

import (
	"crypto/rsa"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zhangshanwen/the_one/initialize/conf"
	l "github.com/zhangshanwen/the_one/initialize/logger"
)

var (
	defaultExpiresTimes = 12 * time.Hour
	defaultTokenType    = "the_one"
	path                = "rsa"
	privateKey          *rsa.PrivateKey
	publicKey           *rsa.PublicKey
	method              = jwt.SigningMethodRS256 //默认256
)

type Payload struct {
	Uid       int
	TokenType string
}

type Claims struct {
	*jwt.StandardClaims
	Payload
}

func load() {
	var err error
	var privateBytes, publicBytes []byte
	path += string(os.PathSeparator)
	privateBytes, err = ioutil.ReadFile(path + "the_one.rsa")
	if err != nil {
		l.Logger.Panic(err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		l.Logger.Fatalf("[initKeys]: %s\n", err)
	}
	publicBytes, err = ioutil.ReadFile(path + "the_one.rsa.pub")
	if err != nil {
		l.Logger.Panic(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		l.Logger.Fatalf("[initKeys]: %s\n", err)
	}

}
func CreateToken(uid int) (token string, err error) {
	var expiresAt int64 //过期时间
	if conf.C.Authorization.ExpireHour == 0 {
		expiresAt = time.Now().Add(defaultExpiresTimes).Unix()
	} else {
		expiresAt = time.Now().Add(time.Duration(conf.C.Authorization.ExpireHour) * time.Hour).Unix()
	}
	t := jwt.NewWithClaims(method, Claims{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		Payload{uid, defaultTokenType},
	})
	return t.SignedString(privateKey)
}
func VerifyToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return publicKey, nil
	})
	if err != nil {
		return
	}
	claims = token.Claims.(*Claims)
	return
}

func init() {
	load()
}
