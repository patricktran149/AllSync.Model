package Model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type SystemUser struct {
	ID                primitive.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty"`
	TenantID          string                    `json:"tenantID" bson:"tenantID"`
	TenantName        string                    `json:"tenantName" bson:"tenantName"`
	SecurityRoleID    string                    `json:"securityRoleID" bson:"securityRoleID"`
	Email             string                    `json:"email" bson:"email"`
	Password          string                    `json:"password,omitempty" bson:"password,omitempty"`
	TokenList         []Token                   `json:"tokenList" bson:"tokenList"`
	IsActive          bool                      `json:"isActive" bson:"isActive"`
	IsAdmin           bool                      `json:"isAdmin" bson:"isAdmin"`
	OTP               Token                     `json:"otp"  bson:"otp"`
	RandomKey         string                    `json:"randomKey" bson:"randomKey"`
	RandomExpiredDate int64                     `json:"randomExpiredDate" bson:"randomExpiredDate"`
	FCMs              []string                  `json:"fcms" bson:"fcms"`
	TwoFA             bool                      `json:"2fa" bson:"2fa"`
	TwoFAType         UserSecFAType             `json:"2faType" bson:"2faType"`
	SecretCode        string                    `json:"secretCode" bson:"secretCode"`
	Configurations    []SystemUserConfiguration `json:"configurations" bson:"configurations"`
	CreatedDate       int64                     `json:"createdDate" bson:"createdDate"`
	CreatedBy         string                    `json:"createdBy" bson:"createdBy"`
	UpdatedDate       int64                     `json:"updatedDate" bson:"updatedDate"`
	UpdatedBy         string                    `json:"updatedBy" bson:"updatedBy"`
}

type SystemUserRequest struct {
	ID             primitive.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty"`
	TenantID       string                    `json:"tenantID" bson:"tenantID"`
	TenantName     string                    `json:"tenantName" bson:"tenantName"`
	SecurityRoleID string                    `json:"securityRoleID" bson:"securityRoleID"`
	Email          string                    `json:"email" bson:"email"`
	Password       string                    `json:"password" bson:"password"`
	IsActive       bool                      `json:"isActive" bson:"isActive"`
	IsAdmin        bool                      `json:"isAdmin" bson:"isAdmin"`
	TwoFA          bool                      `json:"2fa" bson:"2fa"`
	TwoFAType      UserSecFAType             `json:"2faType" bson:"2faType"`
	SecretCode     string                    `json:"secretCode" bson:"secretCode"`
	Configurations []SystemUserConfiguration `json:"configurations" bson:"configurations"`
	CreatedBy      string                    `json:"createdBy" bson:"createdBy"`
	UpdatedBy      string                    `json:"updatedBy" bson:"updatedBy"`
}

type SystemUserConfiguration struct {
	ConfigurationID    string `json:"configurationID" bson:"configurationID"`
	ConfigurationValue string `json:"configurationValue" bson:"configurationValue"`
	Description        string `json:"description" bson:"description"`
	IsActive           bool   `json:"isActive" bson:"isActive"`
	Group              string `json:"group" bson:"group"`
}

type SystemUserConfigurationUpdateRequest struct {
	ConfigurationID    string `json:"configurationID" bson:"configurationID"`
	ConfigurationValue string `json:"configurationValue" bson:"configurationValue"`
	Description        string `json:"description" bson:"description"`
	IsActive           bool   `json:"isActive" bson:"isActive"`
	Group              string `json:"group" bson:"group"`
}

type SystemUserUpdateRequest struct {
	Password       string                                 `json:"password" bson:"password"`
	IsActive       bool                                   `json:"isActive" bson:"isActive"`
	IsAdmin        bool                                   `json:"isAdmin" bson:"isAdmin"`
	SecurityRoleID string                                 `json:"securityRoleID" bson:"securityRoleID"`
	TwoFA          bool                                   `json:"2fa" bson:"2fa"`
	TwoFAType      UserSecFAType                          `json:"2faType" bson:"2faType"`
	SecretCode     string                                 `json:"secretCode" bson:"secretCode"`
	Configurations []SystemUserConfigurationUpdateRequest `json:"configurations" bson:"configurations"`
}

type SystemUserSaveFCM struct {
	Email    string `json:"email"`
	TenantID string `json:"tenantID"`
	FCM      string `json:"fcm"`
}

type SystemUserChangePassword struct {
	//Email           string `json:"email"`
	//TenantID        string `json:"tenantID"`
	ID              primitive.ObjectID `json:"_id"`
	CurrentPassword string             `json:"currentPassword"`
	NewPassword     string             `json:"newPassword"`
}

type SystemUserForgotPassword struct {
	Email    string `json:"email"`
	TenantID string `json:"tenantID"`
}

type SystemUserResetPassword struct {
	Email       string `json:"email"`
	TenantID    string `json:"tenantID"`
	NewPassword string `json:"newPassword"`
	RandomKey   string `json:"randomKey"`
}

type SystemUserLogin struct {
	TenantID string `json:"tenantID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SystemUserOTP struct {
	TenantID string `json:"tenantID"`
	Email    string `json:"email"`
	OTP      string `json:"otp"`
}

type UserSecFAType string

const (
	SecFATypeAPP   UserSecFAType = "APP"
	SecFATypeEmail UserSecFAType = "EMAIL"
)

type Token struct {
	ID         string `json:"id" bson:"id"`
	ExpireDate int64  `json:"expireDate" bson:"expireDate"`
}

func (user SystemUser) GetConfigValue(id string) string {
	for _, config := range user.Configurations {
		if strings.TrimSpace(strings.ToUpper(id)) == strings.TrimSpace(strings.ToUpper(config.ConfigurationID)) {
			return config.ConfigurationValue
		}
	}
	return ""
}

func (user SystemUser) GetConfigValueByGroup(id, group string) string {
	for _, config := range user.Configurations {
		if strings.TrimSpace(strings.ToUpper(id)) == strings.TrimSpace(strings.ToUpper(config.ConfigurationID)) &&
			strings.TrimSpace(strings.ToUpper(group)) == strings.TrimSpace(strings.ToUpper(config.Group)) {
			return config.ConfigurationValue
		}
	}
	return ""
}
