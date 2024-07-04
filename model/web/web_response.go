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
	Errors string `json:"errors" example:"not found"`
}

type WebForbiddenError struct {
	Code   int    `json:"code" example:"403"`
	Errors string `json:"errors" example:"forbidden"`
}

type WebUnauthorizedError struct {
	Code   int    `json:"code" example:"401"`
	Errors string `json:"errors" example:"unauthorized"`
}

type WebBadRequestError struct {
	Code   int    `json:"code" example:"400"`
	Errors string `json:"errors" example:"bad request"`
}
