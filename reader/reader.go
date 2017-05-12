package reader

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DanielDanteDosSantosViana/darth_vader/aws"
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/DanielDanteDosSantosViana/darth_vader/email"
)

type Reader struct {
	FilePath  string
	FileByte  io.ReadSeeker
	FileType  string
	clientAWS *aws.ClientAWS
}

func NewReader(filePath string, clientAWS *aws.ClientAWS) *Reader {
	return &Reader{filePath, nil, "", clientAWS}
}

func (r *Reader) Read() {
	time.Sleep(time.Second * 5)
	fileInfo, err := os.Stat(r.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("arquivo não existe.")
			return
		}
	}
	file, e := os.Open(r.FilePath)
	if e != nil {
		log.Printf("error ao abrir o arquivo : %s", e)
		return
	}

	email.Send(file.Name())
	defer file.Close()
	fileInfo, _ = file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	r.FileByte = fileBytes
	r.FileType = fileType
	params := aws.NewParams(config.Conf.Bucket.Name, r.FilePath, size, r.FileByte, r.FileType)
	r.clientAWS.SendToS3(params)
}
