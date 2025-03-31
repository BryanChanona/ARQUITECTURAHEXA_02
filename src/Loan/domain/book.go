package domain


type Book struct {
	Id int `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Available  bool `json:"available"`
}