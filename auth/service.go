package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func (s *service) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
