/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package jwt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	client = &http.Client{}
	token  string
)

func Client() {
	handleLogin()

	handleOther()
}

func handleOther() {
	u := "http://127.0.0.1:8080/jwt/other"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}

func handleLogin() {
	u, err := url.Parse("http://127.0.0.1:8080/jwt/login")
	if err != nil {
		log.Fatalln(err)
	}

	q := u.Query()
	q.Add("account", "admin")
	q.Add("password", "123456")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	m := make(map[string]string)
	if err := json.Unmarshal(body, &m); err != nil {
		log.Fatalln(err)
	}

	token = m["access_token"]
}
