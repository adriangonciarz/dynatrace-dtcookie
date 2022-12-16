package relation

type Normalization string

var Normalizations = struct {
	Tolowercase   Normalization
	LeavetextasIs Normalization
	Touppercase   Normalization
}{
	Normalization("Tolowercase"),
	Normalization("Leavetextas_is"),
	Normalization("Touppercase"),
}

type RelationType string

var RelationTypes = struct {
	InstanceOf RelationType
	RunsOn     RelationType
	ChildOf    RelationType
	Calls      RelationType
	PartOf     RelationType
	SameAs     RelationType
}{
	RelationType("INSTANCE_OF"),
	RelationType("RUNS_ON"),
	RelationType("CHILD_OF"),
	RelationType("CALLS"),
	RelationType("PART_OF"),
	RelationType("SAME_AS"),
}

type IngestDataSource string

var IngestDataSources = struct {
	Metrics  IngestDataSource
	Logs     IngestDataSource
	Spans    IngestDataSource
	Entities IngestDataSource
	Topology IngestDataSource
	Events   IngestDataSource
}{
	IngestDataSource("Metrics"),
	IngestDataSource("Logs"),
	IngestDataSource("Spans"),
	IngestDataSource("Entities"),
	IngestDataSource("Topology"),
	IngestDataSource("Events"),
}
