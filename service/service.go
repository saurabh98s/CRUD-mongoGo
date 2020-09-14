package service

import ("mongo-golang/repository"
		"mongo-golang/persistent"
)

// Services is the interface for enclosing all the services
type Services interface {
	Auth() repository.AuthRepo
}

type services struct {
	AuthService repository.AuthRepo
}

func (svc *services) Auth() repository.AuthRepo {
	return svc.AuthService
}

func Init() Services {
	return &services{
		AuthService: repository.NewAuthRepo(persistent.DB()),
	}
}
