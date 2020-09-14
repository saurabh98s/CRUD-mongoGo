package models

import "github.com/globalsign/mgo/bson"

//User maps the data which document should hold
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	CreatedAt int64         `json:"created_at" bson:"created_at"`
	UpdateAt  int64         `json:"update_at" bson:"updated_at"`
	DeletedAt int64         `json:"deleted_at" bson:"deleted_at"`
}
