/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package http

import (
	"net/http"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"
)

func TestGockRequest(t *testing.T) {

	c.Convey("gomock-test", t, func() {
		respData := map[string]interface{}{
			"msg": "success",
			"id":  "1",
		}

		gock.New(BaseURL).
			Post("/user").
			MatchType("json").
			Reply(http.StatusOK).
			SetHeader("Authorization", "xxx").
			JSON(respData)

		res, err := GockRequest()
		if err != nil {
			t.Errorf("err: %s", err.Error())
		}
		c.So(true, c.ShouldBeTrue)
		c.So(res, c.ShouldResemble, respData)
	})

}
