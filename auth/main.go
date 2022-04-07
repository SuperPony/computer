/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"log"
	"net/http"
	"sync"

	"go-tmp/src/basic"
	"go-tmp/src/digest"
	"go-tmp/src/jwt"

	"github.com/gin-gonic/gin"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	r := gin.Default()
	// register routes.
	r.GET("/basic", basic.Server)
	r.GET("/digest", digest.Server)
	r.GET("/jwt/login", jwt.Login)
	r.GET("/jwt/other", jwt.Server)

	wg.Add(1)
	go run(r)
	// go basicExample()

	// wg.Add(1)
	// go digestExample()

	wg.Add(1)
	go jwtExample()

	wg.Wait()
}

func run(handle http.Handler) {
	defer wg.Done()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handle,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err.Error())
	}
}

// ClientExample is mock client actions.
func basicExample() {
	defer wg.Done()
	basic.Client()
}

func digestExample() {
	defer wg.Done()
	digest.Client()
}

func jwtExample() {
	defer wg.Done()
	jwt.Client()
}
