package session

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	JsonIter "github.com/json-iterator/go"
	Map "github.com/orcaman/concurrent-map"
	"github.com/pkg/errors"
	Logger "go-fiber-v2/pkg/libs/logger"
	"go.uber.org/zap"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	AppSession = "AppSession"
	ThreadId   = "ThreadId"
)

type Session struct {
	InstitutionID           string
	Map                     Map.ConcurrentMap
	Logger                  *Logger.Logger
	RequestTime             time.Time
	ThreadID                string
	PersonalId              string
	AppName, AppVersion, IP string
	Port                    int
	SrcIP, URL, Method      string
	Header, Request         interface{}
	ErrorMessage            string
	ActionTo                string
	ActionName              string
}

func New(logger *Logger.Logger) *Session {
	//initValidationString()
	sessionID := strconv.Itoa(int(time.Now().UnixNano() / int64(time.Microsecond)))
	session := &Session{
		RequestTime: time.Now(),
		Logger:      logger,
		Map:         Map.New(),
	}
	session.ThreadID = sessionID
	session.Logger.ThreadID = sessionID
	return session
}

func (session *Session) SetThreadID(sessionID string) *Session {
	session.ThreadID = sessionID
	session.Logger.ThreadID = sessionID
	return session
}

func (session *Session) SetMethod(method string) *Session {
	session.Method = method
	return session
}

func (session *Session) SetAppName(appName string) *Session {
	session.AppName = appName
	return session
}

func (session *Session) SetAppVersion(appVersion string) *Session {
	session.AppVersion = appVersion
	return session
}

func (session *Session) SetURL(url string) *Session {
	session.URL = url
	return session
}

func (session *Session) SetIP(ip string) *Session {
	session.IP = ip
	return session
}

func (session *Session) SetPort(port int) *Session {
	session.Port = port
	return session
}

func (session *Session) SetSrcIP(srcIp string) *Session {
	session.SrcIP = srcIp
	return session
}

func (session *Session) SetHeader(header interface{}) *Session {
	session.Header = header
	return session
}

func (session *Session) SetRequest(request interface{}) *Session {
	session.Request = request
	return session
}

func (session *Session) SetErrorMessage(errorMessage string) *Session {
	session.ErrorMessage = errorMessage
	return session
}

func (session *Session) SetInstitutionID(institutionID string) *Session {
	session.InstitutionID = institutionID
	return session
}

func (session *Session) SetActionTo(actionTo string) *Session {
	session.ActionTo = actionTo
	return session
}

func (session *Session) SetActionName(actionName string) *Session {
	session.ActionName = actionName
	return session
}

func (session *Session) SetPersonalIdentifier(phone string) *Session {
	session.PersonalId = phone
	return session
}

func (session *Session) Get(key string) (data interface{}, err error) {
	data, ok := session.Map.Get(key)
	if !ok {
		err = errors.New("not found")
	}
	return
}

func (session *Session) Put(key string, data interface{}) {
	session.Map.Set(key, data)
}

func (session *Session) LogDatabase(sql string, result interface{}, error interface{}) {
	session.Logger.InfoSys("",
		zap.String("request_id", session.ThreadID),
		zap.String("sql", sql),
		zap.Any("result", result),
		zap.Any("error", error),
	)
}

func (session *Session) LogRequest(message ...interface{}) {
	//if session.ActionName != "" {
	//	req := session.NewPublishLog().Request().SetInfo().SetRequestBody(session.Request).SetRequestHeader(session.Header)
	//	go session.Logger.PublishLog(req)
	//}
	if session.Request != nil {
		session.Request = session.Logger.MaskingJson(session.Request)
	}
	session.Logger.InfoSys("",
		zap.String("level", "INFO"),
		zap.String("request_id", session.ThreadID),
		zap.String("method", session.Method),
		zap.String("url", session.URL),
		zap.Any("request", session.Request),
		zap.Any("header", session.Header),
		zap.String("message", formatResponse(message...)),
	)
}

func (session *Session) LogResponse(response interface{}, message ...interface{}) {
	if response != nil {
		response = session.Logger.MaskingJson(response)
	}
	stop := time.Now()
	rt := stop.Sub(session.RequestTime).Milliseconds()
	session.Logger.InfoSys("",
		zap.String("level", "INFO"),
		zap.String("request_id", session.ThreadID),
		zap.Any("personal_id", session.PersonalId),
		zap.String("method", session.Method),
		zap.String("url", session.URL),
		zap.Any("response", response),
		zap.String("response_time", fmt.Sprintf("%d ms", rt)),
		zap.String("message", formatResponse(message...)),
	)
}

func (session *Session) LogRequestHttp(url string, method string, body interface{}, header interface{}, params interface{}) {
	if body != nil {
		//b, _ := json.Marshal(body)
		body = session.Logger.MaskingJson(body)
	}
	session.Logger.InfoSys("",
		zap.String("level", "INFO"),
		zap.String("request_id", session.ThreadID),
		zap.String("personal_id", session.PersonalId),
		zap.String("method", method),
		zap.String("url", url),
		zap.Any("request", body),
		zap.Any("params", params),
		zap.Any("header", header),
	)
}

func (session *Session) LogResponseHttp(responseTime time.Duration, code int, url string, method string, body interface{}, messageError ...string) {
	if body != nil {
		//b, _ := json.Marshal(body)
		body = session.Logger.MaskingJson(body)
	}

	if len(messageError) > 0 {
		msgErr := ""
		msgErr = strings.Join(messageError, ",")
		session.Logger.InfoSys("",
			zap.String("level", "INFO"),
			zap.String("request_id", session.ThreadID),
			zap.String("personal_id", session.PersonalId),
			zap.String("method", method),
			zap.String("url", url),
			zap.Int("http_status", code),
			zap.String("error", msgErr),
			zap.Any("response", body),
			zap.String("process_time", fmt.Sprintf("%d ms", responseTime.Milliseconds())),
		)
	} else {
		session.Logger.InfoSys("",
			zap.String("level", "INFO"),
			zap.String("request_id", session.ThreadID),
			zap.String("personal_id", session.PersonalId),
			zap.String("method", method),
			zap.String("url", url),
			zap.Int("http_status", code),
			zap.Any("response", body),
			zap.String("process_time", fmt.Sprintf("%d ms", responseTime.Milliseconds())),
		)
	}
}

func (session *Session) LogRequestGrpc(url string, method string, body interface{}, header interface{}) {

	if body != nil {
		//b, _ := json.Marshal(body)
		body = session.Logger.MaskingJson(body)
	}
	session.Logger.InfoSys("",
		zap.String("level", "INFO"),
		zap.String("request_id", session.ThreadID),
		zap.String("personal_id", session.PersonalId),
		zap.String("method", method),
		zap.String("url", url),
		zap.Any("request", body),
		zap.Any("header", header),
	)
}

func (session *Session) LogResponseGrpc(startProcessTime time.Time, url string, method string, body interface{}) {
	stop := time.Now()
	if body != nil {
		//b, _ := json.Marshal(body)
		body = session.Logger.MaskingJson(body)
	}
	session.Logger.InfoSys("",
		zap.String("level", "INFO"),
		zap.String("request_id", session.ThreadID),
		zap.String("personal_id", session.PersonalId),
		zap.String("method", method),
		zap.String("url", url),
		zap.Any("response", body),
		zap.String("process_time", fmt.Sprintf("%d ms", stop.Sub(startProcessTime).Milliseconds())),
	)
}

func (session *Session) LogMessage(message interface{}, data interface{}) {
	session.Logger.InfoSys("",
		zap.String("request_id", session.ThreadID),
		zap.Any("message", message),
		zap.Any("data", data),
	)
}

func (session *Session) LogTdr(message interface{}) {
	stop := time.Now()
	rt := stop.Sub(session.RequestTime).Milliseconds()

	session.Logger.InfoSys("",
		zap.String("request_id", session.ThreadID),
		zap.String("method", session.Method),
		zap.String("uri", session.URL),
		zap.Any("response", message),
		zap.String("response_time", fmt.Sprintf("%d ms", rt)),
	)

	//session.Logger.InfoTdr("",
	//	zap.String("request_id", session.ThreadID),
	//	zap.String("method", session.Method),
	//	zap.String("uri", session.URL),
	//	zap.Any("request", session.Request),
	//	zap.Any("response", message),
	//	zap.String("response_time", fmt.Sprintf("%d ms", rt)),
	//	zap.Any("header_request", session.Header),
	//)
}

func (session *Session) Info(message ...interface{}) {
	stop := time.Now()
	rt := stop.Sub(session.RequestTime).Milliseconds()

	session.Logger.InfoSys("",
		zap.String("level", "INFO"),
		zap.String("request_id", session.ThreadID),
		zap.String("method", session.Method),
		zap.String("uri", session.URL),
		zap.Any("message", message),
		zap.String("response_time", fmt.Sprintf("%d ms", rt)),
	)

	//session.Logger.InfoTdr("",
	//	zap.String("level", "INFO"),
	//	zap.String("request_id", session.ThreadID),
	//	zap.String("method", session.Method),
	//	zap.String("uri", session.URL),
	//	zap.Any("request", session.Request),
	//	zap.Any("response", message),
	//	zap.String("response_time", fmt.Sprintf("%d ms", rt)),
	//	zap.Any("header_request", session.Header),
	//)
}

func (session *Session) Error(message interface{}) {
	stop := time.Now()
	rt := stop.Sub(session.RequestTime).Milliseconds()
	_, fn, line, _ := runtime.Caller(1)
	file := fmt.Sprintf("%s:%d", fn, line)

	session.Logger.InfoSys("",
		zap.String("level", "ERROR"),
		zap.String("request_id", session.ThreadID),
		zap.String("personal_id", session.PersonalId),
		zap.String("method", session.Method),
		zap.String("uri", session.URL),
		zap.Any("message", message),
		zap.String("file", file),
		zap.String("response_time", fmt.Sprintf("%d ms", rt)),
	)

	//session.Logger.InfoTdr("",
	//	zap.String("level", "ERROR"),
	//	zap.String("request_id", session.ThreadID),
	//	zap.String("method", session.Method),
	//	zap.String("uri", session.URL),
	//	zap.Any("request", session.Request),
	//	zap.Any("response", message),
	//	zap.String("response_time", fmt.Sprintf("%d ms", rt)),
	//	zap.Any("header_request", session.Header),
	//)
}

var json = JsonIter.ConfigCompatibleWithStandardLibrary

func formatResponse(message ...interface{}) string {
	sb := strings.Builder{}

	for _, msg := range message {
		var m []byte
		if reflect.ValueOf(msg).Kind().String() == "string" {
			m = []byte(msg.(string))
		} else {
			m, _ = json.Marshal(msg)
		}

		sb.Write(m)
	}

	return sb.String()
}

//
//func (session *Session) IsValid(model interface{}) error {
//
//	if isNext(model) {
//		valid := validation.Validation{}
//		b, _ := valid.Valid(model)
//		var notValid []string
//		if !b {
//			for _, err := range valid.Errors {
//				field, _ := reflect.TypeOf(model).Elem().FieldByName(err.Field)
//				name := field.Tag.Get("name")
//				if name == "" {
//					name = field.Tag.Get("json")
//				}
//				notValid = append(notValid, name+" "+err.Tmpl)
//			}
//			er := strings.Join(notValid, ", ")
//			return Error.New(http.StatusBadRequest, "FAILED", er)
//		}
//		return nil
//	} else {
//		session.Error("Cannot perform validation: The type of model does not support")
//		return nil
//	}
//
//}

func isNext(i interface{}) bool {
	return reflect.ValueOf(i).Type().Kind() != reflect.Struct
}

// GetSession For Fiber
func GetSession(c *fiber.Ctx) *Session {
	return c.Locals(AppSession).(*Session)
}

//For Beego
//func GetSession(c *context.Context) *Session {
//	return c.Input.GetData(AppSession).(*Session)
//}
//
//func (session *Session) BindRequest(c *context.Context, requestModel interface{}) error {
//	if err := json.Unmarshal(c.Input.RequestBody, &requestModel); err != nil {
//		session.Error(err.Error())
//		err = Error.New(http.StatusBadRequest, "FAILED", err.Error())
//		return err
//	}
//
//	if err := session.IsValid(requestModel); err != nil {
//		err = Error.New(http.StatusBadRequest, "FAILED", err.Error())
//		return err
//	}
//
//	return nil
//}
//
//func InitValidationString() {
//	validation.SetDefaultMessage(map[string]string{
//		"Required":     "Tidak Boleh Kosong",
//		"Min":          "Telalu Kecil",
//		"Max":          "Terlalu Besar",
//		"Range":        "Harus Antara %d Sampai Dengan %d",
//		"MinSize":      "Ukuran minimum %d",
//		"MaxSize":      "Ukuran maksimum %d",
//		"Length":       "Panjang Maksimum %d",
//		"Alpha":        "Harus terdiri dari huruf",
//		"Numeric":      "Harus terdiri dari angka",
//		"AlphaNumeric": "Harus terdiri dari huruf atau angka",
//		"Match":        "Tidak Valid",
//		"NoMatch":      "Tidak cocok dengan %s",
//		"AlphaDash":    "Harus terdiri dari huruf, angka atau simbol (-_)",
//		"Email":        "Format email salah",
//		"IP":           "IP Tidak valid",
//		"Base64":       "Harus dalam format base64 yang benar",
//		"Mobile":       "Harus nomor ponsel yang benar",
//		"Tel":          "Harus nomor telepon yang benar",
//		"Phone":        "Harus nomor telepon atau ponsel yang benar",
//		"ZipCode":      "Harus kode pos yang valid",
//	})
//}
