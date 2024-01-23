package types

import "github.com/lbrjms/go-curd/v2/db"

func init() {
	db.Register(new(User), new(Password))
}
