package repository

import (
	"errors"
	"mongo-golang/persistent/models"
	"time"

	"github.com/globalsign/mgo"
)

// AuthRepo implements the functions required for Auth
type AuthRepo interface {
	Save(doc *models.User) error
}



// authRepo holds all the dependencies needed for the methods
type authRepo struct {
	db *mgo.Database
}


func (repo *authRepo) Save(doc *models.User) error {
	doc.CreatedAt = time.Now().Unix()
	doc.DeletedAt = 0
	doc.UpdateAt = 0

	if doc == nil {
		return errors.New("[ERROR] nothing to save")
	}
	err:=repo.db.C("auth").Insert(doc)
	if err !=nil {
		return err
	}
	return nil
}


//NewAuthRepo creates a new Instance of the DomainRepo
func NewAuthRepo(db *mgo.Database) AuthRepo {

	return &authRepo{
		db:       db,
	}
}
