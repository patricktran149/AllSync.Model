package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type APILog struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	URL          string             `json:"url" bson:"url"`
	Method       string             `json:"method" bson:"method"`
	SourceIP     string             `json:"sourceIP" bson:"sourceIP"`
	StartDate    int64              `json:"startDate" bson:"startDate"`
	EndDate      int64              `json:"endDate" bson:"endDate"`
	ResponseTime int64              `json:"responseTime" bson:"responseTime"`
	TenantID     string             `json:"tenantID" bson:"tenantID"`
	CreatedDate  int64              `json:"createdDate" bson:"createdDate"`
}

type APILogRequest struct {
	URL          string `json:"url" bson:"url"`
	Method       string `json:"method" bson:"method"`
	SourceIP     string `json:"sourceIP" bson:"sourceIP"`
	StartDate    int64  `json:"startDate" bson:"startDate"`
	EndDate      int64  `json:"endDate" bson:"endDate"`
	ResponseTime int64  `json:"responseTime" bson:"responseTime"`
	TenantID     string `json:"tenantID" bson:"tenantID"`
}
