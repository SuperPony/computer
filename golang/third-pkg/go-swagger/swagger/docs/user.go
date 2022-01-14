package docs

import "go-swagger-example/api/user"

// swagger:route GET /users/{id} user getUserRequest
//
// 获取指定用户.
//
// 这个接口目前已经废弃.
//
//     Security:
//       api_key:
//       basic:
//
//     Deprecated: true
//
//     Responses:
//       default: errResponse
//       200: getUserResponse

// swagger:route POST /users user createUserRequest
// 创建用户.
//
//     Security:
//       api_key:
// responses:
//   200: createUserResponse
//   default: errResponse

// 请求说明.
// swagger:parameters getUserRequest
type getUserParamsWrapper struct {
	// swagger:allOf
	user.UserItem
}

// 成功状态返回数据
// swagger:response getUserResponse
type getUserResponse struct {
	// in:body
	Body user.UserItem
}

// 创建用户
// swagger:parameters createUserRequest
type createUserParamsRequest struct {
	// in:body
	Body user.CreateUser
}

// 成功状态响应
// swagger:response createUserResponse
type createUserResponse struct {
	// in:body
	Body user.CreateUser
}

// 失败状态响应
// swagger:response errResponse
type errResponse struct {
	// Error code.
	Code int `json:"code"`
	// Error message.
	Message string `json:"message"`
}
