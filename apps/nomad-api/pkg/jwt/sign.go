package jwt

import (
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

func Sign(signingKey []byte, id string, roles []string, expiresAt *time.Time) (*string, error) {
	expirationTime := func() int64 {
		if expiresAt != nil {
			return expiresAt.Unix()
		}
		return time.Now().Add(30 * (24 * time.Hour)).Unix()
	}()

	claims := &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: expirationTime,
		Audience:  strings.Join(roles, ","),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
