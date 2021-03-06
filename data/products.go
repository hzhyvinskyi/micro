package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID					int			`json:"id"`
	Name				string	`json:"name"`
	Description	string	`json:"description"`
	Price				float32 `json:"price"`
	SKU					string	`json:"sku"`
	CreatedOn		string	`json:"created_on"`
	UpdatedOn		string	`json:"updated_on"`
	DeletedOn		string	`json:"-"`
}

type Products []*Product

var productList = Products{
	{
		ID:						1,
		Name:				 	"Test Name #1",
		Description:	"Test Description #1",
		Price: 				1.45,
		SKU:					"14308610",
		CreatedOn: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
	{
		ID:						2,
		Name:				 	"Test Name #2",
		Description:	"Test Description #2",
		Price: 				1.00,
		SKU:					"143325958",
		CreatedOn: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
