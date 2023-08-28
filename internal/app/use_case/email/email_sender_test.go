package email_test

import (
	"context"
	"event-calendar/internal/app/domain/schedule/entity"
	"event-calendar/internal/app/use_case/email"
	"testing"
)

type fakeEmailClient struct {
}

func TestEmailGenerator_SendBatchTo(t *testing.T) {
	type args struct {
		ctx      context.Context
		email    []string
		schedule *entity.Schedule
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "failed when get user sender",
			args: args{
				ctx:      context.Background(),
				email:    []string{"testmail@mail.com"},
				schedule: &entity.Schedule{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emailGenerator := email.NewEmailGenerator(new(fakeEmailClient))
			if err := emailGenerator.SendBatchTo(tt.args.ctx, tt.args.email, tt.args.schedule); (err != nil) != tt.wantErr {
				t.Errorf("EmailGenerator.SendBatchTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
