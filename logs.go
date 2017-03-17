/**
 * Created by angelina-zf on 17/2/25.
 */

// yeego 日志相关的功能
// 依赖： "github.com/Sirupsen/logrus"
// 基于logrus的log类,自己搞成了分级的形式
// 可以设置将error层次的log发送到es
package yeego

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v2"
	"gopkg.in/olivere/elastic.v5"
	"os"
	"yeego/yeeFile"
	"time"
	"runtime"
	"yeego/yeeTime"
)

var infoLogS *logrus.Logger
var debugLogS *logrus.Logger
var errorLogS *logrus.Logger
var day map[string]int
var logfile map[string]*os.File

type LogFields logrus.Fields

// LogPath 日志的存储地址，按照时间存储
var LogPath string = "logs/" + yeeTime.DateFormat(time.Now(), "YYYY-MM-DD")
// ESPath es服务器的地址以及端口
var ESPath string = "http://localhost:9200"
// ESFromHost 从哪个host发送过来的
var ESFromHost string = "localhost"
// ESIndex 要存储在哪个index索引下,默认的type为log
var ESIndex string = "testlog"

func Init() {
	if !yeeFile.FileExists(LogPath) {
		if createErr := yeeFile.MkdirFile(LogPath); createErr != nil {
			Print(createErr)
		}
	}
	day = make(map[string]int)
	logfile = make(map[string]*os.File)
	infoLogS = logrus.New()
	setLogSConfig(infoLogS, logrus.InfoLevel)
	debugLogS = logrus.New()
	setLogSConfig(debugLogS, logrus.DebugLevel)
	errorLogS = logrus.New()
	setLogSConfig(errorLogS, logrus.ErrorLevel)
}

// MustInitESErrorLog 为error级别的log注册es
// 有错误则直接panic
func MustInitESErrorLog() {
	//TODO 用的是utc时间需要修改源码为Local
	client, err := elastic.NewClient(elastic.SetURL(ESPath))
	if err != nil {
		panic(err.Error())
	}
	hook, err := elogrus.NewElasticHook(client, ESFromHost, logrus.ErrorLevel, ESIndex)
	if err != nil {
		panic(err.Error())
	}
	errorLogS.Hooks.Add(hook)
}

func setLogSConfig(logger *logrus.Logger, level logrus.Level) {
	var err error
	logger.Level = level
	logger.Formatter = new(logrus.JSONFormatter)
	logfile[level.String()], err = os.OpenFile(getLogFullPath(level), os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		logfile[level.String()], err = os.Create(getLogFullPath(level))
		if err != nil {
			Print(err)
		}
	}
	//TODO 根据配置的运行模式选择是命令行输出还是文件输出
	logger.Out = os.Stdout
	//logger.Out = logfile[level.String()]
	day[level.String()] = yeeTime.Day()
}

// updateLogFile 检测是否跨天了,把记录记录到新的文件目录中
func updateLogFile(level logrus.Level) {
	var err error
	day2 := yeeTime.Day()
	if day2 != day[level.String()] {
		logfile[level.String()].Close()
		logfile[level.String()], err = os.Create(getLogFullPath(level))
		if err != nil {
			Print(err)
		}
		switch level {
		case logrus.InfoLevel:
			infoLogS.Out = logfile[level.String()]
		case logrus.DebugLevel:
			debugLogS.Out = logfile[level.String()]
		case logrus.ErrorLevel:
			errorLogS.Out = logfile[level.String()]
		}
	}
}

// locate 找到是哪个文件的哪个地方打出的log
func locate(fields LogFields) LogFields {
	_, path, line, ok := runtime.Caller(2)
	if ok {
		fields["file"] = path
		fields["line"] = line
	}
	return fields
}

// LogDebug 记录Debug信息
func LogDebug(str interface{}, data LogFields) {
	//TODO 如果是dev环境，则打印，如果是pro环境，则不打印
	updateLogFile(logrus.DebugLevel)
	debugLogS.WithFields(logrus.Fields(locate(data))).Debug(str)
}

// LogInfo 记录Info信息
func LogInfo(str interface{}, data LogFields) {
	//TODO 如果是dev环境，则打印，如果是pro环境，则不打印
	updateLogFile(logrus.InfoLevel)
	infoLogS.WithFields(logrus.Fields(data)).Info(str)
}

// LogError 记录Error信息
func LogError(str interface{}, data LogFields) {
	updateLogFile(logrus.ErrorLevel)
	errorLogS.WithFields(logrus.Fields(data)).Error(str)
}

func getLogFullPath(l logrus.Level) string {
	return LogPath + "." + l.String() + ".log"
}
