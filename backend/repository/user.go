package repository

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetUserById(ctx context.Context, id int) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	UpdateUser(ctx context.Context, user dto.UserUpdateRequest) error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var user []entity.User
	if err := r.db.WithContext(ctx).Select("id", "username", "nama", "bio", "photo_profil").
		Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, userId int) (entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).
		Select("id", "username", "nama", "bio", "photo_profil").
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Select("id", "username", "nama", "bio", "photo_profil").
		Where("username = ?", username).
		Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userReq dto.UserUpdateRequest) error {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, userReq.ID).Error; err != nil {
		return err
	}

	if userReq.Nama != "" {
		user.Nama = userReq.Nama
	}

	if userReq.Username != "" {
		user.Username = userReq.Username
	}

	if userReq.Bio != "" {
		user.Bio = userReq.Bio
	}

	if userReq.OldPhotoProfil != user.PhotoProfil && userReq.OldPhotoProfil != "" {
		user.PhotoProfil = userReq.OldPhotoProfil
	}

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userId int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, &userId).Error; err != nil {
		return err
	}

	return nil
}
