package route

import (
	"CM-SSO/common/middleware"
	"CM-SSO/controller"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func GetRoutes() *gin.Engine  {
	// Default use CheckLoginState and Recovery middleware
	router := gin.New()
	//set middleware
	//CheckLoginState write log to gin.DefaultWriter，even you set GIN_MODE to release
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery recover any panic。if panic present，will write 500。
	router.Use(gin.Recovery())
	cmsso := router.Group("/cmsso")
	{
		//email
		cmsso.PUT("/mail", controller.SendEmail)
		cmsso.GET("/mail", controller.EmailSignUp)
		cmsso.POST("/mail", controller.EmailLogin, middleware.CheckMaliciousLogin())

		//user service
		//single sign on
		cmsso.GET("/", controller.SingleSignOn, middleware.CheckLoginState())
	}
	//设置静态资源
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dir+="/static"
	log.Printf("static file path: %s",dir)
	router.Static("/static/",dir)
	return router
}
