package updatewindows

type TimezoneEnum string

var TimezoneEnums = struct {
	Gmt0400 TimezoneEnum
	Gmt0600 TimezoneEnum
	Gmt1000 TimezoneEnum
	Gmt0700 TimezoneEnum
	Gmt0200 TimezoneEnum
	Gmt0800 TimezoneEnum
	Gmt0300 TimezoneEnum
	Gmt0900 TimezoneEnum
	Gmt1200 TimezoneEnum
	Gmt1100 TimezoneEnum
	Gmt0100 TimezoneEnum
	Gmt0000 TimezoneEnum
	Gmt0500 TimezoneEnum
}{
	TimezoneEnum("GMT_04_00"),
	TimezoneEnum("GMT_06_00"),
	TimezoneEnum("GMT_10_00"),
	TimezoneEnum("GMT_07_00"),
	TimezoneEnum("GMT_02_00"),
	TimezoneEnum("GMT_08_00"),
	TimezoneEnum("GMT_03_00"),
	TimezoneEnum("GMT_09_00"),
	TimezoneEnum("GMT_12_00"),
	TimezoneEnum("GMT_11_00"),
	TimezoneEnum("GMT_01_00"),
	TimezoneEnum("GMT_00_00"),
	TimezoneEnum("GMT_05_00"),
}

type RecurrenceEnum string

var RecurrenceEnums = struct {
	Once    RecurrenceEnum
	Daily   RecurrenceEnum
	Weekly  RecurrenceEnum
	Monthly RecurrenceEnum
}{
	RecurrenceEnum("ONCE"),
	RecurrenceEnum("DAILY"),
	RecurrenceEnum("WEEKLY"),
	RecurrenceEnum("MONTHLY"),
}
