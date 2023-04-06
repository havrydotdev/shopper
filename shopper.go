package shopper

import (
	"errors"
	"time"
)

type Item struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Amount      int     `json:"amount"`
	Keywords    string  `json:"keywords"`
	Rating      float32 `json:"rating"`
	CompanyId   int     `json:"company_id"`
	IsVerified  bool    `json:"isVerified"`
}

type Company struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
		return errors.New("update struct is empty")
	}

	return nil
}

type Property struct {
	Id    int    `json:"id"`
	Name  string `json:"key"`
	Value string `json:"value"`
}

type Comment struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	ItemId int    `json:"item_id"`
}

type Discount struct {
	Id       int       `json:"id"`
	Percent  int       `json:"percent"`
	Relevant time.Time `json:"relevant"`
}

type Notification struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Text      string    `json:"text"`
	UserId    int       `json:"user_id"`
}
