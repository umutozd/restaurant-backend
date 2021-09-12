package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/umutozd/restaurant-backend/types"
)

type DeleteMenuItemReq struct {
	ID string `json:"id,omitempty"`
}

func (req *DeleteMenuItemReq) UnmarshalBody(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return types.Err(types.ERR_BODY_UNREADABLE, err)
	}

	if err = json.Unmarshal(b, req); err != nil {
		return types.Err(types.ERR_UNMARSHAL, err)
	}

	return nil
}
