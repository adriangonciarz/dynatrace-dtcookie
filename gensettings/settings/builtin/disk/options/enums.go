package options

type OsTypeEnum string

var OsTypeEnums = struct {
	OsTypeAix     OsTypeEnum
	OsTypeDarwin  OsTypeEnum
	OsTypeHpux    OsTypeEnum
	OsTypeLinux   OsTypeEnum
	OsTypeSolaris OsTypeEnum
	OsTypeUnknown OsTypeEnum
	OsTypeWindows OsTypeEnum
	OsTypeZos     OsTypeEnum
}{
	OsTypeEnum("OS_TYPE_AIX"),
	OsTypeEnum("OS_TYPE_DARWIN"),
	OsTypeEnum("OS_TYPE_HPUX"),
	OsTypeEnum("OS_TYPE_LINUX"),
	OsTypeEnum("OS_TYPE_SOLARIS"),
	OsTypeEnum("OS_TYPE_UNKNOWN"),
	OsTypeEnum("OS_TYPE_WINDOWS"),
	OsTypeEnum("OS_TYPE_ZOS"),
}
