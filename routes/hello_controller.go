package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetQueryString(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "API": "GetQueryString",
        "name": c.DefaultQuery("name", "Guest"),
        "nickname": c.Query("nickname"),
    })
}

func GetParam(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "API": "GetParam",
        "name": c.Param("name"),
        "action": c.Param("action"),
    })
}

