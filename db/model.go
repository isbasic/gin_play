package db

import (
	"encoding/json"
	"os"
	"time"

	"github.com/golang/glog"
)

type BIN_TEST struct {
	Sampleid uint   `gorm:"primaryKey;column:"sampleid"`
	BData    []byte `gorm:"column:bData"`
}

type FileList struct {
	Id           string `gorm:"primaryKey;column:"id"`
	FileName     string `gorm:"unique"`
	FileData     string
	FileSize     uint64
	FileSizeUnit string
	CreateAt     time.Time `gorm:"autoUpdateTime:nano"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:nano"`
}

func (fl *FileList) HasDB() bool {
	return false
}

func JsonMarshal(src []byte, prefix string, hasIdent bool, ident string) ([]byte, error) {
	if hasIdent {
		res, err := json.MarshalIndent(src, prefix, ident)
		if err != nil {
			glog.Errorf("MarshalIdent Error: %s", err)
		}
		return res, err
	} else {
		res, err := json.Marshal(src)
		if err != nil {
			glog.Errorf("Marshal Error: %s", err)
		}
		return res, err
	}
}

func GraphJPG(b []byte) *os.File {
	// buf := bytes.NewBuffer(b)
	f, _ := os.CreateTemp(".", "s")

	defer os.Remove(f.Name())
	defer f.Close()
	return f
}
