/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BaseURL = "http://api.com"

// response data
type result struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

func GockRequest() (map[string]interface{}, error) {
	// 模拟一次请求，将数据结果转储之 map 返回。
	resp, err := http.Post(fmt.Sprintf("%s%s", BaseURL, "/user"), "application/json", nil)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := map[string]interface{}{"id": "", "msg": ""}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
