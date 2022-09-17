package mock

import "context"

type SendInvitationArgs struct{ Email, Code string }

type Emailer struct {
	SendInvitationFn     func(ctx context.Context, email, code string) error
	SendInvitationCalled struct {
		Count int
		With  SendInvitationArgs
	}
}

func (e *Emailer) SendInvitation(ctx context.Context, email, code string) error {
	if e.SendInvitationFn == nil {
		return mockMethodNotDefined("SendInvitation")
	}
	e.SendInvitationCalled.Count++
	e.SendInvitationCalled.With = SendInvitationArgs{Email: email, Code: code}
	return e.SendInvitationFn(ctx, email, code)
}
