package tokenizer

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Tokenizer struct {
	ServiceTokens       map[string]map[string]string
	ImpersonationTokens map[string]map[string]string
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		ServiceTokens:       make(map[string]map[string]string),
		ImpersonationTokens: make(map[string]map[string]string),
	}
}

func (r *Tokenizer) GenerateJWTToken(username, applicationID string) (string, error) {

	if len(r.ServiceTokens[applicationID]) > 0 {
		delete(r.ServiceTokens[applicationID], username)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()
	claims["appid"] = applicationID

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	r.ServiceTokens[applicationID] = map[string]string{username: tokenString}

	return tokenString, nil

}

func (r *Tokenizer) GetTokenExpiration(tokenString string) (string, error) {
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

func (r *Tokenizer) ValidateJWTToken(applicationID, token string) (bool, string) {
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

	username, usernameOk := r.getUsernameFromToken(parsedToken)
	if !usernameOk {
		fmt.Println("Username not found in token claims")
		return false, ""
	}
	if storedToken, exists := r.ServiceTokens[applicationID][username]; exists && token == storedToken {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			expiration, expOk := claims["exp"].(float64)
			if !expOk {
				fmt.Println("Missing 'exp' claim in the token")
				return false, ""
			}

			fmt.Printf("Expiration: %f\n", expiration)

			expirationTime := time.Unix(int64(expiration), 0)

			return true, expirationTime.Format(time.RFC3339)
		}

		fmt.Println("Token is not valid or claims do not match")
		return false, ""
	}

	fmt.Println("Token not found in ServiceTokens map or does not match")
	return false, ""
}

func (r *Tokenizer) getUsernameFromToken(parsedToken *jwt.Token) (string, bool) {
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}

	username, usernameOk := claims["username"].(string)
	return username, usernameOk
}

func (r *Tokenizer) GenerateImpersonationToken(usertype, value string) (string, error) {
	if len(r.ImpersonationTokens[usertype]) > 0 {
		delete(r.ImpersonationTokens[usertype], value)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userType"] = usertype
	claims["Value"] = value
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))

	r.ImpersonationTokens[usertype] = map[string]string{value: tokenString}

	if err != nil {
		return "", err
	}

	r.ImpersonationTokens[usertype][value] = tokenString

	return tokenString, nil
}

func (r *Tokenizer) GenerateServiceAndImpersonationToken(applicationID, username, usertype, value string) (string, error) {
	if len(r.ImpersonationTokens[usertype]) > 0 {
		delete(r.ImpersonationTokens[usertype], value)
	}

	if len(r.ImpersonationTokens[usertype]) > 0 {
		delete(r.ImpersonationTokens[usertype], value)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["appid"] = applicationID
	claims["userType"] = usertype
	claims["Value"] = value
	claims["exp"] = time.Now().Add(90 * time.Minute).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	r.ServiceTokens[applicationID][username] = tokenString
	r.ImpersonationTokens[usertype][value] = tokenString

	return tokenString, nil
}
