package server

import (
	"whatisyourtime/internal/rest/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
)

func setup(mode string) *gin.Engine {
	gin.SetMode(mode)
	router := gin.New()
	setupMiddelware(router)
	return router
}

func setupMiddelware(router *gin.Engine) {
	gin.Default()
	// router.Use(gin.Logger())1
	router.Use(gin.Recovery())
	router.Use(middleware.RateLimitMiddleware)
	router.Use(Logger(logrus.WithFields(logrus.Fields{})))
	setupCors(router)
	setupSecurityHeaders(router)

}

func setupCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: false,
	}))
}

func setupSecurityHeaders(router *gin.Engine) {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:             true,                 // Clickjacking 방지
		ContentTypeNosniff:    true,                 // MIME 타입 강제
		BrowserXssFilter:      true,                 // 브라우저 XSS 방지
		ContentSecurityPolicy: "default-src 'self'", // CSP 설정
	})
	router.Use(func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.Abort()
			return
		}
		c.Next()
	})
}
