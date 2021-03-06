package route

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/OthloTech/OthloCalendar/server/controllers"
	"github.com/gin-gonic/gin"
)

//var server *gin.Engine
var templates map[string]*template.Template

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func loadTemplates() {
	var baseTemplate = "../../dist/templates/layout/_base.html"
	templates = make(map[string]*template.Template)

	templates["index"] = template.Must(template.ParseFiles(baseTemplate, "../../dist/index.html"))
}

func IndexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "pokohide"})
}

func EmbedRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "embed.html", gin.H{})
}

func Init() *gin.Engine {
	//loadTemplates()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	router.LoadHTMLGlob("dist/*.html")

    event := new(controllers.EventController)
	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.GET("/embed", EmbedRoute)
	router.GET("/", IndexRoute)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/search", event.Search)
		//v1.POST("/style", api.PostStyle)
		//v1.GET("/style/:id", api.GetStyle)
	}
	router.StaticFile("/bundle.js", "dist/bundle.js")
	router.StaticFile("/bundle.js.map", "dist/bundle.js.map")
	return router
}
