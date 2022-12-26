package templates

type MessageResponse struct {
	Message string `json:"message"`
}

func NewResp(msg string) MessageResponse {
	r := MessageResponse{}
	r.Message = msg
	return r
}
