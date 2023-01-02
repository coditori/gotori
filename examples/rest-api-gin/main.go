package main

import (
	"flag"
	"fmt"
	"os"
	_ "rest/docs"
	"rest/server"
)

// @title           Gin Book Service
// @version         1.0
// @description     A book management service API in Go using Gin framework.
// @termsOfService  https://github.com/coditori/javatori/blob/master/LICENSE
// @contact.name   Ario Afrashteh
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @BasePath  /
// @securitydefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	println("environment is " + *environment)
	server.Init()
}
