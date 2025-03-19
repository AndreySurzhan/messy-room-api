package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// PostResource handler implementation
func (i *Service) PostResource(c *gin.Context) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}
