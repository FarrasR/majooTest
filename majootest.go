package main

import(
	"majootest/routing"
	"majootest/database"
	"majootest/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"	
	// "github.com/joho/godotenv"
	// "fmt"
	"log"
	// "os"
)


// func Route(router *gin.Engine){
// 	router.GET("/", hello)
// 	router.POST("/register", function.RegisterNewUser)
// 	router.POST("/login", function.LoginUser)
// }



func main() {

	connection := "root:@tcp(127.0.0.1:3306)/mojootest?charset=utf8mb4&parseTime=True&loc=Local"
	// should have in the form of env but i dont have time

	var err error
  	database.DBConnection, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	database.DBConnection.AutoMigrate(&model.User{})


	router := gin.Default()
	routing.Route(router)
	router.Run()
}
