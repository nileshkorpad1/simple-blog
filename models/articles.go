package models

import (
	"github.com/nileshkorpad1/simple-blog/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// connect db
var collection = config.ConnectDB()

//Article fields
type Article struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" example:"Remove Blog ID"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty" validate:"required" example:"Blog Title"`
	Content string             `json:"content,omitempty" bson:"content,omitempty" validate:"required" example:"Blog Content"`
	Author  string             `json:"author,omitempty" bson:"author,omitempty" validate:"required" example:"Blog author"`
}

//Articles of article
type Articles []Article
