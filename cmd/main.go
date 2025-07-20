package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	/*
		Создаем или открываем файл для записи логов
		Если файл не существует, он будет создан
		Параметры открытия: добавление в конец файла, создание при отсутствии, только запись
	*/
	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Закрываем файл после работы с ним
	defer flog.Close()

	/*
		Настраиваем логгер с дополнительными опциями:
		- префикс "serv" для всех записей
		- стандартная временная метка
		- краткое отображение имени файла и номера строки
	*/
	mylog := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)

	// Инициализируем новый экземпляр сервера, передавая ему настроенный логгер
	srv := server.NewServer(mylog)

	/*
		Запускаем HTTP-сервер на указанном адресе и обработчике
		В случае ошибки при запуске сервер завершится с выводом сообщения в лог
	*/
	if err := http.ListenAndServe(srv.Server.Addr, srv.Server.Handler); err != nil {
		mylog.Fatal(err)
	}
}
