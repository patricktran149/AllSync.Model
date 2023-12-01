package Model

type IntegrationLog struct {
	ServiceName  string `json:"serviceName"`
	FunctionName string `json:"functionName"`
	Log1         string `json:"log1"`
	Log2         string `json:"log2"`
	Log3         string `json:"log3"`
	LogDate      int64  `json:"logDate"`
}

type SocketMsg struct {
	TenantID       string         `json:"tenantID"`
	Token          string         `json:"token"`
	IntegrationLog IntegrationLog `json:"integrationLog"`
}
