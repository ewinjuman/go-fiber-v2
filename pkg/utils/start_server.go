package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	//  syscall.SIGHUP,  // kill -SIGHUP XXXX
	//	syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
	//	syscall.SIGQUIT, // kill -SIGQUIT XXXX

	go func() {
		sigint := make(chan os.Signal)
		//signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		signal.Notify(sigint, syscall.SIGHUP)
		signal.Notify(sigint, syscall.SIGTERM)
		signal.Notify(sigint, syscall.SIGINT)
		<-sigint
		//println(fmt.Printf("caught signal: %v", sigint))
		log.Println("\nWaiting for finish...")
		// Received an interrupt signal, shutdown.
		if err := a.ShutdownWithTimeout(20 * time.Second); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}
		close(idleConnsClosed)
		log.Printf("Done")
	}()
	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

// AnotherWayStartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func AnotherWayStartServerWithGracefulShutdown(a *fiber.App) {
	fiberConnURL, _ := ConnectionURLBuilder("fiber")
	// Listen from a different goroutine
	go func() {
		if err := a.Listen(fiberConnURL); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal)                                                       // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = a.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
