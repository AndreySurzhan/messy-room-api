package prompts

import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"gitlab.stripchat.dev/myclub/go-service/internal/config"
	"strings"
)

func New(cfg *config.Config) *ginprometheus.Prometheus {
	prompts := ginprometheus.NewPrometheus(strings.Replace(cfg.GetString("env.ServiceName"), "-", "_", -1))
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
