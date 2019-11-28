package logger

import (
	"os"
	"fmt"
)

type FileLogger struct {
	level int
	logPath string
	logName string
	file *os.File
	warnFile *os.File
}

func NewFileLogger (level int, logPath string, logName string) *FileLogger {
	logger := &FileLogger{
		level: level,
		logPath: logPath,
		logName: logName,
	}

	logger.init()
	return logger
}

func (f *FileLogger) init(){
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

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}