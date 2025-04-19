package models

type Response struct {
	Valid bool `json:"valid"`
} // @name Response

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
} // @name Error

type ErrorResponse struct {
	Response
	Error Error `json:"error"`
} // @name ErrorResponse
