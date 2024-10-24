package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// DB 是数据库连接的封装
type DB struct {
	*gorm.DB
}

// NewDB 创建一个新的数据库连接
func NewDB(dsn string) (*DB, error) {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 获取底层 *sql.DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(25)                 // 最大打开连接数
	sqlDB.SetMaxIdleConns(25)                 // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // 连接的最大生命周期

	return &DB{db}, nil
}

// Close 关闭数据库连接
func (d *DB) Close() error {
	sqlDB, err := d.DB.DB() // 获取底层的 *sql.DB
	if err != nil {
		return err
	}
	return sqlDB.Close() // 关闭连接
}

func loadMysql() {

}
