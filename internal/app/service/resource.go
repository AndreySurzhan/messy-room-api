package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.stripchat.dev/myclub/go-service/internal/app/api"
)

// PostResource handler implementation
func (i *Service) PostResource(c *gin.Context) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}

// DeleteResourceId handler implementation
func (i *Service) DeleteResourceId(c *gin.Context, id uint64) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}

// GetResourceId handler implementation
func (i *Service) GetResourceId(c *gin.Context, id uint64, params api.GetResourceIdParams) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}

// PutResourceId handler implementation
func (i *Service) PutResourceId(c *gin.Context, id uint64) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}
