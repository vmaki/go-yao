package job

type SendSMSPayload struct {
	Phone string
	Code  int
}

type SendEmailPayload struct {
	Email   string
	Content string
}
