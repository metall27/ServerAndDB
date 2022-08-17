package api

import (
	"net/http"

	"github.com/metall27/ServerAndDB/storage"
	"github.com/sirupsen/logrus"
)

// Пытаемся отконфигурировать наш API инстанс (а конкретнее - поле logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

//Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is REST API!"))
	})
}

// Пытаеся конфигурировать наше хранилище (storage API)
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	// Пытаемся установить соединение, если невозможно - возвращаем ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
