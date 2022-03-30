package handler

import (
	"github.com/gin-gonic/gin"
)

type dataResponse struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type idResponse struct {
	ID interface{} `json:"id"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	//logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
