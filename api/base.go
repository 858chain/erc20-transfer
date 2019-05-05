package api

import (
	"github.com/gin-gonic/gin"
)

func R(data interface{}) gin.H {
	return gin.H{
		"message": data,
	}
}
