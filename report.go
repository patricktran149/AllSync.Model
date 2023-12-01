package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description" bson:"description"`
	Layout      string             `json:"layout" bson:"layout"`
	ReportType  string             `json:"reportType" bson:"reportType"`
	CreatedDate int64              `json:"createdDate" bson:"createdDate"`
	CreatedBy   string             `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64              `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string             `json:"updatedBy" bson:"updatedBy"`
}

type ReportRequest struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description" bson:"description"`
	Layout      string             `json:"layout" bson:"layout"`
	ReportType  string             `json:"reportType" bson:"reportType"`
	CreatedBy   string             `json:"createdBy" bson:"createdBy"`
	UpdatedBy   string             `json:"updatedBy" bson:"updatedBy"`
}

type ReportUpdateRequest struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description" bson:"description"`
	Layout      string             `json:"layout" bson:"layout"`
	ReportType  string             `json:"reportType" bson:"reportType"`
	UpdatedBy   string             `json:"updatedBy" bson:"updatedBy"`
}
