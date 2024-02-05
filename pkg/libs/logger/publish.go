package logger

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

var url string

type PublishRequest struct {
	InstitutionID string `json:"institutionId"`
	ServiceName   string `json:"serviceName"`
	LogType       string `json:"logType"`
	Data          struct {
		PublishTime    time.Time   `json:"publishTime"`
		LogLevel       string      `json:"logLevel"`
		TraceID        string      `json:"traceId"`
		ActionTo       string      `json:"actionTo"`
		ActionName     string      `json:"actionName"`
		EndPoint       string      `json:"endPoint"`
		ErrorDesc      interface{} `json:"errorDesc"`
		FileName       string      `json:"fileName"`
		RequestBody    interface{} `json:"requestBody"`
		RequestHeader  interface{} `json:"requestHeader"`
		ResponseBody   interface{} `json:"responseBody"`
		ResponseHeader interface{} `json:"responseHeader"`
	} `json:"data"`
}

func (log *PublishRequest) SetError() *PublishRequest {
	log.Data.LogLevel = "Error"
	return log
}

func (log *PublishRequest) SetInfo() *PublishRequest {
	log.Data.LogLevel = "Info"
	return log
}

func (log *PublishRequest) SetRequestBody(requestBody interface{}) *PublishRequest {
	b, _ := json.Marshal(requestBody)
	log.Data.RequestBody = string(b)
	return log
}

func (log *PublishRequest) SetResponseBody(responseBody interface{}) *PublishRequest {
	b, _ := json.Marshal(responseBody)
	log.Data.ResponseBody = string(b)
	return log
}

func (log *PublishRequest) SetRequestHeader(requestHeader interface{}) *PublishRequest {
	log.Data.RequestHeader = fmt.Sprintf("%v", requestHeader)
	return log
}

func (log *PublishRequest) SetResponseHeader(responseHeader interface{}) *PublishRequest {
	log.Data.ResponseHeader = fmt.Sprintf("%v", responseHeader)
	return log
}

func (log *PublishRequest) SetErrorDesc(errorDesc interface{}) *PublishRequest {
	log.Data.ErrorDesc = fmt.Sprintf("%v", errorDesc)
	return log
}

func (log *PublishRequest) SetFileName(fileName string) *PublishRequest {
	log.Data.FileName = fileName
	return log
}

func (log *PublishRequest) Request() *PublishRequest {
	log.Data.ActionName = log.Data.ActionName + " Request"
	return log
}

func (log *PublishRequest) Response() *PublishRequest {
	log.Data.ActionName = log.Data.ActionName + " Response"
	return log
}

type RestClient interface {
	Execute(traceId string, log *Logger, url string, payload interface{}) (body []byte, statusCode int, err error)
}

func NewPublish(options PublishOption) RestClient {
	httpClient := resty.New()
	url = options.PublishLogTo
	if options.SkipTLS {
		httpClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	httpClient.SetTimeout(options.Timeout * time.Second)
	httpClient.SetDebug(options.DebugMode)

	return &client{
		options:    options,
		httpClient: httpClient,
	}
}

type client struct {
	options    PublishOption
	httpClient *resty.Client
}

func (c *client) Execute(traceId string, log *Logger, url string, payload interface{}) (body []byte, statusCode int, err error) {
	//url := host + path
	request := c.httpClient.R()
	request.Header.Set("Content-Type", "application/json")
	request.SetBody(payload)
	iTraceId := strconv.Itoa(int(time.Now().UnixNano() / int64(time.Millisecond)))
	log.InfoSys("",
		zap.String("trace_id", traceId+"-"+iTraceId),
		zap.String("action", "Baikal Log Request"),
		//zap.Any("request", request.Body),
	)
	var result *resty.Response
	result, err = request.Post(url)

	if err != nil {
		log.Error(err.Error())
		return
	}

	if result != nil {
		body = result.Body()
	}

	if result != nil && result.StatusCode() != 0 {
		statusCode = result.StatusCode()
	}
	var resultRes HaikalResponse
	var r interface{}
	errr := json.Unmarshal(body, &resultRes)
	if errr != nil {
		r = string(body)
	} else {
		r = resultRes
	}

	log.InfoSys("",
		zap.String("trace_id", traceId+"-"+iTraceId),
		zap.String("action", "Baikal Log Response"),
		zap.Int("http_status", statusCode),
		zap.Any("response", r))
	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, err
}

type HaikalResponse struct {
	SuccessMessage string `json:"successMessage"`
	Code           int    `json:"code"`
	ErrorMessage   string `json:"errorMessage"`
}
