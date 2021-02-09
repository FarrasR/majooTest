package function

import(
	"majootest/database"
	"majootest/model"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"	
	"golang.org/x/crypto/bcrypt"
	"net/http"
	// "errors"
	// "regexp"
)

type UpdateUser struct {
	ID uint `form:"ID" json:"ID"   binding:"required"`
	Username string `form:"Username" json:"Username"   binding:"required"`
	Password string `form:"Password" json:"Password"   binding:"required"`
	Fullname string `form:"Fullname" json:"Fullname" binding:"required"`
}



func UpdateThis(c *gin.Context){
	var request UpdateUser

	if err := c.ShouldBindJSON(&request); err != nil {
		ErrorBadRequest(c, "Request incomplete")
		return
	}


	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
    	ErrorBadRequest(c, "Password Unnaceptable")
        return
 	}

 	user:= model.User{}
    database.DBConnection.Where("ID = ?", request.ID).First(&user)

	user.Username=request.Username
	user.Fullname=request.Fullname
	user.Password=string(passwordHash)

	database.DBConnection.Save(&user)

	UpdateUserSuccess(c, user)

}

func UpdateUserSuccess(c *gin.Context, user model.User){

	c.JSON(http.StatusOK, gin.H{
			"status": "exists",
			"username": user.Username, 
			"fullname": user.Fullname,
			"password": user.Password,
	})	
}


