package maintenancewindow

type DayOfWeekType string

var DayOfWeekTypes = struct {
	Wednesday DayOfWeekType
	Thursday  DayOfWeekType
	Friday    DayOfWeekType
	Saturday  DayOfWeekType
	Sunday    DayOfWeekType
	Monday    DayOfWeekType
	Tuesday   DayOfWeekType
}{
	DayOfWeekType("WEDNESDAY"),
	DayOfWeekType("THURSDAY"),
	DayOfWeekType("FRIDAY"),
	DayOfWeekType("SATURDAY"),
	DayOfWeekType("SUNDAY"),
	DayOfWeekType("MONDAY"),
	DayOfWeekType("TUESDAY"),
}

type SuppressionType string

var SuppressionTypes = struct {
	DetectProblemsAndAlert  SuppressionType
	DetectProblemsDontAlert SuppressionType
	DontDetectProblems      SuppressionType
}{
	SuppressionType("DETECT_PROBLEMS_AND_ALERT"),
	SuppressionType("DETECT_PROBLEMS_DONT_ALERT"),
	SuppressionType("DONT_DETECT_PROBLEMS"),
}

type MaintenanceType string

var MaintenanceTypes = struct {
	Planned   MaintenanceType
	Unplanned MaintenanceType
}{
	MaintenanceType("PLANNED"),
	MaintenanceType("UNPLANNED"),
}

type ScheduleType string

var ScheduleTypes = struct {
	Once    ScheduleType
	Daily   ScheduleType
	Weekly  ScheduleType
	Monthly ScheduleType
}{
	ScheduleType("ONCE"),
	ScheduleType("DAILY"),
	ScheduleType("WEEKLY"),
	ScheduleType("MONTHLY"),
}
