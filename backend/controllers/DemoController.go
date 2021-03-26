package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapLog try uber zap log
func ZapLog(c *gin.Context) {
	url := c.FullPath()

	//git repo: https://github.com/uber-go/zap

	/*
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any
		sugar := logger.Sugar()
		sugar.Infow("failed to fetch URL",
			// Structured context as loosely typed key-value pairs.
			"url", url,
			"attempt", 3,
			"backoff", time.Second,
		)
		sugar.Infof("Failed to fetch URL: %s", url)
	*/

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	Success(c, nil)
}
