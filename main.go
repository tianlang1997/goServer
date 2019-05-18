package main

import (
	"./app/main_app"
	"./config"
	"flag"
)
var env = flag.String("env", "default", "config file type: default|product")
func main() {
	flag.Parse()
	config.Init(*env)
	main_app.Start("8080")
}