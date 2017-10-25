package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/jwt-go"
)

type EncodeInfo struct {
	TestString string `json:"testString"`
	jwt.StandardClaims
}

// genNewToken by jwt-go
func genClairToken(targetID int64, appToken, deviceID string) (string, error) {

	claims := EncodeInfo{}
	claims.TestString = appToken
	claims.Id = fmt.Sprintf("%d", targetID)
	claims.Audience = deviceID
	claims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	signingKey := []byte("AllYourBase")
	tokenString, err := token.SignedString(signingKey)

	fmt.Println(claims)

	return tokenString, err
}

func parseClairToken(clairToken string) (targetID int64, appToken, deviceID string, timeStamp int64, err error) {

	type EncodeInfo struct {
		TestString string `json:"testString"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(clairToken, &EncodeInfo{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return 0, "", "", 0, err
	}

	if claims, ok := token.Claims.(*EncodeInfo); ok && token.Valid {

		targetID, _ := strconv.ParseInt(claims.Id, 10, 64)
		return targetID, claims.TestString, claims.Audience, claims.IssuedAt, nil

	} else {
		err = errors.New("ParseClairToken Error")
	}

	return 0, "", "", 0, err
}
