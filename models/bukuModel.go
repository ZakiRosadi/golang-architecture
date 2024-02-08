package models

type BookInput struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	AuthorKita string `json:"author_kita"`
}