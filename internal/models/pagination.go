package models

// PaginationParams представляет параметры пагинации
type PaginationParams struct {
	Page    int `json:"page"`     // Номер страницы (начиная с 1)
	PerPage int `json:"per_page"` // Колличество записей на странице
}

// PaginatedResponse представляет ответ с пагинацией
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int         `json:"total"`       // Общее количество записей
	Page       int         `json:"page"`        // Текущая страница
	PerPage    int         `json:"per_page"`    // Колличество записей на странице
	TotalPages int         `json:"total_pages"` // Общее количество страниц
}
