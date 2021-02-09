package function


import(
	"majootest/database"
	"majootest/model"
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
	// "path/filepath"
	// "net/http"
	"log"
)


type UploadStruct struct {
	ID uint `form:"ID" json:"ID"   binding:"required"`
	Filename string `form:"Fullname" json:"Fullname" binding:"required"`
}



func UploadPhoto(c *gin.Context){
	var request UploadStruct

	if err := c.ShouldBindJSON(&request); err != nil {
		ErrorBadRequest(c, "Request incomplete")
		return
	}

 	

    file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}
	
	err = c.SaveUploadedFile(file, "saved/"+file.Filename)
	if err != nil {
			log.Fatal(err)
		}

	user:= model.User{}
    database.DBConnection.Where("ID = ?", request.ID).First(&user)

    user.Photo=request.Filename
	database.DBConnection.Save(&user)

	UpdateUserSuccess(c, user)
}