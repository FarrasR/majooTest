package function

import(
	"majootest/database"
	"majootest/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"	
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"errors"
	// "regexp"
)

type LoginUser struct {
	Username string `form:"Username" json:"Username"   binding:"required"`
	Password string `form:"Password" json:"Password"   binding:"required"`	
}



func Login(c *gin.Context){
	var request LoginUser

	if err := c.ShouldBindJSON(&request); err != nil {
		ErrorBadRequest(c, "Request incomplete")
		return
	}

	if CheckUsernameExist(request.Username) == false {
		ErrorBadRequest(c, "No such username")
        return
    }

	status, _ :=checkUsernamePassword(request.Password, request.Username)


	if status == false {
    	ErrorBadRequest(c, "Wrong Password")
        return
    }

    user:= model.User{}
	database.DBConnection.Where("Username = ?", request.Username).First(&user)
    LoginUserSuccess(c , user)
}



func CheckUsernameExist(usernameToCheck string) bool{
	user:= model.User{}

	err:=database.DBConnection.Where("Username = ?", usernameToCheck).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}




func checkUsernamePassword(password string, usernameToCheck string)(bool, int){
	user:= model.User{}

	database.DBConnection.Where("Username = ?", usernameToCheck).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
    	return false,  int(user.ID)
  	} else {
  		return true, int(user.ID)
  	}
}

func LoginUserSuccess(c *gin.Context, user model.User){

	c.JSON(http.StatusOK, gin.H{
			"login": "success",
			"username": user.ID, 
			"fullname": user.Username,
	})	
}