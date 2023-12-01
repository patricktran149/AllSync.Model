package Model

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Object struct {
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	TableName     string `json:"tableName" bson:"tableName"`
	Fields        bson.M `json:"fields" bson:"fields"`
}

type ObjectRequest struct {
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	TableName     string `json:"tableName" bson:"tableName"`
	Fields        bson.M `json:"fields" bson:"fields"`
}

type ObjectUpdateRequest struct {
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	TableName     string `json:"tableName" bson:"tableName"`
	Fields        bson.M `json:"fields" bson:"fields"`
}

type ObjectOverView struct {
	Input      bson.M   `json:"input"`
	MappingIn  []bson.M `json:"mappingIn"`
	MappingOut []bson.M `json:"mappingOut"`
}

type LiquidObjectOverView struct {
	Data     map[string]interface{} `json:"data"`
	Template string                 `json:"template"`
}
