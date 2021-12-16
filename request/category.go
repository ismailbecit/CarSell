package request

type CategoryInsert struct {
	Name  string `json:"name"`
	Year  uint   `json:"year"`
	Price uint   `json:"price"`
}

type CategoryDel struct {
	ID uint `json:"id"`
}
