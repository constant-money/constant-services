package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultController :
type DefaultController struct{}

// Home :
func (d DefaultController) Home(c *gin.Context) {
	resp := JSONResponse{1, "IPFS REST API", nil}
	c.JSON(http.StatusOK, resp)
}

// Send :
func (d DefaultController) Send(c *gin.Context) {
	var jsonData map[string]interface{}

	c.Bind(&jsonData)
	data, ok := jsonData["data"]

	if !ok {
		resp := JSONResponse{0, "Invalid data", nil}
		c.JSON(http.StatusOK, resp)
		c.Abort()
		return
	}

	var status int
	var message string

	result, err := fcmService.Send(data.(map[string]interface{}))

	if result {
		status = 1
	} else {
		status = 0
		if err != nil {
			message = err.Error()
		}
	}

	resp := JSONResponse{status, message, nil}
	c.JSON(http.StatusOK, resp)
}

// NotFound :
func (d DefaultController) NotFound(c *gin.Context) {
	resp := JSONResponse{0, "Page not found", nil}
	c.JSON(http.StatusOK, resp)
}
