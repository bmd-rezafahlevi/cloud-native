package app

import (
	"net/http"

	"github.com/bmd-rezafahlevi/cloud-native/util/logger"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

const (
	appErrDataAccessFailure      = "data access failure"
	appErrJsonCreationFailure    = "json creation failure"
	appErrDataCreationFailure    = "data creation failure"
	appErrFormDecodingFailure    = "form decoding failure"
	appErrDataUpdateFailure      = "data update failure"
	appErrFormErrResponseFailure = "form error response failure"
)

type App struct {
	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func New(logger *logger.Logger, db *gorm.DB, validator *validator.Validate) *App {
	return &App{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}

func (app *App) HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nonsif")

	w.Write([]byte("Hello World!"))
}
