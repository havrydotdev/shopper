package shopper

import "database/sql"

type User struct {
	Id            int           `json:"id"`
	Username      string        `json:"username"`
	Email         string        `json:"email"`
	Password      string        `json:"password"`
	Balance       float32       `json:"balance"`
	IsTempBlocked bool          `json:"is_blocked" db:"istempblocked"`
	CompanyId     sql.NullInt32 `json:"company_id" db:"company_id"`
}

type UpdateUserBalance struct {
	Value int `json:"value"`
}
