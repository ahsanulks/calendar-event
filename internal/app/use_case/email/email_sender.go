package email

import (
	"context"
	"event-calendar/internal/app/constant"
	"event-calendar/internal/app/domain/schedule/entity"
	"fmt"
	"io"
)

type EmailGenerator struct {
	emailClient io.Writer
}

func NewEmailGenerator(emailClient io.Writer) *EmailGenerator {
	return &EmailGenerator{
		emailClient: emailClient,
	}
}

func (emailGenerator EmailGenerator) SendBatchTo(ctx context.Context, email []string, schedule *entity.Schedule) error {
	sender := ctx.Value(constant.ContextEmailKey).(string)
	emailTemplate := "%s invite you Meeting schedule on %s\n" +
		"time: %s\n" +
		"end: %s\n"
	emailBody := fmt.Sprintf(sender, emailTemplate, schedule.Name, schedule.StartTime, schedule.EndTime)
	// better write on message queue then we have background job to send it
	_, err := emailGenerator.emailClient.Write([]byte(emailBody))
	return err
}
