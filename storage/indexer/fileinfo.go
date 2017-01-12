package indexer

import (
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
)

type FileInfo struct {
	ID       int64  `xorm:"pk autoincr 'id'" json:"id"`
	DeviceID string `xorm:"notnull index 'device_id'" json:"device_id"`
	Path     string `xorm:"notnull index 'path'" json:"path"`
	Hash     string `xorm:"notnull index 'hash'" json:"hash"`
	Size     int64  `xorm:"'size'" json:"hash"`

	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

func (fi *FileInfo) Insert(orm *xorm.Engine) error {
	n, err := orm.Insert(fi)
	return checkOrmRet(n, err)
}

func (fi *FileInfo) Update(orm *xorm.Engine) error {
	n, err := orm.Where("id = ?", fi.ID).Update(fi)
	return checkOrmRet(n, err)
}

func (fi *FileInfo) Remove(orm *xorm.Engine) error {
	n, err := orm.Where("id = ?", fi.ID).Delete(fi)
	return checkOrmRet(n, err)
}

func checkOrmRet(n int64, err error) error {
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("control xorm not found")
	}

	return nil
}
