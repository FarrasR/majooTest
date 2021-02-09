package routing

import(
	"github.com/gin-gonic/gin"
	"majootest/function"
	"majootest/auth"
	"net/http"
)

func Route(router *gin.Engine){
	router.GET("/", test)
	router.GET("/GetToken", GetToken)


	api := router.Group("/api", auth.AuthorizeJWT())
	{
		api.GET("/ParseToken", test)
		api.POST("/create", function.Create)
		api.GET("/read", function.ReadThis)
		api.POST("/update", function.UpdateThis)
		api.DELETE("/delete", function.DeleteThis)
		api.POST("/login", function.Login)
		api.POST("/uploadPhoto", function.UploadPhoto)
	}
}


func test(c *gin.Context){
	c.JSON(200, gin.H{
			"message": "hello mr examiner",
		})
}


type TokenAuthUserPassword struct {
	Username string `form:"Username" json:"Username"   binding:"required"`
	Password string `form:"Password" json:"Password"   binding:"required"`
}

func GetToken(c *gin.Context){
	var request TokenAuthUserPassword

	if err := c.ShouldBindJSON(&request); err != nil {
		function.ErrorBadRequest(c, "Request incomplete")
		return
	}


	token, authorized := auth.GenerateToken(request.Username, request.Password)
	
	if(authorized == false) {
		function.ErrorBadRequest(c, "not authorized")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

