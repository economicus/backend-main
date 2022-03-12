package main

import "main/docs"

// @title           Economicus Main Server
// @version         1.0.0
// @description     Economicus 메인 서버
// @termsOfService  https://www.economicus.kr/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  JWT
// @schemes                    https http

func main() {
	docs.SwaggerInfo.Title = "Economicus Main Server API"
	docs.SwaggerInfo.Description = "Economicus 메인 서버 API"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := New()
	app.Run()
}
