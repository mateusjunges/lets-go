package http

type JsonResponse struct {
	HttpCode int
	Data     []byte
	Message  string
}
