package Model

import "go.mongodb.org/mongo-driver/bson"

type AllSyncConfig struct {
	SystemAPIURL    string
	SystemSecretKey string
	TenantID        string
	Token           string
}

type ErrorType struct {
	MongoErrMsg string
	APIErrMsg   string
}

var MapErrorType = map[int]string{
	11000: "Duplicated key",
}

type ResponseResult struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode int         `json:"errorCode"`
	Data      interface{} `json:"data"`
}

type ToAppResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
	Data      Data   `json:"data"`
}

type Sort int

const (
	SortDesc Sort = -1
	SortAsc  Sort = 1
)

type Data struct {
	//------------------Unique field-----------------
	FileURL string `json:"fileURL"`

	//------------------Array-----------------
	ObjectList                        []bson.M                        `json:"objectList"`
	QueueIncomingList                 []Queue                         `json:"queueIncomingList"`
	QueueOutgoingList                 []Queue                         `json:"queueOutgoingList"`
	QueueSFTPList                     []Queue                         `json:"queueSFTPList"`
	QueueWebhookList                  []Queue                         `json:"queueWebhookList"`
	UserDefinedTableList              []UserDefinedTable              `json:"userDefinedTableList"`
	ApplicationList                   []Application                   `json:"applicationList"`
	IntegrationFlowList               []IntegrationFlow               `json:"integrationFlowList"`
	QueueStatusList                   []ReportQueueStatus             `json:"queueStatusList"`
	DataMapperList                    []DataMapper                    `json:"dataMapperList"`
	TenantList                        []Tenant                        `json:"tenantList"`
	MicroServiceList                  []MicroServiceResponse          `json:"microServiceList"`
	MicroServiceStatusList            []MicroServiceStatusResponse    `json:"microServiceStatusList"`
	DirectIntegrationFlowList         []DirectIntegrationFlow         `json:"directIntegrationFlowList"`
	DirectIntegrationExecutionList    []DirectIntegrationExecution    `json:"directIntegrationExecutionList"`
	DirectIntegrationExecutionLogList []DirectIntegrationExecutionLog `json:"directIntegrationExecutionLogList"`

	//------------------Object-----------------

	Object                        bson.M                        `json:"object"`
	QueueIncoming                 Queue                         `json:"queueIncoming"`
	QueueOutgoing                 Queue                         `json:"queueOutgoing"`
	QueueSFTP                     Queue                         `json:"queueSFTP"`
	QueueWebhook                  Queue                         `json:"queueWebhook"`
	UserDefinedTable              UserDefinedTable              `json:"userDefinedTable"`
	Application                   Application                   `json:"application"`
	DataOut                       string                        `json:"dataOut"`
	DataMapper                    DataMapper                    `json:"dataMapper"`
	IntegrationFlow               IntegrationFlow               `json:"integrationFlow"`
	Tenant                        Tenant                        `json:"tenant"`
	MicroService                  MicroServiceResponse          `json:"microService"`
	MicroServiceStatus            MicroServiceStatusResponse    `json:"microServiceStatus"`
	DirectIntegrationFlow         DirectIntegrationFlow         `json:"directIntegrationFlow"`
	DirectIntegrationExecution    DirectIntegrationExecution    `json:"directIntegrationExecution"`
	DirectIntegrationExecutionLog DirectIntegrationExecutionLog `json:"directIntegrationExecutionLog"`
}

type DenominationRequest struct {
	DenominationList []float64 `json:"denominationList"`
	Payment          float64   `json:"payment"`
}

type CollectionPipeline struct {
	Collection string `json:"collection"`
	Pipeline   string `json:"pipeline"`
}
