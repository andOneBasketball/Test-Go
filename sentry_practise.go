package main

import (
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://97cfb5b629f96e8e1a7b33c4e15f4674@o4508256997277696.ingest.us.sentry.io/4508256999768064",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for tracing.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Then create your app
	app := gin.Default()

	// Once it's done, you can attach the handler as one of your middleware
	app.Use(sentrygin.New(sentrygin.Options{}))

	// Set up routes
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world!")
	})

	app.GET("/foo", func(ctx *gin.Context) {
		// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
		// in the middleware before, it will be sent through as well
		panic("y tho")
	})

	// And run it
	app.Run(":3000")
}
