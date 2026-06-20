package service


type SubService struct{}

var subService *SubService

func NewSubscriptionsService() *SubService {
	if subService == nil {
		subService = &SubService{}
	}
	return subService
}






