package route

import (
	"github.com/AthulKrishna2501/The-Furniture-Spot/user"
	"github.com/gin-gonic/gin"
)

func RegisterURL(router *gin.Engine) {
	router.POST("/signup/", user.SignUp)
	router.POST("/verifyotp", user.VerifyOTP)
}