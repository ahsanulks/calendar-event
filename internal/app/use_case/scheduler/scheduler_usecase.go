package scheduler

type SchedulerUsecase struct {
	// need db scheduler
	// need db user
	// client email
	ABC string
	DEF int

	abc2 string
}

func NewSchedulerUsecase(abc string) *SchedulerUsecase {
	return &SchedulerUsecase{
		ABC: abc,
	}
}
