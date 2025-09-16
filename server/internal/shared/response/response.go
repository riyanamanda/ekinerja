package response

func CreatePaginationResponse[T any](data T) map[string]T {
	return map[string]T{
		"content": data,
	}
}

func CreateSuccessResponse(message string) map[string]string {
	return map[string]string{
		"message": message,
	}
}

func CreateErrorResponse[T any](err T) map[string]T {
	return map[string]T{
		"errors": err,
	}
}
