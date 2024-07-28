package models

import (
    "time"
    "gorm.io/gorm"
)

type Usuario struct {
    ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
    Nome      string         `gorm:"type:varchar(100);not null" json:"nome"`
    Senha     string         `gorm:"type:varchar(255);not null" json:"senha"`
    Email     string         `gorm:"type:varchar(100);not null" json:"email"`
    Tipo      string         `gorm:"type:varchar(255);not null" json:"tipo"`
    Tel       string         `gorm:"type:varchar(20)" json:"tel"`
    Endereco  string         `gorm:"type:text" json:"endereco"`
    CPF       string         `gorm:"type:varchar(30)" json:"cpf"`
    CEP       string         `gorm:"type:varchar(45)" json:"cep"`
    CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
