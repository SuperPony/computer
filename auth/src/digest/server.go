/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package digest

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

type WWWAuthenticate struct {
	Realm    string `json:"realm"`
	Nonce    string `json:"nonce"`
	Nc       string `json:"nc"`
	Cnonce   string `json:"cnonce"`
	Qop      string `json:"qop"`
	Uri      string `json:"uri"`
	Response string `json:"response"`
	Account  string `json:"account"`
}

func Server(c *gin.Context) {
	wa := &WWWAuthenticate{}
	wah := c.GetHeader("WWW-Authenticate")
	if wah == "" {
		authHeader, _ := json.Marshal(gin.H{
			"realm": "digest.com",
			"qop":   "auth,auth-int",
			"nonce": createNonce(),
		})
		c.Header("WWW-Authenticate", string(authHeader))
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	if err := json.Unmarshal([]byte(wah), wa); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})

		return
	}

	if err := wa.validate(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})

		return
	}

	pwd, err := mockFindUser(wa.Account)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})

		return
	}

	if ok := wa.comparison(pwd); !ok {
		fmt.Println(wa)
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "authentication failed",
		})

		return
	}

	authInfo, _ := json.Marshal(gin.H{
		"otherData": "hello world",
	})
	c.Header("Authorization-Info", string(authInfo))
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// Validate WWWAuthenticate fields.
func (w *WWWAuthenticate) validate() error {
	if w.Realm != "digest.com" {
		return errors.New("the realm validation failed")
	}
	// more validation...

	return nil
}

func (w *WWWAuthenticate) comparison(pwd string) bool {
	m := md5.New()
	m.Write([]byte(w.Account + w.Nonce + pwd + w.Cnonce))
	if w.Response == base64.StdEncoding.EncodeToString(m.Sum(nil)) {
		return true
	}

	return false
}

// mockFindUser is mock find user action.
func mockFindUser(account string) (pwd string, err error) {
	// mock find user on storage
	return "123456", nil
}

// createNonce create new nonce.
func createNonce() string {
	output := make([]byte, 32)
	randomness := make([]byte, 32)
	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	l := len(Alphabet62)

	for pos := range output {
		random := randomness[pos]
		// random % 64
		randomPos := random % uint8(l)
		output[pos] = Alphabet62[randomPos]
	}

	return string(output)
}
