package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ibhesholihin/hevent/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	myhndlr "github.com/ibhesholihin/hevent/apps/http/handler"
	appMiddleware "github.com/ibhesholihin/hevent/apps/http/middleware"
	myrepo "github.com/ibhesholihin/hevent/apps/repository"
	myserv "github.com/ibhesholihin/hevent/apps/service"

	"github.com/ibhesholihin/hevent/config/db"
	httpRoutes "github.com/ibhesholihin/hevent/routes"
	"github.com/ibhesholihin/hevent/utils"
	"github.com/ibhesholihin/hevent/utils/crypto"
	"github.com/ibhesholihin/hevent/utils/jwt"
	"github.com/ibhesholihin/hevent/utils/paygate"
)

func main() {

	// Load config
	configApp := config.LoadConfig()
	ctxTimeout := time.Duration(configApp.ContextTimeout) * time.Second

	PORT := os.Getenv("APP_PORT")
	payment_route := configApp.BaseURL + PORT + "/payment"

	// Setup utils
	cryptoSvc := crypto.NewCryptoService()
	jwtSvc := jwt.NewJWTService(configApp.JWTSecretKey)
	payService := paygate.NewPaymentService(configApp.MIDTRANS_SERVER_KEY, payment_route)

	// Setup db config
	dbInstance, err := db.Database(configApp)
	utils.PanicIfNeeded(err)

	//migrate model/table
	db.Migrate(dbInstance)

	// Setup repository interface
	//redisRepo := redisRepository.NewRedisRepository(cacheInstance)
	myRepo := myrepo.NewRepository(dbInstance)

	// Setup service
	myService := myserv.NewService(myRepo, cryptoSvc, jwtSvc, payService, ctxTimeout)

	//setup handler
	myHandler := myhndlr.NewHandler(myService)

	// Setup app middleware
	appMiddleware := appMiddleware.NewMiddleware(jwtSvc)

	//echo instances
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(appMiddleware.Logger(nil))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "i am alive")
	})

	//setup routes
	httpRoutes.MyRoutes(e, appMiddleware, myHandler)

	if PORT == "" {
		PORT = "8778"
	}

	// Start server
	go func() {
		if err := e.Start(":" + PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configApp.ContextTimeout)*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
