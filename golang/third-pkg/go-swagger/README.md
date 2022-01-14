# 介绍

Swagger 是一套围绕 OpenAPI 规范构建的开源工具，可以设计、构建、编写和使用 REST API。Swagger 包含很多工具，其中主要的 Swagger 工具包括：

- Swagger 编辑器：基于浏览器的编辑器，可以在其中编写 OpenAPI 规范，并实时预览 API 文档。https://editor.swagger.io 就是一个 Swagger 编辑器，你可以尝试在其中编辑和预览 API 文档。
- Swagger UI：将 OpenAPI 规范呈现为交互式 API 文档，并可以在浏览器中尝试 API 调用。
- Swagger Codegen：根据 OpenAPI 规范，生成服务器存根和客户端代码库；

Swaager 文档 https://goswagger.io/

# Index

- 安装
- 常用命令
- Spec 规则
  - meta;
  - route;
  - parameters;
  - response;
  - field.

# 安装

`go get -u github.com/go-swagger/go-swagger/cmd/swagger`

# 常用命令

- `swagger version`: 查看版本
- `swagger generate spec -o file`: 生成文档，生的格式（json ｜ yaml）取决于 file 的类型；
- `swagger generate markdown -f sepcFile --output target.md`: 根据 sepc 文件生成 markdown 文档；
  - `-f`: 目标 spec 文件;
  - `--output`: 输出 md 文件；
- `swagger serve sepcFile`: 在浏览器中启动交互式文档；
  - `--no-open`: 表示不自动打开浏览器；
  - `-F=redoc|swagger`: 文档 UI 风格，swagger 简洁、redoc 排版更加精美；
  - `--port protNum`: 指定端口号

# Spec 规则

swagger generate 命令会找到 main 函数，然后遍历整个项目源文件，并解析 swagger 相关注解，最后将其输出到指定文件中。

## meta

meta 是整个项目相关的说明，例如，项目简介，版本号，baseurl 等，因此需要入口文件中，引入 meta 注解所在的包，常用项如下：

```
// Package docs awesome. // awesome 表示项目名称
//
// Documentation of our awesome API.
//
//     Schemes: http, https // 默认方案列表，如果有 https 则作为首选
//     Host: localhost // 域名
//     BasePath: / // 基础路径
//     Version: 0.0.1 // 版本
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security: // 项目中用到的鉴权模式
//     - basic
//     - api_key
//
//    SecurityDefinitions:
//    basic:
//      type: basic|apiKey|OAuth // 基础模式，输入账号密码|token|OAuth
//    api_key:
//      type: apiKey
//      name: Authorization
//      in: header
// swagger:meta // 标识着 meta 结束
package docs
```

## route

route 用于定义 api，以 `swagger:route` 注解开头，具体定义以及常用项如下：

- `swagger:route METHOD url tag parametersId`:
  - METHOD: 表示请求类型;
  - url: 如果 url 中有 path 参数，则需要定义为 `url/{param}`;
  - tag: tag 用于为 api 打标签，相同 tag 的 api 为一组；
  - parametersId: parametersId 是一个具有唯一性的标识符，与请求时对应的参数结构体对应，用于让 swagger 找到请求该 api 时，需要的参数；

```
// swagger:route METHOD url tag parametersId
//
// api title.
//
// api description.
//
//     Security: // 鉴权模式，与 meta 中定义的鉴权模式标识符对应
//       api_key:
//       basic:
//
//     Deprecated: true // Deprecated 为 true 表示该 api 已经不推荐使用
//
//     Responses: // 定义状态码关联的响应
//       default: errResponseId
//       200: okResponseId
```

## parameters

parameters 表示 api 请求时对应的参数，以 `swagger:parameters parametersId` 注解表明，`parametersId` 与对应 api 的 `parametersId` 相同。

```
// swagger:parameters parametersId
type getUserParamsWrapper struct {
	// swagger:allOf
	user.UserItem
}
```

## response

response 表示 api 的响应，以 `swagger:response responseId` 注解表明， `responseId` 与对应 api 中 Responses 里的 `responseId` 对应；

```
// 成功状态响应
// swagger:response responseId
type createUserResponse struct {
	// in:body
	Body user.CreateUser
}
```

## field

此处记录一些与参数项有关的注解

- `swagger:allOf` 结构体内的匿名结构体需要标注该注解，否责匿名结构体可能会无法正常的解析；
- `in:body|query|path`: 表示该参数所在的位置；
- `required:true`: 表示该参数是否必须；
- `min:num`: 该参数最小值；
- `max:num`: 该参数最大值；
- `example:value`: 示例值，path 参数不需要；
- `min length:num`: 参数最小长度；
- `max length:num`: 参数最大长度。
