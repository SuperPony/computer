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