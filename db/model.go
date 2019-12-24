package db

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type Error struct {
	Code int
	Msg  string
}

type StudentUpdateRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

type StudentSearchRequest struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
}
