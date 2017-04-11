package reader

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/DanielDanteDosSantosViana/darth_vader/aws"
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
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
func (r *Reader) read() {
	fmt.Printf("LEU")
	time.Sleep(time.Second * 3)
	fileInfo, err := os.Stat(r.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File does not exist.")
			return
		}
	}
	file, e := os.Open(r.FilePath)
	if e != nil {
		fmt.Printf("err opening file: %s", e)
	}
	defer file.Close()
	fileInfo, _ = file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	r.FileByte = fileBytes
	r.FileType = fileType
	fmt.Printf("LEU-2")
	params := aws.NewParams(config.Conf.Bucket.Name, r.FilePath, size, r.FileByte, r.FileType)
	r.clientAWS.SendToS3(params)
}
