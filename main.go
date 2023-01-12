package main

import (
	"fmt"

	db "users/grpc/database"
	"users/grpc/routes"
)

func main() {
	db.ConectaBanco()
	fmt.Println("Iniciando Servidor... ")
	routes.StartServerGRPC()
}
