package maintenancewindow

type DayOfWeekType string

var DayOfWeekTypes = struct {
	Friday    DayOfWeekType
	Monday    DayOfWeekType
	Saturday  DayOfWeekType
	Sunday    DayOfWeekType
	Thursday  DayOfWeekType
	Tuesday   DayOfWeekType
	Wednesday DayOfWeekType
}{
	DayOfWeekType("FRIDAY"),
	DayOfWeekType("MONDAY"),
	DayOfWeekType("SATURDAY"),
	DayOfWeekType("SUNDAY"),
	DayOfWeekType("THURSDAY"),
	DayOfWeekType("TUESDAY"),
	DayOfWeekType("WEDNESDAY"),
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
	Daily   ScheduleType
	Monthly ScheduleType
	Once    ScheduleType
	Weekly  ScheduleType
}{
	ScheduleType("DAILY"),
	ScheduleType("MONTHLY"),
	ScheduleType("ONCE"),
	ScheduleType("WEEKLY"),
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
