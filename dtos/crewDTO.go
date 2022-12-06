package dtos

type CrewDTO struct {
	Name string `bson:"name,omitempty" json:"name"`
	Job  string `bson:"job,omitempty" json:"job"`
}
