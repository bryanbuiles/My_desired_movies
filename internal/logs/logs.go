package logs

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO:", 0)
	errorLogger = log.New(os.Stdout, "", log.Lshortfile|log.Ltime) // Lshortfile muestra el numbre final y la li
	// Ltime muestra el tiempo
)

// Info To display Info in stdout
func Info(s interface{}) {
	infoLogger.Println(s)
}

// Error log para el error
func Error(s interface{}) {
	errorLogger.Println(s)
}
