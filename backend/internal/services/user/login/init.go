package login

var (
	localService Service
)

func Use() Service {
	if localService == nil {
		panic("Implement localService User Login not found for interface Service")
	}
	return localService
}

func Init(s Service) {
	localService = s
}
