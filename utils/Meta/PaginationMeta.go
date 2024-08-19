package Meta

type PaginationMeta struct {
	TotalRecords   int `json:"total_records"`
	TotalPages     int `json:"total_pages"`
	CurrentPage    int `json:"current_page"`
	RecordsPerPage int `json:"records_per_page"`
}
