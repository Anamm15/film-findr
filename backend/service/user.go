package service

import (
	"context"
	"mime/multipart"

	"ReviewPiLem/dto"
	"ReviewPiLem/entity"
	"ReviewPiLem/helpers"
	"ReviewPiLem/repository"
	"ReviewPiLem/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UserService interface {
	GetAllUser(ctx context.Context) ([]dto.UserResponse, error)
	GetUserById(ctx context.Context, id int) (dto.UserResponse, error)
	RegisterUser(ctx context.Context, user dto.UserCreateRequest, photoProfil *multipart.FileHeader) (dto.UserResponse, error)
	LoginUser(ctx context.Context, req dto.UserLoginRequest) (entity.User, error)
	UpdateUser(ctx context.Context, user dto.UserUpdateRequest) error
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	cloudinary     *cloudinary.Cloudinary
	userRepository repository.UserRepository
}

func NewUserService(cloudinary *cloudinary.Cloudinary, userRepository repository.UserRepository) UserService {
	return &userService{
		cloudinary:     cloudinary,
		userRepository: userRepository,
	}
}

func (s *userService) GetAllUser(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.userRepository.GetAllUser(ctx)
	if err != nil {
		return nil, dto.ErrGetAllUser
	}

	var userResponse []dto.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, dto.UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			Nama:        user.Nama,
			Bio:         user.Bio,
			PhotoProfil: user.PhotoProfil,
		})
	}
	return userResponse, nil
}

func (s *userService) GetUserById(ctx context.Context, id int) (dto.UserResponse, error) {
	user, err := s.userRepository.GetUserById(ctx, id)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserByID
	}

	return dto.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Nama:        user.Nama,
		Bio:         user.Bio,
		PhotoProfil: user.PhotoProfil,
	}, nil
}

func (s *userService) RegisterUser(ctx context.Context, userCreateRequest dto.UserCreateRequest, photoProfil *multipart.FileHeader) (dto.UserResponse, error) {
	var uploadResult *uploader.UploadResult
	if photoProfil != nil {
		src, err := photoProfil.Open()
		if err != nil {
			return dto.UserResponse{}, err
		}
		defer src.Close()

		uniqueFileName := utils.GenerateUniqueImageName(userCreateRequest.Username, photoProfil.Filename)
		uploadResult, err = s.cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{
			Folder:   "user",
			PublicID: uniqueFileName,
		})
		if err != nil {
			return dto.UserResponse{}, err
		}
	}

	var photoURL string
	if uploadResult != nil {
		photoURL = uploadResult.SecureURL
	}

	user := entity.User{
		Username:    userCreateRequest.Username,
		Nama:        userCreateRequest.Nama,
		Password:    userCreateRequest.Password,
		Bio:         userCreateRequest.Bio,
		PhotoProfil: photoURL,
		Role:        helpers.ENUM_ROLE_USER,
	}

	userRepspone, err := s.userRepository.RegisterUser(ctx, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:          userRepspone.ID,
		Username:    userRepspone.Username,
		Nama:        userRepspone.Nama,
		Bio:         userRepspone.Bio,
		PhotoProfil: userRepspone.PhotoProfil,
	}, nil
}

func (s *userService) LoginUser(ctx context.Context, req dto.UserLoginRequest) (entity.User, error) {
	userResponse, err := s.userRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return entity.User{}, dto.ErrCheckUsername
	}

	checkPassword, _ := helpers.CheckPassword(userResponse.Password, []byte(req.Password))

	if userResponse.Username == req.Username && checkPassword {
		return userResponse, nil
	}

	return entity.User{}, dto.ErrEmailOrPassword
}

func (s *userService) UpdateUser(ctx context.Context, user dto.UserUpdateRequest) error {
	userRequest := entity.User{
		ID:       user.ID,
		Username: user.Username,
		Nama:     user.Nama,
		Bio:      user.Bio,
	}
	err := s.userRepository.UpdateUser(ctx, userRequest)
	if err != nil {
		return dto.ErrUpdateUser
	}

	return nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	if err := s.userRepository.DeleteUser(ctx, id); err != nil {
		return dto.ErrDeleteUser
	}

	return nil
}
