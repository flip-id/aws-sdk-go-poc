package v1

import (
	"flip/aws-sdk-go-poc/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/go-playground/validator"
)

func SendEmail(request models.RequestSendEmail) error {

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return err
	}

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String("ap-southeast-1")})
	if err != nil {
		return err
	}

	bodyEmail := &ses.Body{}
	if request.Type == models.HTMLTypeEmail {
		bodyEmail.Html = &ses.Content{
			Charset: aws.String(models.CHARSET),
			Data:    aws.String(request.Body),
		}
	} else if request.Type == models.TEXTTypeEmail {
		bodyEmail.Text = &ses.Content{
			Charset: aws.String(models.CHARSET),
			Data:    aws.String(request.Body),
		}
	}

	sesClient := ses.New(awsSession)
	sesClient.SendEmail(&ses.SendEmailInput{
		Source: aws.String(request.From),
		Destination: &ses.Destination{
			ToAddresses:  convertStrings(request.To),
			CcAddresses:  convertStrings(request.Cc),
			BccAddresses: convertStrings(request.Bcc),
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Charset: aws.String(models.CHARSET),
				Data:    aws.String(request.Subject),
			},
			Body: bodyEmail,
		},
	})

	return nil
}

func convertStrings(req []string) (res []*string) {

	for _, i := range req {
		res = append(res, aws.String(i))
	}

	return
}
