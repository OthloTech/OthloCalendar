package main

import (
    "github.com/Sirupsen/logrus"
    "gtihub.com/labstack/echo/engine/fasthttp"
    //"github.com/OthloTech/OthloCalendar/server/route"
    "./route"
)

func init() {
    logrus.SetLevel(logrus.DebugLevel)
    logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
    router := route.Init()
    router.Run(fasthttp.New(":8888"))
}
