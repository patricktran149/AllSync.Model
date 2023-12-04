package Model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type DataMapper struct {
	ID            primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID string                 `json:"applicationID" bson:"applicationID"`
	TableName     string                 `json:"tableName" bson:"tableName"`
	Template      string                 `json:"template" bson:"template"`
	Direction     DataMapperDirection    `json:"direction" bson:"direction"`
	Data          map[string]interface{} `json:"data" bson:"data"`
	CreatedDate   int64                  `json:"createdDate" bson:"createdDate"`
	CreatedBy     string                 `json:"createdBy" bson:"createdBy"`
	UpdatedDate   int64                  `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy     string                 `json:"updatedBy" bson:"updatedBy"`
}

type DataMapperRequest struct {
	ID            primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID string                 `json:"applicationID" bson:"applicationID"`
	TableName     string                 `json:"tableName" bson:"tableName"`
	Template      string                 `json:"template" bson:"template"`
	Direction     DataMapperDirection    `json:"direction" bson:"direction"`
	Data          map[string]interface{} `json:"data" bson:"data"`
	CreatedBy     string                 `json:"createdBy" bson:"createdBy"`
}

type DataMapperUpdateRequest struct {
	ID            primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID string                 `json:"applicationID" bson:"applicationID"`
	TableName     string                 `json:"tableName" bson:"tableName"`
	Template      string                 `json:"template" bson:"template"`
	Direction     DataMapperDirection    `json:"direction" bson:"direction"`
	Data          map[string]interface{} `json:"data" bson:"data"`
	UpdatedBy     string                 `json:"updatedBy" bson:"updatedBy"`
}

type DataMapperDirection string

const (
	DataMapperDirectionIn  DataMapperDirection = "IN"
	DataMapperDirectionOut DataMapperDirection = "OUT"
)

func (dmReq DataMapperRequest) Validate() error {
	if !IsItemExistsInArray(DataMapperDirection(strings.ToUpper(string(dmReq.Direction))), []DataMapperDirection{DataMapperDirectionIn, DataMapperDirectionOut}) {
		return errors.New("Direction is invalid ")
	}

	return nil
}
