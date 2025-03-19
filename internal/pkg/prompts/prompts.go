package prompts

import (
	"github.com/AndreySurzhan/messy-room-api/internal/config"
	"github.com/gin-gonic/gin"
	"strings"
)

func New(cfg *config.Config) *ginprometheus.Prometheus {
	prompts := ginprometheus.NewPrometheus(strings.Replace(cfg.GetString(config.ServiceName), "-", "_", -1))
	prompts.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "name" {
				url = strings.Replace(url, p.Value, ":name", 1)
				break
			}
		}
		return url
	}

	return prompts
}
