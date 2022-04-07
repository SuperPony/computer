/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package jwt

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"go-tmp/src/jwt/sh256"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

type Payload struct {
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
	Exp int64  `json:"exp"`
	Aud string `json:"aud"`
	Iat int64  `json:"iat"`
	Nbf int64  `json:"nbf"`
	Jti int64  `json:"jti,omitempty"`
	// 非标准字段
	User
}

type User struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type Jwt struct {
	Header
	Payload
	Signature string `json:"signature"`
}

func Login(c *gin.Context) {
	var u User

	if err := c.ShouldBindQuery(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})

		return
	}

	jwt := NewJwt(u.Account, u.Password)
	token := jwt.generateToken()
	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

func Server(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "authentication failed",
		})

		return
	}

	header, payload, sign, err := splitToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	j := &Jwt{}
	if err := j.applyTo(header, payload); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "authentication failed",
		})

		return
	}

	if err := j.validateExp(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	res := j.comparison(sign)
	if !res {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "authentication failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func NewJwt(account, password string) *Jwt {
	now := time.Now()

	return &Jwt{
		Header: Header{
			Typ: "JWT",
			Alg: "HS256",
		},

		Payload: Payload{
			Iss: "jwt-demo",
			Exp: now.Add(time.Hour * 72).Unix(),
			Aud: "jwt-demo.com",
			Iat: now.Unix(),
			Nbf: now.Unix(),
			User: User{
				Account:  account,
				Password: password,
			},
		},
	}
}

func (j *Jwt) generateToken() string {
	encode := j.generateEncode()
	j.Signature = sh256.Encryption([]byte(encode))
	return encode + "." + j.Signature
}

func (j *Jwt) generateSign() string {
	encode := j.generateEncode()
	j.Signature = sh256.Encryption([]byte(encode))
	return j.Signature
}

func (j *Jwt) generateEncode() string {
	h, _ := json.Marshal(j.Header)
	p, _ := json.Marshal(j.Payload)

	return string(sh256.Encode(h)) + "." + string(sh256.Encode(p))
}

func (j *Jwt) comparison(sign string) bool {
	res := j.generateSign() == sign
	return res
}

func (j *Jwt) applyTo(header, payload string) error {
	h, err := sh256.Decode([]byte(header))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(h, &j.Header); err != nil {
		return err
	}

	p, err := sh256.Decode([]byte(payload))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(p, &j.Payload); err != nil {
		return err
	}

	return nil
}

func (j *Jwt) validateExp() error {
	if time.Now().Unix() > j.Exp {
		return errors.New("exp invalid")
	}

	return nil
}

func splitToken(token string) (header, payload, sign string, err error) {
	res := strings.Split(token, ".")
	if len(res) != 3 {
		return "", "", "", errors.New("authentication failed")
	}

	return res[0], res[1], res[2], nil
}
