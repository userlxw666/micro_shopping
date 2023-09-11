package api_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Msg string
}

func HandleError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Msg: err.Error(),
	})
}
