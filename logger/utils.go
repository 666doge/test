package logger

import (
	"runtime"
)

func GetLineInfo() (fileName string, fnName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fileName = file
		lineNo = line
		fnName = runtime.FuncForPC(pc).Name()
	}
	return
}