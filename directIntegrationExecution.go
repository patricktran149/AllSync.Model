package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DirectIntegrationExecution struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FlowID        primitive.ObjectID `json:"flowID" bson:"flowID"`
	Status        string             `json:"status" bson:"status"`
	ExecutionDate int64              `json:"executionDate" bson:"executionDate"`
	Details       []ExecutionDetail  `json:"details" bson:"details"`
}

type DirectIntegrationExecutionRequest struct {
	FlowID        primitive.ObjectID `json:"flowID" bson:"flowID"`
	Status        string             `json:"status" bson:"status"`
	ExecutionDate int64              `json:"executionDate" bson:"executionDate"`
	Details       []ExecutionDetail  `json:"details" bson:"details"`
}

type DirectIntegrationExecutionUpdateRequest struct {
	Status        string            `json:"status" bson:"status"`
	ExecutionDate int64             `json:"executionDate" bson:"executionDate"`
	Details       []ExecutionDetail `json:"details" bson:"details"`
}

type ExecutionDetail struct {
	Sequence        int            `json:"sequence" bson:"sequence"`
	ConnectionName  string         `json:"connectionName" bson:"connectionName"`
	Type            ConnectionType `json:"type" bson:"type"`
	Status          string         `json:"status" bson:"status"`
	RawData         string         `json:"rawData" bson:"rawData"`
	SendData        string         `json:"sendData" bson:"sendData"`
	ResponseData    string         `json:"responseData" bson:"responseData"`
	ConsolidateData string         `json:"consolidateData" bson:"consolidateData"`
	Message         string         `json:"message" bson:"message"`
}

type DirectIntegrationExecutionLog struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FlowID         primitive.ObjectID `json:"flowID" bson:"flowID"`
	ExecutionID    primitive.ObjectID `json:"executionID" bson:"executionID"`
	ConnectionName string             `json:"connectionName" bson:"connectionName"`
	Sequence       int                `json:"sequence" bson:"sequence"`
	LogMessage     string             `json:"logMessage" bson:"logMessage"`
	LogDate        int64              `json:"logDate" bson:"logDate"`
}

type DirectIntegrationExecutionLogRequest struct {
	FlowID         primitive.ObjectID `json:"flowID" bson:"flowID"`
	ExecutionID    primitive.ObjectID `json:"executionID" bson:"executionID"`
	ConnectionName string             `json:"connectionName" bson:"connectionName"`
	Sequence       int                `json:"sequence" bson:"sequence"`
	LogMessage     string             `json:"logMessage" bson:"logMessage"`
}
