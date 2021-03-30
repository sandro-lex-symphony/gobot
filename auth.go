package gobot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// gen JWT from private key
// the jwt is used to auth with pod
func GenAuthJWT(keyPath, botname string) string {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   botname,
	}

	alg := jwt.GetSigningMethod("RS512")
	token := jwt.NewWithClaims(alg, claims)

	signBytes, err := ioutil.ReadFile(keyPath)
	Fatal(err)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	Fatal(err)
	out, err := token.SignedString(signKey)
	Fatal(err)
	return out
}

// interacts with POD to authenticated
// gets session token and keyManagerToken
func Auth(jwt, path string) string {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	jsonValue, _ := json.Marshal(map[string]string{
		"token": jwt,
	})
	request, err := http.NewRequest("POST", path, bytes.NewBuffer(jsonValue))
	Fatal(err)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	Fatal(err)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("failed %s", resp.Status)
		os.Exit(1)
	}
	var result TokenResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		Fatal(err)
		os.Exit(1)
	}

	resp.Body.Close()
	return result.Token
}
