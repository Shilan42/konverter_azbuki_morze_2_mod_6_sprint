package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

/*
Функция DetectMorseOrText определяет тип текста (морзе или обычный текст)
и выполняет соответствующее преобразование
*/
func DetectMorseOrText(text string) (string, error) {
	/*
		Удаляем пробелы и проверяем, что строка не пустая.
		Если строка пустая, возвращаем ошибку
	*/
	if text = strings.TrimSpace(text); text == "" {
		return "", errors.New("func DetectMorseOrText: empty string")
	}
	/*
		Инициализируем счётчики для символов морзе (точки и тире)
		и буквенных символов обычного текста
	*/
	var morseCount, textCount int
	/*
		Анализируем каждый символ в строке:
		увеличиваем morseCount для точек и тире,
		textCount - для буквенных символов
	*/
	for _, char := range text {
		switch {
		case char == '.' || char == '-':
			morseCount++
		case char == ' ':
			continue // Пропускаем пробелы
		case unicode.IsLetter(char):
			textCount++
		}
	}
	/*
		На основе соотношения символов определяем тип текста
		и вызываем соответствующую функцию преобразования
	*/
	if morseCount > textCount {
		// Если больше символов Морзе - преобразуем в обычный текст
		return morse.ToText(text), nil
	}
	// Иначе преобразуем обычный текст в морзе
	return morse.ToMorse(text), nil
}
