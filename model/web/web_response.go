package web

type WebSuccess[T any] struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
	Data    T      `json:"data"`
}

type WebError struct {
	Code   int `json:"code" example:"400"`
	Errors any `json:"errors"`
}
