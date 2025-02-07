package domain

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/acarl005/stripansi"
)

type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data"`
}

type CustomFileWriter struct {
	File *os.File
}

func (cfw *CustomFileWriter) Write(p []byte) (n int, err error) {
	//Пример модификации текста: добавляем префикс "[MODIFIED] "
	modifiedText, err := formatTextToLog(string(p))
	if err != nil {
		panic(err)
	}

	// Добавляем символ новой строки
	space := string(modifiedText) + "\n"

	return cfw.File.Write([]byte(space))
}

func formatTextToLog(logText string) ([]byte, error) {
	logSplit := strings.Split(logText, " ")

	// Timestamp
	date := strings.Replace(strings.Replace(logSplit[0], "[", "", -1), "]", "", -1)

	// Level
	logLevel := stripansi.Strip(strings.Replace(logSplit[1], ":", "", -1))

	// Message
	msgSplit := logSplit[2:]
	message := stripansi.Strip(strings.Split(strings.Join(msgSplit, " "), "{")[0])

	// Data
	jsonData := ParseAndFormatJSON(logText)

	data := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return nil, err
	}

	entry := LogEntry{
		Timestamp: date,
		Level:     logLevel,
		Message:   message,
		Data:      data,
	}

	logBytes, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	return logBytes, nil
}

func ParseAndFormatJSON(text string) string {
	// Находим первую фигурную скобку и последнюю
	firstBrace := strings.Index(text, "{")
	lastBrace := strings.LastIndex(text, "}")

	if firstBrace == -1 || lastBrace == -1 || lastBrace < firstBrace {
		// JSON не найден
		return "{}"
	}

	// Извлекаем JSON-строку
	jsonPart := text[firstBrace : lastBrace+1]

	// Форматируем JSON
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonPart), &jsonData)
	if err != nil {
		panic(err)
	}

	// Кодируем обратно в форматированный JSON
	formattedJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		panic(err)
	}

	// Возвращаем отформатированный JSON
	return string(formattedJSON)
}
