package repositories

import (
	"template-service-go/internal/users/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (p *UserRepository) ListAll() ([]*models.UserModel, error) {
	var users []*models.UserModel
	err := p.db.Find(&users).Error
	return users, err
}

// func (p *UserRepository) Load(id uint) (*models.UserModel, error) {
// 	user := &models.UserModel{}
// 	err := p.db.Where(`id = ?`, id).First(user).Error
// 	return user, err
// }
