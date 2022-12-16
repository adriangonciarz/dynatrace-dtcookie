package diskrules

type DiskMetric string

var DiskMetrics = struct {
	LowDiskSpace       DiskMetric
	LowInodes          DiskMetric
	ReadTimeExceeding  DiskMetric
	WriteTimeExceeding DiskMetric
}{
	DiskMetric("LOW_DISK_SPACE"),
	DiskMetric("LOW_INODES"),
	DiskMetric("READ_TIME_EXCEEDING"),
	DiskMetric("WRITE_TIME_EXCEEDING"),
}

type DiskNameFilterOperator string

var DiskNameFilterOperators = struct {
	Contains         DiskNameFilterOperator
	DoesNotContain   DiskNameFilterOperator
	DoesNotEqual     DiskNameFilterOperator
	DoesNotStartWith DiskNameFilterOperator
	Equals           DiskNameFilterOperator
	StartsWith       DiskNameFilterOperator
}{
	DiskNameFilterOperator("CONTAINS"),
	DiskNameFilterOperator("DOES_NOT_CONTAIN"),
	DiskNameFilterOperator("DOES_NOT_EQUAL"),
	DiskNameFilterOperator("DOES_NOT_START_WITH"),
	DiskNameFilterOperator("EQUALS"),
	DiskNameFilterOperator("STARTS_WITH"),
}
