package schedule

import (
	"event-calendar/internal/app/domain/schedule/entity"
)

type ScheduleRepository interface {
}

// repository -> call dao

func (r *scheduleRepository) getUserSchedule() []entity.Schedule {
	// get user detail

	schedules, err := r.scheduleRepoInner.GetSchedulesByUseId()

	eventRepoInner.GetEventBySchedule()
	// di loop + masukin user detail

	event = enventRepository.getEvent()

	return response
}

func (r *ScheduleRepository) save(entity.Schedule) {

	entity.Schedule.save()
	entity.Schedule.remove()
	entity.Event = entity.Schedule.event()
	entity.Schedule.findBy("asdadas")
}
