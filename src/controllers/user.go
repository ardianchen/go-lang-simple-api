package controllers

import (
	"fmt"

	res "github.com/ardianchen/simple-api/src/apiHelpers"
	userRepo "github.com/ardianchen/simple-api/src/repository"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var userModel, err = userRepo.GetUser()
	// database.GetDB()
	// var userService serV1.UserService
	// var user = []models
	//decode the request body into struct and failed if any error occur
	fmt.Println("msg:", err)
	// for _, productInfo := range userModel {
	// 	fmt.Println(productInfo.Email)
	// 	fmt.Println("--------------------")
	// }
	// err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	// if err != nil {
	// 	res.Respond(c.Writer, res.Message(1, "Invalid request"))
	// 	return
	// }

	// //call service
	// resp := userService.UserList()
	response := res.Message(0, "This is from version 1 api")
	response["data"] = userModel
	// // //return response using api helper
	res.Respond(c.Writer, response)

}
