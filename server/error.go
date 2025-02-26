package server

type Error struct {
	Code        string            `json:"code"`
	Description string            `json:"description"`
	Params      map[string]string `json:"params,omitempty"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}
