package algho

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetAverage(a ...float64) (result float64, err error) { // Расчет среднего арифметического
	sum := 0.0
	if len(a) > 10 || len(a) <= 1 {
		return 0, errors.New("количество измерений не может быть менее двух и более десяти")
	}

	digitsAfterDot := GetDigitsAfterDot(a[0])
	for _, num := range a {
		digitsAfterDotInLoop := GetDigitsAfterDot(num)
		if digitsAfterDotInLoop > digitsAfterDot {
			digitsAfterDot = digitsAfterDotInLoop
		}
	}
	for i := 0; i < len(a); i++ {
		formatedString := fmt.Sprintf("%.*f", digitsAfterDot, a[i])
		formatedResult, _ := strconv.ParseFloat(formatedString, 64)
		sum += formatedResult
	}
	return GetRoundedFloat(sum/float64(len(a)), digitsAfterDot), nil
}

func GetAverageMinusEl(a ...float64) (result []float64, err error) { // Расчет разности среднего арифметического и i изм.
	var res []float64
	digitsAfterDot := GetDigitsAfterDot(a[0])
	avg, err := GetAverage(a...)
	if err != nil {
		return nil, err
	}
	for _, num := range a {
		res = append(res, GetRoundedFloat(avg-num, digitsAfterDot))
	}
	return res, nil
}

func GetSquare(a ...float64) (result []float64, err error) { // Расчет квадрата разности
	var res []float64
	digitsAfterDot := GetDigitsAfterDot(a[0])
	avg, err := GetAverage(a...)
	if err != nil {
		return nil, err
	}
	for i := range a {
		res = append(res, GetRoundedFloat(math.Pow(avg-a[i], 2), digitsAfterDot*2))
	}
	return res, nil
}

func GetSO(a ...float64) (result float64, err error) { // Расчет среднеквадратичного отклонения
	digitsAfterDot := GetDigitsAfterDot(a[0])
	squaredArr, err := GetSquare(a...)
	if err != nil {
		return 0, err
	}
	sum := 0.0
	n := len(a) - 1
	for _, num := range squaredArr {
		sum += num
	}
	return GetRoundedFloat(math.Sqrt(sum/float64(len(a))/float64(n)), digitsAfterDot*3), nil
}

func GetRandErr(a ...float64) (result float64, err error) { // Расчет случайной погрешности
	digitsAfterDot := GetDigitsAfterDot(a[0])
	SO, err := GetSO(a...)
	var coefStudent float64
	if err != nil {
		return 0.0, err
	}
	switch len(a) {
	case 2:
		coefStudent = 12.706
	case 3:
		coefStudent = 4.303
	case 4:
		coefStudent = 3.182
	case 5:
		coefStudent = 2.776
	case 6:
		coefStudent = 2.571
	case 7:
		coefStudent = 2.447
	case 8:
		coefStudent = 2.365
	case 9:
		coefStudent = 2.306
	case 10:
		coefStudent = 2.262
	}

	return GetRoundedFloat(SO*coefStudent, digitsAfterDot*3), nil
}

// func GetInstrErr(unit float64, a ...float64) (result float64) { // Расчет приборной погрешности
// 	digitsAfterDot := GetDigitsAfterDot(a[0])
// 	return GetRoundedFloat(1.960*unit/3, digitsAfterDot*3)
// }

func GetInstrErr(a ...float64) (result float64) { // Расчет приборной погрешности
	digitsAfterDot := GetDigitsAfterDot(a[0])
	return GetRoundedFloat(1.960*math.Pow(0.1, float64(digitsAfterDot))/3, digitsAfterDot*3)
}

// func GetFullErr(unit float64, a ...float64) (result float64, err error) { // Расчет полной погрешности
// 	digitsAfterDot := GetDigitsAfterDot(a[1])
// 	instrErr := GetInstrErr(unit, a...)
// 	randErr, err := GetRandErr(a...)
// 	if err != nil {
// 		return 0.0, err
// 	}
// 	return GetRoundedFloat(math.Sqrt(math.Pow(instrErr, 2)+math.Pow(randErr, 2)), digitsAfterDot), nil
// }

func GetFullErr(a ...float64) (result float64, err error) { // Расчет полной погрешности
	digitsAfterDot := GetDigitsAfterDot(a[1])
	instrErr := GetInstrErr(a...)
	randErr, err := GetRandErr(a...)
	if err != nil {
		return 0.0, err
	}
	return GetRoundedFloat(math.Sqrt(math.Pow(instrErr, 2)+math.Pow(randErr, 2)), digitsAfterDot+1), nil
}

func GetDigitsAfterDot(num float64) int {
	formatedFloat := strconv.FormatFloat(num, 'f', -1, 64)
	if !strings.Contains(formatedFloat, ".") {
		return 1
	}
	digitsAfterPoint := (len(formatedFloat) - strings.Index(formatedFloat, ".") - 1)
	return digitsAfterPoint

}

func GetRoundedFloat(num float64, digitsAfterDot int) float64 {
	formatedString := fmt.Sprintf("%.*f", digitsAfterDot, num)
	formatedResult, _ := strconv.ParseFloat(formatedString, 64)
	return formatedResult
}
