package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/umutozd/restaurant-backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// First, items are added, then removed.
type UpdateCartReq struct {
	// ID of the Cart to update
	ID string `json:"id,omitempty"`
	// MenuItem ID's to be added
	Add []string `json:"add,omitempty"`
	// CartItem ID's to be removed
	Remove []string `json:"remove,omitempty"`
}

func (req *UpdateCartReq) UnmarshalBody(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return types.Err(types.ERR_BODY_UNREADABLE, err)
	}

	if err = json.Unmarshal(b, req); err != nil {
		return types.Err(types.ERR_UNMARSHAL, err)
	}

	return nil
}

// GetCartItems returns the MenuItemIDs in req with a status = ORDERED
// attached to it.
func (req *UpdateCartReq) GetCartItems() ([]*types.CartItem, []string) {
	add := make([]*types.CartItem, 0, len(req.Add))
	for _, id := range req.Add {
		add = append(add, &types.CartItem{
			ID:         primitive.NewObjectID().Hex(),
			MenuItemID: id,
			Status:     types.CART_ITEM_ORDERED,
		})
	}

	if req.Remove == nil {
		req.Remove = []string{}
	}

	return add, req.Remove
}
