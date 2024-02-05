package utils

import (
	"encoding/json"
	Error "go-fiber-v2/pkg/libs/error"
	"net/http"
	"strings"
)

func ObjectToObject(in interface{}, out interface{}) {
	dataByte, _ := json.Marshal(in)
	json.Unmarshal(dataByte, &out)
}

func ObjectToString(data interface{}) string {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(dataByte)
}

func StringToObject(in string, out interface{}) {
	json.Unmarshal([]byte(in), &out)
	return
}

func ConvertPhoneNumber(mobilePhoneNumber string) (newMobilePhoneNumber string, err error) {
	phoneNumber := strings.Replace(mobilePhoneNumber, " ", "", -1)
	if strings.HasPrefix(phoneNumber, "62") {
		newMobilePhoneNumber = strings.Replace(phoneNumber, "62", "0", 1)
	} else if strings.HasPrefix(phoneNumber, "+62") {
		println(phoneNumber)
		newMobilePhoneNumber = strings.Replace(phoneNumber, "+62", "0", 1)
	} else if strings.HasPrefix(phoneNumber, "0") {
		newMobilePhoneNumber = phoneNumber
	} else {
		newMobilePhoneNumber = "0" + phoneNumber
	}
	valid := NewValidator()
	if err = valid.Var(newMobilePhoneNumber, "numeric"); err != nil {
		err = Error.New(http.StatusBadRequest, "FAILED", "Mobile Phone Number Tidak Valid")
		return
	}

	return
}
