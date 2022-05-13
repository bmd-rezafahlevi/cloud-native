package app

import (
	"net/http"

	"github.com/bmd-rezafahlevi/cloud-native/util/logger"
	"github.com/jinzhu/gorm"
)

const (
	appErrDataAccessFailure   = "data access failure"
	appErrJsonCreationFailure = "json creation failure"
	appErrDataCreationFailure = "data creation failure"
	appErrFormDecodingFailure = "form decoding failure"
	appErrDataUpdateFailure   = "data update failure"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{
		logger: logger,
		db:     db,
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
