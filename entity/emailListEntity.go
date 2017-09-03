package entity

type EmailList struct {
	Emails []string `json:"emails" binding:"required"`
}
