package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MicroServiceResponse struct {
	ID                  primitive.ObjectID            `json:"_id,omitempty" bson:"_id,omitempty"`
	ServiceID           string                        `json:"serviceID" bson:"serviceID"`
	Description         string                        `json:"description" bson:"description"`
	FileURL             string                        `json:"fileURL" bson:"fileURL"`
	VersionID           string                        `json:"versionID" bson:"versionID"`
	Status              string                        `json:"status" bson:"status"`
	LastRun             int64                         `json:"lastRun" bson:"lastRun"`
	Message             string                        `json:"message" bson:"message"`
	CreatedBy           string                        `json:"createdBy" bson:"createdBy"`
	UpdatedBy           string                        `json:"updatedBy" bson:"updatedBy"`
	CustomCode          string                        `json:"customCode" bson:"customCode"`
	LanguageCustomCode  string                        `json:"languageCustomCode" bson:"languageCustomCode"`
	IsDisable           bool                          `json:"isDisable" bson:"isDisable"`
	ApplicationTypeList []MicroServiceApplicationType `json:"applicationTypeList" bson:"applicationTypeList"`
	CreatedDate         int64                         `json:"createdDate" bson:"createdDate"`
	UpdatedDate         int64                         `json:"updatedDate" bson:"updatedDate"`
}

type MicroServiceRequest struct {
	ID                  primitive.ObjectID            `json:"_id,omitempty" bson:"_id,omitempty"`
	ServiceID           string                        `json:"serviceID" bson:"serviceID"`
	Description         string                        `json:"description" bson:"description"`
	FileURL             string                        `json:"fileURL" bson:"fileURL"`
	VersionID           string                        `json:"versionID" bson:"versionID"`
	Status              string                        `json:"status" bson:"status"`
	LastRun             int64                         `json:"lastRun" bson:"lastRun"`
	Message             string                        `json:"message" bson:"message"`
	CustomCode          string                        `json:"customCode" bson:"customCode"`
	LanguageCustomCode  string                        `json:"languageCustomCode" bson:"languageCustomCode"`
	IsDisable           bool                          `json:"isDisable" bson:"isDisable"`
	ApplicationTypeList []MicroServiceApplicationType `json:"applicationTypeList" bson:"applicationTypeList"`
	CreatedBy           string                        `json:"createdBy" bson:"createdBy"`
	UpdatedBy           string                        `json:"updatedBy" bson:"updatedBy"`
}

type MicroServiceUpdateRequest struct {
	ServiceID          string `json:"serviceID" bson:"serviceID"`
	Description        string `json:"description" bson:"description"`
	FileURL            string `json:"fileURL" bson:"fileURL"`
	VersionID          string `json:"versionID" bson:"versionID"`
	Status             string `json:"status" bson:"status"`
	LastRun            int64  `json:"lastRun" bson:"lastRun"`
	Message            string `json:"message" bson:"message"`
	CustomCode         string `json:"customCode" bson:"customCode"`
	LanguageCustomCode string `json:"languageCustomCode" bson:"languageCustomCode"`
	IsDisable          bool   `json:"isDisable" bson:"isDisable"`
	UpdatedBy          string `json:"updatedBy" bson:"updatedBy"`
}

type MicroServiceApplicationType struct {
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	Type          string `json:"type" bson:"type"`
}
