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
	LogInfoChan chan *LogInfo
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
		LogInfoChan: make(chan *LogInfo, 500),
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

	go f.writeLog()
}

func (f *FileLogger) writeLog() {
	for logInfo := range f.LogInfoChan {
		var file *os.File
		if logInfo.IsWarn {
			file = f.warnFile
		} else {
			file = f.file
		}
		fmt.Fprintf(file, logInfo.LogMsg)
	}
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
	logInfo := GetlogInfo(LogLevelDebug, format, args...)
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logInfo := GetlogInfo(LogLevelTrace, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logInfo := GetlogInfo(LogLevelInfo, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logInfo := GetlogInfo(LogLevelWarn, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logInfo := GetlogInfo(LogLevelError, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logInfo := GetlogInfo(LogLevelFatal, format, args...)
	
	select {
	case f.LogInfoChan <- logInfo:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}