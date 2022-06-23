package v2

import (
	"context"
	"flip/aws-sdk-go-poc/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/go-playground/validator"
)

func SendEmail(request models.RequestSendEmail) error {

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return err
	}

	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		return err
	}

	bodyEmail := &types.Body{}
	if request.Type == models.HTMLTypeEmail {
		bodyEmail.Html = &types.Content{
			Charset: aws.String(models.CHARSET),
			Data:    aws.String(request.Body),
		}
	} else if request.Type == models.TEXTTypeEmail {
		bodyEmail.Text = &types.Content{
			Charset: aws.String(models.CHARSET),
			Data:    aws.String(request.Body),
		}
	}

	sesClient := ses.NewFromConfig(awsConfig)
	if _, err = sesClient.SendEmail(context.TODO(), &ses.SendEmailInput{
		Source: aws.String(request.From),
		Destination: &types.Destination{
			ToAddresses:  request.To,
			CcAddresses:  request.Cc,
			BccAddresses: request.Bcc,
		},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String(models.CHARSET),
				Data:    aws.String(request.Subject),
			},
			Body: bodyEmail,
		},
	}); err != nil {
		return err
	}

	return nil
}
