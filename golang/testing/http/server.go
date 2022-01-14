/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpHandle(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"msg": "hello world",
	}
	s, _ := json.Marshal(data)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(s)
}

func GinHandle(c *gin.Context) {
	id := c.Query("id")
	if id == "1" {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", GinHandle)
	return r
}
