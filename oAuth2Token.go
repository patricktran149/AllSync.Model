package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OAuth2Token struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID     string             `json:"applicationID" bson:"applicationID"`
	Token             Token              `json:"token" bson:"token"`
	FreshToken        Token              `json:"freshToken" bson:"freshToken"`
	TokenExpiryEmail  string             `json:"tokenExpiryEmail" bson:"tokenExpiryEmail"`
	LastSentEmailDate int64              `json:"lastSentEmailDate" bson:"lastSentEmailDate"`
}

type OAuth2TokenRequest struct {
	ApplicationID    string `json:"applicationID" bson:"applicationID"`
	Token            Token  `json:"token" bson:"token"`
	FreshToken       Token  `json:"freshToken" bson:"freshToken"`
	TokenExpiryEmail string `json:"tokenExpiryEmail" bson:"tokenExpiryEmail"`
}

type OAuth2TokenUpdateRequest struct {
	ApplicationID     string `json:"applicationID" bson:"applicationID"`
	Token             Token  `json:"token" bson:"token"`
	FreshToken        Token  `json:"freshToken" bson:"freshToken"`
	TokenExpiryEmail  string `json:"tokenExpiryEmail" bson:"tokenExpiryEmail"`
	LastSentEmailDate int64  `json:"lastSentEmailDate" bson:"lastSentEmailDate"`
}
