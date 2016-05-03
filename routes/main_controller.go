package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gin-skeleton/utils"
)

func Index(c *gin.Context) {
    config, _ := utils.ReadConfig()
    
    c.HTML(http.StatusOK, "index.tpl.html", gin.H{
        "appName": config.App.Name,
        "appVersion": config.App.Version,
    })
}