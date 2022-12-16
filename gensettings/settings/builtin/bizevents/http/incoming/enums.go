package incoming

type ComparisonEnum string

var ComparisonEnums = struct {
	NExists     ComparisonEnum
	NEquals     ComparisonEnum
	Contains    ComparisonEnum
	Equals      ComparisonEnum
	NStartsWith ComparisonEnum
	EndsWith    ComparisonEnum
	NContains   ComparisonEnum
	StartsWith  ComparisonEnum
	NEndsWith   ComparisonEnum
	Exists      ComparisonEnum
}{
	ComparisonEnum("N_EXISTS"),
	ComparisonEnum("N_EQUALS"),
	ComparisonEnum("CONTAINS"),
	ComparisonEnum("EQUALS"),
	ComparisonEnum("N_STARTS_WITH"),
	ComparisonEnum("ENDS_WITH"),
	ComparisonEnum("N_CONTAINS"),
	ComparisonEnum("STARTS_WITH"),
	ComparisonEnum("N_ENDS_WITH"),
	ComparisonEnum("EXISTS"),
}

type DataSourceEnum string

var DataSourceEnums = struct {
	ResponseBody       DataSourceEnum
	ResponseHeaders    DataSourceEnum
	ResponseStatuscode DataSourceEnum
	RequestPath        DataSourceEnum
	RequestMethod      DataSourceEnum
	RequestHeaders     DataSourceEnum
	RequestParameters  DataSourceEnum
	RequestBody        DataSourceEnum
}{
	DataSourceEnum("Response_body"),
	DataSourceEnum("Response_headers"),
	DataSourceEnum("Response_statusCode"),
	DataSourceEnum("Request_path"),
	DataSourceEnum("Request_method"),
	DataSourceEnum("Request_headers"),
	DataSourceEnum("Request_parameters"),
	DataSourceEnum("Request_body"),
}

type DataSourceWithStaticStringEnum string

var DataSourceWithStaticStringEnums = struct {
	RequestBody        DataSourceWithStaticStringEnum
	ResponseStatuscode DataSourceWithStaticStringEnum
	ConstantString     DataSourceWithStaticStringEnum
	RequestHeaders     DataSourceWithStaticStringEnum
	RequestMethod      DataSourceWithStaticStringEnum
	ResponseBody       DataSourceWithStaticStringEnum
	RequestPath        DataSourceWithStaticStringEnum
	RequestParameters  DataSourceWithStaticStringEnum
	ResponseHeaders    DataSourceWithStaticStringEnum
}{
	DataSourceWithStaticStringEnum("Request_body"),
	DataSourceWithStaticStringEnum("Response_statusCode"),
	DataSourceWithStaticStringEnum("Constant_string"),
	DataSourceWithStaticStringEnum("Request_headers"),
	DataSourceWithStaticStringEnum("Request_method"),
	DataSourceWithStaticStringEnum("Response_body"),
	DataSourceWithStaticStringEnum("Request_path"),
	DataSourceWithStaticStringEnum("Request_parameters"),
	DataSourceWithStaticStringEnum("Response_headers"),
}
