package storage

import (
	"context"
	"time"

	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
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
