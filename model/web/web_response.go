package web

type WebSuccess[T any] struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
	Data    T      `json:"data"`
}

type WebError struct {
	Code   int `json:"code"`
	Errors any `json:"errors"`
}

// for swagger documentation
type WebNotFoundError struct {
	Code   int    `json:"code" example:"404"`
	Errors string `json:"errors" example:"Not Found"`
}

type WebForbiddenError struct {
	Code   int    `json:"code" example:"403"`
	Errors string `json:"errors" example:"Forbidden"`
}

type WebUnauthorizedError struct {
	Code   int    `json:"code" example:"401"`
	Errors string `json:"errors" example:"Unauthorized"`
}

type WebBadRequestError struct {
	Code   int    `json:"code" example:"400"`
	Errors string `json:"errors" example:"Bad Request"`
}

type WebInternalServerError struct {
	Code   int    `json:"code" example:"500"`
	Errors string `json:"errors" example:"Internal Server Error"`
}
