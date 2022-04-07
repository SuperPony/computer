/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package digest

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	client  = &http.Client{}
	wwwAuth = &WWWAuthenticate{}
)

type User struct {
	Account  string
	Password string
}

func Client() {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/digest", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	setup1(req)
	setup2(req)
}

// Mock first request, fetch WWW-Authenticate header.
func setup1(req *http.Request) {
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()

	auth := resp.Header.Get("WWW-Authenticate")

	if err := json.Unmarshal([]byte(auth), wwwAuth); err != nil {
		log.Fatalln(err)
	}
}

// Mock second request, send WWW-Authenticate header.
func setup2(req *http.Request) {
	u := &User{
		Account:  "admin",
		Password: "123456",
	}
	u.createClientWWWAuth()
	wwwAuthHeader, _ := json.Marshal(wwwAuth)
	req.Header.Set("WWW-Authenticate", string(wwwAuthHeader))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("body: %s --- header: %s", string(body), resp.Header.Get("Authorization-Info"))
}

func (u *User) encrypt(cNonce string) string {
	m := md5.New()
	m.Write([]byte(u.Account + wwwAuth.Nonce + u.Password + cNonce))

	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}

func (u *User) createClientWWWAuth() {
	cNonce := createNonce()
	wwwAuth.Cnonce = cNonce
	wwwAuth.Account = u.Account
	wwwAuth.Qop = "auth"
	wwwAuth.Response = u.encrypt(cNonce)
}
