package main

import (
    "fmt"
    "runtime"
    "github.com/gin-gonic/gin"
    "github.com/fvbock/endless"
    "gin-skeleton/routes"
)

func main() {
    ConfigRuntime()
    StartWebServer()
}

func ConfigRuntime() {
    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)
    fmt.Printf("Running with %d CPUs\n", numCPU)
}

func routeEngine() *gin.Engine {
    gin.SetMode(gin.ReleaseMode)
    
    r := gin.Default()
    r.LoadHTMLGlob("views/*.tpl.html")
    r.Static("/public", "public")
    
    apiv1 := r.Group("/api/v1")
    {
        apiv1.GET("/ping", routes.Ping)
        apiv1.GET("/version", routes.GetVersion)
        apiv1.GET("/config", routes.GetConfig)
        apiv1.GET("/hello", routes.GetQueryString)
        apiv1.GET("/hello/:name/:action", routes.GetParam)
    }
    
    r.GET("/", routes.Index)
    
    return r
}

func StartWebServer() {
    endless.ListenAndServe(":9000", routeEngine())
}