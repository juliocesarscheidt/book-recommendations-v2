package httpmodule

// HttpResponse - simple HttpResponse structure
type HttpResponse struct {
	Data interface{} `json:"data"`
}

type HttpResponseError struct {
	Message interface{} `json:"message"`
}
