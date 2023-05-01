package repository

import (
	"errors"
	"mockerydb/models"

	"gorm.io/gorm"
)

type IPayementRepository interface {
	UpdatePayment(id string, payment models.Payment) (models.Payment, error)
	Delete(id string) (int64, error)
	SelectPaymentWithId(id string) (models.Payment, error)
	CreatePayment(payment models.Payment) int64
}

type Repository struct {
	Database *gorm.DB
}

func (repo Repository) UpdatePayment(id string, payment models.Payment) (models.Payment, error) {
	var updatePayment models.Payment
	result := repo.Database.Model(&updatePayment).Where("id = ?", id).Updates(payment)
	if result.RowsAffected == 0 {
		return models.Payment{}, errors.New("payment data not update")
	}
	return updatePayment, nil
}

func (repo Repository) DeletePayment(id string) (int64, error) {
	var deletePayment models.Payment
	result := repo.Database.Where("id = ?", id).Delete(&deletePayment)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment data not delete")
	}
	return result.RowsAffected, nil
}

func (repo Repository) SelectPaymentWithId(id string) (models.Payment, error) {
	var payment models.Payment
	result := repo.Database.Where("id = ?", id).Select(&payment)
	if result.RowsAffected == 0 {
		return models.Payment{}, errors.New("payment data not found")
	}
	return payment, nil
}
func (repo Repository) CreatePayment(payment models.Payment) (int64, error) {
	result := repo.Database.Create(&payment)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment not created")
	}
	return result.RowsAffected, nil
}
