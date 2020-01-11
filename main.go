package main

import "CM-SSO/model"

func main() {
	readConfig()
	configureLog()
	model.ConfigureDB()
	startServer()
}

