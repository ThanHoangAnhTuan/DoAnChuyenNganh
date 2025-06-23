package facilitydetail

var (
	localService Service
)

func Use() Service {
	if localService == nil {
		panic("Implement localService Facility Detail not found for interface Service")
	}
	return localService
}

func Init(s Service) {
	localService = s
}
