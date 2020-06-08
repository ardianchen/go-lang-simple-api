package main

import (
	"fmt"
	"os"

	"github.com/ardianchen/simple-api/src/database"
	"github.com/ardianchen/simple-api/src/routers"
	"github.com/joho/godotenv"
)

//Execution starts from main function
func main() {

	//Db Connect and Close
	database.InitDb()
	// defer database.CloseDb()

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := routers.SetupRouter()

	port := os.Getenv("port")

	// For run on requested port
	// if len(os.Args) > 1 {
	// 	reqPort := os.Args[1]
	// 	if reqPort != "" {
	// 		port = reqPort
	// 	}
	// }

	if port == "" {
		port = "8080" //localhost
	}
	type Job interface {
		Run()
	}

	r.Run(":" + port)

}
