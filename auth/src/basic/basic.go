/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package basic

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Server is mock basic authentication server.
func Server(c *gin.Context) {
	str := c.Request.Header.Get("authorization")
	if str == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Authentication failed.",
		})

		return
	}

	user, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})

		return
	}

	// analysis client's data.
	data := strings.Split(string(user), ":")
	if len(data) <= 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Authentication failed.",
		})

		return
	}

	// mock storage comparison.
	if data[0] != "admin" || data[1] != "admin@123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Authentication failed.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": string(user),
	})
}

// Client is mock client request.
func Client() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/basic", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// authentication success
	user := fmt.Sprintf("%s:%s", "admin", "admin@123")
	// authentication failed
	// 	user := fmt.Sprintf("%s:%s", "admin2", "admin@1234")

	req.Header.Set("authorization", base64.StdEncoding.EncodeToString([]byte(user)))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("The body value is: %s \n", string(body))
}
