/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package passwordmode

import (
	"log"
	"net/http"
	"net/url"
)

var (
	client = &http.Client{}
)

type user struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

func Client() {
	u := &user{
		Account:  "admin",
		Password: "123456",
	}

	req, err := u.generateRequest()
	if err != nil {
		log.Fatalln(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer resp.Body.Close()

}

func (f *user) generateRequest() (*http.Request, error) {
	u, err := url.Parse("http://127.0.0.1:8080/oauth/password")
	if err != nil {
		log.Fatalln(err)
	}

	q := u.Query()
	q.Add("account", f.Account)
	q.Add("password", f.Password)
	u.RawQuery = q.Encode()

	return http.NewRequest("GET", u.String(), nil)
}
