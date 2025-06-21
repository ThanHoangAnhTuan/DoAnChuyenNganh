package accommodationroom

var (
	localService Service
)

func Use() Service {
	if localService == nil {
		panic("Implement localService Accommodation Room not found for interface Service")
	}
	return localService
}

func Init(s Service) {
	localService = s
}
