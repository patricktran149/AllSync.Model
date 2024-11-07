package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SystemResource struct {
	ID                      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ServerName              string             `json:"serverName" bson:"serverName"`
	TotalCPUCores           string             `json:"totalCPUCores" bson:"totalCPUCores"`
	TotalCPUCoresPercent    string             `json:"totalCPUCoresPercent" bson:"totalCPUCoresPercent"`
	TotalCPUUSagePercent    string             `json:"totalCPUUSagePercent" bson:"totalCPUUSagePercent"`
	TotalMemory             string             `json:"totalMemory" bson:"totalMemory"`
	TotalMemoryUsage        string             `json:"totalMemoryUsage" bson:"totalMemoryUsage"`
	TotalMemoryUsagePercent string             `json:"totalMemoryUsagePercent" bson:"totalMemoryUsagePercent"`
	UpdatedDate             int64              `json:"updatedDate" bson:"updatedDate"`
	ServiceList             []SystemService    `json:"serviceList" bson:"serviceList"`
}

type SystemResourceRequest struct {
	ServerName              string          `json:"serverName" bson:"serverName"`
	TotalCPUCores           string          `json:"totalCPUCores" bson:"totalCPUCores"`
	TotalCPUCoresPercent    string          `json:"totalCPUCoresPercent" bson:"totalCPUCoresPercent"`
	TotalCPUUSagePercent    string          `json:"totalCPUUSagePercent" bson:"totalCPUUSagePercent"`
	TotalMemory             string          `json:"totalMemory" bson:"totalMemory"`
	TotalMemoryUsage        string          `json:"totalMemoryUsage" bson:"totalMemoryUsage"`
	TotalMemoryUsagePercent string          `json:"totalMemoryUsagePercent" bson:"totalMemoryUsagePercent"`
	ServiceList             []SystemService `json:"serviceList" bson:"serviceList"`
}

type SystemService struct {
	Name               string `json:"name" bson:"name"`
	PID                string `json:"pid" bson:"pid"`
	CPUUsagePercent    string `json:"cpuUsagePercent" bson:"cpuUsagePercent"`
	MemoryUsagePercent string `json:"memoryUsagePercent" bson:"memoryUsagePercent"`
}
