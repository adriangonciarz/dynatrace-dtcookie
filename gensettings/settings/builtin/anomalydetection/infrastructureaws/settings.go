package infrastructureaws

import "github.com/dtcookie/hcl"

type Settings struct {
	Ec2CandidateHighCpuDetection     *Ec2CandidateHighCpuDetectionConfig     `json:"ec2CandidateHighCpuDetection"`
	ElbHighConnectionErrorsDetection *ElbHighConnectionErrorsDetectionConfig `json:"elbHighConnectionErrorsDetection"`
	LambdaHighErrorRateDetection     *LambdaHighErrorRateDetectionConfig     `json:"lambdaHighErrorRateDetection"`
	RdsHighCpuDetection              *RdsHighCpuDetectionConfig              `json:"rdsHighCpuDetection"`
	RdsHighMemoryDetection           *RdsHighMemoryDetectionConfig           `json:"rdsHighMemoryDetection"`
	RdsHighWriteReadLatencyDetection *RdsHighWriteReadLatencyDetectionConfig `json:"rdsHighWriteReadLatencyDetection"`
	RdsLowStorageDetection           *RdsLowStorageDetectionConfig           `json:"rdsLowStorageDetection"`
	RdsRestartsSequenceDetection     *RdsRestartsSequenceDetectionConfig     `json:"rdsRestartsSequenceDetection"`
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ec_2_candidate_high_cpu_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(Ec2CandidateHighCpuDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"elb_high_connection_errors_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ElbHighConnectionErrorsDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"lambda_high_error_rate_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(LambdaHighErrorRateDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rds_high_cpu_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsHighCpuDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rds_high_memory_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsHighMemoryDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rds_high_write_read_latency_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsHighWriteReadLatencyDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rds_low_storage_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsLowStorageDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rds_restarts_sequence_detection": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsRestartsSequenceDetectionConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"ec_2_candidate_high_cpu_detection":     me.Ec2CandidateHighCpuDetection,
		"elb_high_connection_errors_detection":  me.ElbHighConnectionErrorsDetection,
		"lambda_high_error_rate_detection":      me.LambdaHighErrorRateDetection,
		"rds_high_cpu_detection":                me.RdsHighCpuDetection,
		"rds_high_memory_detection":             me.RdsHighMemoryDetection,
		"rds_high_write_read_latency_detection": me.RdsHighWriteReadLatencyDetection,
		"rds_low_storage_detection":             me.RdsLowStorageDetection,
		"rds_restarts_sequence_detection":       me.RdsRestartsSequenceDetection,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"ec_2_candidate_high_cpu_detection":     &me.Ec2CandidateHighCpuDetection,
		"elb_high_connection_errors_detection":  &me.ElbHighConnectionErrorsDetection,
		"lambda_high_error_rate_detection":      &me.LambdaHighErrorRateDetection,
		"rds_high_cpu_detection":                &me.RdsHighCpuDetection,
		"rds_high_memory_detection":             &me.RdsHighMemoryDetection,
		"rds_high_write_read_latency_detection": &me.RdsHighWriteReadLatencyDetection,
		"rds_low_storage_detection":             &me.RdsLowStorageDetection,
		"rds_restarts_sequence_detection":       &me.RdsRestartsSequenceDetection,
	})
}
