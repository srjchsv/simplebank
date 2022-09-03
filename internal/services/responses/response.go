package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorResponse(err error) gin.H {
	logrus.Warn(err)
	return gin.H{"error": err.Error()}
}
