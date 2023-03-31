package main

import (
	"lmarrero/dop-exchange-api/server"
)

// @title DOP Exchange API
// @version 1.0
// @description Sample API to exchange DOP to other currencies
// @termsOfService http://swagger.io/terms/

// @contact.name Luis Marrero
// @contact.url https://github.com/luisjmarrero
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	server.Init()
}
