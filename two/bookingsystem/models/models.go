package models

import (
	"github.com/jimweng/bookingsystem/conf"
	"github.com/jimweng/bookingsystem/models/mysqldb"
	"github.com/jimweng/bookingsystem/models/redisdb"
)

/*
use facade design pattern to new db object
1. initDB
2. closeDB
*/

func RetriveRedisAccessModel() RedisImplement {
	return redisdb.RetriveRedisPoolObj()
}

type RedisImplement interface {
	GetKeyValue(string) (interface{}, error)
	SetKey(string, interface{}) error
	SetExpiredKey(string, interface{}, int) error
	SetNonExistedKey(string, interface{}) error
	DeleteKey(string) error
	Close() error
}

func RetriveMySqlDbAccessModel() MySqlImplement {
	return mysqldb.RetriveMySQLDBAccessObj()
}

/*
	TODO: map convert to interface{}?

	type LanguageRecordImp interface {
		GetRowMsg() []byte
	}
*/

type MySqlImplement interface {
	ExecSql(string) error
	Close() error
	// CreateRecord add an extra record with (user, language1, language2, language3, ...)
	CreateRecord(string, ...string) error
	QueryRecord(string) (*[]map[string]interface{}, error) // talbe1
	// QueryRecord(string) (*[]mysqldb.Language, error)

	CreateGameRecords(string) error
}

// InitDb init db
func InitDb(mysqlconf *conf.DbConf, redisconf *conf.RedisConf) error {
	var err error
	if err = initMySql(mysqlconf); err != nil {
		return err
	}

	if err = initRedis(redisconf); err != nil {
		return err
	}

	return nil
}

// Close close db
func Close() error {
	RetriveMySqlDbAccessModel().Close()
	RetriveRedisAccessModel().Close()
	return nil
}

func initMySql(c *conf.DbConf) error {
	mysqldb.LoadMySQLDBConfig(
		c.DbName,
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPassword,
		c.DbLogEnable,
		c.DbMaxConnect,
		c.DbIdleConnect,
	)

	return mysqldb.StartMySQLDB()
}

func initRedis(c *conf.RedisConf) error {
	redisdb.LoadRedisDBConfig(
		c.RedisAddr,
		c.RedisAuth,
		c.RedisMaxIdle,
		c.RedisMaxActive,
		c.RedisIdleTimeout,
	)
	return redisdb.StartRedisPool()
}
