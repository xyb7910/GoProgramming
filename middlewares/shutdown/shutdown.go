package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		c.String(200, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// use go routine to start server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal to gracefully shut down the server and set a timeout of 10 seconds
	quit := make(chan os.Signal, 1)                      // create channel to receive OS signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // notify channel to receive OS signals
	<-quit                                               // wait for OS signal to shut down server
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 10 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
