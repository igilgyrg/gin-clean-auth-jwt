package usecase

import (
	"context"
	"github.com/igilgyrg/gin-todo/internal/entity"
	internalError "github.com/igilgyrg/gin-todo/internal/error"
)

type userUseCase struct {
	repo entity.UserRepository
}

func NewUserCase(repo entity.UserRepository) entity.UserUseCase {
	return &userUseCase{repo: repo}
}

func (u userUseCase) Get(ctx context.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u userUseCase) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	return nil, internalError.NewUserIsExistsWithEmail(nil, "user have not founded")
}

func (u userUseCase) Update(ctx context.Context, user *entity.User) error {
	return nil
}

func (u userUseCase) Create(ctx context.Context, user *entity.User) error {
	return nil
}

func (u userUseCase) Delete(ctx context.Context, id entity.ID) error {
	return nil
}
