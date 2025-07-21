package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	/*
		Настраиваем логгер с дополнительными опциями:
		- префикс "serv" для всех записей
		- стандартная временная метка
		- краткое отображение имени файла и номера строки
	*/
	log := log.New(os.Stdout, `serv `, log.LstdFlags|log.Lshortfile)

	// Инициализируем новый экземпляр сервера, передавая ему настроенный логгер
	srv := server.NewServer(log)

	/*
		Запускаем HTTP-сервер на указанном адресе и обработчике
		В случае ошибки при запуске сервер завершится с выводом сообщения в лог
	*/
	if err := http.ListenAndServe(srv.Server.Addr, srv.Server.Handler); err != nil {
		log.Fatal(err)
	}
}
