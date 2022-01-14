


# awesome.
Documentation of our awesome API.
  

## Informations

### Version

0.0.1

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### api_key (header: Authorization)



> **Type**: apikey

#### basic



> **Type**: basic

## All endpoints

###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /users | [create user request](#create-user-request) | 创建用户. |
| GET | /users/{id} | [get user request](#get-user-request) | 获取指定用户. |
  


## Paths

### <span id="create-user-request"></span> 创建用户. (*createUserRequest*)

```
POST /users
```

#### Security Requirements
  * api_key

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [CreateUser](#create-user) | `models.CreateUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-user-request-200) | OK | 成功状态响应 |  | [schema](#create-user-request-200-schema) |
| [default](#create-user-request-default) | | 失败状态响应 | ✓ | [schema](#create-user-request-default-schema) |

#### Responses


##### <span id="create-user-request-200"></span> 200 - 成功状态响应
Status: OK

###### <span id="create-user-request-200-schema"></span> Schema
   
  

[CreateUser](#create-user)

##### <span id="create-user-request-default"></span> Default Response
失败状态响应

###### <span id="create-user-request-default-schema"></span> Schema
empty schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| code | int64 (formatted integer) | `int64` |  |  | Error code. |
| message | string | `string` |  |  | Error message. |

### <span id="get-user-request"></span> 获取指定用户. (*getUserRequest*)

```
GET /users/{id}
```

这个接口目前已经废弃.

#### Security Requirements
  * api_key
  * basic

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | 用户 ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-user-request-200) | OK | 成功状态返回数据 |  | [schema](#get-user-request-200-schema) |
| [default](#get-user-request-default) | | 失败状态响应 | ✓ | [schema](#get-user-request-default-schema) |

#### Responses


##### <span id="get-user-request-200"></span> 200 - 成功状态返回数据
Status: OK

###### <span id="get-user-request-200-schema"></span> Schema
   
  

[UserItem](#user-item)

##### <span id="get-user-request-default"></span> Default Response
失败状态响应

###### <span id="get-user-request-default-schema"></span> Schema
empty schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| code | int64 (formatted integer) | `int64` |  |  | Error code. |
| message | string | `string` |  |  | Error message. |

## Models

### <span id="create-user"></span> CreateUser


> 添加用户
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | int64 (formatted integer)| `int64` | ✓ | | 用户 ID | `123` |
| Name | string| `string` | ✓ | | 用户姓名 | `张三` |
| Sex | uint64 (formatted integer)| `uint64` | ✓ | | 用户性别，1男、2女、3未知 | `1` |



### <span id="user-item"></span> UserItem


> 获取指定用户
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Id | int64 (formatted integer)| `int64` | ✓ | | 用户 ID
in: path |  |


