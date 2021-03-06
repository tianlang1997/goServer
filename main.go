package main

import (
	"flag"
	"goServer.com/m/app/main_app"
	"goServer.com/m/config"
)
var env = flag.String("env", "default", "config file type: default|product")
func main() {
	flag.Parse()
	config.Init(*env)
	main_app.Start("8080")
}