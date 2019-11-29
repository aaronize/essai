package models

import (
    "github.com/aaronize/essai/pkg/db"
    "github.com/jinzhu/gorm"
)

type Accounts struct {
    BaseModel

    UserName    string      `json:"user_name" gorm:"type:varchar(255) column:user_name not null comment '用户名';"`
    FirstName   string      `json:"first_name" gorm:"type:varchar(255) column:first_name comment '';"`
    LastName    string      `json:"last_name" gorm:"type:varchar(255) column:last_name comment '';"`
    Gender      string      `json:"gender" gorm:"type:varchar(255) column:gender comment '';"`
    Email       string      `json:"email" gorm:"type:varchar(255) column:email comment '';"`
    IsActive    bool        `json:"is_active" gorm:"type:tinyint column:is_active comment ''"`
    Password    string      `json:"password" gorm:"type:varchar(255) column:password"`
    Level       int         `json:"level" gorm:"type:int(3) column:level"`
    IsSuperUser bool        `json:"is_superuser" gorm:"type:tinyint column:is_superuser"`
}

func (a *Accounts) Add() error {
    return a.add(db.WDB())
}

func (a *Accounts) Update() error {
    return a.update(db.WDB())
}

func (a *Accounts) Delete() error {
    return a.delete(db.WDB())
}

func (a *Accounts) List() error {
    return a.list(db.RDB())
}

func (a *Accounts) AddWithTransaction(tx *gorm.DB) error {
    return a.add(tx)
}

func (a *Accounts) DeleteWithTransaction(tx *gorm.DB) error {
    return a.delete(tx)
}

func (a *Accounts) UpdateWithTransaction(tx *gorm.DB) error {
    return a.update(tx)
}

func (a *Accounts) ListWithTransaction(tx *gorm.DB) error {
    return a.list(tx)
}

func (a *Accounts) add(db *gorm.DB) error {

    return nil
}

func (a *Accounts) update(db *gorm.DB) error {

    return nil
}

func (a *Accounts) delete(db *gorm.DB) error {

    return nil
}

func (a *Accounts) list(db *gorm.DB) error {

    return nil
}
