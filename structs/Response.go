package structs

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (response *Response) Success(code int, message interface{}, data interface{}) {
	response.Code = code
	response.Status = "success"
	response.Message = message.(string)
	response.Data = data
}

func (response *Response) Error(code int, message interface{}, data interface{}) {
	response.Code = code
	response.Status = "error"
	response.Message = message.(string)
	response.Data = data
}
