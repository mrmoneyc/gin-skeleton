package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gin-skeleton/utils"
    "gin-skeleton/models"
)

func Ping(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Pong",
    })
}

func GetVersion(c *gin.Context) {
    config, err := utils.ReadConfig()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "API": "GetVersion",
            "Error": err,
        })    
    }
    
    c.JSON(http.StatusOK, gin.H{
        "API": "GetVersion",
        "Version": config.App.Version,
    })
}

func GetConfig(c *gin.Context) {
    config, err := models.GetConfig()
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "API": "GetConfig",
            "Error": err,
        })    
    }
    
    c.JSON(http.StatusOK, gin.H{
        "API": "GetConfig",
        "Config": config,
    })
}