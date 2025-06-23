package stats

var (
	localService Service
)

func Use() Service {
	if localService == nil {
		panic("Implement localService Stats not found for interface Service")
	}
	return localService
}

func Init(s Service) {
	localService = s
}
