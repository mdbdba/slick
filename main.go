package main

import (
	"fmt"
	"github.com/mdbdba/slick/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/expfmt"
	"reflect"
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

			fmt.Println("Comparison Outcome:", comparisonOutcome)

			// define Prom metric names from base
			tLName := fmt.Sprintf("%s_last_run_timestamp", testdefs.Tests[i].MetricBaseName)
			tSName := fmt.Sprintf("%s_success_total", testdefs.Tests[i].MetricBaseName)
			tFName := fmt.Sprintf("%s_failure_total", testdefs.Tests[i].MetricBaseName)
			tHName := fmt.Sprintf("%s_duration_nanoseconds", testdefs.Tests[i].MetricBaseName)

			// define Prom help strings from base
			lastStr := "Timestamp when this job was last run."
			successStr := fmt.Sprintf("SLIck created Counter of successes for %s.",
				testdefs.Tests[i].MetricBaseName)
			failureStr := fmt.Sprintf("SLIck created Counter of failures for %s.",
				testdefs.Tests[i].MetricBaseName)
			durationStr := fmt.Sprintf("SLIck created Histogram for NS duration of %s.",
				testdefs.Tests[i].MetricBaseName)

			// define Prom metrics
			lastGuage := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: tLName,
				Help: lastStr,
			})
			successCntr := prometheus.NewCounter(prometheus.CounterOpts{
				Name: tSName,
				Help: successStr,
				// ConstLabels: prometheus.Labels{"foo": "bar"},
			})
			failureCntr := prometheus.NewCounter(prometheus.CounterOpts{
				Name: tFName,
				Help: failureStr,
				// ConstLabels: prometheus.Labels{"foo": "bar"},
			})
			durationHist := prometheus.NewHistogram(
				prometheus.HistogramOpts{
					Name: tHName,
					Help: durationStr,
				})

			lastGuage.SetToCurrentTime()
			durationHist.Observe(durationSec)

			if comparisonOutcome == true {
				successCntr.Inc()
			} else {
				failureCntr.Inc()
			}

			if err := push.New("http://localhost:8888/metrics", "slick").
				Format(expfmt.FmtText).
				Collector(lastGuage).
				Grouping("slickGroup", testdefs.Tests[i].MetricBaseName).
				Push(); err != nil {
				fmt.Println("Could not push lastGuage to Pushgateway:", err)
			}

			if err := push.New("http://localhost:8888/metrics", "slick").
				Format(expfmt.FmtText).
				Collector(successCntr).
				Grouping("slickGroup", testdefs.Tests[i].MetricBaseName).
				Push(); err != nil {
				fmt.Println("Could not push successCntr to Pushgateway:", err)
			}

			if err := push.New("http://localhost:8888/metrics", "slick").
				Format(expfmt.FmtText).
				Collector(failureCntr).
				Grouping("slickGroup", testdefs.Tests[i].MetricBaseName).
				Push(); err != nil {
				fmt.Println("Could not push failureCntr to Pushgateway:", err)
			}

			if err := push.New("http://localhost:8888/metrics", "slick").
				Format(expfmt.FmtText).
				Collector(durationHist).
				Grouping("slickGroup", testdefs.Tests[i].MetricBaseName).
				Push(); err != nil {
				fmt.Println("Could not push durationHist to Pushgateway:", err)
			}
			fmt.Println("")
		}
	}
}
