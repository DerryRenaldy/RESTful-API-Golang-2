package main

import "restful-api2/server"

func main() {
	sv := server.Register()

	sv.Start()
}
