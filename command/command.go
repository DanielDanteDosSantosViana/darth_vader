package command

import (
	"github.com/DanielDanteDosSantosViana/darth_vader/aws"
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/DanielDanteDosSantosViana/darth_vader/email"
	"github.com/DanielDanteDosSantosViana/darth_vader/models"
	"github.com/DanielDanteDosSantosViana/darth_vader/reader"
	"path"
	"log"
	"strings"
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
	fileName:= r.FilePath
	fileName = strings.TrimPrefix(fileName,path.Dir(r.FilePath)+"/")
	file:= models.NewFile(fileName,path.Dir(r.FilePath), r.Status)
	if filter.S3 {
		params := aws.NewParams(config.Conf.Bucket.Name, r.FilePath, r.Size, r.FileByte, r.FileType)
		cmd.clientAWS.SendToS3(params)
	}
	if filter.Email {
		if filter.Status {
			email.SendPersonlization(file, filter.TemplateMail)
		} else {
			email.Send(file)
		}
	}
	fileModel :=models.NewFileModel()
	err:= fileModel.Create(file)
	if err!=nil{
	  log.Println("Error ao persistir o arquivo :  %s",err)
	  return
	}
}
