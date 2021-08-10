package main

import (
	"fmt"
	"github.com/mdbdba/slick/utils"
	"reflect"
)

func main() {
	//var comparisonOutcome bool

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
		fmt.Println("Comparison Value: ", testdefs.Tests[i].ComparisonValue)
		f := reflect.ValueOf(testdefs.Tests[i].ComparisonValue)
		fmt.Println("Comparison Value datatype:", f.Kind())
		fmt.Println("Comparison Value datatype 2:", f.String())
		if testdefs.Tests[i].ExecutionMethod == "url" {
			testCall := testdefs.Tests[i].Url + "/" + testdefs.Tests[i].ExecutionDefinition
			fmt.Println("Testing :", testCall)

			bodyArray := utils.GetTestResult(testCall)
			if bodyArray != nil {
				fmt.Println("Body Array:", bodyArray)
			}

			comparisonOutcome := utils.IntComparison(testdefs.Tests[i].ComparisonKey,
				testdefs.Tests[i].ComparisonValue, testdefs.Tests[i].ResponseEvaluationIdentifier,
				bodyArray)

			fmt.Println(comparisonOutcome)
		}

	}
}
