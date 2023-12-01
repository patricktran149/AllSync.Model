package Model

import "html/template"

type SendEmail struct {
	Subject   string                 `json:"subject"`
	Variables map[string]interface{} `json:"variables"`
	ToEmails  []string               `json:"toEmails"`
	Data      template.HTML          `json:"data"`
}

type CallOutboundAPIConfig struct {
	Url         string            `json:"url"`
	Method      string            `json:"method"`
	Params      map[string]string `json:"params"`
	Headers     map[string]string `json:"headers"`
	Payload     string            `json:"payload"`
	ContentType string            `json:"contentType"`
}
