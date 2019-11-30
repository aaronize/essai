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
    return a.add(db.WDB(), nil)
}

func (a *Accounts) Update() error {
    return a.update(db.WDB(), nil)
}

func (a *Accounts) Delete() error {
    return a.delete(db.WDB(), nil)
}

func (a *Accounts) List() error {
    return a.list(db.RDB(), nil)
}

func (a *Accounts) AddWithTransaction(tx *gorm.DB) error {
    return a.add(tx, nil)
}

func (a *Accounts) DeleteWithTransaction(tx *gorm.DB) error {
    return a.delete(tx, nil)
}

func (a *Accounts) UpdateWithTransaction(tx *gorm.DB) error {
    return a.update(tx, nil)
}

func (a *Accounts) ListWithTransaction(tx *gorm.DB) error {
    return a.list(tx, nil)
}

func (a *Accounts) add(db *gorm.DB, mp *AccountMap) error {
    var i Model
    if mp == nil {
        i = mp
    } else {
        i = a
    }

    if err := db.Model(&Accounts{}).Create(i).Error; err != nil {
        return err
    }

    return nil
}

func (a *Accounts) update(db *gorm.DB, mp *AccountMap) error {
    var i Model
    if mp == nil {
        i = mp
    } else {
        i = a
    }

    if err := db.Model(&Accounts{}).Updates(i).Error; err != nil {
        return err
    }

    return nil
}

func (a *Accounts) delete(db *gorm.DB, mp *AccountMap) error {
    var i Model
    if mp == nil {
        i = mp
    } else {
        i = a
    }

    if err := db.Model(&Accounts{}).Delete(i).Error; err != nil {
        return err
    }

    return nil
}

func (a *Accounts) list(db *gorm.DB, mp *AccountMap) error {

    return nil
}

// map
type AccountMap map[string]interface{}

func (am *AccountMap) Add() error {
    return (&Accounts{}).add(db.WDB(), am)
}

func (am *AccountMap) Update() error {
    return (&Accounts{}).update(db.WDB(), am)
}

func (am *AccountMap) Delete() error {
    return (&Accounts{}).delete(db.WDB(), am)
}

func (am *AccountMap) List() error {
    return (&Accounts{}).list(db.RDB(), am)
}

func (am *AccountMap) AddWithTransaction(tx *gorm.DB) error {
    return (&Accounts{}).add(tx, am)
}

func (am *AccountMap) DeleteWithTransaction(tx *gorm.DB) error {
    return (&Accounts{}).delete(tx, am)
}

func (am *AccountMap) UpdateWithTransaction(tx *gorm.DB) error {
    return (&Accounts{}).update(tx, am)
}

func (am *AccountMap) ListWithTransaction(tx *gorm.DB) error {
    return (&Accounts{}).list(tx, am)
}