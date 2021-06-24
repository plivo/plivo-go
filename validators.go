package plivo

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func MultipleValidIntegers(paramname string,paramvalue string, lowerbound int,upperbound int){
	values := strings.SplitN(paramvalue,"<",-1)
	for i :=0;i<len(values);i++{
		val, err := strconv.Atoi(values[i])
		if err!=nil{
			logrus.Fatal(paramname + " values passed in the string must be integers")
		}else{
			if val < lowerbound || val >upperbound{
				error := paramname + " values passed must be in the range [" + strconv.Itoa(lowerbound) + " , " + strconv.Itoa(upperbound) + "]"
				logrus.Fatal(error)
			}
		}
	}
}