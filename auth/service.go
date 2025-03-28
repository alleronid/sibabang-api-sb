package auth

import (
	"errors"
	"lanaya/api/app/merchant"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(merchant merchant.Merchant) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("sibabang1945!@#")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(merchant merchant.Merchant) (string, error) {
	payload := jwt.MapClaims{}
	payload["merchant"] = merchant
	payload["exp"] = jwt.NewNumericDate(time.Now().Add(time.Hour * 1)) // Changed to 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	tkn, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		_, ok := tkn.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("Token expired")
		}

		return nil, err
	}

	return tkn, nil
}
