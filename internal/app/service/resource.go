package service

import (
	"github.com/AndreySurzhan/messy-room-api/internal/app/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// PostResource handler implementation
func (i *Service) PostResource(c *gin.Context) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}

// DeleteResourceID handler implementation
func (i *Service) DeleteResourceID(c *gin.Context, id uint64) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}

// GetResourceID handler implementation
func (i *Service) GetResourceID(c *gin.Context, id uint64, params api.GetResourceIDParams) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}

// PutResourceID handler implementation
func (i *Service) PutResourceID(c *gin.Context, id uint64) {
	_ = c.AbortWithError(500, errors.New("not implemented"))
}
