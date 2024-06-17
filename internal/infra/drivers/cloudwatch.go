package drivers

import (
	"botwhatsapp/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"time"
)

type Cloudwatch struct{}

func NewCloudwatchDriver() *Cloudwatch {
	return &Cloudwatch{}
}

func (cw *Cloudwatch) CommonCloudwatchZapCore(streamName string, level zapcore.Level) zapcore.Core {
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == level
	})

	//awsKey := os.Getenv("AWS_ACCESS_KEY_ID")
	//awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	//awsCred := credentials.NewStaticCredentials(awsKey, awsSecretKey, "")
	//awsCfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(awsCred)
	awsCfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(nil)
	enc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	cloudWatchParams := util.NewCloudwatchCoreParams{
		GroupName:    "deepbot",
		StreamName:   streamName,
		IsAsync:      true,
		Config:       awsCfg,
		Level:        level,
		LevelEnabler: debugPriority,
		Enc:          enc,
		Out:          zapcore.AddSync(io.Discard),
	}
	core, err := util.NewCloudwatchCore(&cloudWatchParams)
	if err != nil {
		log.Printf("can't initialize cloudwatch logger: %v", err)
	}
	return core
}

func (cw *Cloudwatch) InitializeLoggers() *zap.Logger {
	cdw := NewCloudwatchDriver()

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	enc := zapcore.NewConsoleEncoder(cfg)
	zapCore := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	core := zapcore.NewTee(
		zapCore,
		cdw.CommonCloudwatchZapCore("usecases", zapcore.DebugLevel),
		cdw.CommonCloudwatchZapCore("usecases", zapcore.ErrorLevel),

		cdw.CommonCloudwatchZapCore("interface-http", zapcore.DebugLevel),
		cdw.CommonCloudwatchZapCore("interface-http", zapcore.ErrorLevel),
	)
	return zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

type LoggerCloudwatch struct {
	log *zap.Logger
}

func NewLoggerCloudwatch(l *zap.Logger) *LoggerCloudwatch {
	return &LoggerCloudwatch{log: l}
}

func (cw *LoggerCloudwatch) Debug(msg string, payload any, sttCode int) {
	zt := zap.Duration("time", time.Since(time.Now()))
	zs := zap.Int("status", sttCode)
	ztp := zap.Reflect("payload", payload)
	zip := zap.String("reqId", uuid.New().String())
	cw.log.Debug(msg, zt, zs, ztp, zip)
}

func (cw *LoggerCloudwatch) Info(msg string, payload any, sttCode int) {
	zt := zap.Duration("time", time.Since(time.Now()))
	zs := zap.Int("status", sttCode)
	ztp := zap.Reflect("payload", payload)
	zip := zap.String("reqId", uuid.New().String())
	cw.log.Info(msg, zt, zs, ztp, zip)
}

func (cw *LoggerCloudwatch) Warn(msg string, payload any, sttCode int) {
	zt := zap.Duration("time", time.Since(time.Now()))
	zs := zap.Int("status", sttCode)
	ztp := zap.Reflect("payload", payload)
	zip := zap.String("reqId", uuid.New().String())
	cw.log.Warn(msg, zt, zs, ztp, zip)
}

func (cw *LoggerCloudwatch) Error(msg string, payload any, sttCode int) {
	zt := zap.Duration("time", time.Since(time.Now()))
	zs := zap.Int("status", sttCode)
	ztp := zap.Reflect("payload", payload)
	zip := zap.String("reqId", uuid.New().String())
	cw.log.Error(msg, zt, zs, ztp, zip)
}
