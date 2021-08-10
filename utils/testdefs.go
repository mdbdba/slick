package utils

import (
	"encoding/json"
	"io/ioutil"
)

type TestDefs struct {
	Tests []struct {
		MetricBaseName               string      `json:"metric_base_name"`
		MetricType                   string      `json:"metric_type"`
		MetricDesc                   string      `json:"metric_desc"`
		ExecutionMethod              string      `json:"execution_method"`
		Url                          string      `json:"url"`
		ExecutionDefinition          string      `json:"execution_definition"`
		ResponseType                 string      `json:"response_type"`
		ComparisonKey                string      `json:"comparison_key"`
		ResponseEvaluationIdentifier string      `json:"response_evaluation_identifier"`
		ComparisonValue              interface{} `json:"comparison_value"`
	} `json:"tests"`
}

func GetTestDefs(fp string) TestDefs {
	file, _ := ioutil.ReadFile(fp)

	data := TestDefs{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}
