package logger

import (
	"runtime"
	"os"
	"time"
	"fmt"
)

func GetLineInfo() (fileName string, fnName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		lineNo = line
		fnName = runtime.FuncForPC(pc).Name()
	}
	return
}

func  WriteLog(file *os.File, level int, format string, args ...interface{}) {
	nowString := time.Now().Format("2006-01-02 15:04:05")
	fileName, funcName, lineNo := GetLineInfo()

	fmt.Fprintf(
		file, "%s [%s] %s %s:%d ",
		nowString,
		LogLevelText(level),
		fileName,
		funcName,
		lineNo,
	)
	fmt.Fprintf(file, format, args...)
	fmt.Fprintln(file)
}