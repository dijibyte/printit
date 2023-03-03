package printit

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

func Error(message string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println("\033[31m", time.Now().Format("2006.01.02 15:04:05.000Z"), file+":"+fmt.Sprint(line), message, "\033[0m")
	} else {
		fmt.Println("\033[31m", time.Now().Format("2006.01.02 15:04:05.000Z"), message, "\033[0m")
	}
}

func Success(message string) {
	fmt.Println("\033[32m", message, "\033[0m")
}

func Warning(message string) {
	fmt.Println("\033[33m", message, "\033[0m")
}

func Info(message string) {
	fmt.Println("\033[34m", message, "\033[0m")
}

func JSON(v interface{}) {
	jsonbyte, _ := json.Marshal(v)
	println(string(jsonbyte))
}
