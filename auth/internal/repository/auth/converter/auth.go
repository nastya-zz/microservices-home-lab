package converter

import "auth/internal/model"
import modelRepo "auth/internal/repository/auth/model"

//func ToCreateUserFromRepo(createUser *modelRepo.CreateUser) *model.CreateUser {
//	return &model.CreateUser{
//		Email:    createUser.Email,
//		Password: createUser.Password,
//		Name:     createUser.Name,
//		Role:     createUser.Role,
//	}
//}

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

//func ToUpdateUserFromRepo(user *modelRepo.UpdateUser) *model.UpdateUser {
//	return &model.UpdateUser{
//		Email: user.Email,
//		Name:  user.Name,
//		ID:    user.ID,
//	}
//}
