package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/ystar2020/lib1"
	"os"
)

//升级日志专用类 自动远程发送

//全局调用变量
var(
 Logger_global zap.Logger
 Path_value string
)

/*
    //"go.uber.org/zap"
	zaplog.Logger_global.Info("MSG",zap.String("Decrypt_msg",err.Error()))

    //多参数
   logger.Info("网址",
   zap.String("url", "http://www.baidu.com"),
   zap.Int("attempt", 3),
   zap.Duration("backoff", time.Second))

 */

//专用日志初始化类
//Log_init("eth","eth")
func Log_init(Service string,Value string) {


	var hook lumberjack.Logger


	//Path_value,_=lib1.Create_path_os()
	//自动新建文件夹 获取当前目录
	lib1.Create_New_File("log")


	Path_value=lib1.Create_Format_time("flie_time")[0:13]

	hook = lumberjack.Logger{
		Filename:   "./log/" + Path_value + ".log", // 日志文件路径
		MaxSize:    10,                             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,                              // 日志文件最多保存多少个备份
		MaxAge:     30,                             // 文件最多保存多少天
		Compress:   true,                           // 是否压缩
	}


	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String(Service, Value))
	// 构造日志
	logger := zap.New(core, caller, development, filed)

	//var list []string

	//list=append(list,"123")
	//list=append(list,"123")

	logger.Info("log 初始化成功")

	Logger_global=*logger


}
