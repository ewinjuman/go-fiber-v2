package example

type ValidateSessionRequest struct {
	Token string `json:"token"`
}

type ValidateSessionResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		OldID             int    `json:"oldId"`
		MobilePhoneNumber string `json:"mobilePhoneNumber"`
	} `json:"data"`
}
