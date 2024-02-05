package logger

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	Error "go-fiber-v2/pkg/libs/error"
	"go-fiber-v2/pkg/libs/helper/convert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
	"time"

	RotateLogs "github.com/lestrrat-go/file-rotatelogs"
)

type Logger struct {
	loggerSys   *zap.Logger
	publish     bool
	publishRest RestClient
	InstID      string
	Options     Options
	ThreadID    string
	//CentralLogIsEnable      bool
	//loggerTdr *zap.Logger
}
type Fields map[string]interface{}

type Interface interface {
	InfoSys(message string, fields ...zap.Field)
	InfoTdr(message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
}

func (l *Logger) Printf(s string, v ...interface{}) {
	if len(v) == 4 {
		l.loggerSys.Info("",
			zap.String("level", "INFO"),
			zap.String("request_id", l.ThreadID),
			zap.String("query", v[3].(string)),
			zap.String("duration ", fmt.Sprintf("%.3fms", v[1].(float64))),
			zap.Int64("affected-rows", v[2].(int64)),
			zap.String("source", v[0].(string)),
		)
	} else {
		l.loggerSys.Info("",
			zap.String("request_id", l.ThreadID),
			zap.Any("value", v),
		)
	}
}
func (l *Logger) Print(v ...interface{}) {
	if len(v) < 2 {
		return
	}
	switch v[0] {
	case "sql":
		delimiter := "/"
		rightOfDelimiter := strings.Join(strings.Split(v[1].(string), delimiter)[4:], delimiter)
		l.loggerSys.Info("",
			zap.String("level", "INFO"),
			zap.String("request_id", l.ThreadID),
			zap.String("query", v[3].(string)),
			zap.Any("values", v[4]),
			zap.Float64("duration", float64(v[2].(time.Duration))/float64(time.Millisecond)),
			zap.Int64("affected-rows", v[5].(int64)),
			zap.String("source", rightOfDelimiter),
		)
	default:
		delimiter := "/"
		rightOfDelimiter := strings.Join(strings.Split(v[1].(string), delimiter)[4:], delimiter)
		l.loggerSys.Info("",
			zap.String("request_id", l.ThreadID),
			zap.Any("values", v[2:]),
			zap.String("source", rightOfDelimiter),
		)
	}
}

func getRotateWriter(config Options) zapcore.WriteSyncer {
	rotate, err := RotateLogs.New(
		config.FileLocation+"%Y-%m-%d."+config.FileName,
		RotateLogs.WithMaxAge(config.FileMaxAge*24*time.Hour),
		RotateLogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(rotate)
}

func New(config Options) *Logger {
	var cores []zapcore.Core
	var writer zapcore.WriteSyncer

	if config.Stdout {
		writer = zapcore.AddSync(os.Stdout)
	} else {
		writer = getRotateWriter(config)
	}

	core := zapcore.NewCore(getEncoder(), writer, zapcore.InfoLevel)
	cores = append(cores, core)

	combinedCore := zapcore.NewTee(cores...)

	loggerSys := zap.New(combinedCore,
		zap.AddCallerSkip(3),
		zap.AddCaller(),
	)

	l := &Logger{
		loggerSys: loggerSys,
	}
	if config.PublishLog {
		l.publishRest = NewPublish(config.PublishOption)
		l.publish = true
	}
	l.InstID = config.PublishOption.InstId
	l.Options = config
	return l
}

type LogTdrModel struct {
	AppName        string      `json:"app"`
	AppVersion     string      `json:"ver"`
	IP             string      `json:"ip"`
	Port           int         `json:"port"`
	SrcIP          string      `json:"srcIP"`
	RespTime       int64       `json:"rt"`
	Path           string      `json:"path"`
	Header         interface{} `json:"header"`
	Request        interface{} `json:"req"`
	Response       interface{} `json:"resp"`
	Error          string      `json:"error"`
	ThreadID       string      `json:"threadID"`
	AdditionalData interface{} `json:"addData"`
}

func getEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		MessageKey:     "message",
		EncodeDuration: MillisDurationEncoder,
		EncodeTime:     TDRLogTimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}
	return zapcore.NewConsoleEncoder(config)
}

func TDRLogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	enc.AppendString(t.In(location).Format("2006-01-02 15:04:05.999"))
}

func MillisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Milliseconds())
}

func (l *Logger) Error(message string, fields ...zap.Field) {
	_, fn, line, _ := runtime.Caller(1)
	file := fmt.Sprintf("%s:%d - ", fn, line)
	l.loggerSys.Error(file+message, fields...)
}

func (l *Logger) InfoSys(message string, fields ...zap.Field) {
	l.loggerSys.Info(message, fields...)
}

func (l *Logger) PublishLog(traceId string, request interface{}) (response interface{}, err error) {
	if !l.publish {
		return
	}
	result, httpStatus, err := l.publishRest.Execute(traceId, l, url, request)
	if err != nil {
		return
	}
	if httpStatus != 200 {
		err = Error.New(httpStatus, "FAILED", "Error publish")
		return
	}

	var resp interface{}
	json.Unmarshal(result, &resp)
	response = resp
	return
}

func (l *Logger) MaskingJson(data interface{}) interface{} {
	jsonString := convert.ObjectToString(data)
	path := strings.Split(l.Options.MaskingLogJsonPath, "|")

	for _, key := range path {
		value := gjson.Get(jsonString, key)
		if value.String() != "" && value.Type.String() != "Null" {
			switch value.Type.String() {
			case "String":
				jsonString, _ = sjson.Set(jsonString, key, "******")
			case "JSON":
				jsonString, _ = sjson.Set(jsonString, key, "***Mask JSON***")
			case "False":
			case "True":
			default:
				jsonString, _ = sjson.Set(jsonString, key, 00000)
			}
		}
	}

	var mapData interface{}
	convert.StringToObject(jsonString, &mapData)
	return mapData
}

func (l *Logger) MaskingJsonWithPath(data interface{}, jsonPath string) interface{} {
	jsonString := convert.ObjectToString(data)
	path := strings.Split(jsonPath, "|")

	for _, key := range path {
		value := gjson.Get(jsonString, key)
		if value.String() != "" && value.Type.String() != "Null" {
			switch value.Type.String() {
			case "String":
				jsonString, _ = sjson.Set(jsonString, key, "******")
			case "JSON":
				jsonString, _ = sjson.Set(jsonString, key, "***Mask JSON***")
			case "False":
			case "True":
			default:
				jsonString, _ = sjson.Set(jsonString, key, 00000)
			}
		}
	}

	var mapData interface{}
	convert.StringToObject(jsonString, &mapData)
	return mapData
}

//func (l *Logger) InfoTdr(message string, fields ...zap.Field) {
//	l.loggerTdr.Info(message, fields...)
//}

//func (l *Logger) MaskingData(jsonByte []byte) (jsonString string) {
//	b := new(bytes.Buffer)
//	json.Compact(b, jsonByte)
//	body := fmt.Sprintf(`%v`, b)
//	path := strings.Split(Config.Config.Logger.MaskingLogJsonPath,"|")
//	//[]string{"idCard", "selfie", "data.id"}
//	var replaceString []string
//	for _, p := range path {
//		value := gjson.Get(fmt.Sprintf(`%v`, b), p)
//		if value.String() != "" && value.String() != "null" {
//			replaceString = append(replaceString, value.String())
//		}
//	}
//
//	for _, rc := range replaceString {
//		body = strings.Replace(body, rc, "********", -1)
//	}
//
//	jsonString = body
//	return
//}
