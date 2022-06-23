package models

type TypeEmail string

const (
	CHARSET = "UTF-8"

	HTMLTypeEmail = TypeEmail("html")
	TEXTTypeEmail = TypeEmail("text")
)

type RequestSendEmail struct {
	To      []string `validate:"min=1,dive,email"`
	Cc      []string `validate:"dive,email"`
	Bcc     []string `validate:"dive,email"`
	From    string   `validate:"required"`
	Subject string   `validate:"required"`
	Body    string   `validate:"required"`
	Type    TypeEmail
}
