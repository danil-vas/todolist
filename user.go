package todolist

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" building:"required"`
	Username string `json:"username" building:"required"`
	Password string `json:"password" building:"required"`
}
