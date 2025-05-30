package model

type Response struct {
	Code		int			`json:"code"`
	Message	string	`json:"message"`
	Data		any			`json:"data,omitempty"`
	Error		any			`json:"error,omitempty"`
}