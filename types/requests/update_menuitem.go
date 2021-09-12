package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/umutozd/restaurant-backend/types"
)

type UpdateMenuItemField int32

const (
	CATEGORY_ID UpdateMenuItemField = 0
	ITEM_NAME   UpdateMenuItemField = 1
	PRICE       UpdateMenuItemField = 2
	IMAGE       UpdateMenuItemField = 3
)

type UpdateMenuItemReq struct {
	Fields []UpdateMenuItemField `json:"fields,omitempty"`
	Item   *types.MenuItem       `json:"item,omitempty"`
}

func (req *UpdateMenuItemReq) UnmarshalBody(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return types.Err(types.ERR_BODY_UNREADABLE, err)
	}

	if err = json.Unmarshal(b, req); err != nil {
		return types.Err(types.ERR_UNMARSHAL, err)
	}

	return nil
}
