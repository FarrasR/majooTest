package function

import(
	"majootest/database"
	"majootest/model"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"	
	// "golang.org/x/crypto/bcrypt"
	"net/http"
	// "errors"
	// "regexp"
)

type DeleteUser struct {
	ID uint `form:"ID" json:"ID"   binding:"required"`
	
}


func DeleteThis(c *gin.Context){
	var request DeleteUser

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

    database.DBConnection.Delete(&user)

    DeleteUserSuccess(c)
}



func DeleteUserSuccess(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
			"delete": "success",
	})	
}
