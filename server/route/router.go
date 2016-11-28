package route

import (
    "github.com/OthloTech/OthloCalendar/server/api"
    "github.com/OthloTech/OthloCalendar/server/db"
    "github.com/OthloTech/OthloCalendar/server/handler"
    othloMw "github.com/OthloTech/OthloCalendar/midleware"
    "github.com/labstack/echo"
    echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
    e := echo.New()

    e.Debug()

    // Set Bundle MiddleWare
    e.Use(echoMw.Logger())
    e.Use(echoMw.Gzip())
    e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
    }))
    e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

    // Set Custom MiddleWare
    e.Use(othloMw.TransactionHandler(db.Init()))

    // Routes
    v1 := e.Group("/api/v1")
    {
        v1.GET("/search", api.Search())
        v1.POST("/style", api.PostStyle())
        v1.GET("/style/:id", api.GetStyle())
    }

    return e   
}
