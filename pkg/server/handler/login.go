package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type loginInput struct {
	UserId int64 `json:"user-id" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	var input loginInput
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

}
