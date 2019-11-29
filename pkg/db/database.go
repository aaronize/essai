package db

import (
    "github.com/jinzhu/gorm"
    "sync"
)

var (
    once            sync.Once
    wDbInstance     *gorm.DB
    rDbInstance     *gorm.DB
)

type Config struct {

}

func NewConfig() *Config {
    return &Config{}
}

func InitializeDB(conf *Config) error {
    return nil
}

func WDB() *gorm.DB {
    if wDbInstance == nil {
        //
    }
    return wDbInstance
}

func RDB() *gorm.DB {
    if rDbInstance == nil {

    }
    return rDbInstance
}

func Close() error {

    return nil
}
