package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type Admin struct {
	Id  int8 `gorm:"type:int(10);primary_key;AUTO_INCREMENT"`
	Account string  `gorm:"type:varchar(100);not null;default:'';comment:账户" `
	Pwd    string	`gorm:"type:varchar(100);comment:密码"`
	LastTime time.Time `gorm:"comment:最后登录时间"`
	Level int8 `gorm:"type:int(1);default:1"`
	Status int8 `gorm:"type:int(1);default:1;comment:状态1正常0异常"`
}

func init(){

}

// QueryAdminByAccountAndPassword 根据账户和密码查询后台用户
func (db *DB) QueryAdminByAccountAndPassword(account, password string) (admin Admin, err error) {
	// 获取md5加密后的密码
	h:= md5.New()
	h.Write([]byte(password))
	md5password := hex.EncodeToString(h.Sum(nil))
	// gorm的魅力，一句代码即可，对开发友好。
	return admin, db.db.Model(&Admin{}).Where("account = ? AND pwd = ?", account, md5password ).Take(&admin).Error
}

// UpdatePwdByAccount 修改密码
func (db *DB) UpdatePwdByAccount(adm *Admin, account string) (err error) {
	return db.db.Model(&Admin{}).Where("account = ?", account).Updates(&adm).Error
}

