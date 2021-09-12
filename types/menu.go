package types

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Menu is the top level struct that contains every category and
// menu item within. It is not stored in the database, but rather
// created when sending the whole menu.
type Menu []*CategoryAndItems

type CategoryAndItems struct {
	Category *Category   `json:"category,omitempty"`
	Items    []*MenuItem `json:"items,omitempty"`
}

// MenuItem is the most basic type that contains information
// about a menu item. It must belong to a category.
type MenuItem struct {
	ID         string  `json:"id,omitempty" bson:"_id"`
	CategoryID string  `json:"category_id,omitempty" bson:"category_id"`
	Name       string  `json:"name,omitempty" bson:"name"`
	Price      float64 `json:"price,omitempty" bson:"price"`
	Image      string  `json:"image,omitempty" bson:"image"`
}

func (item *MenuItem) UnmarshalBody(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return Errf(ERR_BODY_UNREADABLE, "error reading request body: %v", err)
	}

	if err = json.Unmarshal(b, item); err != nil {
		return Errf(ERR_UNMARSHAL, "error unmarshaling request body: %v", err)
	}

	return nil
}
