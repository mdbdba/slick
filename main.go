package main

import (
	"fmt"
	"github.com/mdbdba/slick/utils"
	"reflect"
)

func main() {
	var successCnt int
	var failureCnt int

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

			bodyArray, durationSec := utils.GetTestResult(testCall)
			if bodyArray != nil {
				fmt.Println("Body Array:", bodyArray)
			}

			comparisonOutcome := utils.IntComparison(testdefs.Tests[i].ComparisonKey,
				testdefs.Tests[i].ComparisonValue, testdefs.Tests[i].ResponseEvaluationIdentifier,
				bodyArray)

			fmt.Println(comparisonOutcome)

			if comparisonOutcome == true {
				successCnt = 1
				failureCnt = 0
			} else {
				successCnt = 0
				failureCnt = 1
			}
			tSName := fmt.Sprintf("%s_success_total", testdefs.Tests[i].MetricBaseName)
			tFName := fmt.Sprintf("%s_failure_total", testdefs.Tests[i].MetricBaseName)
			tBName := fmt.Sprintf("%s_seconds_bucket", testdefs.Tests[i].MetricBaseName)
			tB1Name := fmt.Sprintf("%s{le=\"1\"}", tBName)
			tB2Name := fmt.Sprintf("%s{le=\"2\"}", tBName)
			tB3Name := fmt.Sprintf("%s{le=\"4\"}", tBName)
			tB4Name := fmt.Sprintf("%s{le=\"8\"}", tBName)
			tB5Name := fmt.Sprintf("%s{le=\"+Inf\"}", tBName)
			tSumName := fmt.Sprintf("%s_seconds_sum", testdefs.Tests[i].MetricBaseName)
			tCntName := fmt.Sprintf("%s_seconds_count", testdefs.Tests[i].MetricBaseName)
			fmt.Println("# HELP", tSName)
			fmt.Println("# TYPE", tSName, "counter")
			fmt.Println(tSName, successCnt)
			fmt.Println("# HELP", tFName)
			fmt.Println("# TYPE", tFName, "counter")
			fmt.Println(tFName, failureCnt)
			fmt.Println("# HELP", tFName)
			fmt.Println("# TYPE", tFName, "histogram")
			var tCnt int
			if durationSec < 1.0 {
				tCnt = 1
			} else {
				tCnt = 0
			}
			fmt.Println(tB1Name, tCnt)
			if durationSec < 2.0 {
				tCnt = 1
			} else {
				tCnt = 0
			}
			fmt.Println(tB2Name, tCnt)
			if durationSec < 4.0 {
				tCnt = 1
			} else {
				tCnt = 0
			}
			fmt.Println(tB3Name, tCnt)
			if durationSec < 8.0 {
				tCnt = 1
			} else {
				tCnt = 0
			}
			fmt.Println(tB4Name, tCnt)
			fmt.Println(tB5Name, "1")
			fmt.Println(tSumName, durationSec)
			fmt.Println(tCntName, 1)
		}
	}
}
