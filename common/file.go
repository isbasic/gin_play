package common

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func Exist(fp string) bool {
	_, err := os.Stat(fp)

	return os.IsExist(err)
}

func GetFileTime(fp string) (time.Time, error) {
	if !Exist(fp) {
		return time.Now(), errors.New("File not exists, please check %s is valid.")
	}

	ff, err := os.Stat(fp)
	if err != nil {
		return time.Now(), err
	}

	res := ff.ModTime()
	return res, err
}

func GetSep(osName string) string {
	if osName != "" {
		if osName == "windows" {
			return "\\"
		} else {
			return "/"
		}
	} else {
		o := runtime.GOOS
		return GetSep(o)
	}
}

func DirScan(fp string) (map[string]string, error) {
	var fl map[string]string
	var e error

	if Exist(fp) {
		abs, err := filepath.Abs(fp)
		fs := os.DirFS(fp)
		filepath.WalkDir(fs,".",FWalkDir(path string,d os.DirEntry,e error){
			if err != nil {
            log.Fatal(err)
        }
        fmt.Println(path)
        return nil
		})
	}

	return fl, e
}

