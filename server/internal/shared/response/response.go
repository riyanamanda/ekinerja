package response

type PaginatedResponse[T any] struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	IsFirst     bool  `json:"is_first"`
	IsLast      bool  `json:"is_last"`
	Total       int64 `json:"total"`
	Content     T     `json:"content"`
}

func CreatePaginationResponse[T any](content T, currentPage, perPage int, total int64) PaginatedResponse[T] {
	isFirst := currentPage == 1
	lastPage := int((total + int64(perPage) - 1) / int64(perPage))
	isLast := currentPage >= lastPage
	return PaginatedResponse[T]{
		CurrentPage: currentPage,
		PerPage:     perPage,
		IsFirst:     isFirst,
		IsLast:      isLast,
		Total:       total,
		Content:     content,
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
