package compress

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/isbasic/gin_play/common"
)

func GetDirBase(fp string) (string, error) {
	if common.Exist(fp) {
		abs, err := filepath.Abs(fp)
		if err != nil {
			return "", err
		}
		dir := filepath.Dir(fp)
		base := filepath.Base(dir)
		return base, err
	}

	return "", nil
}

func GetCompressName(fp string) (string, error) {
	var e error
	var res string
	if Exist(fp) {
		t := GetFileTime(fp)
		tail := FormatTime(t) + ".7z"
		dir := filepath.Dir(fp)
		name, err := GetDirBase(fp)
		if err != nil {
			return "", err
		}
		filename := fmt.Sprintf("%s%s%s%s%s", dir, common.GetSep(), name, "_", tail)
		return filename, err
	} else {
		return "", errors.New("File not exist.")
	}
}
