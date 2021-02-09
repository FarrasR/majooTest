package function

import(
	"majootest/database"
	"majootest/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)


type Register struct {
	Username string `form:"Username" json:"Username"   binding:"required"`
	Password string `form:"Password" json:"Password"   binding:"required"`
	Fullname string `form:"Fullname" json:"Fullname" binding:"required"`
}

func Create(c *gin.Context){
	var request Register

	if err := c.ShouldBindJSON(&request); err != nil {
		ErrorBadRequest(c, "Request incomplete")
		return
	}

	if CheckUsernameExist(request.Username) == true {
		ErrorBadRequest(c, "Username Already exists")
        return
    }

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)


	if err != nil {
    	ErrorBadRequest(c, "Password Unnaceptable")
        return
 	}

 	newUser := model.User{
		Username: request.Username,
		Password: string(passwordHash),
		Fullname: request.Fullname}

	database.DBConnection.Create(&newUser)

	user:= model.User{}
    database.DBConnection.Where("Username = ?", request.Username).First(&user)

    CreateUserSuccess(c, user.ID)
}



func ErrorBadRequest(c *gin.Context, reason string){
	c.JSON(http.StatusBadRequest, gin.H{"Status": reason})
}

func CreateUserSuccess(c *gin.Context, id uint){
	c.JSON(http.StatusOK, gin.H{
			"status": "Register Success",
			"id": id,
	})
}