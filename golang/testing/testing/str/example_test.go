/*
 * Copyright (c) 2022. SuperPony <superponyyy@gmail.com>. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package str

import "fmt"

func ExampleSplit() {
	fmt.Printf("%q\n", Split("hello world", " "))
	fmt.Printf("%q\n", Split("nice,to,meet,you", ","))
	// Output:
	// ["hello" "world"]
	// ["nice" "to" "meet" "you"]
}

func ExampleSplit_second() {
	fmt.Printf("%q\n", Split("my,name,is,jack", ","))
	// Output:
	// ["my" "name" "is" "jack"]
}
