package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MicroServiceStatusResponse struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ServiceID     string             `json:"serviceID" bson:"serviceID"`
	ApplicationID string             `json:"applicationID" bson:"applicationID"`
	TableName     string             `json:"tableName" bson:"tableName"`
	IsRunning     bool               `json:"isRunning" bson:"isRunning"`
}

type MicroServiceStatusRequest struct {
	ServiceID     string `json:"serviceID" bson:"serviceID"`
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	TableName     string `json:"tableName" bson:"tableName"`
	IsRunning     bool   `json:"isRunning" bson:"isRunning"`
}

type MicroServiceStatusUpdateRequest struct {
	ServiceID     string `json:"serviceID" bson:"serviceID"`
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	TableName     string `json:"tableName" bson:"tableName"`
	IsRunning     bool   `json:"isRunning" bson:"isRunning"`
}
