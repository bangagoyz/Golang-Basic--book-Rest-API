package main

import "chapter2_1/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
