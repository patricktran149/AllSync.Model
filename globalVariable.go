package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GlobalVariable struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VariableID    string             `json:"variableID" bson:"variableID"`
	VariableValue string             `json:"variableValue" bson:"variableValue"`
	CreatedDate   int64              `json:"createdDate" bson:"createdDate"`
	CreatedBy     string             `json:"createdBy" bson:"createdBy"`
	UpdatedDate   int64              `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy     string             `json:"updatedBy" bson:"updatedBy"`
}

type GlobalVariableRequest struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VariableID    string             `json:"variableID" bson:"variableID"`
	VariableValue string             `json:"variableValue" bson:"variableValue"`
	CreatedBy     string             `json:"createdBy" bson:"createdBy"`
}

type GlobalVariableUpdateRequest struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VariableID    string             `json:"variableID" bson:"variableID"`
	VariableValue string             `json:"variableValue" bson:"variableValue"`
	UpdatedBy     string             `json:"updatedBy" bson:"updatedBy"`
}
