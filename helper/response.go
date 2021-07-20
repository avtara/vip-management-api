package helper

import "strings"

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"ok"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

//EmptyObj return empty json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
	return res
}

func BuildResponseWithoutData(status bool, message string) Response {
	res := Response{
		Status:  status,
		Message: message,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic err response
func BuildErrorResponse(err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Data:    data,
		Message: splittedError,
	}
	return res
}
