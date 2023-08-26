package common

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	_ "github.com/google/uuid"
	"github.com/pborman/uuid"
)

const FP = "conf/conf.json"

type ConfigError struct {
	name string
	info string
	err  error
}

func (ce *ConfigError) Set(name, info string) {
	ce.info = info
	ce.name = name
	ce.err = errors.New(info)
}

func (ce *ConfigError) Error() string {
	return fmt.Sprintf("%s Error: %s", ce.name, ce.info)
}

func Exists(fp string) bool {
	_, err := os.Stat(fp)
	return os.IsExist(err)
}

func ReadJson(fp string) ([]byte, error) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("Error:", err)
		return data, err
	}
	return data, err
}

func WriteJson(fp string, v []byte) error {
	err := ioutil.WriteFile(fp, v, 0644)
	return err
}

type Config interface {
	Read() ([]byte, error)
	Set(v any) ([]byte, error)
}

type PgConfig struct {
	Alias    string `json:"Alias"`
	Host     string `json:"Host"`
	Port     int32  `json:"Port"`
	UserName string `json:"UserName"`
	Passwd   string `json:"PassWd"`
	SSL      bool   `json:"SSL"`
	DBName   string `json:"DBName"`
	TimeZone string `json:"TimeZone"`
}

func (p *PgConfig) init(name string) error {
	now := time.Now()
	tzName, _ := now.Zone()
	alias := name
	p.Alias, p.TimeZone = alias, tzName
	return nil
}

func (p *PgConfig) Read(name string) ([]byte, error) {
	if Exists(FP) {
		res, err := ReadJson(FP)
		return res, err
	} else {
		err := p.init(name)
		return []byte(""), err
	}
}

func (p *PgConfig) Set(key string, value any) error {
	var keys = []string{"Alias", "Host", "Port", "UserName", "Passwd", "SSL", "DBName", "TimeZone"}
	var hasFlag = false

	for _, k := range keys {
		if k == key {
			hasFlag = true
		}
	}

	if !hasFlag {
		var flagError ConfigError
		flagError.Set("Key Error", fmt.Sprintf("The Config struct hasn't key named: %s", key))
		return flagError.err
	}
	return nil
}

func (p *PgConfig) DSN() string {
	tmplt := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%t TimeZone=%s"
	dsn := fmt.Sprintf(tmplt, p.Host, p.UserName, p.Passwd, p.DBName, p.Port, p.SSL, p.TimeZone)
	return dsn
}

func B64Encode(b []byte) string {
	str := base64.StdEncoding.EncodeToString(b)
	return str
}

func B64Decode(in string) ([]byte, error) {
	sDec, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		glog.Errorf("Error decoding string: %s ", err.Error())
	}

	return sDec, err
}

func UnknowTypeAssert(name string, val interface{}) error {
	var e error
	switch t := val.(type) {
	default:
		e = errors.New(fmt.Sprintf("%s need string data, the val type is %T", name, t))
		return e
		break
	}
	return e
}

func GetUUID(hyphen bool) string {
	uuidWithHyphen := uuid.NewRandom()

	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	if hyphen {
		return uuidWithHyphen
	} else {
		return uuid
	}
}
