package models

//Buat validator/verifikator object

type UserModel struct {
	Id       string      `json:"_id"`
	Rev      string      `json:"_rev"`
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     string      `json:"role" binding:"required"`
	Contact  ContactUser `json:"contact"`
	Isactive bool        `json:"isactive" binding:"required"`
	Table    string      `json:"table" binding:"required"`
}
type ContactUser struct {
	Email  string `json:"email"`
	Mobile string `json:"mobile" binding:"required"`
}
