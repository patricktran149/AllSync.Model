package Model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Queue struct {
	ID                primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID     string                   `json:"applicationID" bson:"applicationID"`
	ObjectID          string                   `json:"objectID" bson:"objectID"`
	Type              string                   `json:"type" bson:"type"`
	Method            FlowMethod               `json:"method" bson:"method"`
	SFTPConfig        SFTPConfig               `json:"sftpConfig" bson:"sftpConfig"`
	APIConfig         APIConfig                `json:"apiConfig" bson:"apiConfig"`
	SQLConfig         SQLConfig                `json:"sqlConfig" bson:"sqlConfig"`
	Status            QueueStatus              `json:"status" bson:"status"`
	RawData           string                   `json:"rawData" bson:"rawData"`
	SendData          string                   `json:"sendData" bson:"sendData"`
	ResponseData      string                   `json:"responseData" bson:"responseData"`
	Message           string                   `json:"message" bson:"message"`
	Mapping           []map[string]interface{} `json:"mapping" bson:"mapping"`
	Ref1              string                   `json:"ref1" bson:"ref1"`
	Ref2              string                   `json:"ref2" bson:"ref2"`
	Ref3              string                   `json:"ref3" bson:"ref3"`
	Ref4              string                   `json:"ref4" bson:"ref4"`
	Ref5              string                   `json:"ref5" bson:"ref5"`
	Flag              bool                     `json:"flag" bson:"flag"`
	IsSkip            bool                     `json:"isSkip" bson:"isSkip"`
	RetryTimes        int                      `json:"retryTimes" bson:"retryTimes"`
	RetryDays         int                      `json:"retryDays" bson:"retryDays"`
	IntegrationFlowID primitive.ObjectID       `json:"integrationFlowID" bson:"integrationFlowID"`
	Logs              []QueueLog               `json:"logs" bson:"logs"`
	Date              time.Time                `json:"date" json:"date"`
	CreatedDate       int64                    `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
	CreatedBy         string                   `json:"createdBy" bson:"createdBy"`
	UpdatedDate       int64                    `json:"updatedDate,omitempty" bson:"updatedDate,omitempty"`
	UpdatedDateMs     int64                    `json:"updatedDateMs" bson:"updatedDateMs"`
	UpdatedBy         string                   `json:"updatedBy" bson:"updatedBy"`
}

func (q *Queue) CalcRetry(limitTimes, limitDays int) {
	createdDate := time.Unix(q.CreatedDate, 0)

	if !time.Now().Truncate(24 * time.Hour).Equal(createdDate.Truncate(24 * time.Hour)) {
		if q.RetryTimes >= limitTimes && q.RetryDays < limitDays {
			q.RetryTimes = 0
			q.RetryDays += 1
		}
	}
}

type QueueRequest struct {
	//ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID     string                   `json:"applicationID" bson:"applicationID"`
	ObjectID          string                   `json:"objectID" bson:"objectID"`
	Type              string                   `json:"type" bson:"type"`
	Method            FlowMethod               `json:"method" bson:"method"`
	SFTPConfig        SFTPConfig               `json:"sftpConfig" bson:"sftpConfig"`
	APIConfig         APIConfig                `json:"apiConfig" bson:"apiConfig"`
	SQLConfig         SQLConfig                `json:"sqlConfig" bson:"sqlConfig"`
	Status            QueueStatus              `json:"status" bson:"status"`
	RawData           string                   `json:"rawData" bson:"rawData"`
	SendData          string                   `json:"sendData" bson:"sendData"`
	ResponseData      string                   `json:"responseData" bson:"responseData"`
	Message           string                   `json:"message" bson:"message"`
	Mapping           []map[string]interface{} `json:"mapping" bson:"mapping"`
	Ref1              string                   `json:"ref1" bson:"ref1"`
	Ref2              string                   `json:"ref2" bson:"ref2"`
	Ref3              string                   `json:"ref3" bson:"ref3"`
	Ref4              string                   `json:"ref4" bson:"ref4"`
	Ref5              string                   `json:"ref5" bson:"ref5"`
	Flag              bool                     `json:"flag" bson:"flag"`
	IsSkip            bool                     `json:"isSkip" bson:"isSkip"`
	RetryTimes        int                      `json:"retryTimes" bson:"retryTimes"`
	RetryDays         int                      `json:"retryDays" bson:"retryDays"`
	IntegrationFlowID primitive.ObjectID       `json:"integrationFlowID" bson:"integrationFlowID"`
	Logs              []QueueLog               `json:"logs" bson:"logs"`
	CreatedBy         string                   `json:"createdBy" bson:"createdBy"`
	UpdatedBy         string                   `json:"updatedBy" bson:"updatedBy"`
}

type QueueUpdateRequest struct {
	ApplicationID     string                   `json:"applicationID" bson:"applicationID"`
	ObjectID          string                   `json:"objectID" bson:"objectID"`
	Type              string                   `json:"type" bson:"type"`
	Method            FlowMethod               `json:"method" bson:"method"`
	SFTPConfig        SFTPConfig               `json:"sftpConfig" bson:"sftpConfig"`
	APIConfig         APIConfig                `json:"apiConfig" bson:"apiConfig"`
	SQLConfig         SQLConfig                `json:"sqlConfig" bson:"sqlConfig"`
	Status            string                   `json:"status" bson:"status"`
	RawData           string                   `json:"rawData" bson:"rawData"`
	SendData          string                   `json:"sendData" bson:"sendData"`
	ResponseData      string                   `json:"responseData" bson:"responseData"`
	Message           string                   `json:"message" bson:"message"`
	Mapping           []map[string]interface{} `json:"mapping" bson:"mapping"`
	Ref1              string                   `json:"ref1" bson:"ref1"`
	Ref2              string                   `json:"ref2" bson:"ref2"`
	Ref3              string                   `json:"ref3" bson:"ref3"`
	Ref4              string                   `json:"ref4" bson:"ref4"`
	Ref5              string                   `json:"ref5" bson:"ref5"`
	Flag              bool                     `json:"flag" bson:"flag"`
	IsSkip            bool                     `json:"isSkip" bson:"isSkip"`
	RetryTimes        int                      `json:"retryTimes" bson:"retryTimes"`
	RetryDays         int                      `json:"retryDays" bson:"retryDays"`
	IntegrationFlowID primitive.ObjectID       `json:"integrationFlowID" bson:"integrationFlowID"`
	Logs              []QueueLog               `json:"logs" bson:"logs"`
	CreatedBy         string                   `json:"createdBy" bson:"createdBy"`
	UpdatedBy         string                   `json:"updatedBy" bson:"updatedBy"`
}

func (qReq QueueRequest) Validate() error {
	if qReq.ApplicationID == "" {
		return errors.New("Application ID is empty ")
	}

	if qReq.ObjectID == "" {
		return errors.New("Object ID is empty ")
	}

	if qReq.Type == "" {
		return errors.New("Type is empty ")
	}

	if qReq.RawData == "" {
		return errors.New("Raw Data is empty ")
	}
	return nil
}

type QueueStatus string

const (
	QueueStatusSuccess QueueStatus = "SUCCESS"
	QueueStatusPending QueueStatus = "PENDING"
	QueueStatusFailed  QueueStatus = "FAILED"
)

type ReportQueueStatus struct {
	ApplicationID string `json:"applicationID"`
	IncomingQueue int64  `json:"incomingQueue"`
	OutgoingQueue int64  `json:"outgoingQueue"`
	Status        string `json:"status"`
}

type QueueObject struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Direction     string             `json:"direction"`
	ApplicationID string             `json:"applicationID"`
	CreatedDate   int64              `json:"createdDate"`
	UpdatedDate   int64              `json:"updatedDate"`
	Status        QueueStatus        `json:"status"`
	Message       string             `json:"message"`
}

type QueueLog struct {
	FunctionName string `json:"functionName" bson:"functionName"`
	Log1         string `json:"log1" bson:"log1"`
	Log2         string `json:"log2" bson:"log2"`
	Log3         string `json:"log3" bson:"log3"`
	LogDate      int64  `json:"logDate" bson:"logDate"`
}

func (q *Queue) Log(funcName, log1, log2, log3 string) {
	q.Logs = append(q.Logs, QueueLog{
		FunctionName: funcName,
		Log1:         log1,
		Log2:         log2,
		Log3:         log3,
		LogDate:      time.Now().Unix(),
	})
}

func (q *Queue) AddLogs(logs []QueueLog) {
	q.Logs = append(q.Logs, logs...)
}
