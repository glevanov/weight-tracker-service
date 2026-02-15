package handlers

type SuccessResult struct {
	IsSuccess bool `json:"isSuccess"`
	Data      any  `json:"data,omitempty"`
}

type ErrorResult struct {
	IsSuccess bool   `json:"isSuccess"`
	Error     string `json:"error,omitempty"`
}
