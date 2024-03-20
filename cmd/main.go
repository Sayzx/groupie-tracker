package main

import (
	"main/internal/handler"
	"main/internal/routes"
)

func main() {
	routes.Run()
	handler.Proxy()
}
