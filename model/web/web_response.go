package web

type WebSuccess[T any] struct {
	Code     int       `json:"code" example:"200" extensions:"x-order=0"`
	Message  string    `json:"message" example:"success" extensions:"x-order=1"`
	Payload  T         `json:"payload" extensions:"x-order=2"`
	Metadata *Metadata `json:"metadata" extensions:"x-order=3"`
}

type Metadata struct {
	Page       *int   `json:"page" form:"limit" extensions:"x-order=0"`
	Limit      *int   `json:"limit" form:"page" extensions:"x-order=1"`
	TotalPages *int   `json:"total_pages" extensions:"x-order=2"`
	TotalData  *int64 `json:"total_data" extensions:"x-order=3"`
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
