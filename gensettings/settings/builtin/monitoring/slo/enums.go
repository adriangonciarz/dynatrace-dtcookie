package slo

type SloEvaluationType string

var SloEvaluationTypes = struct {
	Aggregate SloEvaluationType
}{
	SloEvaluationType("AGGREGATE"),
}
