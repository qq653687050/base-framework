package mysql

import (
	"base-framework/inf/conf"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	slog "log"
	"os"
	"time"
)

/**
 * @description:
 * @author:xy
 * @date:2022/8/30 14:16
 * @Version: 1.0
 */

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {

	if c.Database.Enable {
		// 终端打印输入 sql 执行记录
		newLogger := logger.New(
			slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢查询 SQL 阈值
				Colorful:      true,        // 禁用彩色打印
				//IgnoreRecordNotFoundError: false,
				LogLevel: logger.Info, // Log lever
			},
		)
		log.Info("-----database.source: %v", c.Database.Source)
		db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
			Logger:                                   newLogger,
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名是否加 s
			},
		})

		if err != nil {
			log.Errorf("failed opening connection to sqlite: %v", err)
			panic("failed to connect database")
		}
		sqlDB, err := db.DB()
		if err != nil {
			log.Errorf("Failed to connect mysql %s", err.Error())
		}
		sqlDB.SetMaxIdleConns(int(c.Database.Pool.Max))
		sqlDB.SetMaxOpenConns(int(c.Database.Pool.Min))
		sqlDB.SetConnMaxLifetime(time.Minute)
		return db
	}
	return nil
}
