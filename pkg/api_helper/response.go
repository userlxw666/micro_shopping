package api_helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg interface{}
}

func ResponseHandler(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, msg)
}
