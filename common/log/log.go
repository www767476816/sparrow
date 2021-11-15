package log

import (
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)
type LogLevel int32
const (
	Debug = 5
	Info = 4
	Warning = 3
	Error = 2
	Panic = 0
)
var g_Str2LogLevel=map[string]LogLevel{"debug":Debug,"info":Info,"warning":Warning,"error":Error,"panic":Panic}

func ConvertToLogLevel(level string) LogLevel{
	result,err :=g_Str2LogLevel[level]
	if !err{
		return LogLevel(-1)
	}

	return result
}
type LogInfo struct {
	*logrus.Logger
}
func (this* LogInfo)Init(level LogLevel){
	this.Logger=logrus.New()
	this.AddHook(createMyHook())
	this.SetReportCaller(true)
	this.SetLevel(logrus.Level(level))
}

type myHook struct {
	writer io.Writer
}

func createMyHook() *myHook {
	file_name := fmt.Sprintf("%d%d%d%d%d%d", time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second())

	log_file := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", file_name),
		MaxSize:    10,
		MaxAge:     0,
		MaxBackups: 0,
		LocalTime:  true,
		Compress:   false,
	}

	return &myHook{
		writer: log_file,
	}
}

func (hook *myHook) Fire(entry *logrus.Entry) error {
	var message string
	switch entry.Level {
	case logrus.InfoLevel:
		{
			message = fmt.Sprintf("[%s][%s][%s:%d][%s]\n", "Info", entry.Time.Format("2006-01-02 15:04:05"), entry.Caller.File, entry.Caller.Line, entry.Message)
			break
		}
	case logrus.WarnLevel:
		{
			message = fmt.Sprintf("[%s][%s][%s:%d][%s]\n", "Warning", entry.Time.Format("2006-01-02 15:04:05"), entry.Caller.File, entry.Caller.Line, entry.Message)
			break
		}
	case logrus.DebugLevel:
		{
			message = fmt.Sprintf("[%s][%s][%s:%d][%s]\n", "Debug", entry.Time.Format("2006-01-02 15:04:05"), entry.Caller.File, entry.Caller.Line, entry.Message)
			break
		}
	case logrus.ErrorLevel:
		{
			message = fmt.Sprintf("[%s][%s][%s:%d][%s]\n", "Error", entry.Time.Format("2006-01-02 15:04:05"), entry.Caller.File, entry.Caller.Line, entry.Message)
			break
		}
	case logrus.PanicLevel:
		{
			message = fmt.Sprintf("[%s][%s][%s:%d][%s]\n", "Panic", entry.Time.Format("2006-01-02 15:04:05"), entry.Caller.File, entry.Caller.Line, entry.Message)
			break
		}
	default:
		{
			return errors.New(fmt.Sprintf("logrus level error:%d", entry.Level))
		}
	}

	if _, write_err := hook.writer.Write([]byte(message)); write_err != nil {
		return write_err
	}
	return nil
}

func (hook *myHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.PanicLevel,
	}
}