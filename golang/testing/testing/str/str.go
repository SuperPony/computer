/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package str

import "strings"

func Split(str, sep string) []string {
	if i := index(str, sep); i < 0 {
		return nil
	}
	return strings.Split(str, sep)
}

func index(str, substr string) int {
	return strings.Index(str, substr)
}
