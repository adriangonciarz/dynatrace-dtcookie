package traffic

type OsTypeEnum string

var OsTypeEnums = struct {
	OsTypeWindows OsTypeEnum
	OsTypeZos     OsTypeEnum
	OsTypeUnknown OsTypeEnum
	OsTypeAix     OsTypeEnum
	OsTypeDarwin  OsTypeEnum
	OsTypeHpux    OsTypeEnum
	OsTypeLinux   OsTypeEnum
	OsTypeSolaris OsTypeEnum
}{
	OsTypeEnum("OS_TYPE_WINDOWS"),
	OsTypeEnum("OS_TYPE_ZOS"),
	OsTypeEnum("OS_TYPE_UNKNOWN"),
	OsTypeEnum("OS_TYPE_AIX"),
	OsTypeEnum("OS_TYPE_DARWIN"),
	OsTypeEnum("OS_TYPE_HPUX"),
	OsTypeEnum("OS_TYPE_LINUX"),
	OsTypeEnum("OS_TYPE_SOLARIS"),
}
