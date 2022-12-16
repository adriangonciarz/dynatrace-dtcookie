package diskrules

type DiskNameFilterOperator string

var DiskNameFilterOperators = struct {
	DoesNotContain   DiskNameFilterOperator
	Equals           DiskNameFilterOperator
	DoesNotEqual     DiskNameFilterOperator
	StartsWith       DiskNameFilterOperator
	DoesNotStartWith DiskNameFilterOperator
	Contains         DiskNameFilterOperator
}{
	DiskNameFilterOperator("DOES_NOT_CONTAIN"),
	DiskNameFilterOperator("EQUALS"),
	DiskNameFilterOperator("DOES_NOT_EQUAL"),
	DiskNameFilterOperator("STARTS_WITH"),
	DiskNameFilterOperator("DOES_NOT_START_WITH"),
	DiskNameFilterOperator("CONTAINS"),
}

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
