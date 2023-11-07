package repository

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) IsValidCredentials(applicationID, username, password string) bool {
	return applicationID == "1" && username == "Dias" && password == "dias111"
}

func (r *Repository) GenerateJWTToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *Repository) ValidateJWTToken(applicationID, token string) (bool, string) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return false, ""
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		username, usernameOk := claims["username"].(string)
		expiration, expOk := claims["exp"].(float64)

		if !usernameOk || !expOk {
			fmt.Println("Missing claims in the token")
			return false, ""
		}

		if applicationID != "1" {
			return false, ""
		}

		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Expiration: %f\n", expiration)

		expirationTime := time.Unix(int64(expiration), 0)

		return true, expirationTime.Format(time.RFC3339)
	}

	fmt.Println("Token is not valid or claims do not match")
	return false, ""
}
