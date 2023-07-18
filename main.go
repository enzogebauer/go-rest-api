package main

import (
	"fmt"

	"github.com/enzogebauer/go-rest-api.git/routes"
)


	
func main (){
	fmt.Println("Iniciando servidor rest com GO")
	routes.HandleRequest()
}