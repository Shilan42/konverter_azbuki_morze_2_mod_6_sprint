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
	// Используем http.ServeFile для открытия, чтения и отправки файла
	http.ServeFile(w, r, `C:\Users\ppv20\Desktop\Dev\Projects of the 2nd mod\konverter_azbuki_morze_2_mod_6_sprint\index.html`)
}

// processUploadRequest обрабатывает загрузку файла и конвертацию данных
func ProcessUploadRequest(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса и если это не MethodPost, то возвращаем ошибку
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP method not allowed. Only POST requests are supported", http.StatusMethodNotAllowed)
		return
	}

	// Получение файла из поля формы index.html = "name"
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to retrieve file from request form", http.StatusBadRequest)
		return
	}
	// Закрываем файл после работы с ним
	defer file.Close()

	// Использование прямого чтения в буфер фиксированного размера
	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, file); err != nil {
		http.Error(w, "failed to read file content", http.StatusInternalServerError)
		return
	}

	// Передаем данные в функцию автоопределения из пакета service, которую мы создали, чтобы получить переконвертируемую строку
	str, err := service.DetectMorseOrText(buf.String())
	if err != nil {
		http.Error(w, "failed to convert data", http.StatusInternalServerError)
		return
	}

	// Формируем имя файла для сохранения
	timestamp := time.Now().UTC().Format("2006_01_02_15_04_05")
	ext := filepath.Ext(handler.Filename)
	outputFilename := timestamp + ext

	// Создаем локальный файл для записи результатов конвертации строки
	resultFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "failed to create file", http.StatusInternalServerError)
		return
	}
	// Закрываем файл после работы с ним
	defer resultFile.Close()

	// Записываем результаты в файл
	if _, err := fmt.Fprintf(resultFile, "%s:\n", str); err != nil {
		http.Error(w, "failed to write data to output file", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/plain")

	// Возвращаем результат конвертации в ответ
	_, err = w.Write([]byte(str))
	if err != nil {
		http.Error(w, "failed to write response to client", http.StatusInternalServerError)
		return
	}
}
