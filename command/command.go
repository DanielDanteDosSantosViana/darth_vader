package command

import (
	"github.com/DanielDanteDosSantosViana/darth_vader/aws"
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/DanielDanteDosSantosViana/darth_vader/email"
	"github.com/DanielDanteDosSantosViana/darth_vader/models"
	"github.com/DanielDanteDosSantosViana/darth_vader/reader"
)

type Command struct {
	clientAWS *aws.ClientAWS
}

func NewCommand(clientAWS *aws.ClientAWS) *Command {
	return &Command{clientAWS}
}
func (cmd *Command) Exec(filter *models.Filter) {
	r := reader.NewReader(filter.FilePath)
	r.Read()
	if filter.S3 {
		params := aws.NewParams(config.Conf.Bucket.Name, r.FilePath, r.Size, r.FileByte, r.FileType)
		cmd.clientAWS.SendToS3(params)
	}
	if filter.Email {
		if filter.Status {
			email.SendPersonlization(r.Status, filter.TemplateMail)
		} else {
			email.Send(r.Status)
		}
	}

}
