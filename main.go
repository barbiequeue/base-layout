package main

import "github.com/gin-gonic/gin"

func main() {
	var err error

	if err = conf.loadFromEnv(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	initLogger(conf.logLevel.String())

	engine := gin.Default()

	engine.Use(authMiddleware())
	setRoutes(engine)

	if err := engine.Run(conf.serverAddress()); err != nil {
		log.Fatalf("The web server stopped working unexpectedly: %v", err)
	}
}
