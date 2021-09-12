package storage

import (
	"context"
	"time"

	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
	"github.com/umutozd/restaurant-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *storage) UpdateCart(ctx context.Context, req *requests.UpdateCartReq) (*types.Cart, error) {
	now := time.Now()
	toAdd, toRemove := req.GetCartItems()
	if req.ID == "" {
		req.ID = primitive.NewObjectID().Hex()
	}
	filter := bson.M{
		"_id": req.ID,
	}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	var result types.Cart

	// add
	add := bson.M{
		"$set": bson.M{
			"updated_at": now,
		},
		"$push": bson.M{"items": bson.M{"$each": toAdd}},
		"$setOnInsert": bson.M{
			"created_at": now,
		},
	}
	err := s.carts().FindOneAndUpdate(ctx, filter, add, opts).Err()
	if err != nil {
		return nil, types.Errf(types.ERR_DB_UPDATE, "error adding items to cart: %v", err)
	}

	// remove
	remove := bson.M{
		"$set": bson.M{
			"updated_at": now,
		},
		"$pull": bson.M{"items": bson.M{"_id": bson.M{"$in": toRemove}}},
	}
	err = s.carts().FindOneAndUpdate(ctx, filter, remove, opts).Decode(&result)
	if err != nil {
		return nil, types.Errf(types.ERR_DB_UPDATE, "error removing items from cart: %v", err)
	}

	return &result, nil
}

func (s *storage) UpdateCartItemsStatus(ctx context.Context, req *requests.UpdateCartItemsStatusReq) (*types.Cart, error) {
	var cart, result types.Cart
	err := s.carts().FindOne(ctx, bson.M{"_id": req.CartID}).Decode(&cart)
	if err != nil {
		return nil, types.Errf(types.ERR_DB_GET, "error getting cart with id=%s: %v", req.CartID, err)
	}

	for _, item := range cart.Items {
		if utils.StringSliceContains(req.ItemIDs, item.ID) {
			item.Status = req.Status
		}
	}

	update := bson.M{
		"$set": bson.M{
			"items": cart.Items,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = s.carts().FindOneAndUpdate(ctx, bson.M{"_id": req.CartID}, update, opts).Decode(&result)
	if err != nil {
		return nil, types.Errf(types.ERR_DB_UPDATE, "error setting items in cart: %v", err)
	}

	return &result, nil
}
