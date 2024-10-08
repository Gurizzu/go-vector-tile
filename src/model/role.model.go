package model

import (
	"git.blackeye.id/Aldi.Rismawan/centrotil/db/umongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generated by https://quicktype.io

type Role struct {
	MetadataWithID `bson:",inline"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Privileges  string `json:"privileges" bson:"privileges"`
}

type Role_View struct {
	Role `bson:",inline"`
}

type Role_Search struct {
	//? Regex
	Search string `json:"search"`

	umongo.Request
}

func (o *Role_Search) HandleFilter(listFilterAnd *[]bson.M) {
	if search := o.Search; search != "" {
		*listFilterAnd = append(*listFilterAnd, bson.M{"name": primitive.Regex{Pattern: search, Options: "i"}})
	}
}
