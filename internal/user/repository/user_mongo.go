package repository

import (
	"github.com/igilgyrg/gin-todo/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongoRepository struct {
	db *mongo.Database
}

func NewUserMongoRepository(db *mongo.Database) entity.UserRepository {
	return &userMongoRepository{db: db}
}

func (u userMongoRepository) Get(id entity.ID) (*entity.User, error){
	return nil, nil
}

func (u userMongoRepository) GetByUsername(username string) (*entity.User, error){
	return nil, nil
}

func (u userMongoRepository) Update(user entity.User) error{
	return nil
}

func (u userMongoRepository) Save(user entity.User) error{
	return nil
}

func (u userMongoRepository) Delete(id entity.ID) error{
	return nil
}
