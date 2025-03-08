package domain


type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Available  bool `json:"available"`
}