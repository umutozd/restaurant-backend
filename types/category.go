package types

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Category is the type that stores menu item ids.
type Category struct {
	ID    string   `json:"id,omitempty" bson:"_id"`
	Name  string   `json:"name,omitempty" bson:"name"`
	Items []string `json:"items,omitempty" bson:"items"`
}

func (c *Category) UnmarshalBody(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return Err(ERR_BODY_UNREADABLE, err)
	}

	if err = json.Unmarshal(b, c); err != nil {
		return Err(ERR_UNMARSHAL, err)
	}

	return nil
}
