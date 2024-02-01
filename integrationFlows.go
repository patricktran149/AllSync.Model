package Model

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type IntegrationFlow struct {
	ID          primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	FromApp     string                   `json:"fromApplicationID" bson:"fromApplicationID"`
	ToApp       string                   `json:"toApplicationID" bson:"toApplicationID"`
	TableName   string                   `json:"tableName" bson:"tableName"`
	Source      IntegrationFlowConfig    `json:"source" bson:"source"`
	Target      IntegrationFlowConfig    `json:"target" bson:"target"`
	TargetList  []IntegrationFlowConfig  `json:"targetList" bson:"targetList"`
	Mapping     []map[string]interface{} `json:"mapping" bson:"mapping"`
	CreatedDate int64                    `json:"createdDate" bson:"createdDate"`
	CreatedBy   string                   `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64                    `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string                   `json:"updatedBy" bson:"updatedBy"`
}

type IntegrationFlowConfig struct {
	ApplicationID   string                     `json:"applicationID" bson:"applicationID"`
	Condition       string                     `json:"condition" bson:"condition"`
	Mode            FlowMode                   `json:"mode" bson:"mode"`
	Method          FlowMethod                 `json:"method" bson:"method"`
	S3              S3                         `json:"s3" bson:"s3"`
	SFTP            SFTPConfig                 `json:"sftp" bson:"sftp"`
	API             APIConfig                  `json:"api" bson:"api"`
	OAuth           APIConfig                  `json:"oAuth" bson:"oAuth"`
	SQL             SQLConfig                  `json:"sql" bson:"sql"`
	Oracle          OracleConfig               `json:"oracle" bson:"oracle"`
	Data            map[string]interface{}     `json:"data" bson:"data"`
	Template        string                     `json:"template" bson:"template"`
	ResponseMapping string                     `json:"responseMapping" bson:"responseMapping"`
	SuccessTarget   IntegrationFlowConfigAfter `json:"successTarget" bson:"successTarget"`
	FailedTarget    IntegrationFlowConfigAfter `json:"failedTarget" bson:"failedTarget"`
}

type IntegrationFlowConfigAfter struct {
	ApplicationID   string                 `json:"applicationID" bson:"applicationID"`
	Condition       string                 `json:"condition" bson:"condition"`
	Mode            FlowMode               `json:"mode" bson:"mode"`
	Method          FlowMethod             `json:"method" bson:"method"`
	S3              S3                     `json:"s3" bson:"s3"`
	SFTP            SFTPConfig             `json:"sftp" bson:"sftp"`
	API             APIConfig              `json:"api" bson:"api"`
	OAuth           APIConfig              `json:"oAuth" bson:"oAuth"`
	SQL             SQLConfig              `json:"sql" bson:"sql"`
	Oracle          OracleConfig           `json:"oracle" bson:"oracle"`
	Data            map[string]interface{} `json:"data" bson:"data"`
	Template        string                 `json:"template" bson:"template"`
	ResponseMapping string                 `json:"responseMapping" bson:"responseMapping"`
}

type IntegrationFlowRequest struct {
	FromApp     string                   `json:"fromApplicationID" bson:"fromApplicationID"`
	ToApp       string                   `json:"toApplicationID" bson:"toApplicationID"`
	TableName   string                   `json:"tableName" bson:"tableName"`
	Source      IntegrationFlowConfig    `json:"source" bson:"source"`
	TargetList  []IntegrationFlowConfig  `json:"targetList" bson:"targetList"`
	Mapping     []map[string]interface{} `json:"mapping" bson:"mapping"`
	CreatedDate int64                    `json:"createdDate" bson:"createdDate"`
	CreatedBy   string                   `json:"createdBy" bson:"createdBy"`
	UpdatedDate int64                    `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string                   `json:"updatedBy" bson:"updatedBy"`
}

type IntegrationFlowUpdateRequest struct {
	FromApp     string                   `json:"fromApplicationID" bson:"fromApplicationID"`
	ToApp       string                   `json:"toApplicationID" bson:"toApplicationID"`
	TableName   string                   `json:"tableName" bson:"tableName"`
	Condition   string                   `json:"condition" bson:"condition"`
	Source      IntegrationFlowConfig    `json:"source" bson:"source"`
	TargetList  []IntegrationFlowConfig  `json:"targetList" bson:"targetList"`
	Mode        FlowMode                 `json:"mode" bson:"mode"`
	Mapping     []map[string]interface{} `json:"mapping" bson:"mapping"`
	CreatedDate int64                    `json:"createdDate" bson:"createdDate"`
	UpdatedDate int64                    `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy   string                   `json:"updatedBy" bson:"updatedBy"`
}

type IntegrationFlowMapping struct {
	In  []map[string]interface{} `json:"in" bson:"in"`
	Out []map[string]interface{} `json:"out" bson:"out"`
}

type SFTPConfig struct {
	Host     string `json:"host" bson:"host"`
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
	Port     int    `json:"port" bson:"port"`
	Path     string `json:"path" bson:"path"`
	Ext      string `json:"ext" bson:"ext"`
	CSVComma string `json:"comma" bson:"comma"`
	FileName string `json:"fileName" bson:"fileName"`
}

type SQLConfig struct {
	Server       string `json:"server" bson:"server"`
	User         string `json:"user" bson:"user"`
	Password     string `json:"password" bson:"password"`
	DatabaseName string `json:"databaseName" bson:"databaseName"`
	TableName    string `json:"tableName" bson:"tableName"`
}

type OracleConfig struct {
	Server       string `json:"server" bson:"server"`
	User         string `json:"user" bson:"user"`
	Password     string `json:"password" bson:"password"`
	DatabaseName string `json:"databaseName" bson:"databaseName"`
	TableName    string `json:"tableName" bson:"tableName"`
}

type APIConfig struct {
	URL                string                 `json:"url" bson:"url"`
	AuthenticationType AuthenticationType     `json:"authenticationType" bson:"authenticationType"`
	HMAC               HMACConfig             `json:"hmac" bson:"hmac"`
	OAuth1             OAuth1                 `json:"oAuth1" bson:"oAuth1"`
	OAuth2             OAuth2                 `json:"oAuth2" bson:"oAuth2"`
	Token              string                 `json:"token" bson:"token"`
	Method             string                 `json:"method" bson:"method"`
	BasicAuth          BasicAuth              `json:"basicAuth" bson:"basicAuth"`
	IsMandatory        bool                   `json:"isMandatory" bson:"isMandatory"`
	Body               map[string]interface{} `json:"body" bson:"body"`
	Params             map[string]interface{} `json:"params" bson:"params"`
	Headers            map[string]interface{} `json:"headers" bson:"headers"`
	Template           string                 `json:"template" bson:"template"`
}

type OAuth1 struct {
	AddData         string `json:"addData" bson:"addData"`
	SignatureMethod string `json:"signatureMethod" bson:"signatureMethod"`
	ConsumerKey     string `json:"consumerKey" bson:"consumerKey"`
	ConsumerSecret  string `json:"consumerSecret" bson:"consumerSecret"`
	AccessToken     string `json:"accessToken" bson:"accessToken"`
	TokenSecret     string `json:"tokenSecret" bson:"tokenSecret"`
	Nonce           string `json:"nonce" bson:"nonce"`
	Realm           string `json:"realm" bson:"realm"`
	Timestamp       string `json:"timestamp" bson:"timestamp"`
}

type OAuth2 struct {
	SignatureMethod      SignatureMethod `json:"signatureMethod" bson:"signatureMethod"`
	AuthorizeUrl         string          `json:"authorizeUrl" bson:"authorizeUrl"`
	AccessTokenUrl       string          `json:"accessTokenUrl" bson:"accessTokenUrl"`
	FreshTokenUrl        string          `json:"freshTokenUrl" bson:"freshTokenUrl"`
	ClientID             string          `json:"clientID" bson:"clientID"`
	ClientSecret         string          `json:"clientSecret" bson:"clientSecret"`
	Scope                string          `json:"scope" bson:"scope"`
	ClientAuthentication string          `json:"clientAuthentication" bson:"clientAuthentication"`
	HeaderPrefix         string          `json:"headerPrefix" bson:"headerPrefix"`
}

type S3 struct {
	HostURL    string `json:"hostUrl"`
	Region     string `json:"region"`
	BucketName string `json:"bucketName"`
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	Path       string `json:"path"`
	Ext        string `json:"ext"`
}

//addData: BODY, HEADER
//- signatureMethod: HMAC-SHA1, HMAC-SHA256, HMAC-SHA512, RSA-SHA1, RSA-SHA256, RSA-SHA512, PLAINTEXT
//- consumerKey
//- consumerSecret
//- accessToken
//- tokenSecret
//- nonce
//- realm
//- timestamp

type HMACConfig struct {
	AppIDName       string        `json:"appIDName" bson:"appIDName"`
	AppIDValue      string        `json:"appIDValue" bson:"appIDValue"`
	AppKeyName      string        `json:"appKeyName" bson:"appKeyName"`
	AppKeyValue     string        `json:"appKeyValue" bson:"appKeyValue"`
	SignatureName   string        `json:"signatureName" bson:"signatureName"`
	BodyName        string        `json:"bodyName" bson:"bodyName"`
	BaseFields      string        `json:"baseFields" bson:"baseFields"`
	Algorithm       HMACAlgorithm `json:"algorithm" bson:"algorithm"`
	TimestampName   string        `json:"timestampName" bson:"timestampName"`
	TimestampFormat string        `json:"timestampFormat" bson:"timestampFormat"`
}

func (ifr *IntegrationFlowRequest) Validate() (err error) {
	//if !ifr.Source.Method.isValid() {
	//	return errors.New(fmt.Sprintf("Method [%s] is invalid ", ifr.Source.Method))
	//}

	if !ifr.Source.API.HMAC.Algorithm.isValidType() && ifr.Source.API.AuthenticationType == AuthenticationTypeHmac {
		return errors.New(fmt.Sprintf("HMAC Algorithm [%s] is invalid ", ifr.Source.API.HMAC.Algorithm))
	}

	//if !ifr.Source.Mode.isValid() {
	//	return errors.New(fmt.Sprintf("source.mode [%s] is invalid ", ifr.Source.Mode))
	//}
	//
	//if !ifr.Source.Method.isValid() {
	//	return errors.New(fmt.Sprintf("source.method [%s] is invalid ", ifr.Source.Method))
	//}
	//
	//if !ifr.Source.API.HMAC.Algorithm.isValidType() {
	//	return errors.New(fmt.Sprintf("source.api.hmac.algorithm [%s] is invalid ", ifr.Source.API.HMAC.Algorithm))
	//}
	//
	//if !ifr.Source.API.AuthenticationType.isValidType() {
	//	return errors.New(fmt.Sprintf("source.api.authenticationType [%s] is invalid ", ifr.Source.API.AuthenticationType))
	//}
	//
	//if !ifr.Source.API.OAuth2.SignatureMethod.isValidType() {
	//	return errors.New(fmt.Sprintf("source.api.oAuth2.signatureMethod [%s] is invalid ", ifr.Source.API.OAuth2.SignatureMethod))
	//}

	return
}

// -------------------------------------------------------
type HMACAlgorithm string

const (
	HMACAlgorithmSHA1       HMACAlgorithm = "SHA1"
	HMACAlgorithmSHA256     HMACAlgorithm = "SHA256"
	HMACAlgorithmSHA512     HMACAlgorithm = "SHA512"
	HMACAlgorithmHMACSHA1   HMACAlgorithm = "HMAC-SHA1"
	HMACAlgorithmHMACSHA256 HMACAlgorithm = "HMAC-SHA256"
	HMACAlgorithmHMACSHA512 HMACAlgorithm = "HMAC-SHA512"
)

func (a *HMACAlgorithm) isValidType() bool {
	*a = HMACAlgorithm(strings.ToUpper(strings.TrimSpace(string(*a))))
	if !IsItemExistsInArray(*a, []HMACAlgorithm{
		HMACAlgorithmSHA1,
		HMACAlgorithmSHA256,
		HMACAlgorithmSHA512,
		HMACAlgorithmHMACSHA1,
		HMACAlgorithmHMACSHA256,
		HMACAlgorithmHMACSHA512}) {
		return false
	}

	return true
}

// -------------------------------------------------------
type AuthenticationType string

func (t *AuthenticationType) isValidType() bool {
	*t = AuthenticationType(strings.ToUpper(strings.TrimSpace(string(*t))))
	if !IsItemExistsInArray(*t, []AuthenticationType{
		AuthenticationTypeBearer,
		AuthenticationTypeBasic,
		AuthenticationTypeHmac,
		AuthenticationTypeOAuth1,
		AuthenticationTypeOAuth2,
		AuthenticationTypeNone,
	}) {
		return false
	}

	return true
}

const (
	AuthenticationTypeBearer AuthenticationType = "BEARER"
	AuthenticationTypeBasic  AuthenticationType = "BASIC"
	AuthenticationTypeHmac   AuthenticationType = "HMAC"
	AuthenticationTypeOAuth1 AuthenticationType = "OAUTH1"
	AuthenticationTypeOAuth2 AuthenticationType = "OAUTH2"
	AuthenticationTypeNone   AuthenticationType = "NONE"
)

// -------------------------------------------------------
type FlowMethod string

const (
	FlowMethodAPI    FlowMethod = "API"
	FlowMethodSFTP   FlowMethod = "SFTP"
	FlowMethodQueue  FlowMethod = "QUEUE"
	FlowMethodSQL    FlowMethod = "SQL"
	FlowMethodOracle FlowMethod = "ORACLE"
	FlowMethodS3     FlowMethod = "S3"
)

func (fm *FlowMethod) isValid() bool {
	*fm = FlowMethod(strings.ToUpper(strings.ReplaceAll(strings.Trim(string(*fm), " "), "_", "")))
	if !IsItemExistsInArray(*fm, []FlowMethod{FlowMethodAPI, FlowMethodSFTP, FlowMethodQueue, FlowMethodS3, FlowMethodSQL, FlowMethodOracle}) {
		return false
	}

	return true
}

//-------------------------------------------------------

type FlowMode string

const (
	FlowModeAdd    FlowMode = "ADD"
	FlowModeUpdate FlowMode = "UPDATE"
	FlowModeBoth   FlowMode = "BOTH"
)

func (fMod *FlowMode) isValid() bool {
	*fMod = FlowMode(strings.ToUpper(strings.ReplaceAll(strings.Trim(string(*fMod), " "), "_", "")))
	if !IsItemExistsInArray(*fMod, []FlowMode{FlowModeAdd, FlowModeUpdate, FlowModeBoth}) {
		return false
	}

	return true
}

// -------------------------------------------------------
type SignatureMethod string

const (
	SignatureMethodHMACSHA1   SignatureMethod = "HMAC-SHA1"
	SignatureMethodHMACSHA256 SignatureMethod = "HMAC-SHA256"
	SignatureMethodHMACSHA512 SignatureMethod = "HMAC-SHA512"
	SignatureMethodRSASHA1    SignatureMethod = "RSA-SHA1"
	SignatureMethodRSASHA256  SignatureMethod = "RSA-SHA256"
	SignatureMethodRSASHA512  SignatureMethod = "RSA-SHA512"
	SignatureMethodPlainText  SignatureMethod = "PLAINTEXT"
)

func (t *SignatureMethod) isValidType() bool {
	*t = SignatureMethod(strings.ToUpper(strings.TrimSpace(string(*t))))
	if !IsItemExistsInArray(*t, []SignatureMethod{
		SignatureMethodHMACSHA1,
		SignatureMethodHMACSHA256,
		SignatureMethodHMACSHA512,
		SignatureMethodRSASHA1,
		SignatureMethodRSASHA256,
		SignatureMethodRSASHA512,
		SignatureMethodPlainText,
	}) {
		return false
	}

	return true
}
