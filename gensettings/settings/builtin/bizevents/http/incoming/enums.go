package incoming

type ComparisonEnum string

var ComparisonEnums = struct {
	Contains    ComparisonEnum
	EndsWith    ComparisonEnum
	Equals      ComparisonEnum
	Exists      ComparisonEnum
	NContains   ComparisonEnum
	NEndsWith   ComparisonEnum
	NEquals     ComparisonEnum
	NExists     ComparisonEnum
	NStartsWith ComparisonEnum
	StartsWith  ComparisonEnum
}{
	ComparisonEnum("CONTAINS"),
	ComparisonEnum("ENDS_WITH"),
	ComparisonEnum("EQUALS"),
	ComparisonEnum("EXISTS"),
	ComparisonEnum("N_CONTAINS"),
	ComparisonEnum("N_ENDS_WITH"),
	ComparisonEnum("N_EQUALS"),
	ComparisonEnum("N_EXISTS"),
	ComparisonEnum("N_STARTS_WITH"),
	ComparisonEnum("STARTS_WITH"),
}

type DataSourceEnum string

var DataSourceEnums = struct {
	RequestBody        DataSourceEnum
	RequestHeaders     DataSourceEnum
	RequestMethod      DataSourceEnum
	RequestParameters  DataSourceEnum
	RequestPath        DataSourceEnum
	ResponseBody       DataSourceEnum
	ResponseHeaders    DataSourceEnum
	ResponseStatuscode DataSourceEnum
}{
	DataSourceEnum("Request_body"),
	DataSourceEnum("Request_headers"),
	DataSourceEnum("Request_method"),
	DataSourceEnum("Request_parameters"),
	DataSourceEnum("Request_path"),
	DataSourceEnum("Response_body"),
	DataSourceEnum("Response_headers"),
	DataSourceEnum("Response_statusCode"),
}

type DataSourceWithStaticStringEnum string

var DataSourceWithStaticStringEnums = struct {
	ConstantString     DataSourceWithStaticStringEnum
	RequestBody        DataSourceWithStaticStringEnum
	RequestHeaders     DataSourceWithStaticStringEnum
	RequestMethod      DataSourceWithStaticStringEnum
	RequestParameters  DataSourceWithStaticStringEnum
	RequestPath        DataSourceWithStaticStringEnum
	ResponseBody       DataSourceWithStaticStringEnum
	ResponseHeaders    DataSourceWithStaticStringEnum
	ResponseStatuscode DataSourceWithStaticStringEnum
}{
	DataSourceWithStaticStringEnum("Constant_string"),
	DataSourceWithStaticStringEnum("Request_body"),
	DataSourceWithStaticStringEnum("Request_headers"),
	DataSourceWithStaticStringEnum("Request_method"),
	DataSourceWithStaticStringEnum("Request_parameters"),
	DataSourceWithStaticStringEnum("Request_path"),
	DataSourceWithStaticStringEnum("Response_body"),
	DataSourceWithStaticStringEnum("Response_headers"),
	DataSourceWithStaticStringEnum("Response_statusCode"),
}
