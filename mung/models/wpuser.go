package models

// Wpuser 微信登录用户
type Wpuser struct {
	Model
	Sessionkey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Openid     string `json:"openid"`
	Nickname   string `json:"nickName"`
	Avatar     string `json:"avatarUrl"`
	Gender     int    `json:"gender"`
	Language   string `json:"language"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
}

func init() {

}

// SaveWpUser 新建用户
func (db *DB) SaveWpUser(wpUser *Wpuser) (err error) {
	return db.db.Create(wpUser).Error
}

// SaveWpUserByStruct 新建用户 用map
func (db *DB) SaveWpUserByStruct(userMap map[string]interface{}) (err error) {
	return db.db.Model(&Wpuser{}).Create(userMap).Error
}

// QueryWpUsersByArgs 查询微信用户列表
func (db *DB) QueryWpUsersByArgs(name string, page int, limit int) (wpusers []*Wpuser, err error) {
	// 先创建一个查询
	tx := db.db.Model(&Wpuser{}).Where("status = 1 OR status = 2").Order("id desc")

	if name != "" {
		tx.Where("nickname LIKE ?", "%"+name+"%")
	}
	// 分页
	if page > 0 && limit > 0 {
		tx = tx.Limit(limit).Offset((page - 1) * limit)
	}
	// 返回
	return wpusers, tx.Find(&wpusers).Error
}

// QueryWpUserCountByArgs 根据条件计算总数
func (db *DB) QueryWpUserCountByArgs(name string) (count int64, err error) {
	//先创建一个查询
	tx := db.db.Model(&Wpuser{})

	if name != "" {
		tx.Where("nickname LIKE ?", "%"+name+"%")
	}
	return count, tx.Count(&count).Error
}

// QueryOneWpUserById 获取一条wpuser
func (db *DB) QueryOneWpUserById(id int) (wpuser Wpuser, err error) {
	return wpuser, db.db.Where("id = ?", id).Take(&wpuser).Error
}

// DeleteWpUserById 删除某条数据
func (db *DB) DeleteWpUserById(uid int) error {
	return db.db.Delete(&Wpuser{}, "id = ? ", uid).Error
	//db.Delete(&User{}, "10")
}

// User 绑定表名
//func (v User) TableName() string {
//	//return "user"
//}
