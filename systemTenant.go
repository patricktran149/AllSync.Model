package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tenant struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	TenantID       string             `json:"tenantID" bson:"tenantID"`
	TenantName     string             `json:"tenantName" bson:"tenantName"`
	Description    string             `json:"description" bson:"description"`
	Token          string             `json:"token" bson:"token"`
	Logo           string             `json:"logo" bson:"logo"`
	IsActive       bool               `json:"isActive" bson:"isActive"`
	DatabaseServer string             `json:"databaseServer" bson:"databaseServer"`
	AdminEmail     string             `json:"adminEmail" bson:"adminEmail"`
	AdminPassword  string             `json:"adminPassword,omitempty" bson:"adminPassword,omitempty"`
	SMTP           SMTP               `json:"smtp" bson:"smtp"`
	CreatedDate    int64              `json:"createdDate" bson:"createdDate"`
	CreatedBy      string             `json:"createdBy" bson:"createdBy"`
	UpdatedDate    int64              `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy      string             `json:"updatedBy" bson:"updatedBy"`
}

type TenantRequest struct {
	TenantName     string `json:"tenantName" bson:"tenantName"`
	Description    string `json:"description" bson:"description"`
	Logo           string `json:"logo" bson:"logo"`
	IsActive       bool   `json:"isActive" bson:"isActive"`
	DatabaseServer string `json:"databaseServer" bson:"databaseServer"`
	AdminEmail     string `json:"adminEmail" bson:"adminEmail"`
	AdminPassword  string `json:"adminPassword" bson:"adminPassword"`
	SMTP           SMTP   `json:"smtp" bson:"smtp"`
	CreatedBy      string `json:"createdBy" bson:"createdBy"`
	UpdatedBy      string `json:"updatedBy" bson:"updatedBy"`
}

type TenantUpdateRequest struct {
	TenantName     string `json:"tenantName" bson:"tenantName"`
	Description    string `json:"description" bson:"description"`
	Logo           string `json:"logo" bson:"logo"`
	IsActive       bool   `json:"isActive" bson:"isActive"`
	DatabaseServer string `json:"databaseServer" bson:"databaseServer"`
	SMTP           SMTP   `json:"smtp" bson:"smtp"`
	UpdatedBy      string `json:"updatedBy" bson:"updatedBy"`
}

type SMTP struct {
	Server    string `json:"server" bson:"server"`
	User      string `json:"user" bson:"user"`
	Password  string `json:"password" bson:"password"`
	Port      int    `json:"port" bson:"port"`
	HasSSL    bool   `json:"hasSSL" bson:"hasSSL"`
	FromEmail string `json:"fromEmail" bson:"fromEmail"`
	BccEmail  string `json:"bccEmail" bson:"bccEmail"`
}

type TenantDuplicateRequest struct {
	TenantName     string `json:"tenantName"`
	DatabaseServer string `json:"databaseServer"`
	TenantID       string `json:"tenantID"`
}

type TenantSchedule struct {
	TenantID string `json:"tenantID"`
	Schedule string `json:"schedule"`
}
