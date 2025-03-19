package controller

import (
	"github.com/gin-gonic/gin"
)

func (ctr *Controller) PostGetRoomCleanlinessStatus(c *gin.Context) {
	picture := c.PostForm("picture")

	if picture == "" {
		c.JSON(400, gin.H{"error": "picture is required"})
		return
	}

	status, err := ctr.service.GetRoomCleanlinessStatus(picture)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		200,
		gin.H{
			"items": []string{status},
			"prompt": gin.H{
				"speech": "",
				"text":   status,
			},
		},
	)
}
