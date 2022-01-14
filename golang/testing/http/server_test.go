/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHttpHandle(t *testing.T) {
	handleMux := http.NewServeMux()
	handleMux.HandleFunc("/", HttpHandle)
	ts := httptest.NewServer(handleMux) // 启动 test 专用的 server，测试完毕后需要关闭
	defer ts.Close()

	t.Run("http-test", func(t *testing.T) {
		// mock 一个响应记录器，后续通过响应记录器的内容作为请求成功与否的凭证。
		w := httptest.NewRecorder()
		// mock 一次请求
		r := httptest.NewRequest("GET", "/", nil)
		HttpHandle(w, r)

		// 通过对比响应是否返回 200 状态作为测试成功与否的依据。
		if !reflect.DeepEqual(w.Code, http.StatusOK) {
			t.Errorf("Request is error, error code: %v", w.Code)
		}
	})
}

func TestGinHandle(t *testing.T) {
	t.Run("gin-test", func(t *testing.T) {
		r := InitRouter()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/?id=%s", "2")
		req := httptest.NewRequest("GET", url, nil)
		r.ServeHTTP(w, req)
		if !reflect.DeepEqual(w.Code, http.StatusOK) {
			t.Errorf("Request is error, http code:%v, id want: 1, got:%s", w.Code, "2")
		}
	})
}
