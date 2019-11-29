package logger

import (
	"os"
	"fmt"
	"time"
)

type FileLogger struct {
	level int
	logPath string
	logName string
	file *os.File
	warnFile *os.File
}

func NewFileLogger (config map[string]string) (logger LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("config not found : log_level")
		return
	}

	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("config not found : log_path")
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("config not found : log_name")
		return
	}

	level := GetLogLevel(logLevel)
	logger = &FileLogger{
		level: level,
		logPath: logPath,
		logName: logName,
	}

	logger.Init()
	return
}

func (f *FileLogger) Init(){
	// debug, trace, info 共用一个文件
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err: %v", filename, err))
	}
	f.file = file

	// warn, err, fatal 共用一个文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err: %v", filename, err))
	}
	f.warnFile = file
}

func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		f.level = LogLevelDebug
		return
	}
	f.level = level
}

func (f *FileLogger) writeLog(level int, format string, args ...interface{}) {
	var file *os.File
	if (level >= LogLevelWarn){
		file = f.warnFile
	} else {
		file = f.file
	}
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

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	f.writeLog(LogLevelDebug, format, args...)
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	f.writeLog(LogLevelTrace, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	f.writeLog(LogLevelInfo, format, args...)
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	f.writeLog(LogLevelWarn, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	f.writeLog(LogLevelError, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	f.writeLog(LogLevelFatal, format, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}