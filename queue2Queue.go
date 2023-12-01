package Model

import (
	"errors"
	helper "github.com/patricktran149/Helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Queue2Queue struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FromQueue         string             `json:"fromQueue" bson:"fromQueue"`
	FromApplicationID string             `json:"fromApplicationID" bson:"fromApplicationID"`
	FromStatus        string             `json:"fromStatus" bson:"fromStatus"`
	FromType          string             `json:"fromType" bson:"fromType"`
	FromMode          string             `json:"fromMode" bson:"fromMode"`
	ToQueue           string             `json:"toQueue" bson:"toQueue"`
	ToApplicationID   string             `json:"toApplicationID" bson:"toApplicationID"`
	ToStatus          string             `json:"toStatus" bson:"toStatus"`
	ToType            string             `json:"toType" bson:"toType"`
	ToRawDataType     string             `json:"toRawDataType" bson:"toRawDataType"`
	CreatedDate       int64              `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
	CreatedBy         string             `json:"createdBy" bson:"createdBy"`
	UpdatedDate       int64              `json:"updatedDate,omitempty" bson:"updatedDate,omitempty"`
	UpdatedBy         string             `json:"updatedBy" bson:"updatedBy"`
}

type Queue2QueueRequest struct {
	//ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FromQueue         string `json:"fromQueue" bson:"fromQueue"`
	FromApplicationID string `json:"fromApplicationID" bson:"fromApplicationID"`
	FromStatus        string `json:"fromStatus" bson:"fromStatus"`
	FromType          string `json:"fromType" bson:"fromType"`
	FromMode          string `json:"fromMode" bson:"fromMode"`
	ToQueue           string `json:"toQueue" bson:"toQueue"`
	ToApplicationID   string `json:"toApplicationID" bson:"toApplicationID"`
	ToStatus          string `json:"toStatus" bson:"toStatus"`
	ToType            string `json:"toType" bson:"toType"`
	ToRawDataType     string `json:"toRawDataType" bson:"toRawDataType"`
	CreatedBy         string `json:"createdBy" bson:"createdBy"`
	UpdatedBy         string `json:"updatedBy" bson:"updatedBy"`
}

type Queue2QueueUpdateRequest struct {
	FromQueue         string `json:"fromQueue" bson:"fromQueue"`
	FromApplicationID string `json:"fromApplicationID" bson:"fromApplicationID"`
	FromStatus        string `json:"fromStatus" bson:"fromStatus"`
	FromType          string `json:"fromType" bson:"fromType"`
	FromMode          string `json:"fromMode" bson:"fromMode"`
	ToQueue           string `json:"toQueue" bson:"toQueue"`
	ToApplicationID   string `json:"toApplicationID" bson:"toApplicationID"`
	ToStatus          string `json:"toStatus" bson:"toStatus"`
	ToType            string `json:"toType" bson:"toType"`
	ToRawDataType     string `json:"toRawDataType" bson:"toRawDataType"`
	CreatedBy         string `json:"createdBy" bson:"createdBy"`
	UpdatedBy         string `json:"updatedBy" bson:"updatedBy"`
}

func (qReq *Queue2QueueRequest) Validate() error {
	if !helper.IsStringInArray(qReq.FromQueue, []string{"Incoming", "Outgoing"}) {
		return errors.New("Invalid From Queue ")
	}

	qReq.FromQueue = strings.Title(strings.ToLower(qReq.FromQueue))

	if !helper.IsStringInArray(qReq.FromStatus, []string{"PENDING", "FAILED", "SUCCESS"}) {
		return errors.New("Invalid From Status ")
	}

	qReq.FromStatus = strings.ToUpper(qReq.FromStatus)

	if !helper.IsStringInArray(qReq.FromMode, []string{"ADD", "UPDATE", "BOTH"}) {
		return errors.New("Invalid From Mode ")
	}

	qReq.FromMode = strings.ToUpper(qReq.FromMode)

	toRawDataType, ok := helper.IsStringInArrayAndGet(qReq.ToRawDataType, []string{"rawData", "sendData", "responseData"})
	if !ok {
		return errors.New("Invalid To Raw Data Type ")
	}

	qReq.ToRawDataType = toRawDataType

	return nil
}
