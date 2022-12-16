package generictype

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
