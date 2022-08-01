package models

type User struct {
	Model
	Nickname string `gorm:"type:varchar(150);not null;default:'';comment:用户名" json:"name"`
	Phone    string `gorm:"type:varchar(150);not null;default:'';comment:手机" json:"phone"`
	Account string `gorm:"type:varchar(150);not null;default:'';comment:微信号" json:"account"`
	Avatar string `gorm:"type:varchar(200);not null;default:'';comment:微信号" json:"avatar"`
	Status int `gorm:"type:int(1);not null;default:1;comment:状态" json:"status"`
}

func init() {

}

// QueryUserByArgs 根据条件查询用户
func (db *DB) QueryUserByArgs(username ,phone string, page int, limit int) (users []*User,err error) {

	//先创建一个查询
	tx := db.db.Model(&User{}).Where("status = 1 OR status = 0").Order("id desc")

	if username != "" {
		tx.Where("nickname LIKE ?",  "%"+ username + "%")
	}

	if phone != "" {
		tx.Where("phone LIKE ?", phone + "%")
	}
	// 分页
	if page >0 && limit >0 {
		tx = tx.Limit(limit).Offset((page - 1) * limit)
	}
	// 返回
	return users,tx.Find(&users).Error

}

// QueryUserCount 根据条件计算总行数
func (db *DB) QueryUserCount(username,phone string)(count int64,err error)  {
	//先创建一个查询
	tx := db.db.Model(&User{})

	if username != "" {
		tx.Where("nickname LIKE ?",  "%"+ username + "%")
	}

	if phone != "" {
		tx.Where("phone LIKE ?", phone + "%")
	}
	return count,tx.Count(&count).Error
}

// QueryOneUserById 获取一条user
func (db *DB) QueryOneUserById(id int) (user User, err error) {
	return user, db.db.Where("id = ?", id).Take(&user).Error
}

// SaveUser 新建用户
func (db *DB) SaveUser(user *User) (err error) {
	return db.db.Create(user).Error
}
// SaveUserByStruct 新建用户 用map
func (db *DB) SaveUserByStruct(userMap map[string]interface{}) (err error) {
	return db.db.Model(&User{}).Create(userMap).Error
}

// UpdateUserById 根据ID更新User, 存在着0值不更新
func (db *DB) UpdateUserById(user *User,id int) ( err error) {
	return db.db.Model(&User{}).Where("id = ?", id).Updates(&user).Error
}
// UpdateUserStructById 根据ID更新User,用map 解决gorm 0值不更新的坑
func (db *DB) UpdateUserStructById(userMap map[string]interface{},id int) ( err error) {
	return db.db.Model(&User{}).Where("id = ?", id).Updates(userMap).Error
}


// DeleteUserById 删除某条数据
func (db *DB) DeleteUserById(uid int) error {
	return db.db.Delete(&User{}, "id = ? ",uid).Error
	//db.Delete(&User{}, "10")
}

// User 绑定表名
//func (v User) TableName() string {
//	//return "user"
//}
