package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"mattermost-tools/internal/config"
	"mattermost-tools/internal/core/handler"
	"mattermost-tools/pkg/environment"
	"mattermost-tools/pkg/mattermost"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const (
	defaultIdleTimeout  = time.Minute
	defaultReadTimeout  = 1 * time.Minute
	defaultWriteTimeout = 1 * time.Minute
)

func main() {
	cfg, _ := config.InitConfig()
	r := gin.New()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	client := mattermost.NewClient(environment.Environment("local"), cfg.Mattermost)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":4000"),
		Handler:      r,
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	h := handler.NewMattermostHandler(&logger, client)
	r.Use(
		gin.Recovery(),
		JsonLoggerMiddleware(&logger),
	)

	r.POST("/mattermost/emoji", h.PostEmojiOnPost)
	//r.POST("/mattermost/emoji/all", h.PostAllEmoji)
	r.DELETE("/mattermost/emoji", h.DeleteAllEmojiOnPost)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func JsonLoggerMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()
		logger.Info().Str("client_ip", c.ClientIP()).
			Int64("duration(ms)", time.Now().Sub(start).Milliseconds()).
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).Int("status", c.Writer.Status()).
			Msg("")
	}
}
