package route

import (
    //"github.com/OthloTech/OthloCalendar/server/api"
    //"github.com/OthloTech/OthloCalendar/server/db"
    "github.com/OthloTech/OthloCalendar/server/handler"å
    //othloMw "github.com/OthloTech/OthloCalendar/midleware"
    "github.com/gin-gonic/gin"
)

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
    var baseTemplate = "dist/templates/layout/_base.html"
    templates = make(map[string]*template.Template)

    templates["index"] = template.Must(template.ParseFiles(baseTemplate, "dist/index.html"))
}

func IndexRoute(g *gin.Context) {
    server.SetHTMLTemplate(templates["index"])
    g.HTML(200, "_base.html", nil)
}

var server *gin.Engine
var templates = map[string]*template.Template

func Init() *echo.Echo {
    loadTemplates()

    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    router.Use(CORSMiddleware())

    //health := new(controllers.HealthController)
    //router.GET("/health", health.Status)
    router.Use(AuthMiddleware())

    v1 := router.Group("/api/v1")
    {
        //v1.GET("/search", api.Search)
        //v1.POST("/style", api.PostStyle)
        //v1.GET("/style/:id", api.GetStyle)
    }

    //router.Static("/dist", "../dist")
    router.Static("/bundle.js", "../dist/bundle.js")
    router.Static("/style.css", "../dist/style.css")
    router.GET("/", IndexRoute)

    return router
}

