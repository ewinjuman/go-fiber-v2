package example

import "go-fiber-v2/pkg/repository"

func GetCode(rc string) int {
	if code, ok := rcMap[rc]; ok {
		return code
	}
	return repository.UndefinedCode
}

var rcMap = map[string]int{
	"00": repository.SuccessCode,
	"01": repository.BadRequestCode,
	"06": repository.PendingCode,
}
