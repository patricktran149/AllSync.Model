package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IntegrationFlow2 struct {
	ID          primitive.ObjectID         `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string                     `json:"name" bson:"name"`
	RequestID   primitive.ObjectID         `json:"requestID" bson:"requestID"`
	Success     []IntegrationFlowResultSet `json:"success" bson:"success"`
	Failure     []IntegrationFlowResultSet `json:"failure" bson:"failure"`
	Designer    string                     `json:"designer" bson:"designer"`
	CreatedDate int64                      `json:"createdDate" bson:"createdDate"`
	CreatedBy   string                     `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64                      `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string                     `json:"updatedBy" bson:"updatedBy"`
}

type IntegrationFlowResultSet struct {
	RequestID primitive.ObjectID         `json:"requestID" bson:"requestID"`
	Success   []IntegrationFlowResultSet `json:"success" bson:"success"`
	Failure   []IntegrationFlowResultSet `json:"failure" bson:"failure"`
}

type IntegrationFlow2Request struct {
	ID        primitive.ObjectID         `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string                     `json:"name" bson:"name"`
	RequestID primitive.ObjectID         `json:"requestID" bson:"requestID"`
	Success   []IntegrationFlowResultSet `json:"success" bson:"success"`
	Failure   []IntegrationFlowResultSet `json:"failure" bson:"failure"`
	Designer  string                     `json:"designer" bson:"designer"`
	CreatedBy string                     `json:"createdBy" bson:"createdBy"`
	UpdatedBy string                     `json:"updatedBy" bson:"updatedBy"`
}

type IntegrationFlow2UpdateRequest struct {
	Name      string                     `json:"name" bson:"name"`
	RequestID primitive.ObjectID         `json:"requestID" bson:"requestID"`
	Success   []IntegrationFlowResultSet `json:"success" bson:"success"`
	Failure   []IntegrationFlowResultSet `json:"failure" bson:"failure"`
}
