package models

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/driver/mysql"
)
type DB struct {
	db *gorm.DB
}

// Model 定义公共的model `json:"name"` 这句是定义json输出的格式
type Model struct {
	ID        uint `gorm:"type:int(10);primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

var (
	db *gorm.DB
	sqlDB *sql.DB
	err error
)
// NewDB 初始化DB
func NewDB() *DB {
	return &DB{db: db}
}

func init() {
	// 获取配置文件
	root,_ := beego.AppConfig.String("mysqluser")
	pwd,_ := beego.AppConfig.String("mysqlpass")
	myurl,_ := beego.AppConfig.String("mysqlurl")
	mydb,_ := beego.AppConfig.String("mysqldb")

	// 链接mysql数据库
	myconn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",root,pwd,myurl,mydb)

	db,err = gorm.Open(mysql.Open(myconn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "m_", // 加前缀
			SingularTable: true, //去s
		},
		DisableForeignKeyConstraintWhenMigrating: true, //设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)
	})
	// gorm 提供了 DB 方法 可用于从当前 *gorm.DB 返回一个通用的数据库接口 *sql.DB
	sqlDB,_ = db.DB()
	if err != nil {
		fmt.Printf(err.Error())
		defer sqlDB.Close()
	}else {
		fmt.Printf("mysql数据库 已连上\n")
		sqlDB.SetMaxIdleConns(10) // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxOpenConns(100) // SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。
	}

	// 自动同步表结构
	//err = DB.AutoMigrate(&Admin{},&Message{},&Art{}, &User{})
	//if err != nil {
	//	fmt.Printf("同步数据库出错, %s \n", err)
	//}else{
	//	fmt.Printf("数据库表自动更新成功! \n")
	//}

	//// 如果admin表为空就赋值一个初始化用户
	//var count int64
	//// (&User{})查询用户表, Count(&count) 将用户表的数据赋值给count字段。
	//if err := DB.Model(&Admin{}).Count(&count).Error; err == nil && count == 0 {
	//	//新增
	//	DB.Create(&Admin{
	//		Account: "admin",
	//		//密码
	//		Pwd: "31f18760af93c60f021c2bc28ff50f21",
	//		//角色
	//		Level: 1,
	//	})
	//}

}
