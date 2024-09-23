package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryManager struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	QueryName   string             `json:"queryName" bson:"queryName"`
	SQLQuery    string             `json:"sqlQuery" bson:"sqlQuery"`
	MongoQuery  string             `json:"mongoQuery" bson:"mongoQuery"`
	CreatedBy   string             `json:"createdBy" bson:"createdBy"`
	CreatedDate int64              `json:"createdDate" bson:"createdDate"`
	UpdatedBy   string             `json:"updatedBy" bson:"updatedBy"`
	UpdatedDate int64              `json:"updatedDate" bson:"updatedDate"`
}

type QueryManagerRequest struct {
	QueryName  string `json:"queryName" bson:"queryName"`
	SQLQuery   string `json:"sqlQuery" bson:"sqlQuery"`
	MongoQuery string `json:"mongoQuery" bson:"mongoQuery"`
	CreatedBy  string `json:"createdBy" bson:"createdBy"`
	UpdatedBy  string `json:"updatedBy" bson:"updatedBy"`
}

type QueryManagerUpdateRequest struct {
	QueryName  string `json:"queryName" bson:"queryName"`
	SQLQuery   string `json:"sqlQuery" bson:"sqlQuery"`
	MongoQuery string `json:"mongoQuery" bson:"mongoQuery"`
	CreatedBy  string `json:"createdBy" bson:"createdBy"`
	UpdatedBy  string `json:"updatedBy" bson:"updatedBy"`
}
