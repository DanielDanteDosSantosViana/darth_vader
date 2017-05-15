package models

type Filter struct {
	Type         string `toml:"type"`
	S3           bool   `toml:"s3"`
	Email        bool   `toml:"email"`
	Status       bool   `toml:"status"`
	TemplateMail string `toml:"templatemail"`
	FilePath     string `toml:"filePath"`
}

func NewFilter(s3, email, status bool, templateMail string, types string, filePath string) *Filter {
	return &Filter{types, s3, email, status, templateMail, filePath}
}
