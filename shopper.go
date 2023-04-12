package shopper

import (
	"database/sql"
	"errors"
	"time"
)

var (
	updErr = errors.New("update struct is empty")
)

type Item struct {
	Id                int     `json:"id"`
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description"`
	Price             float32 `json:"price" binding:"required"`
	Amount            int     `json:"amount" binding:"required"`
	Keywords          string  `json:"keywords"`
	Rating            float32 `json:"rating"`
	CompanyId         int     `json:"company_id" db:"company_id"`
	IsVerified        bool    `json:"isVerified" db:"isverified"`
	PriceWithDiscount float32 `json:"priceWithDiscount" db:"price_with_discount"`
}

type Company struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Logo        string `json:"logo"`
	IsVerified  bool   `json:"isVerified"`
}

type UpdateCompanyInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Logo        *string `json:"logo"`
}

func (u UpdateCompanyInput) Validate() error {
	if u.Name == nil && u.Logo == nil && u.Description == nil {
		return updErr
	}

	return nil
}

type UpdateUserInput struct {
	Username      *string `json:"username"`
	Email         *string `json:"email"`
	Password      *string `json:"password"`
	IsTempBlocked *bool   `json:"isTempBlocked"`
}

func (u UpdateUserInput) Validate() error {
	if u.Username == nil && u.Email == nil && u.IsTempBlocked == nil && u.Password == nil {
		return updErr
	}

	return nil
}

type UpdateItemInput struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float32 `json:"price"`
	Amount      *int     `json:"amount"`
	Keywords    *string  `json:"keywords"`
}

func (u *UpdateItemInput) Validate() error {
	if u.Name == nil && u.Amount == nil && u.Keywords == nil && u.Price == nil && u.Description == nil {
		return updErr
	}

	return nil
}

type UpdateCommentInput struct {
	Text *string `json:"text" example:"updated text :)"`
}

func (u *UpdateCommentInput) Validate() error {
	if u.Text == nil {
		return updErr
	}

	return nil
}

type Property struct {
	Id    int    `json:"id"`
	Name  string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type Comment struct {
	Id     int    `json:"id" example:"1"`
	Text   string `json:"text" example:"Such an amazing item!"`
	ItemId int    `json:"item_id" db:"item_id" example:"1"`
	UserId int    `json:"user_id" db:"user_id" example:"1"`
}

type Discount struct {
	Id       int       `json:"id"`
	Percent  int       `json:"percent" binding:"required"`
	Relevant time.Time `json:"relevant" binding:"required"`
}

type SignUpInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=26"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Notification struct {
	Id        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	Text      string    `json:"text" binding:"required"`
	UserId    int       `json:"user_id" db:"user_id"`
}

type Rate struct {
	Value float32 `json:"value"`
}

type User struct {
	Id            int           `json:"id"`
	Username      string        `json:"username" binding:"required"`
	Email         string        `json:"email" binding:"required,email"`
	Password      string        `json:"password" binding:"required,max=26"`
	Balance       float32       `json:"balance"`
	IsTempBlocked bool          `json:"is_blocked" db:"istempblocked"`
	CompanyId     sql.NullInt32 `json:"company_id" db:"company_id"`
}

type UpdateUserBalance struct {
	Value int `json:"value"`
}
