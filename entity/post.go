package entity

type Post struct {
	ID    int64  `jsonn:"id"`
	Title string `jsonn:"title"`
	Text  string `jsonn:"text"`
}
