package models

import (
    "time"
    "gorm.io/gorm"
)

type Venda struct {
    ID           int            `gorm:"primaryKey;autoIncrement" json:"id"`
    VendedorID   int            `gorm:"not null" json:"vendedor_id"`
    UserID       int            `gorm:"not null" json:"user_id"`
    ProdutoID    int            `gorm:"not null" json:"produto_id"`
    Total        float64        `gorm:"type:decimal(10,2);not null" json:"total"`
    Quantidade   int            `gorm:"not null" json:"quantidade"`
    CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
    Endereco     string         `gorm:"type:text" json:"endereco"`
    NumResidencia int           `gorm:"not null" json:"num_residencia"`
    CPF          string         `gorm:"type:varchar(30)" json:"cpf"`
    CEP          string         `gorm:"type:varchar(45)" json:"cep"`
    MtdPay       string         `gorm:"type:varchar(45)" json:"mtd_pay"`
    StsVenda     string         `gorm:"type:varchar(45);default:'Confirmada'" json:"sts_venda"`
}
