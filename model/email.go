package model

type EmailRequest struct {
	Subject   string `json:"subject"`
	UserEmail string `json:"user_email"`
	Message   string `json:"message"`
}
