package models

import (
	"gorm.io/gorm"
	"time"
)

type Produto struct {
	ID         int            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int            `gorm:"not null" json:"user_id"`
	Nome       string         `gorm:"type:varchar(100);not null" json:"nome"`
	Descricao  string         `gorm:"type:text" json:"descricao"`
	Imagem     string         `gorm:"type:longtext" json:"imagem"`
	Preco      float64        `gorm:"type:decimal(10,2);not null" json:"preco"`
	Quantidade int            `gorm:"default:0" json:"quantidade"`
	Codigo     string         `gorm:"type:varchar(255)" json:"codigo"`
	Garantia   int            `json:"garantia"`
	Categoria  string         `gorm:"type:varchar(255)" json:"categoria"`
	Marca      string         `gorm:"type:varchar(255)" json:"marca"`
	CreatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
