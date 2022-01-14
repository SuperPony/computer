/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package mock

//go:generate mockgen -destination ./mock_db.go -package mock  . Store

type Store interface {
	Find(id int64) (*User, error)
	Create(user *User) error
}

// User user model.
type User struct {
	ID   int64
	Name string
	Age  int8
}

func GetFromDB(s Store, id int64) (*User, error) {
	u, err := s.Find(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
