package models

type User struct{
	Password		string					`json:"Password" validate:"required,min=6"`
	Email			string					`json:"email" validate:"email,required"`
}
