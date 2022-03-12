package jwt

import (
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

var cnf = newConf()

type conf struct {
	AccessDuration  time.Duration
	RefreshDuration time.Duration
	Alg             jwt.SigningMethod

	secret string
}

func newConf() *conf {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatalf("error while getting secret key: empty secret")
	}
	return &conf{
		AccessDuration:  time.Hour,
		RefreshDuration: time.Hour * 24 * 7,
		Alg:             jwt.SigningMethodHS512,
		secret:          secret,
	}
}

func (c *conf) GetSecret() string {
	return c.secret
}
