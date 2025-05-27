package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email  string				`json:"email" bson:"email"`
	AiKey    string              `bson:"ai_key"`
	AiType   string             `bson:"ai_type"`
}

type Chat struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title    string				`json:"title" bson:"title"`
}


type ChatMessage struct {
    Sender    string `json:"sender"`
    Message   string `json:"message"`
    // Timestamp string `json:"timestamp"`
}