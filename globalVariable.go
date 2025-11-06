package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GlobalVariable struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VariableID    string             `json:"variableID" bson:"variableID"`
	VariableValue string             `json:"variableValue" bson:"variableValue"`
}

type GlobalVariableRequest struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VariableID    string             `json:"variableID" bson:"variableID"`
	VariableValue string             `json:"variableValue" bson:"variableValue"`
}

type GlobalVariableUpdateRequest struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VariableID    string             `json:"variableID" bson:"variableID"`
	VariableValue string             `json:"variableValue" bson:"variableValue"`
}
