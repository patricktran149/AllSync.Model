package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OAuth2Token struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ApplicationID string             `json:"applicationID" bson:"applicationID"`
	Token         Token              `json:"token" bson:"token"`
	FreshToken    Token              `json:"freshToken" bson:"freshToken"`
}

type OAuth2TokenRequest struct {
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	Token         Token  `json:"token" bson:"token"`
	FreshToken    Token  `json:"freshToken" bson:"freshToken"`
}

type OAuth2TokenUpdateRequest struct {
	ApplicationID string `json:"applicationID" bson:"applicationID"`
	Token         Token  `json:"token" bson:"token"`
	FreshToken    Token  `json:"freshToken" bson:"freshToken"`
}
