package freshchat

type ResponseCode int

const (
	Success             ResponseCode = 200
	Accepted            ResponseCode = 202
	BadRequest          ResponseCode = 400
	Unauthenticated     ResponseCode = 401
	Forbidden           ResponseCode = 403
	NotFound            ResponseCode = 404
	TooManyRequests     ResponseCode = 429
	InternalServerError ResponseCode = 500
	ServiceUnavailable  ResponseCode = 503
)