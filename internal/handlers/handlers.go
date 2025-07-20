package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

/*
handleRootRequest - основная функция-обработчик корневого запроса
обрабатывает GET-запросы к корневому URL и возвращает HTML-страницу
*/
func HandleRootRequest(w http.ResponseWriter, r *http.Request) {
	// Открываем файл index.html для чтения. Если произошла ошибка при открытии файла, возвращаем сообщение об ошибке
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "func handleRootRequest: failed to open file", http.StatusInternalServerError)
		return
	}
	// Закрываем файл после завершения работы
	defer file.Close()

	// Читаем содержимое файла в память. Если произошла ошибка при чтении файла, возвращаем сообщение об ошибке
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "func handleRootRequest: failed to read file", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type для корректного отображения HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Отправляем содержимое файла в ответ
	w.Write(data)
}

// processUploadRequest обрабатывает загрузку файла и конвертацию данных
func ProcessUploadRequest(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса и если это не MethodPost, то возвращаем ошибку
	if r.Method != http.MethodPost {
		http.Error(w, "func processUploadRequest: HTTP method not allowed. Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}

	// Получение файла из поля формы index.html = "name"
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "func processUploadRequest: failed to retrieve file from request form", http.StatusBadRequest)
		return
	}
	// Закрываем файл после работы с ним
	defer file.Close()

	// Использование прямого чтения в буфер фиксированного размера
	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, file); err != nil {
		http.Error(w, "func processUploadRequest: failed to read file content", http.StatusInternalServerError)
		return
	}

	// Передаем данные в функцию автоопределения из пакета service, которую мы создали, чтобы получить переконвертируемую строку
	str, err := service.DetectMorseOrText(buf.String())
	if err != nil {
		http.Error(w, "func processUploadRequest: DetectMorseOrText: failed to convert data", http.StatusInternalServerError)
		return
	}

	// Формируем имя файла для сохранения
	timestamp := time.Now().UTC().String()
	ext := filepath.Ext(handler.Filename)
	outputFilename := timestamp + ext

	// Создаем локальный файл для записи результатов конвертации строки
	resultFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "func processUploadRequest: failed to create file", http.StatusInternalServerError)
		return
	}
	// Закрываем файл после работы с ним
	defer resultFile.Close()

	// Записываем результаты в файл
	if _, err := fmt.Fprintf(resultFile, "%s:\n", str); err != nil {
		http.Error(w, "func processUploadRequest: failed to write data to output file.", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат конвертации в ответ
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(str))
}
