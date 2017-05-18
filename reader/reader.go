package reader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DanielDanteDosSantosViana/darth_vader/models"
)

type Reader struct {
	FilePath string
	FileByte io.ReadSeeker
	FileType string
	Size     int64
	Status   *models.StatusFile
}

func NewReader(filePath string) *Reader {
	return &Reader{filePath, nil, "", 0, &models.StatusFile{}}
}

func (r *Reader) Read() {
	time.Sleep(time.Second * 5)
	readDataToS3(r)
	readStatus(r)

}

func readDataToS3(r *Reader) {
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

	defer file.Close()
	fileInfo, _ = file.Stat()
	r.Size = fileInfo.Size()
	buffer := make([]byte, r.Size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	r.FileByte = fileBytes
	r.FileType = fileType
}

func readStatus(r *Reader) {
	_, err := os.Stat(r.FilePath)
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

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		target := string([]rune(line)[0])
		//6 - total de faturas  e 17 total dinheiro
		if target == "Z" {
			r.setStatusFile(line)
		}
	}
}

func (r *Reader) setStatusFile(line string) {
	status := models.NewStatusFile(r.getTotalInvoice(line), r.getTotalMoney(line), r.FilePath)
	r.Status = status
}

func (r *Reader) getTotalInvoice(line string) string {
	total := line[1:7]
	value, err := strconv.ParseInt(total, 10, 64)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return strconv.FormatInt(value-2, 10)
}
func (r *Reader) getTotalMoney(line string) string {
	total := line[7:24]
	value, err := strconv.ParseFloat(total, 64)
	if err != nil {
		fmt.Print(err)
		return ""
	}

	value = value / 100
	return strconv.FormatFloat(value, 'f', 2, 64)
}
