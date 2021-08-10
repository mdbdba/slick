package main

import (
	"fmt"
	"github.com/mdbdba/slick/utils"
)

// type T struct {
// 	Tests []struct {
// 		MetricBaseName               string      `json:"metric_base_name"`
// 		MetricType                   string      `json:"metric_type"`
// 		MetricDesc                   string      `json:"metric_desc"`
// 		ExecutionMethod              string      `json:"execution_method"`
// 		Url                          string      `json:"url"`
// 		ExecutionDefinition          string      `json:"execution_definition"`
// 		ResponseType                 string      `json:"response_type"`
// 		ComparisonKey                string      `json:"comparison_key"`
// 		ResponseEvaluationIdentifier string      `json:"response_evaluation_identifier"`
// 		ComparisonValue              interface{} `json:"comparison_value"`
// 	} `json:"tests"`
// }

func main() {
	// 	file, _ := ioutil.ReadFile("./examples/compose/tests/sre-roller-tests.json")

	// data := T{}

	// _ = json.Unmarshal([]byte(file), &data)

	testdefs := utils.GetTestDefs("./examples/compose/tests/sre-roller-tests.json")
	for i := 0; i < len(testdefs.Tests); i++ {
		fmt.Println("Metric Base Name: ", testdefs.Tests[i].MetricBaseName)
		fmt.Println("Metric Type: ", testdefs.Tests[i].MetricType)
		fmt.Println("Metric Description: ", testdefs.Tests[i].MetricDesc)
		fmt.Println("Execution Method: ", testdefs.Tests[i].ExecutionMethod)
		fmt.Println("Url: ", testdefs.Tests[i].Url)
		fmt.Println("Execution Definition: ", testdefs.Tests[i].ExecutionMethod)
		fmt.Println("ResponseType: ", testdefs.Tests[i].ResponseType)
		fmt.Println("Comparison Key: ", testdefs.Tests[i].ComparisonKey)
		fmt.Println("Response Eval Identifier: ", testdefs.Tests[i].ResponseEvaluationIdentifier)
		fmt.Println("Comparision Value ", testdefs.Tests[i].ComparisonValue)
	}
}
