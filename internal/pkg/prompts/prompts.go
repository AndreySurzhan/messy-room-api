package prompts

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zsais/go-gin-prometheus"
	"gitlab.stripchat.dev/myclub/go-service/internal/config"
	"strings"
)

func New(cfg *viper.Viper) *ginprometheus.Prometheus {
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
