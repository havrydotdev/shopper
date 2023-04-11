package shopper

import (
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
	Text *string `json:"text"`
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
	Id     int    `json:"id"`
	Text   string `json:"text"`
	ItemId int    `json:"item_id"`
}

type Discount struct {
	Id       int       `json:"id"`
	Percent  int       `json:"percent" binding:"required"`
	Relevant time.Time `json:"relevant" binding:"required"`
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
