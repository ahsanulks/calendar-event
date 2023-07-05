package schedule

type ScheduleRepository struct {
	db *Database
}

func NewScheduleRepository(db *Database) *ScheduleRepository {
  return &ScheduleRepository{
		db: db,
  }
}

func (r *scheduleRepository) Create(ctx context.Context, schedule *entity.Schedule) (int, error) {
	
}
