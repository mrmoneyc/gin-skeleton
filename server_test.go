package main

import (
    "testing"
    "net/http"
    "github.com/appleboy/gofight"
    "github.com/stretchr/testify/assert"
    "github.com/buger/jsonparser"
)

func TestPing(t *testing.T) {
    r := gofight.New()
    
    r.GET("/api/v1/ping").SetDebug(false).Run(routeEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
        expected, _ := jsonparser.GetString([]byte(res.Body.String()), "message")
        assert.Equal(t, http.StatusOK, res.Code)
        assert.Equal(t, "Pong", expected)
    })
}

func TestGetVersion(t *testing.T) {
    r := gofight.New()
    
    r.GET("/api/v1/version").SetDebug(false).Run(routeEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
        expected, _ := jsonparser.GetString([]byte(res.Body.String()), "Version")
        assert.Equal(t, http.StatusOK, res.Code)
        assert.Equal(t, "0.0.1", expected)
    })
}

func TestGetConfig(t *testing.T) {
    r := gofight.New()
    
    r.GET("/api/v1/config").SetDebug(false).Run(routeEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
        expected, _ := jsonparser.GetString([]byte(res.Body.String()), "Config", "AnotherCfg")
        assert.Equal(t, "true", expected)
        assert.Equal(t, http.StatusOK, res.Code)
    })
}

func TestGetQueryString(t *testing.T) {
    r := gofight.New()
    
    r.GET("/api/v1/hello?name=Gin&nickname=webframework").SetDebug(false).Run(routeEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
        expectedName, _ := jsonparser.GetString([]byte(res.Body.String()), "name")
        expectedNickName, _ := jsonparser.GetString([]byte(res.Body.String()), "nickname")
        assert.Equal(t, "Gin", expectedName)
        assert.Equal(t, "webframework", expectedNickName)
        assert.Equal(t, http.StatusOK, res.Code)
    })
}

func TestGetParam(t *testing.T) {
    r := gofight.New()
    
    r.GET("/api/v1/hello/Gin/Run").SetDebug(false).Run(routeEngine(), func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
        expectedName, _ := jsonparser.GetString([]byte(res.Body.String()), "name")
        expectedAction, _ := jsonparser.GetString([]byte(res.Body.String()), "action")
        assert.Equal(t, "Gin", expectedName)
        assert.Equal(t, "Run", expectedAction)
        assert.Equal(t, http.StatusOK, res.Code)
    })
}