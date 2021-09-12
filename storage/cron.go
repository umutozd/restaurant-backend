package storage

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *storage) startCron() {
	ticker := time.NewTicker(s.cronInterval)
	for {
		<-ticker.C
		s.deleteExpiredCarts()
	}
}

func (s *storage) deleteExpiredCarts() {
	logrus.Info("Storage: deleting expired carts")
	filter := bson.M{
		"created_at": bson.M{"$lt": time.Now().AddDate(0, 0, -14)},
	}
	res, err := s.carts().DeleteMany(context.Background(), filter)
	if err != nil {
		err = types.Err(types.ERR_DB_DELETE, err)
		logrus.WithError(err).Error("Storage: error deleting expired carts")
	}
	logrus.Infof("Storage: deleted %d expired carts", res.DeletedCount)
}
