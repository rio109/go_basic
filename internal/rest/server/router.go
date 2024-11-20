package server

import (
	"net/http"
	"whatisyourtime/internal/rest/handler"

	"github.com/gin-gonic/gin"
)

func NewEngine(mode string) *gin.Engine {
	// 서버 모드 설정
	// gin.SetMode(mode)
	// router := gin.Default()
	router := setup(mode)

	// 템플릿 및 정적 파일 설정
	router.LoadHTMLGlob("internal/template/*.html")
	router.Static("/static", "./internal/static")

	// 라우팅 설정
	api := router.Group("/")
	registerRoutes(api)

	return router
}

func registerRoutes(api *gin.RouterGroup) {
	api.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	{
		api.GET("/api", handler.GetCurrentTime)
		api.GET("/api/world", handler.GetWorldTimes)
		api.GET("/api/target", handler.GetTimeInZone)
	}

	// api.GET("/api/time", func(c *gin.Context) {
	// 	tz := c.DefaultQuery("timezone", "UTC")
	// 	fmt.Println(tz)
	// 	location, err := time.LoadLocation(tz)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timezone"})
	// 		return
	// 	}

	// 	currentTime := time.Now().In(location)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"time":     currentTime.Format("2006-01-02 15:04:05"),
	// 		"timezone": location.String(),
	// 	})
	// })
}
