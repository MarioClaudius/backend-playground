package response

type APIResponse struct {
	Data   interface{} `json:"data"`
	Errors string      `json:"errors"`
}
