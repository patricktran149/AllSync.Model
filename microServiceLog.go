package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MicroServiceLogResponse struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ServiceID    string             `json:"serviceID" bson:"serviceID"`
	VersionID    string             `json:"versionID" bson:"versionID"`
	FunctionName string             `json:"functionName" bson:"functionName"`
	Message      string             `json:"message" bson:"message"`
	Date         int64              `json:"date" bson:"date"`
}

type MicroServiceLogRequest struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ServiceID    string             `json:"serviceID" bson:"serviceID"`
	VersionID    string             `json:"versionID" bson:"versionID"`
	FunctionName string             `json:"functionName" bson:"functionName"`
	Message      string             `json:"message" bson:"message"`
}
