package export

import "github.com/alecthomas/log4go"

const (
	//LogError 错误tag
	LogError = "error"
	LogDebug = "debug"
	LogInfo  = "info"
	LogTrace = "trace"
)

//Log 日志写入
type Log struct {
}

var log *Log

//NewLog 初始化日志类
func NewLog() *Log {
	if nil == log {
		log = &Log{}
		log.Init()
	}

	return log
}

//Init 初始化日志写入工具
func (log *Log) Init() *Log {

	//输出到控制台,级别为DEBUG, 系统默认配置
	// log4go.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())

	//输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
	log4go.AddFilter("file", log4go.DEBUG|log4go.INFO|log4go.WARNING|log4go.TRACE, log4go.NewFileLogWriter("test.log", false))
	log4go.AddFilter("file", log4go.ERROR, log4go.NewFileLogWriter("error.log", false))
	return log
}

//Close 关闭日志写入
func (log *Log) Close() {

	log4go.Close()
}

//Write
func (log *Log) Write(cate string, s ...string) bool {
	switch cate {
	case "debug":
		log4go.Debug(s)
	case "info":
		log4go.Info(s)
	case "error":
		log4go.Error(s)
	case "warn":
		log4go.Warn(s)
	case "trace":
		log4go.Trace(s)
		break
	default:
		panic("the key names cat(" + cate + ") what you set is not supported!")
	}

	return true
}
