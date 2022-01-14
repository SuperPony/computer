# GoConvey

GoConvey 是一款针对 Golang 的测试框架，相比内置的测试包，具备了友好输出、断言、以及 web 界面的特性，并且可以很好的与内置的测试包集成使用，因此更加易用、强大。

安装 `$ go get github.com/smartystreets/goconvey/convey` 

# 基本使用

`Convey(items ...interface{})` 方法代替了 testing 包中的 `Run`方法作为子测试的调用，通常，第一个参数为子测试名，第二个参数为 `*Test.T`, 第三个参数是一个函数，里面存储了测试的逻辑；

```
// str.go

package conveydemo

import "strings"

func Split(s, sep string) []string {
    return strings.Split(s, sep)
}



// str_test.go

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
			c.SoMsg("err-message:", got, c.ShouldResemble, tt.want)
			// _, _ = c.Println("append some message")
		})
	}
}
```

# 断言

GoConvey 提供了许多断言函数，用于判断输出的值与期望的值是否相符，并根据是否相符自动将用例置为成功或失败的状态；

## 一般断言：

- `So(thing1, ShouldEqual, thing2)`
- `So(thing1, ShouldNotEqual, thing2)`
- `So(thing1, ShouldResemble, thing2)`:	用于数组、切片、map和结构体相等
- `So(thing1, ShouldNotResemble, thing2)`
- `So(thing1, ShouldPointTo, thing2)`
- `So(thing1, ShouldNotPointTo, thing2)`
- `So(thing1, ShouldBeNil)`
- `So(thing1, ShouldNotBeNil)`
- `So(thing1, ShouldBeTrue)`
- `So(thing1, ShouldBeFalse)`
- `So(thing1, ShouldBeZeroValue)`

## 数字数量断言：

`So(1, ShouldBeGreaterThan, 0)`
`So(1, ShouldBeGreaterThanOrEqualTo, 0)`
`So(1, ShouldBeLessThan, 2)`
`So(1, ShouldBeLessThanOrEqualTo, 2)`
`So(1.1, ShouldBeBetween, .8, 1.2)`
`So(1.1, ShouldNotBeBetween, 2, 3)`
`So(1.1, ShouldBeBetweenOrEqual, .9, 1.1)`
`So(1.1, ShouldNotBeBetweenOrEqual, 1000, 2000)`
`So(1.0, ShouldAlmostEqual, 0.99999999, .0001) `  // tolerance is optional; default 0.0000000001
`So(1.0, ShouldNotAlmostEqual, 0.9, .0001)`

## 包含断言：

`So([]int{2, 4, 6}, ShouldContain, 4)`
`So([]int{2, 4, 6}, ShouldNotContain, 5)`
`So(4, ShouldBeIn, ...[]int{2, 4, 6})`
`So(4, ShouldNotBeIn, ...[]int{1, 3, 5})`
`So([]int{}, ShouldBeEmpty)`
`So([]int{1}, ShouldNotBeEmpty)`
`So(map[string]string{"a": "b"}, ShouldContainKey, "a")`
`So(map[string]string{"a": "b"}, ShouldNotContainKey, "b")`
`So(map[string]string{"a": "b"}, ShouldNotBeEmpty)`
`So(map[string]string{}, ShouldBeEmpty)`
`So(map[string]string{"a": "b"}, ShouldHaveLength, 1)` // supports map, slice, chan, and string

## 字符串断言

`So("asdf", ShouldStartWith, "as")`
`So("asdf", ShouldNotStartWith, "df")`
`So("asdf", ShouldEndWith, "df")`
`So("asdf", ShouldNotEndWith, "df")`
`So("asdf", ShouldContainSubstring, "稍等一下")`		// optional 'expected occurences' arguments?
`So("asdf", ShouldNotContainSubstring, "er")`
`So("adsf", ShouldBeBlank)`
`So("asdf", ShouldNotBeBlank)`

## panic 断言

`So(func(), ShouldPanic)`
`So(func(), ShouldNotPanic)`
`So(func(), ShouldPanicWith, "")`		// or errors.New("something")
`So(func(), ShouldNotPanicWith, "")`	// or errors.New("something")

## 时间、间隔断言

`So(time.Now(), ShouldHappenBefore, time.Now())`
`So(time.Now(), ShouldHappenOnOrBefore, time.Now())`
`So(time.Now(), ShouldHappenAfter, time.Now())`
`So(time.Now(), ShouldHappenOnOrAfter, time.Now())`
`So(time.Now(), ShouldHappenBetween, time.Now(), time.Now())`
`So(time.Now(), ShouldHappenOnOrBetween, time.Now(), time.Now())`
`So(time.Now(), ShouldNotHappenOnOrBetween, time.Now(), time.Now())`
`So(time.Now(), ShouldHappenWithin, duration, time.Now())`
`So(time.Now(), ShouldNotHappenWithin, duration, time.Now())`



# 友好输出

`Println, Printf`: 当断言为真，通过测试后，则调用此类函数进行一些自定义内容的输出；
`SoMsg(msg string, actual interface{}, assert Assertion, expected ...interface{})`: 该函数与`So` 作用一致，区别在于，当断言失败，导致单元测试失败后，`msg` 作为自定义的错误信息进行输出。 

# WebUI

在项目目录中执行 `$ goconvey`命令时，GoConvey 会递归搜集整个项目中的测试文件进行测试，并以 WebUI 的形式进行展示。（启动的 http server 具有热更新的功能）

![](/Applications/MAMP/htdocs/localhost/go-tmp/img/5.jpg)