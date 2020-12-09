package Models

//Article => returns article struct
type Article struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"description"`
}

//TableName => returns tablename
func (aa *Article) TableName() string {
	return "article"
}
