package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Model interface {
    Add() error
    Update() error
    Delete() error
    List() error

    // with transaction
    AddWithTransaction(tx *gorm.DB) error
    DeleteWithTransaction(tx *gorm.DB) error
    UpdateWithTransaction(tx *gorm.DB) error
    ListWithTransaction(tx *gorm.DB) error
}

type BaseModel struct {
    Id          string      `json:"id" gorm:"primary_key column:id"`
    CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
    UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
    DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at" sql:"index"`
}
