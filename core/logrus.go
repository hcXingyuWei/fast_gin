package core

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/sirupsen/logrus"
)

type MyLog struct {
}

type MyHook struct {
	file     *os.File   //当前打开的日志
	errFile  *os.File   //错误日志日志
	fileDate string     //当前日志的时间
	logPath  string     //当前日志目录
	mu       sync.Mutex //预防冲突
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func InitLogger() {
	logrus.SetLevel(logrus.DebugLevel)
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&MyLog{})
	logrus.SetReportCaller(true)
	logrus.AddHook(&MyHook{})
}

func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	var leverColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		leverColor = gray
	case logrus.WarnLevel:
		leverColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		leverColor = red
	default:
		leverColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, leverColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

// Fire 写入到文件 按时间分片 错误日志单独存放
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	hook.mu.Lock()
	defer hook.mu.Unlock()
	hook.logPath = "logs"
	timer := entry.Time.Format("2006-01-02")
	line, err := entry.String()
	if err != nil {
		return fmt.Errorf("failed to format log enty : %v", err)
	}

	if hook.fileDate != timer {
		if err := hook.rotateFiles(timer); err != nil {
			return err
		}
	}

	if _, err := hook.file.Write([]byte(line)); err != nil {
		return fmt.Errorf("failed to write to log file: %v", err)
	}

	if entry.Level <= logrus.ErrorLevel {
		if _, err := hook.errFile.Write([]byte(line)); err != nil {
			return fmt.Errorf("failed to write to error log file: %v", err)
		}
	}

	return nil
}

func (hook *MyHook) rotateFiles(timer string) error {
	if hook.file != nil {
		if err := hook.file.Close(); err != nil {
			return fmt.Errorf("failed to close file: %v", err)
		}
	}

	if hook.errFile != nil {
		if err := hook.errFile.Close(); err != nil {
			return fmt.Errorf("failed to close file: %v", err)
		}
	}

	dirName := fmt.Sprintf("%s/%s", hook.logPath, timer)
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	infoFileName := fmt.Sprintf("%s/info.log", dirName)
	errFileName := fmt.Sprintf("%s/err.log", dirName)

	var err error
	hook.file, err = os.OpenFile(infoFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	hook.errFile, err = os.OpenFile(errFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	hook.fileDate = timer

	return nil
}

func (hook *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
