package response

import (
	"github.com/gin-gonic/gin"
)

type dataResponse struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}
type DataResponse struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
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

func NewResponse(c *gin.Context, statusCode int, message string) {
	//logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
func JsonResponse(c *gin.Context, message interface{}, statusCode int) {
	c.JSON(statusCode, message)
}
