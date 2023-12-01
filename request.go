package Model

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name"`
	ApplicationID    string             `json:"applicationID" bson:"applicationID"`
	ObjectType       string             `json:"objectType" bson:"objectType"`
	Direction        string             `json:"direction" bson:"direction"`
	LiquidTemplate   string             `json:"liquidTemplate" bson:"LiquidTemplate"`
	LiquidSampleData string             `json:"liquidSampleData" bson:"liquidSampleData"`
	RequestType      string             `json:"requestType" bson:"requestType"`
	SFTP             SFTPConfig         `json:"sftp" bson:"sftp"`
	API              APIConfig          `json:"api" bson:"api"`
	OAuth            APIConfig          `json:"oAuth" bson:"oAuth"`
	SQL              SQLConfig          `json:"sql" bson:"sql"`
	ResponseMapping  string             `json:"responseMapping" bson:"responseMapping"`
	CreatedDate      int64              `json:"createdDate" bson:"createdDate"`
	CreatedBy        string             `json:"createdBy" bson:"createdBy"`
	UpdatedDate      int64              `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy        string             `json:"updatedBy" bson:"updatedBy"`
}

type RequestRequest struct {
	Name             string     `json:"name" bson:"name"`
	ApplicationID    string     `json:"applicationID" bson:"applicationID"`
	ObjectType       string     `json:"objectType" bson:"objectType"`
	Direction        string     `json:"direction" bson:"direction"`
	LiquidTemplate   string     `json:"liquidTemplate" bson:"LiquidTemplate"`
	LiquidSampleData string     `json:"liquidSampleData" bson:"liquidSampleData"`
	RequestType      string     `json:"requestType" bson:"requestType"`
	SFTP             SFTPConfig `json:"sftp" bson:"sftp"`
	API              APIConfig  `json:"api" bson:"api"`
	OAuth            APIConfig  `json:"oAuth" bson:"oAuth"`
	SQL              SQLConfig  `json:"sql" bson:"sql"`
	ResponseMapping  string     `json:"responseMapping" bson:"responseMapping"`
	CreatedBy        string     `json:"createdBy" bson:"createdBy"`
	UpdatedBy        string     `json:"updatedBy" bson:"updatedBy"`
}

type RequestUpdateRequest struct {
	Name             string     `json:"name" bson:"name"`
	ApplicationID    string     `json:"applicationID" bson:"applicationID"`
	ObjectType       string     `json:"objectType" bson:"objectType"`
	Direction        string     `json:"direction" bson:"direction"`
	LiquidTemplate   string     `json:"liquidTemplate" bson:"LiquidTemplate"`
	LiquidSampleData string     `json:"liquidSampleData" bson:"liquidSampleData"`
	RequestType      string     `json:"requestType" bson:"requestType"`
	SFTP             SFTPConfig `json:"sftp" bson:"sftp"`
	API              APIConfig  `json:"api" bson:"api"`
	OAuth            APIConfig  `json:"oAuth" bson:"oAuth"`
	SQL              SQLConfig  `json:"sql" bson:"sql"`
	ResponseMapping  string     `json:"responseMapping" bson:"responseMapping"`
	CreatedBy        string     `json:"createdBy" bson:"createdBy"`
	UpdatedBy        string     `json:"updatedBy" bson:"updatedBy"`
}

func (req *RequestRequest) Validate() (err error) {
	//if !ifr.Source.Method.isValid() {
	//	return errors.New(fmt.Sprintf("Method [%s] is invalid ", ifr.Source.Method))
	//}

	if !req.API.HMAC.Algorithm.isValidType() && req.API.AuthenticationType == AuthenticationTypeHmac {
		return errors.New(fmt.Sprintf("HMAC Algorithm [%s] is invalid ", req.API.HMAC.Algorithm))
	}

	return
}
