package routers

import (
	"github.com/amirhasanpour/bama-clone-web-api/src/api/handlers"
	"github.com/amirhasanpour/bama-clone-web-api/src/api/middlewares"
	"github.com/amirhasanpour/bama-clone-web-api/src/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}
