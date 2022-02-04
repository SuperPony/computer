/*
 * Copyright (c) 2022. KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package server

import (
	"fmt"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world \n")

	fmt.Println(r.URL.RawFragment)
}
