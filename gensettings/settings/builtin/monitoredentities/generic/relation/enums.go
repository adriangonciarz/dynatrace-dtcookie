package relation

type IngestDataSource string

var IngestDataSources = struct {
	Entities IngestDataSource
	Events   IngestDataSource
	Logs     IngestDataSource
	Metrics  IngestDataSource
	Spans    IngestDataSource
	Topology IngestDataSource
}{
	IngestDataSource("Entities"),
	IngestDataSource("Events"),
	IngestDataSource("Logs"),
	IngestDataSource("Metrics"),
	IngestDataSource("Spans"),
	IngestDataSource("Topology"),
}

type Normalization string

var Normalizations = struct {
	LeavetextasIs Normalization
	Tolowercase   Normalization
	Touppercase   Normalization
}{
	Normalization("Leavetextas_is"),
	Normalization("Tolowercase"),
	Normalization("Touppercase"),
}

type RelationType string

var RelationTypes = struct {
	Calls      RelationType
	ChildOf    RelationType
	InstanceOf RelationType
	PartOf     RelationType
	RunsOn     RelationType
	SameAs     RelationType
}{
	RelationType("CALLS"),
	RelationType("CHILD_OF"),
	RelationType("INSTANCE_OF"),
	RelationType("PART_OF"),
	RelationType("RUNS_ON"),
	RelationType("SAME_AS"),
}
