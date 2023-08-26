package db

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/isbasic/gin_play/common"
)

// dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
//
//	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
func IsNil(v interface{}) bool {

	valueOf := reflect.ValueOf(v)

	k := valueOf.Kind()

	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return valueOf.IsNil()
	default:
		return v == nil || v == "" || v == 0 || v == false
	}
}

type Database interface {
	Set(name string, val any) (bool, error)
	// Get(name string) (any, error)
	DSN() string
}

type Pg struct {
	alias    string `json:"alias"`
	host     string `json:"host"`
	user     string `json:"user"`
	password string `json:"password"`
	dbname   string `json:"dbname"`
	port     int    `json:"port"`
	ssl      string `json:"ssl"`
	tz       string `json:"tz"`
	dsn      string `json:"dsn"`
	fmtStr   string `json:"fmtStr"`
}

func (p *Pg) Set(name string, val interface{}) (bool, error) {
	keys := []string{"alias", "host", "user", "password", "dbname", "port", "ssl", "tz", "dsn", "fmtStr"}
	bFlag := false

	var e error

	for _, k := range keys {
		if k == name {
			bFlag = true
			switch {
			case k == "alias":
				alias, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.alias = alias
				break
			case k == "host":
				host, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.host = host
				break
			case k == "user":
				userv, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.user = userv
				break
			case k == "password":
				password, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.password = password
				break
			case k == "dbname":
				dbname, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.dbname = dbname
				break
			case k == "port":
				port, ok := val.(int)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.port = port
				break
			case k == "ssl":
				ssl, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.ssl = ssl
				break
			case k == "tz":
				tz, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.tz = tz
				break
			case k == "dsn":
				dsn, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.dsn = dsn
				break
			case k == "fmtStr":
				fmtStr, ok := val.(string)
				if !ok {
					return ok, common.UnknowTypeAssert(k, val)
				}
				p.fmtStr = fmtStr
				break
			}
		} else {
			e = errors.New(fmt.Sprintf("There's no key named %s.", name))
			return bFlag, e
		}
	}

	return bFlag, e
}

func (p *Pg) DSN() string {
	alias := p.alias
	host := p.host
	user := p.user
	password := p.password
	dbname := p.dbname
	port := p.port
	ssl := p.ssl
	tz := p.tz

	if IsNil(alias) {
		alias = "default"
	}

	if IsNil(host) {
		host = "localhost"
	}

	if IsNil(user) {
		user = "mac"
	}

	if IsNil(password) {
		password = "Yc_19860717"
	}

	if IsNil(dbname) {
		dbname = "mac"
	}

	if IsNil(port) {
		port = 5432
	}

	if IsNil(ssl) {
		ssl = "disable"
	}

	if IsNil(tz) {
		tz = "Asia/Shanghai"
	}

	dsnFormat := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s"

	p.fmtStr = dsnFormat

	dsn := fmt.Sprintf(dsnFormat, host, user, password, dbname, port, ssl, tz)

	p.dsn = common.B64Encode([]byte(dsn))

	r, err := common.B64Decode(p.dsn)

	if err != nil {
		fmt.Println(p.dsn)
	}

	return string(r)
}
