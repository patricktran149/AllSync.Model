package Model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type DirectIntegrationFlow struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FlowName    string             `json:"flowName" bson:"flowName"`
	IsActive    bool               `json:"isActive" bson:"isActive"`
	Flows       []FlowSequence     `json:"flows" bson:"flows"`
	CreatedDate int64              `json:"createdDate" bson:"createdDate"`
	CreatedBy   string             `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64              `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string             `json:"updatedBy" bson:"updatedBy"`
}

type DirectIntegrationFlowRequest struct {
	FlowName  string         `json:"flowName" bson:"flowName"`
	IsActive  bool           `json:"isActive" bson:"isActive"`
	Flows     []FlowSequence `json:"flows" bson:"flows"`
	CreatedBy string         `json:"createdBy" bson:"createdBy"`
}

type DirectIntegrationFlowUpdateRequest struct {
	FlowName  string         `json:"flowName" bson:"flowName"`
	IsActive  bool           `json:"isActive" bson:"isActive"`
	Flows     []FlowSequence `json:"flows" bson:"flows"`
	UpdatedBy string         `json:"updatedBy" bson:"updatedBy"`
}

type DirectIntegrationVersion struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FlowID      primitive.ObjectID `json:"flowID" bson:"flowID"`
	FlowName    string             `json:"flowName" bson:"flowName"`
	IsActive    bool               `json:"isActive" bson:"isActive"`
	Flows       []FlowSequence     `json:"flows" bson:"flows"`
	CreatedDate int64              `json:"createdDate" bson:"createdDate"`
	CreatedBy   string             `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64              `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string             `json:"updatedBy" bson:"updatedBy"`
}

type FlowSequence struct {
	Sequence         int                     `json:"sequence" bson:"sequence"`
	IsActive         bool                    `json:"isActive" bson:"isActive"`
	Code             string                  `json:"code" bson:"code"`
	SuccessCondition string                  `json:"successCondition" bson:"successCondition"`
	Connection       *FlowSequenceConnection `json:"connection" bson:"connection"`
}

type FlowSequenceConnection struct {
	Name           string                   `json:"name" bson:"name"`
	Type           ConnectionType           `json:"type" bson:"type"`
	ImmediateRetry int                      `json:"immediateRetry" bson:"immediateRetry"`
	Sleep          int64                    `json:"sleep" bson:"sleep"`
	API            *ConnectionAPIConfig     `json:"api" bson:"api"`
	SQL            *ConnectionSQLConfig     `json:"sql" bson:"sql"`
	MySQL          *ConnectionMySQLConfig   `json:"mysql" json:"mysql"`
	Oracle         *ConnectionOracleConfig  `json:"oracle" bson:"oracle"`
	SAPHana        *ConnectionSAPHanaConfig `json:"sapHana" bson:"sapHana"`
	SFTP           *ConnectionSFTPConfig    `json:"sftp" bson:"sftp"`
	S3             *ConnectionS3Config      `json:"s3" bson:"s3"`
	Email          *ConnectionEmailConfig   `json:"email" bson:"email"`
}

type ConnectionAPIConfig struct {
	URL                string                 `json:"url" bson:"url"`
	AuthenticationType APIAuthenticationType  `json:"authenticationType" bson:"authenticationType"`
	BearerToken        string                 `json:"bearerToken" bson:"bearerToken"`
	Method             string                 `json:"method" bson:"method"`
	BasicAuth          BasicAuth              `json:"basicAuth" bson:"basicAuth"`
	Body               map[string]interface{} `json:"body" bson:"body"`
	Params             map[string]interface{} `json:"params" bson:"params"`
	Headers            map[string]interface{} `json:"headers" bson:"headers"`
	P12                P12                    `json:"p12" bson:"p12"`
}

type ConnectionSQLConfig struct {
	Server       string `json:"server" bson:"server"`
	User         string `json:"user" bson:"user"`
	Password     string `json:"password" bson:"password"`
	DatabaseName string `json:"databaseName" bson:"databaseName"`
	Query        string `json:"query" bson:"query"`
}

type ConnectionOracleConfig struct {
	Server       string `json:"server" bson:"server"`
	User         string `json:"user" bson:"user"`
	Password     string `json:"password" bson:"password"`
	DatabaseName string `json:"databaseName" bson:"databaseName"`
	Query        string `json:"query" bson:"query"`
}

type ConnectionSAPHanaConfig struct {
	Server       string `json:"server" bson:"server"`
	User         string `json:"user" bson:"user"`
	Password     string `json:"password" bson:"password"`
	DatabaseName string `json:"databaseName" bson:"databaseName"`
	Schema       string `json:"schema" bson:"schema"`
	Query        string `json:"query" bson:"query"`
}

type ConnectionMySQLConfig struct {
	Server       string `json:"server" bson:"server"`
	User         string `json:"user" bson:"user"`
	Password     string `json:"password" bson:"password"`
	DatabaseName string `json:"databaseName" bson:"databaseName"`
	Query        string `json:"query" bson:"query"`
}

type ConnectionSFTPConfig struct {
	Host      string `json:"host" bson:"host"`
	UserName  string `json:"userName" bson:"userName"`
	Password  string `json:"password" bson:"password"`
	Port      int    `json:"port" bson:"port"`
	Path      string `json:"path" bson:"path"`
	Ext       string `json:"ext" bson:"ext"`
	CSVComma  string `json:"comma" bson:"comma"`
	FileName  string `json:"fileName" bson:"fileName"`
	IsOneFile bool   `json:"isOneFile" bson:"isOneFile"`
}

type ConnectionEmailConfig struct {
	Server    string `json:"server" bson:"server"`
	User      string `json:"user" bson:"user"`
	Password  string `json:"password" bson:"password"`
	Port      int    `json:"port" bson:"port"`
	EnableSSL bool   `json:"enableSSL" bson:"enableSSL"`
	IsDraft   bool   `json:"isDraft" bson:"isDraft"`
}

type ConnectionS3Config struct {
	HostURL    string `json:"hostUrl"`
	Region     string `json:"region"`
	BucketName string `json:"bucketName"`
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	Path       string `json:"path"`
	Ext        string `json:"ext"`
}

type ConnectionType string

const (
	ConnectionTypeAPI     ConnectionType = "API"
	ConnectionTypeSQL     ConnectionType = "SQL"
	ConnectionTypeMySQL   ConnectionType = "MYSQL"
	ConnectionTypeOracle  ConnectionType = "ORACLE"
	ConnectionTypeSAPHana ConnectionType = "SAP_HANA"
	ConnectionTypeSFTP    ConnectionType = "SFTP"
	ConnectionTypeS3      ConnectionType = "S3"
	ConnectionTypeEmail   ConnectionType = "EMAIL"
	ConnectionTypeWebhook ConnectionType = "WEBHOOK"
)

func (t *ConnectionType) isValidType() bool {
	*t = ConnectionType(strings.ToUpper(strings.TrimSpace(string(*t))))
	if !IsItemExistsInArray(*t, []ConnectionType{
		ConnectionTypeAPI,
		ConnectionTypeSQL,
		ConnectionTypeMySQL,
		ConnectionTypeOracle,
		ConnectionTypeSAPHana,
		ConnectionTypeSFTP,
		ConnectionTypeS3,
		ConnectionTypeEmail,
		ConnectionTypeWebhook,
	}) {
		return false
	}

	return true
}

type APIAuthenticationType string

const (
	APIAuthenticationTypeBearer APIAuthenticationType = "BEARER"
	APIAuthenticationTypeBasic  APIAuthenticationType = "BASIC"
	APIAuthenticationTypeNone   APIAuthenticationType = "NONE"
)

func (t *APIAuthenticationType) isValidType() bool {
	*t = APIAuthenticationType(strings.ToUpper(strings.TrimSpace(string(*t))))
	if !IsItemExistsInArray(*t, []APIAuthenticationType{
		APIAuthenticationTypeBearer,
		APIAuthenticationTypeBasic,
		APIAuthenticationTypeNone,
	}) {
		return false
	}

	return true
}

func (difReq *DirectIntegrationFlowRequest) Validate() (err error) {
	for _, flow := range difReq.Flows {
		if !flow.Connection.Type.isValidType() {
			return errors.New("Validate type ERROR - " + err.Error())
		}
	}

	return
}
