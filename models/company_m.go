package models

//Buat validator/verifikator object
type Contact struct {
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
type CompanyObjt struct {
	Id      string   `json:"_id"`
	Rev     string   `json:"_rev"`
	Level   string   `json:"level"`
	Contact Contact  `json:"contact"`
	Users   []string `json:"users"`
}
type UserObjt struct {
	Id        string `json:"_id"`
	Rev       string `json:"_rev"`
	Table     string `json:"table"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Access1   string `json:"access1"`
	Access2   string `json:"access2"`
	Idcompany string `json:"idcompany"`
}
type CredDBObjt struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Type     string `json:"type"`
}
