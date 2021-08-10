package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func IntComparison(ComparisonKey string, ComparisonValue interface{},
	ResponseEvaluation string, resultSet []responseT) bool {

	var returnValue bool
	var compValueEQ int
	var compValueLower int
	var compValueUpper int
	var srcValue int
	var rangeSearch bool = false

	if reflect.ValueOf(ComparisonValue).Kind() == reflect.Float64 {
		compValueEQ = int(ComparisonValue.(float64))
		fmt.Println("compValue: ", compValueEQ)
	} else if reflect.ValueOf(ComparisonValue).Kind() == reflect.String {
		rangeSearch = true
		tmpArray := strings.Split(reflect.ValueOf(ComparisonValue).String(), "-")
		tLower, err := strconv.ParseInt(tmpArray[0], 10, 64)
		if err != nil {
			fmt.Println("issue converting lower bounds number.")
		}
		compValueLower = int(tLower)
		tUpper, err := strconv.ParseInt(tmpArray[1], 10, 64)
		if err != nil {
			fmt.Println("issue converting upper bounds number.")
		}
		compValueUpper = int(tUpper)
		fmt.Println("compValueLower: ", compValueLower)
		fmt.Println("compValueUpper: ", compValueUpper)
	}

	if ComparisonKey == "result" {
		srcValue = resultSet[0].Result
	}

	if ResponseEvaluation == "equals" {
		if compValueEQ == srcValue {
			returnValue = true
		} else {
			returnValue = false
		}
	} else if ResponseEvaluation == "between" {
		if rangeSearch == true {
			if srcValue >= compValueLower && srcValue <= compValueUpper {
				returnValue = true
			} else {
				returnValue = false
			}
		}
	}
	return returnValue
}
