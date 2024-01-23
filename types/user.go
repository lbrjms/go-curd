package types

import "time"

type User struct {
	Id        string    `json:"id" xorm:"pk"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CellPhone string    `json:"cellPhone,omitempty"`
	Admin     bool      `json:"admin,omitempty"`
	Disable   bool      `json:"disable,omitempty"`
	Create    time.Time `json:"create,omitempty" xorm:"created"`
}

type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}
