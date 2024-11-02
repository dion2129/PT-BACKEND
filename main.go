package main

import (
	"api-test/database"
	"api-test/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	
	db := database.ConnectToDb()
	routes.Routes(r, db)
	
	fmt.Println("coba jalan")		
	r.Run(":8080")

}
