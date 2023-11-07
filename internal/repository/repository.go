package repository

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Repository struct {
	UserTokens map[string]map[string]string
}

func NewRepository() *Repository {
	return &Repository{
		UserTokens: make(map[string]map[string]string),
	}
}

func (r *Repository) GenerateJWTToken(username, applicationID string) (string, error) {
	if userTokens, ok := r.UserTokens[applicationID]; ok {
		if existingToken, ok := userTokens[username]; ok {
			expiration, err := r.GetTokenExpiration(existingToken)
			if err != nil {
				return "", err
			}

			expirationTime, _ := time.Parse(time.RFC3339, expiration)
			if time.Now().Before(expirationTime) {
				return existingToken, nil
			}
		}
	} else {
		r.UserTokens[applicationID] = make(map[string]string)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))

	if err != nil {
		return "", err
	}

	r.UserTokens[applicationID][username] = tokenString

	return tokenString, nil
}

func (r *Repository) GetTokenExpiration(tokenString string) (string, error) {
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		expiration, expOk := claims["exp"].(float64)
		if !expOk {
			return "", fmt.Errorf("missing 'exp' claim in the token")
		}

		expirationTime := time.Unix(int64(expiration), 0)
		return expirationTime.Format(time.RFC3339), nil
	}

	return "", fmt.Errorf("invalid token claims")
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

		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Expiration: %f\n", expiration)

		expirationTime := time.Unix(int64(expiration), 0)

		return true, expirationTime.Format(time.RFC3339)
	}

	fmt.Println("Token is not valid or claims do not match")
	return false, ""
}
