package main

import (
    "github.com/Sirupsen/logrus"
    "github.com/gin-gonic/gin"
    "github.com/OthloTech/OthloCalendar/server/config"
    "github.com/OthloTech/OthloCalendar/server/route"
)

func init() {
    logrus.SetLevel(logrus.DebugLevel)
    logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
    //config := config.GetConfig()
    r := route.Init()
    r.Run(":8080")
    //r.Run(config.GetString("server.port"))
}
