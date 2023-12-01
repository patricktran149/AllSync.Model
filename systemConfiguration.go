package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Configuration struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ConfigurationID    string             `json:"configurationID" bson:"configurationID"`
	ConfigurationValue string             `json:"configurationValue" bson:"configurationValue"`
	Description        string             `json:"description" bson:"description"`
	Group              string             `json:"group" bson:"group"`
}

type ConfigurationRequest struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ConfigurationID    string             `json:"configurationID" bson:"configurationID"`
	ConfigurationValue string             `json:"configurationValue" bson:"configurationValue"`
	Description        string             `json:"description" bson:"description"`
	Group              string             `json:"group" bson:"group"`
}

type ConfigurationUpdateRequest struct {
	ConfigurationID    string `json:"configurationID" bson:"configurationID"`
	ConfigurationValue string `json:"configurationValue" bson:"configurationValue"`
	Description        string `json:"description" bson:"description"`
	Group              string `json:"group" bson:"group"`
}
