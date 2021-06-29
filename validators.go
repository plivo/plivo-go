package plivo

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

func MultipleValidIntegers(paramname string, paramvalue interface{}, lowerbound int, upperbound int) {
	if reflect.TypeOf(paramvalue).Kind() == reflect.Int {
		paramvalue := paramvalue.(int)
		if paramvalue < lowerbound || paramvalue > upperbound {
			error := paramname + " values must be in the range [" + strconv.Itoa(lowerbound) + " , " + strconv.Itoa(upperbound) + "]"
			logrus.Fatal(error)
		}
	} else if reflect.TypeOf(paramvalue).Kind() == reflect.String {
		paramvalue := paramvalue.(string)
		values := strings.SplitN(paramvalue, "<", -1)
		for i := 0; i < len(values); i++ {
			val, err := strconv.Atoi(values[i])
			if err != nil {
				logrus.Fatal(paramname + " Destination values in the string must be integers")
			} else {
				if val < lowerbound || val > upperbound {
					error := paramname + " Destination values must be in the range [" + strconv.Itoa(lowerbound) + " , " + strconv.Itoa(upperbound) + "]"
					logrus.Fatal(error)
				}
			}
		}
	} else {
		logrus.Fatal(paramname + " must be either string or integer")
	}
}
