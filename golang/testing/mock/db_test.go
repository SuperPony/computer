/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mock

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockStore(ctrl)
	m.EXPECT().Find(gomock.Eq(int64(2))).Return(&User{
		ID:   2,
		Name: "jack",
		Age:  22,
	}, nil).Times(1)

	u, err := GetFromDB(m, 2)
	if err != nil {
		t.Fatal("data not found")
	}

	t.Logf("user: %v", u)
}
