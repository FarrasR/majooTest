package function

import(
	"majootest/database"
	"majootest/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"	
	"net/http"
	"errors"
)

type ReadID struct {
	ID uint `form:"ID" json:"ID"   binding:"required"`
}


func ReadThis(c *gin.Context){
	var request ReadID


	if err := c.ShouldBindJSON(&request); err != nil {
		ErrorBadRequest(c, "Request incomplete")
		return
	}

	status :=CheckIDExists(request.ID)

    if status == false {
    	ErrorBadRequest(c, "no such id")
        return
    }

    user:= model.User{}
	database.DBConnection.Where("ID = ?", request.ID).First(&user)

    ReadUserSuccess(c, user)
}

func CheckIDExists(idToCheck uint) bool{
	user:= model.User{}

	err:=database.DBConnection.Where("ID = ?", idToCheck).First(&user).Error


	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}




func ReadUserSuccess(c *gin.Context, user model.User){

	c.JSON(http.StatusOK, gin.H{
			"status": "exists",
			"username": user.Username, 
			"fullname": user.Fullname,
	})	
}
