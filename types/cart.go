package types

import "time"

type Cart struct {
	ID        string      `json:"id,omitempty" bson:"_id"`
	Items     []*CartItem `json:"items,omitempty" bson:"items"`
	CreatedAt time.Time   `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" bson:"updated_at"`
}

type CartItemStatus int32

const (
	// Customer ordered, registry not approved yet.
	CART_ITEM_ORDERED CartItemStatus = 0
	// Registry approved the order.
	CART_ITEM_APPROVED CartItemStatus = 1
	// Customer paid for the order.
	CART_ITEM_PAID CartItemStatus = 2
)

type CartItem struct {
	ID         string         `json:"id,omitempty" bson:"_id"`
	MenuItemID string         `json:"menu_item_id,omitempty" bson:"menu_item_id"`
	Status     CartItemStatus `json:"status,omitempty" bson:"status"`
}
