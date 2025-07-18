package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server - структура, представляющая HTTP-сервер
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// NewServer - функция для создания нового экземпляра сервера
func NewServer(l *log.Logger) *Server {

	mux := http.NewServeMux()

	// Регистрация хендлеров
	mux.HandleFunc("/", handlers.HandleRootRequest)
	mux.HandleFunc("/upload", handlers.ProcessUploadRequest)

	// Настройка параметров сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Возврат нового экземпляра сервера с настроенными параметрами
	return &Server{
		Logger: l,
		Server: server,
	}
}
