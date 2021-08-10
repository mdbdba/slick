package main

import (
	"fmt"
	"github.com/mdbdba/slick/utils"
	"net/http"
)

func main() {

	testdefs := utils.GetTestDefs("./examples/compose/tests/sre-roller-tests.json")
	for i := 0; i < len(testdefs.Tests); i++ {
		fmt.Println("Metric Base Name: ", testdefs.Tests[i].MetricBaseName)
		fmt.Println("Metric Type: ", testdefs.Tests[i].MetricType)
		fmt.Println("Metric Description: ", testdefs.Tests[i].MetricDesc)
		fmt.Println("Execution Method: ", testdefs.Tests[i].ExecutionMethod)
		fmt.Println("Url: ", testdefs.Tests[i].Url)
		fmt.Println("Execution Definition: ", testdefs.Tests[i].ExecutionDefinition)
		fmt.Println("ResponseType: ", testdefs.Tests[i].ResponseType)
		fmt.Println("Comparison Key: ", testdefs.Tests[i].ComparisonKey)
		fmt.Println("Response Eval Identifier: ", testdefs.Tests[i].ResponseEvaluationIdentifier)
		fmt.Println("Comparision Value ", testdefs.Tests[i].ComparisonValue)
		if testdefs.Tests[i].ExecutionMethod == "url" {
			test_call := testdefs.Tests[i].Url + "/" + testdefs.Tests[i].ExecutionDefinition
			fmt.Println("Testing :", test_call)
			resp, err := http.Get(test_call)
			fmt.Println("Response: ", resp)
			fmt.Println("Error: ", err)
		}

	}
}
