package scheduler

type CreateSchedulerUsecase interface {
	Create(map[string]interface{}) error
}

type CreateEventUsecase interface {
	Create(map[string]interface{}) error
}

type SchedulerUsecase struct {
	// need db scheduler
	// need db user
	// client email
	// ABC        string // repositoryA
	// ASD        string // repositoryB
	// ServX      string // service X
	// MicroServX string // client handler microserviceA
	// DEF        int    // S3, GCS

	// abc2 string
	createScheduler CreateSchedulerUsecase
	createEvent     CreateEventUsecase
}

/* apply high cohesion principle
cohesion tinggi = dipake semua dependency-nya
*/

/* usecase dependecy-nya apa aja?
service, clientSvc, dao access
*/

func NewSchedulerUsecase(createScheduler CreateSchedulerUsecase, createEvent CreateEventUsecase) *SchedulerUsecase {
	return &SchedulerUsecase{
		createScheduler: createScheduler,
		createEvent:     createEvent,
	}
}

// facade pattern https://refactoring.guru/design-patterns/facade
// facade -> clean architecture https://medium.com/@mammimia/clean-architecture-part-iv-architecture-5eee3cb5656a
// func (su *SchedulerUsecase) GetSchedule() {
// 	a := su.ABC
// 	b := su.MicroServX
// 	return a + b
// }

// handling transaction db pattern related to business logic (?)
func (su *SchedulerUsecase) CreateSchedule(params map[string]interface{}) error {
	err := su.createScheduler.Create(params)
	if err != nil {
		return err
	}
	return su.createEvent.Create(params)
}

// func (su *SchedulerUsecase) Rechedule() {
// 	fmt.Println(su.ABC)
// 	fmt.Println(su.ASD)
// 	fmt.Println(su.DEF)
// }
