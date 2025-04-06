// @title       E-Commerce API
// @version     1.0
// @description This is the API documentation for E-Commerce
// @host        localhost:8080
// @BasePath    /
// @schemes     http

// @contact.name  Vaibhav Jaiswal
// @contact.email vaibhav.jaiswal@gmail.com
// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	configs "e-commerce/config"
	"e-commerce/database/connections"
	"e-commerce/database/migrations"
	_ "e-commerce/docs"
	"e-commerce/middleware/auth"
	"e-commerce/middleware/compression"
	"e-commerce/middleware/ratelimiting"
	"e-commerce/middleware/requestlog"
	"e-commerce/modules/user_management/route"
	"e-commerce/services"
	configdata "e-commerce/utils/config_data"
	"e-commerce/utils/helper"
)

func main() {

	// Load config.env file
	if err := godotenv.Load("config.env"); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to load .env file: %s\n", err)
		os.Exit(1)
	}

	// Get all config data from config.json file (including config.env data)
	configData, err := configs.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to read config data: %s\n", err)
		os.Exit(1)
	}

	// init db connection & initialize a global db connection variable to use for db operations
	err = connections.InitDB(&configData.DBConnection)
	if err != nil {
		fmt.Printf("\ncan not connect to db...! \nError: %s", err.Error())
		fmt.Fprintf(os.Stderr, "❌ Failed to connect to DB: %s\n", err)
		os.Exit(1)
	}

	// init redis connection & initialize a global redis connection variable to use for redis operations
	err = connections.InitRedis(&configData.RedisConnection)
	if err != nil {
		fmt.Printf("\ncan not connect to redis...! \nError: %s", err.Error())
		fmt.Fprintf(os.Stderr, "❌ Failed to connect to Redis: %s\n", err)
		os.Exit(1)
	}

	// init email smtp connection & initialize a global smtp connection variable to use for email notification
	services.InitSmtpServer(configData.SmtpServer)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = configData.AllowedOrigins
	config.AllowMethods = configData.AllowedMethods

	fmt.Printf("CORS configured for: %v\n", configData.AllowedOrigins)

	router.Use(cors.New(config))

	helper.InitiateHelper(*configData)

	router.Use(requestlog.Logger(configData.Logger.Request.LogDir))

	router.Use(compression.Compression())

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/api-docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api-docs/index.html")
	})

	router.GET("/load-data", configdata.PreLoadDataHandler)

	router.Use(auth.Auth())

	router.Use(
		ratelimiting.RateLimiter(
			configData.RateLimit.MaxRequest,
			time.Duration(int(time.Minute)*configData.RateLimit.Duration),
			connections.GetRedisClient(),
		),
	)

	if err := migrations.RunMigrations(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Migration error: %s\n", err)
		os.Exit(1)
	}

	// add modules
	route.UserManagementRoutes(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", configData.Server.Port),
		Handler: router,
	}

	fmt.Printf("Server running on port: %s\n", configData.Server.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Err : %s\n", err)
		}
	}()

	// This code is for cleanup after the graceful shutdown
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)
	<-gracefulShutdown

	fmt.Printf("\nShutting down servers...!")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := connections.DeInitDB(); err != nil {
		fmt.Printf("\nClosing connections : %s\n", err)
	}

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("\nServer forced to shutdown : %s\n", err)
	}

	fmt.Printf("\nServers shut down...!")
}
