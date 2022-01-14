/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package str

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("test is begin")
	m.Run()
	fmt.Println("test is over")
}

func TestSplit(t *testing.T) {
	// if testing.Short() {
	// 	// go test --short
	// 	t.Skip("short 模式下，跳过该测试用例")
	// }

	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestSplit-first",
			args: args{
				str: "hello world",
				sep: " ",
			},
			want: []string{"hello", "world"},
		},
		{
			name: "TestSplit-second",
			args: args{
				str: "nice,to,meet,you",
				sep: ",",
			},
			want: []string{"nice", "to", "meet", "you"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 表明该函数并发执行
			// t.Parallel()
			if got := Split(tt.args.str, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			} else {
				t.Logf("name: %s is success", tt.name)
			}
		})
	}
}

func Test_index(t *testing.T) {
	type args struct {
		str    string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_index-first",
			args: args{
				str:    "hello world",
				substr: "h",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		// Go 1.7 加入了子测试，从而避免了需要为不同的测试数据，定义多个测试函数的麻烦。
		t.Run(tt.name, func(t *testing.T) {
			if got := index(tt.args.str, tt.args.substr); got != tt.want {
				t.Errorf("index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("hello world", " ")
	}
}

func Benchmark_index(b *testing.B) {
	// 停止计时
	b.StopTimer()
	// do something
	// 启动计时
	b.StartTimer()
	// do something
	// 重制计时
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index("hello world", "r")
	}
}
