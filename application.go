package Model

import (
	"errors"
	helper "github.com/patricktran149/Helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type Application struct {
	ID                  primitive.ObjectID         `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID       string                     `json:"applicationID" bson:"applicationID"`
	ApplicationName     string                     `json:"applicationName" bson:"applicationName"`
	Token               Token                      `json:"token" bson:"token"`
	ImmediateRetryTimes int                        `json:"immediateRetryTimes" bson:"immediateRetryTimes"`
	AutoRetry           bool                       `json:"autoRetry" bson:"autoRetry"`
	CreatedDate         int64                      `json:"createdDate" bson:"createdDate"`
	CreatedBy           string                     `json:"createdBy" bson:"createdBy"`
	UpdatedDate         int64                      `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy           string                     `json:"updatedBy" bson:"updatedBy"`
	Configurations      []ApplicationConfiguration `json:"configurations" bson:"configurations"`
	Mappings            []ApplicationMapping       `json:"mappings" bson:"mappings"`
}

type ApplicationRequest struct {
	ID                  primitive.ObjectID         `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID       string                     `json:"applicationID" bson:"applicationID"`
	ApplicationName     string                     `json:"applicationName" bson:"applicationName"`
	Token               Token                      `json:"token" bson:"token"`
	ImmediateRetryTimes int                        `json:"immediateRetryTimes" bson:"immediateRetryTimes"`
	AutoRetry           bool                       `json:"autoRetry" bson:"autoRetry"`
	CreatedDate         int64                      `json:"createdDate" bson:"createdDate"`
	CreatedBy           string                     `json:"createdBy" bson:"createdBy"`
	UpdatedDate         int64                      `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy           string                     `json:"updatedBy" bson:"updatedBy"`
	Configurations      []ApplicationConfiguration `json:"configurations,omitempty" bson:"configurations,omitempty"`
	Mappings            []ApplicationMapping       `json:"mappings" bson:"mappings"`
}

type ApplicationUpdateRequest struct {
	ID                  primitive.ObjectID                      `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID       string                                  `json:"applicationID" bson:"applicationID"`
	ApplicationName     string                                  `json:"applicationName" bson:"applicationName"`
	Token               Token                                   `json:"token" bson:"token"`
	ImmediateRetryTimes int                                     `json:"immediateRetryTimes" bson:"immediateRetryTimes"`
	AutoRetry           bool                                    `json:"autoRetry" bson:"autoRetry"`
	UpdatedBy           string                                  `json:"updatedBy" bson:"updatedBy"`
	UpdatedDate         int64                                   `json:"updatedDate" bson:"updatedDate"`
	Configurations      []ApplicationConfigurationUpdateRequest `json:"configurations,omitempty" bson:"configurations,omitempty"`
	Mappings            []ApplicationMapping                    `json:"mappings" bson:"mappings"`
}

type ApplicationMapping struct {
	TableName string                      `json:"tableName" bson:"tableName"`
	Mapping   []map[string]interface{}    `json:"mapping" bson:"mapping"`
	SFTP      ApplicationObjectSFTPConfig `json:"sftp" bson:"sftp"`
}

type ApplicationConfiguration struct {
	ConfigurationID    string `json:"configurationID" bson:"configurationID"`
	ConfigurationValue string `json:"configurationValue" bson:"configurationValue"`
	Description        string `json:"description" bson:"description"`
	IsActive           bool   `json:"isActive" bson:"isActive"`
	Group              string `json:"group" bson:"group"`
}

type ApplicationConfigurationUpdateRequest struct {
	ConfigurationID    string `json:"configurationID" bson:"configurationID"`
	ConfigurationValue string `json:"configurationValue" bson:"configurationValue"`
	Description        string `json:"description" bson:"description"`
	IsActive           bool   `json:"isActive" bson:"isActive"`
	Group              string `json:"group" bson:"group"`
}

type ApplicationObject struct {
	//ID       primitive.ObjectID          `json:"_id" bson:"_id"`
	ObjectID string                      `json:"objectID" bson:"objectID"`
	Mapping  ApplicationObjectMapping    `json:"mapping" bson:"mapping"`
	Queue    ApplicationObjectQueue      `json:"queue" bson:"queue"`
	WebHook  ApplicationObjectWebHook    `json:"webhook" bson:"webhook"`
	SFTP     ApplicationObjectSFTPConfig `json:"sftp" bson:"sftp"`
}

type ApplicationObjectMapping struct {
	In  []bson.M `json:"in" bson:"in"`
	Out []bson.M `json:"out" bson:"out"`
}

type ApplicationObjectQueue struct {
	Enable bool `json:"enable" bson:"enable"`
}

type ApplicationObjectUpdateRequest struct {
	//ID      primitive.ObjectID          `json:"_id" bson:"_id"`
	Queue   ApplicationObjectQueue      `json:"queue" bson:"queue"`
	WebHook ApplicationObjectWebHook    `json:"webhook" bson:"webhook"`
	SFTP    ApplicationObjectSFTPConfig `json:"sftp" bson:"sftp"`
}

type ApplicationObjectWebHook struct {
	URL                string             `json:"url" bson:"url"`
	AuthenticationType AuthenticationType `json:"authenticationType" bson:"authenticationType"`
	Token              string             `json:"token" bson:"token"`
	Method             string             `json:"method" bson:"method"`
	BasicAuth          BasicAuth          `json:"basicAuth" bson:"basicAuth"`
	IsMandatory        bool               `json:"isMandatory" bson:"isMandatory"`
	Params             bson.M             `json:"params" bson:"params"`
	Headers            bson.M             `json:"headers" bson:"headers"`
	Enable             bool               `json:"enable" bson:"enable"`
}

type ApplicationObjectSFTPConfig struct {
	Host     string `json:"host" bson:"host"`
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
	Port     int    `json:"port" bson:"port"`
	Path     string `json:"path" bson:"path"`
	Enable   bool   `json:"enable" bson:"enable"`
}

type BasicAuth struct {
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
}

func (appReq *ApplicationRequest) Validate() (err error) {

	return
}

func (app *Application) GetObjectByID(id string) (mapping ApplicationMapping) {
	for _, m := range app.Mappings {
		if m.TableName == id {
			mapping = m
			break
		}
	}

	return mapping
}

func (hook *ApplicationObjectWebHook) Validate() (err error) {
	if !hook.Enable {
		return
	}

	//Validate Authentication Type
	if !hook.AuthenticationType.isValidType() {
		return errors.New("Invalid Authentication Type! ")
	}

	//Validate method
	hook.Method = strings.ToUpper(strings.TrimSpace(hook.Method))
	if !helper.IsItemExistsInArray(hook.Method, []string{http.MethodPost, http.MethodPut, http.MethodPatch}) {
		return errors.New("Invalid Method! ")
	}

	return
}

func (app *Application) GetConfigValue(id string) string {
	for _, config := range app.Configurations {
		if strings.TrimSpace(strings.ToUpper(id)) == strings.TrimSpace(strings.ToUpper(config.ConfigurationID)) {
			return config.ConfigurationValue
		}
	}
	return ""
}

func (app *Application) UpdateConfigValue(group, id, updateValue string) {
	isExists := false
	for i, config := range app.Configurations {
		if strings.TrimSpace(strings.ToUpper(id)) == strings.TrimSpace(strings.ToUpper(config.ConfigurationID)) {
			app.Configurations[i].ConfigurationValue = updateValue

			isExists = true
		}
	}

	if !isExists {
		app.Configurations = append(app.Configurations, ApplicationConfiguration{
			ConfigurationID:    id,
			ConfigurationValue: updateValue,
			Description:        "",
			IsActive:           true,
			Group:              group,
		})
	}

	return
}

func (app *Application) GetConfig(id string) (appConfig ApplicationConfiguration) {
	for _, config := range app.Configurations {
		if strings.ToUpper(id) == strings.ToUpper(config.ConfigurationID) {
			return config
		}
	}
	return appConfig
}

/*func (object ApplicationObject) MakeRequest(data interface{}) ( err error) {
	var (
		hook       = object.WebHook
		k          *kazaam.Kazaam
		specConfig string
		client     = &http.Client{}
		query      = url.Values{}
		req        *http.Request
		res        *http.Response
		sendData, responseData string
	)

	var bt []byte

	rawData := helper.JSONToString(data)

	specConfig = helper.JSONToString(object.Mapping.Out)

	k, err = kazaam.NewKazaam(specConfig)
	if err != nil {
		err = errors.New("New Kazaam ERROR - " + err.Error())
		return
	}

	sendData, err = k.TransformJSONStringToString(rawData)
	if err != nil {
		err = errors.New("Transform ERROR - " + err.Error())
		return
	}

	if hook.Method != http.MethodGet {
		bt = []byte(sendData)
	}

	req, err = http.NewRequest(hook.Method, hook.URL, bytes.NewBuffer(bt))
	if err != nil {
		err = errors.New("Make request ERROR - " + err.Error())
		return
	}

	// Add Params
	for k, v := range hook.Params {
		query.Add(k, fmt.Sprintf("%v", v))
	}
	req.URL.RawQuery = query.Encode()

	req.Headers.Set("Content-Type", "application/json")

	// For authorization header
	switch hook.AuthenticationType {
	case AuthenticationTypeBearer:
		{
			req.Headers.Set("Authorization", fmt.Sprintf("Bearer %s", hook.Token))
			break
		}
	case AuthenticationTypeBasic:
		{
			auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", hook.BasicAuth.UserName, hook.BasicAuth.Password)))
			req.Headers.Set("Authorization", fmt.Sprintf("Basic %s", auth))
			break
		}
	default:
		break
	}

	// For User input header
	for k, v := range hook.Headers {
		if hook.AuthenticationType != AuthenticationTypeNone && k == "Authorization" {
			continue
		}
		req.Headers.Set(k, fmt.Sprintf("%v", v))
	}

	res, err = client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	responseData = string(body)
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		if hook.IsMandatory {
			err = errors.New(fmt.Sprintf("Request ERROR - Status [%v] - Code [%v] - Response [%s]", res.Status, res.StatusCode, string(body)))
			return
		}
	}

	return
}*/
