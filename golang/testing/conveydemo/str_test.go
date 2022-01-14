/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package conveydemo

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test-1",
			args: args{
				s:   "hello,world",
				sep: ",",
			},
			want: []string{"hello", "world"},
		},
		{
			name: "test-2",
			args: args{
				s:   "nice to meet you",
				sep: "w",
			},
			want: []string{"nice", "to", "meet", "you"},
		},
	}
	for _, tt := range tests {
		c.Convey(tt.name, t, func() {
			got := Split(tt.args.s, tt.args.sep)
			// c.So(got, c.ShouldResemble, tt.want)
			c.SoMsg("err message:", got, c.ShouldResemble, tt.want)
			_, _ = c.Println("success message")
		})
	}
}
