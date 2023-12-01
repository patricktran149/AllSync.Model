package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SecurityRole struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SecurityID    string             `json:"securityID" bson:"securityID"`
	SecurityRight []struct {
		TableName string `json:"tableName" bson:"tableName"`
		Role      string `json:"role" bson:"role"`
	} `json:"securityRight" bson:"securityRight"`
	CreatedDate int64  `json:"createdDate" bson:"createdDate"`
	CreatedBy   string `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64  `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string `json:"updatedBy" bson:"updatedBy"`
}

type SecurityRoleRequest struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SecurityID    string             `json:"securityID" bson:"securityID"`
	SecurityRight []struct {
		TableName string `json:"tableName" bson:"tableName"`
		Role      string `json:"role" bson:"role"`
	} `json:"securityRight" bson:"securityRight"`
	CreatedDate int64  `json:"createdDate" bson:"createdDate"`
	CreatedBy   string `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64  `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string `json:"updatedBy" bson:"updatedBy"`
}

type SecurityRoleUpdateRequest struct {
	SecurityID    string `json:"securityID" bson:"securityID"`
	SecurityRight []struct {
		TableName string `json:"tableName" bson:"tableName"`
		Role      string `json:"role" bson:"role"`
	} `json:"securityRight" bson:"securityRight"`
	UpdatedBy   string `json:"updatedBy" bson:"updatedBy"`
	UpdatedDate int64  `json:"updatedDate" bson:"updatedDate"`
}
