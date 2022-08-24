package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func errorResponse(err error) gin.H {
	logrus.Warn(err)
	return gin.H{"error": err.Error()}
}
