package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/umutozd/restaurant-backend/types"
)

type UpdateCategoryField int32

const (
	CAGETORY_NAME UpdateCategoryField = 0
)

type UpdateCategoryReq struct {
	Fields   []UpdateCategoryField `json:"fields,omitempty"`
	Category *types.Category       `json:"category,omitempty"`
}

func (req *UpdateCategoryReq) UnmarshalBody(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return types.Err(types.ERR_BODY_UNREADABLE, err)
	}

	if err = json.Unmarshal(b, req); err != nil {
		return types.Err(types.ERR_UNMARSHAL, err)
	}

	return nil
}
