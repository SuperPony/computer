/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import (
	"log"
	"net/http"
)

// Http is a Http Server.
func Http() {
	Init()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func Init() {
	http.HandleFunc("/", Index)
}
