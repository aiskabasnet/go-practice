package Models

//Article => returns article struct
type Article struct {
	ID    uint   `json:"id"`
	title string `json:"title"`
	desc  string `json:"description"`
}

//TableName => returns tablename
func (aa *Article) TableName() string {
	return "article"
}
