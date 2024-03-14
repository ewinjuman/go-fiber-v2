package example

import (
	"encoding/json"
	Error "gitlab.pede.id/otto-library/golang/share-pkg/error"
	Rest "gitlab.pede.id/otto-library/golang/share-pkg/http"
	Session "gitlab.pede.id/otto-library/golang/share-pkg/session"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/pkg/repository"
	"net/http"
)

type (
	UserHttpService interface {
		TokenSessionValidation(request ValidateSessionRequest) (response ValidateSessionResponse, err error)
	}

	userHttp struct {
		session         *Session.Session
		ottoUsersRest   Rest.RestClient
		ottoUsersConfig configs.Ottouser
	}
)

// New Create new http request to Users first
func NewUserHttp(session *Session.Session) UserHttpService {
	return &userHttp{
		session:         session,
		ottoUsersRest:   Rest.New(configs.Config.Ottouser.Option),
		ottoUsersConfig: configs.Config.Ottouser,
	}
}

// TokenSessionValidation Execute Request
func (o *userHttp) TokenSessionValidation(request ValidateSessionRequest) (response ValidateSessionResponse, err error) {

	//How to create header:
	//exampleHeaders := http.Header{} //common way
	//or
	//exampleHeaders := o.ottoUsersRest.DefaultHeader("userORid", "password") //create header include basic auth

	//How to create query param:
	//exampleQueryParam := map[string]string{ //using map
	//	"customerId":         customerId,
	//	"insuranceProductId": strconv.Itoa(insuranceProductId),
	//}

	result, httpStatus, err := o.ottoUsersRest.Execute(o.session, o.ottoUsersConfig.Host, o.ottoUsersConfig.Path.TokenValidation, http.MethodPost, nil, request, nil, nil)
	//do something if err is not nil
	if err != nil {
		if Error.IsTimeout(err) {
			//.... do something if needed
		}
		return
	}
	//do something if http code is not 200
	if httpStatus != 200 {
		err = Error.New(httpStatus, repository.FailedStatus, "Error validation")
		return
	}

	//Bind Response Body
	json.Unmarshal(result, &response)

	//Do Something with response body if needed
	if response.Code != 200 {
		err = Error.New(response.Code, response.Status, response.Message)
		//if response.Code == 400 {
		//	err = Error.New(fiber.StatusNotFound, response.Status, "Not Found")
		//}
		return
	}
	return
}
