package controllers

import (
	"log"

	res "github.com/ardianchen/simple-api/src/apiHelpers"
	"github.com/ardianchen/simple-api/src/database"
	userRepo "github.com/ardianchen/simple-api/src/repository"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var userModel = userRepo.GetUser()
	// //call service
	// resp := userService.UserList()
	response := res.Message(20, "success")
	response["data"] = userModel
	// // //return response using api helper
	res.Respond(c.Writer, response)

}

func CreateUser(c *gin.Context) {
	var users userRepo.User
	c.ShouldBindJSON(&users)
	// fmt.Printf("%+v", &users)
	// name := c.PostFormMap("name")
	// email := c.PostFormMap("email")
	// users.Name = name
	// users.Email = email
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Printf(users)
	// fmt.Printf("name: %v; email: %v", name, email)
	// log.Fatalln(users)
	// a := map[string]interface{}{
	// 	"name":  users.Name,
	// 	"email": users.Email,
	// }
	db := database.InitDb()
	log.Fatalln(db.Create(&users))
	// db.Create(&userData)
	// a := User{Name: UserData.Name, Email: userData.email}

	// log.Fatalln(a)
	// userRepo.InsertUser(a)
	// c.JSON(http.StatusOK, gin.H{
	// 	"name":  user.Name,
	// 	"email": user.Email,
	// })
	// var userModel = userRepo.GetUser()
	// response := res.Message(20, "success")
	// response["data"] = userModel
	// // // //return response using api helper
	// res.Respond(c.Writer, response)

}
