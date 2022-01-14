# Index
- 前言
- GoMock
  - 安装
  - mockgen
    - 源码模式、反射模式 
    - 通过注释使用 mockgen
    - 常用选项说明
  - 使用 Mock 代码编写测试用例
    - 模拟接口方法行为的相关方法
  
- 打桩

# 前言

原则上，单元测试中是不允许有外部依赖的，例如 `MySQL`、`Redis`...因此，这些外部依赖以及接口的实现都需要被模拟，故而需要借助各类的 Mock 工具来实现和模拟相关依赖。


# GoMock

GoMock 是 Golang 官方团队开发、维护的一款 Mock 工具，能与内置的 testing 包良好集成，也能用于其他测试环境中；GoMock 测试框架包含了 GoMock 包和 mockgen 工具两部分；GoMock 包用于完成和管理对象的生成周期，而 mockgen 工具则用来生成 interface 对应的 Mock 类源文件。

## 安装

1. GoMock: `$ go get github.com/golang/mock/gomock`;
2. mockgen: `$ go install github.com/golang/mock/mockgen`.

## mockgen

mockgen 工具用于生成 Mock 代码，生成的方式有两种，源码模式、反射模式；

### 源码模式、反射模式

- 源码模式：如果有接口文件，则可以直接使用 `$ mockgen -destination outfile  -package pkgname -source file` 的方式，将指定路径下接口文件组织成 pkgname 包，并存放到 outfile 位置；
- 反射模式：mockgen 同样支持反射的方式生成 mock 代码，通过两个非标志参数，即倒入的路径和逗号分割的接口列表来生成 Mock 数据，其他参数与源码模式公用, `$ mockgen -destination outfile -package pkgname file MyInterface1,MyInterface2`。

### 通过注释使用 mockgen

如果有多个散落在不同文件的接口需要生成 mock 数据（假定包名不同），如果对每个文件执行一次 mockgen命令，则过于繁琐，因此 mockgen 提供了一种通过注释生成 Mock 文件的方式，该方式需要借助 go generate 工具。

```
// order.go

//go:generate mockgen -destination ./mock_order.go -package order . Store

type Store interface {
	Find(id int64) (*User, error)
	Create(user *User) error
}
```

接下来只需要在对应目录中执行`go generate` 命令就可以自动生成 Mock 代码,`$ go generate path`.

### 常用选项说明

| 参数           | 说明                             |
|--------------|--------------------------------|
| -source      | 指定需要模拟的接口文件                    |
| -destination | 指定 Mock 文件输出的地方，若不设置则打印到标准输出   |
| -package     | 指定 Mock 文件的包名，如果不指定则为 mock_文件名 |
| -imports     | 依赖的包，逗号分割                      |
| -aux_files   | 接口文件不止一个时的附加文件，逗号分割            |
| -build_flags | 传递给 build 工具的参数                |


## 使用 Mock 代码编写测试用例

在 testing 中使用 Mock 代码的大致流程如下：

1. 创建一个 Mock 控制器，该对象控制了整个 Mock 的过程： `ctrl := gomock.NewController(t)`;
2. 延迟回收控制器： `defer ctrl.Finish()`;
3. 创建模拟接口的 Mock 对象，需要传入 Mock 控制器：`m := pkg.NewMockXXX(ctrl)`;
    - 此处的 pkg 表示 Mock 文件所在的包， `NewMockXXX`，其中`NewMock`是固定格式，`XXX`则是实现的接口名；
4. 通过调用 Mock 实例的断言方法 `EXPEXT()`， 其返回的 Call 对象就可以以链式调用法，模拟和约束接口中方法的行为，从而满足测试需要的依赖。

```
type DB interface {
	InterFaceMethod(id int) (id int, error)
}

func XXX(d DB, id int) (id, error) {
  return d.InterFaceMethod(id)
}


Func TestInterFaceMethod(t *testing.T) {
  ctrl := gomock.NewController(t)
  defer ctrl.Finish()
  
  m := NewMockXXX(ctrl)
  
  // 模拟接口中指定方法的行为
  m.EXPECT().InterFaceMethod(gmock.Eq(1)).Return(1, nil).Times(1)
  
  id, err := XXX(m, 1)
  if err != nil {
    t.Fatal("error")
  }
  t.Log("success")
}
```

## 模拟接口方法行为的相关方法

- 参数匹配: 
  - gomock.Any(): 用来表示接口方法的允许任何入参；
  - gomock.Eq(value): 表示接口方法仅接受与 value 相等的入参；
  - gomock.Not(value): 表示接口方法仅接受与 value 意外的入参；
  - gomock.Nil(): 表示接口方法仅接受 nil 入参。

- 指定返回值：
  - `m.EXPECT().InterfaceMethod().Return(rets ...interface{})`: Return 方法用于指定返回值。

- 指定执行次数：
  - `AnyTimes()`: 执行 0 到多次；
  - `MaxTimes(n int)`: 表示如果没有设置，则最多执行 n 次；
  - `MinTimes(n int)`: 表示如果没有设置，最少执行 n 次;
  - `Times(n int)`: 执行次数。

- 指定执行顺序:
- `After(func *Call)`: 当前方法在执行完 func 方法后，才能执行；
```
initCall := mockSpider.EXPECT().Init()
mockSpider.EXPECT().Recv().After(initCall)
```