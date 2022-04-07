/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package passwordmode

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server(c *gin.Context) {
	var uf *user

	if err := c.ShouldBindQuery(uf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// mock find user
	// u := SELECT * FROM user WHERE account = u.Account AND password = u.Password;
	u := true

	if !u {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not find",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": generateUserToken(),
	})
}

func generateUserToken() string {
	return "123214"
}
