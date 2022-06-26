package handler

import (
	"fmt"
	"local/wwc_cocksize_bot/pkg/transport/rest/handler/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// bindData is helper function, returns false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {
	log.Println(c.ContentType())
	if c.ContentType() != "application/json" && c.ContentType() != "multipart/form-data" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		response.NewResponse(c, http.StatusBadRequest, msg)
		return false
	}
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
			// could probably extract this, it is also in middleware_auth_user
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			response.NewResponse(c, http.StatusBadRequest, "Invalid request parameters. See invalidArgs")
			return false
		}

		// later we'll add code for validating max body size here!

		// if we aren't able to properly extract validation errors,
		// we'll fallback and return an internal server error
		response.NewResponse(c, http.StatusHTTPVersionNotSupported, err.Error())
		return false
	}

	return true
}
