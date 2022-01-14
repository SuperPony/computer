package main

import (
	_ "go-swagger-example/swagger/docs"
	"log"
	"os/exec"
)

func main() {
	// startSwaggerServe()
	createSwagger()
	// r := gin.Default()
	// r.GET("/users/:id", user.Get)
	// r.POST("/users", user.Create)
	// r.Run(":8080")
}

// 生成 api 文档
// 	`swagger generate spec -o demo.yaml`
// 	`swagger generate markdown -f ./swagger/swagger.yaml --output ./swagger.markdown`
func createSwagger() {
	yaml := "./swagger/swagger.yaml"
	cmd := exec.Command("swagger", "generate", "spec", "-o", yaml)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error())
	}

	// 生成 markdow 文档
	md := "./swagger/swagger.md"
	cmd2 := exec.Command("swagger", "generate", "markdown", "-f", yaml, "--output", md)
	if err := cmd2.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

// 启用 swagger ui
//	`swagger serve --no-open -F=swagger|redoc --port 36666 swagger.yaml`
//	localhost:36666/docs
func startSwaggerServe() {
	args := []string{
		"serve",
		// "--no-open", // 是否打开浏览器
		"-F=redoc", // ui 风格，个人 swagger 简洁一点，redoc 排版更好看
		"--port",
		"36666",
		"./swagger/swagger.yaml",
	}

	cmd := exec.Command("swagger", args...)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
