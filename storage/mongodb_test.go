package storage

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/umutozd/restaurant-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newTestStorage(t *testing.T) *storage {
	dbURL := os.Getenv("RESTAURANT_TEST_DB_URL")
	dbName := os.Getenv("RESTAURANT_TEST_DB_NAME")
	if dbURL == "" || dbName == "" {
		t.Fatalf("both dbURL and dbName must be specified as env: RESTAURANT_TEST_DB_URL=%s, RESTAURANT_TEST_DB_NAME=%s", dbURL, dbName)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(dbURL))
	if err != nil {
		t.Fatalf("error creating mongo client: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		t.Fatalf("error connecting to mongo server: %v", err)
	}
	s := &storage{client: client, dbName: dbName}

	// drop existing collections
	for _, coll := range []string{categoriesCollection, menuItemsCollection} {
		res, err := s.client.Database(s.dbName).Collection(coll).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			t.Fatalf("error deleting all documents in %s collection: %v", coll, err)
		} else {
			t.Logf("deleted %d documents from collection %s", res.DeletedCount, coll)
		}
	}
	return s
}

func compareCategories(c1, c2 *types.Category) string {
	var differences []string
	if c1 == nil && c2 == nil {
		return "both c1 and c2 are nil"
	}
	if c1 == nil {
		return "c1 is nil"
	}
	if c2 == nil {
		return "c2 is nil"
	}

	if c1.ID != c2.ID {
		differences = append(differences, fmt.Sprintf("ID's are different: c1.ID=%s, c2.ID=%s", c1.ID, c2.ID))
	}
	if c1.Name != c2.Name {
		differences = append(differences, fmt.Sprintf("Names's are different: c1.Name=%s, c2.Name=%s", c1.Name, c2.Name))
	}

	if len(differences) != 0 {
		return strings.Join(differences, "\n")
	}
	return ""
}

func compareMenuItems(i1, i2 *types.MenuItem) string {
	var differences []string
	if i1 == nil && i2 == nil {
		return "both i1 and i2 are nil"
	}
	if i1 == nil {
		return "i1 is nil"
	}
	if i2 == nil {
		return "i2 is nil"
	}

	if i1.ID != i2.ID {
		differences = append(differences, fmt.Sprintf("ID's are different: i1.ID=%s, i2.ID=%s", i1.ID, i2.ID))
	}
	if i1.Name != i2.Name {
		differences = append(differences, fmt.Sprintf("Names's are different: i1.Name=%s, i2.Name=%s", i1.Name, i2.Name))
	}
	if i1.CategoryID != i2.CategoryID {
		differences = append(differences, fmt.Sprintf("CategoryID's are different: i1.CategoryID=%s, i2.CategoryID=%s", i1.CategoryID, i2.CategoryID))
	}
	if i1.Price != i2.Price {
		differences = append(differences, fmt.Sprintf("Prices are different: i1.Price=%f, i2.Price=%f", i1.Price, i2.Price))
	}
	if i1.Image != i2.Image {
		differences = append(differences, fmt.Sprintf("Images are different: i1.Image=%s, i2.Image=%s", i1.Image, i2.Image))
	}

	if len(differences) != 0 {
		return strings.Join(differences, "\n")
	}
	return ""
}

func compareMenus(m1, m2 *types.Menu) bool {
	if m1 == nil || m2 == nil {
		return false
	}

	for _, ci1 := range m1.All {
		// search ci1.Category in m2
		for _, ci2 := range m2.All {
			diff := compareCategories(ci1.Category, ci2.Category)
			if diff == "" {
				// found the item with same category, all items must match
				for _, i1 := range ci1.Items {
					diffMI := "not matched yet"
					for _, i2 := range ci2.Items {
						diffMI = compareMenuItems(i1, i2)
						if diffMI == "" {
							break
						}
					}

					if diffMI == "" {
						// matched, search for other items
						continue
					} else {
						// i1 does not match any items in ci2
						return false
					}
				}
			}
		}
	}

	return true
}
