package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	dbConn "github.com/bmd-rezafahlevi/cloud-native/adapter/gorm"
	"github.com/bmd-rezafahlevi/cloud-native/app/app"
	"github.com/bmd-rezafahlevi/cloud-native/app/router"
	"github.com/bmd-rezafahlevi/cloud-native/config"
	lr "github.com/bmd-rezafahlevi/cloud-native/util/logger"
	vr "github.com/bmd-rezafahlevi/cloud-native/util/validator"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}

	validator := vr.New()

	application := app.New(logger, db, validator)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
