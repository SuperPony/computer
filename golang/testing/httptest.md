# HttpTest

Golang 中内置了用于 httptest 的包，可以很好的进行相关测试，以及对于第三方 web 框架的测试。
常用的方法说明：

- `NewRecorder() *ResponseRecorder`: 模拟一个响应记录器，可以直接作为响应来使用，以及后续通过参照响应记录器中的状态作为请求是否成功的依据；
- `NewRequest(method string, target string, body io.Reader) *http.Request`: 模拟一次请求；
- `NewServer(handler http.Handler) *Server` 启动一个用于测试的服务，必须在测试完毕后将其关闭。

## Example

```
// server.go

func HttpHandle(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"msg": "hello world",
	}
	s, _ := json.Marshal(data)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(s)
}

// server_test.go

func TestHttpHandle(t *testing.T) {
	handleMux := http.NewServeMux()
	handleMux.HandleFunc("/", HttpHandle)
	ts := httptest.NewServer(handleMux)
	defer ts.Close()

	t.Run("http-test", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		HttpHandle(w, r)

		// 通过对比响应是否返回 200 状态作为测试成功与否的依据。
		if !reflect.DeepEqual(w.Code, http.StatusOK) {
			t.Errorf("Request is error, error code: %v", w.Code)
		}
	})
}
```

## Gin Web Example

```
// server.go

func GinHandle(c *gin.Context) {
	id := c.Query("id")
	if id == "1" {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", GinHandle)
	return r
}

// server_test.go

func TestGinHandle(t *testing.T) {
	t.Run("gin-test", func(t *testing.T) {
		r := InitRouter()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/?id=%s", "2")
		req := httptest.NewRequest("GET", url, nil)
		r.ServeHTTP(w, req)
		if !reflect.DeepEqual(w.Code, http.StatusOK) {
			t.Errorf("Request is error, http code:%v, id want: 1, got:%s", w.Code, "2")
		}
	})
}
```


# GoMock

GoMock 是一款用于模拟被访问的外部请求地址的库，GoMock 中有两个重要的类型`*Request`,`*Response`。

安装 GoMock，`$ go get gopkg.in/h2non/gock.v1`;

## 常用方法

- `New(baseurl string)`: 返回 Mock 实例，用于进行 API 的声明；
- `IsDone()`: 断言是否成功出发所有已经注册的模拟，为真则返回 true；
- `Off()`: 关闭 gomock 的 http 监听器，并删除所有已经注册的模拟，通常在测试完毕后进行关闭。
- `*Request`: 用于声明 API 的请求相关条件，例如，URL 地址、请求时必须匹配的 Header 头，Content-Type 等等；
    - `GET(path string), POST(path string), PUT(path string), DELETE(path string)`: 对应4种 http 请求动作；
    - `Json(data interface{})`: 设置请求数据；
    - `MatchHeader(key, value string)`: 设置匹配的头部，在请求时，必须传入该声明的头部；
    - `MatchType(kind string)`: 设置 Content-Type, kind 对应值如下:
        - "html": "text/html",
        - "text": "text/plain",
        - "json": "application/json",
        - "xml":  "application/xml",
        - "form": "multipart/form-data",
        - "url": "application/x-www-form-urlencoded"
    - `Reply(code int) *Response`: 设置请求返回的状态码。
- `*Response`: 用于声明 API 的返回，例如 Header 头，响应状态码，返回数据等;
    - `SetHeader(key, value string)`: 设置响应的头部；
    - `Json(data interface{})`: 设置返回的数据；