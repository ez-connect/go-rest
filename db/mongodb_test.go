package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetSort(t *testing.T) {
	sort := getSort(FindOption{
		Sort:  "myField",
		Order: "desc",
	})

	assert.Equal(t, bson.M{"myField": -1}, sort)

	sort = getSort(FindOption{
		Sort: "myField1",
	})

	assert.Equal(t, bson.M{"myField1": 1}, sort)
}

func getSort(option FindOption) interface{} {
	var sort bson.M
	if option.Sort != "" {
		if option.Order == "desc" {
			sort = bson.M{option.Sort: -1}
		} else {
			sort = bson.M{option.Sort: 1}
		}
	}

	return sort
}
