package responses

type MapResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONWebResponse(msg string, data interface{}) MapResponse {
	return MapResponse{
		Message: msg,
		Data:    data,
	}
}
