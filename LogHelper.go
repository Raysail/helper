package helper

import (
	"errors"
	"fmt"
	"os"
	"time"
)

import (
	log "github.com/Sirupsen/logrus"
)

type KLog struct {
	logrusInstance *log.Logger
	outPutPath     string
	logName        string
	handler        *os.File
	handlerKey     string
	rotationMode   int
}

const (
	ROTATION_TYPE_NONE    = 0
	ROTATION_TYPE_HOUR    = 1 //TODO
	ROTATION_TYPE_DAILY   = 2
	ROTATION_TYPE_MONTHLY = 3 //TODO
)

func NewLogger(filePath, fileName string, rotationMode int) (logger *KLog) {
	logger = &KLog{
		logrusInstance: log.New(),
		outPutPath:     filePath,
		logName:        fileName,
		rotationMode:   rotationMode,
	}
	logger.logrusInstance.Formatter = &log.JSONFormatter{}
	return logger
}

func (logger *KLog) Close() {
	logger.logrusInstance.Writer().Close()
	logger.handler.Close()
}

func (logger *KLog) Debug(content string) {
	logger.init()
	logger.logrusInstance.Debugln(content)
}

func (logger *KLog) Debugf(format string, args ...interface{}) {
	logger.init()
	logger.logrusInstance.Debugf(format+"\n\r", args...)
}

func (logger *KLog) Info(content string) {
	logger.init()
	logger.logrusInstance.Infoln(content)
}

func (logger *KLog) Infof(format string, args ...interface{}) {
	logger.init()
	logger.logrusInstance.Infof(format+"\n\r", args...)
}

func (logger *KLog) Warn(content string) {
	logger.init()
	logger.logrusInstance.Warnln(content)
}

func (logger *KLog) Warnf(format string, args ...interface{}) {
	logger.init()
	logger.logrusInstance.Warnf(format+"\n\r", args...)
}

func (logger *KLog) Error(content string) {
	logger.init()
	logger.logrusInstance.Errorln(content)
}

func (logger *KLog) Errorf(format string, args ...interface{}) {
	logger.init()
	logger.logrusInstance.Errorf(format+"\n\r", args...)
}

func (logger *KLog) Fatal(content string) {
	logger.init()
	logger.logrusInstance.Fatalln(content)
}

func (logger *KLog) Fatalf(format string, args ...interface{}) {
	logger.init()
	logger.logrusInstance.Fatalf(format+"\n\r", args...)
}

func (logger *KLog) init() (err error) {
	if logger.outPutPath == "" || logger.logName == "" {
		return errors.New("invalid path")
	}
	logKey := time.Now().Format("2006-01-02")
	if logKey != logger.handlerKey {
		logger.handler.Close()
		logger.handler, err = os.OpenFile(logger.outPutPath+logger.logName+"_"+logKey+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		logger.handlerKey = logKey
		logger.logrusInstance.Out = logger.handler
		logger.logrusInstance.WriterLevel(log.DebugLevel)
	}
	return err
}

func (logger *KLog) Print(v ...interface{}) {
	fmt.Print(v...)
}
func (logger *KLog) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func (logger *KLog) Println(v ...interface{}) {
	fmt.Println(v...)
}
