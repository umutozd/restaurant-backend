package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/umutozd/restaurant-backend/types"
)

type UpdateCartItemsStatusReq struct {
	CartID  string               `json:"cart_id,omitempty"`
	ItemIDs []string             `json:"item_ids,omitempty"`
	Status  types.CartItemStatus `json:"status,omitempty"`
}

func (req *UpdateCartItemsStatusReq) UnmarshalBody(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return types.Err(types.ERR_BODY_UNREADABLE, err)
	}

	if err = json.Unmarshal(b, req); err != nil {
		return types.Err(types.ERR_UNMARSHAL, err)
	}

	return nil
}
