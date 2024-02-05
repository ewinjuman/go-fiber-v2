package convert

import (
	"bufio"
	"encoding/json"
	"os"
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

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//
//func ConvertPhoneNumber(mobilePhoneNumber string) (newMobilePhoneNumber string, err error) {
//	phoneNumber := strings.Replace(mobilePhoneNumber, " ", "", -1)
//	if strings.HasPrefix(phoneNumber, "62") {
//		newMobilePhoneNumber = strings.Replace(phoneNumber, "62", "0", 1)
//	} else if strings.HasPrefix(phoneNumber, "+62") {
//		println(phoneNumber)
//		newMobilePhoneNumber = strings.Replace(phoneNumber, "+62", "0", 1)
//	} else if strings.HasPrefix(phoneNumber, "0") {
//		newMobilePhoneNumber = phoneNumber
//	} else {
//		newMobilePhoneNumber = "0"+phoneNumber
//	}
//	valid := validation.Validation{}
//	if v := valid.Numeric(newMobilePhoneNumber, "Mobile Phone Number"); !v.Ok {
//		println(newMobilePhoneNumber)
//		println(v.Error.Message)
//		err = Error.New(http.StatusBadRequest, "FAILED", "Mobile Phone Number Tidak Valid")
//		return
//	}
//
//	return
//}
