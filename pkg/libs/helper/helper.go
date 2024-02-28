package helper

import (
	"strings"
)

func ContainsInArr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsInArrNoCaseSensitive(s []string, e string) bool {
	for _, a := range s {
		if strings.ToLower(a) == strings.ToLower(e) {
			return true
		}
	}
	return false
}

func BetweenString(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

// BeforeString Get substring before a string.
func BeforeString(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

// AfterString Get substring after a string.
func AfterString(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:]
}

//func GetId(me string) (int, error) {
//	idStr := strings.Split(me, UtilsConstant.TOKEN_SEPARATOR)
//	id, err := strconv.Atoi(idStr[0])
//	if err != nil {
//		return 0, err
//	}
//
//	return id, nil
//}

//func GetIdV2(c *context.Context) (int, error) {
//	idStr := c.ResponseWriter.Header().Get("ID")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		return 0, err
//	}
//
//	return id, nil
//}
//
//func GetIdV3(session *Session.Session, c *context.Context) (int, error) {
//	idStr := c.ResponseWriter.Header().Get("ID")
//	phone := c.ResponseWriter.Header().Get("PHONE")
//	if phone != ""{
//		session.PersonalId = phone
//	}
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		return 0, err
//	}
//
//	return id, nil
//}
