package generictype

type IngestDataSource string

var IngestDataSources = struct {
	Logs     IngestDataSource
	Spans    IngestDataSource
	Entities IngestDataSource
	Topology IngestDataSource
	Events   IngestDataSource
	Metrics  IngestDataSource
}{
	IngestDataSource("Logs"),
	IngestDataSource("Spans"),
	IngestDataSource("Entities"),
	IngestDataSource("Topology"),
	IngestDataSource("Events"),
	IngestDataSource("Metrics"),
}
