package models

/********** MODELS **********/

type User struct {
	ID       int
	UserName string
	Email    string
	Password string
}
type Posts struct {
	ID       int
	Nom      string
	Contenue string
}

type Category struct {
	ID    int
	Title string
}

/********** TEMPLATES **********/

type IndexTemplate struct {
	categories []Category
}

type RegisterTemplate struct {
	Message string
}

type LoginTemplate struct {
	Message string
}

type NouveauTemplate struct {
	Message string
}
type DebatTemplate struct {
	Message string
}
type SujetTemplate struct {
	Message string
}
