package type_utils

import "strconv"

func ConvertStringToInt(data string) int {
	intParse, err := strconv.Atoi(data)
	if err != nil {
		return 0
	}
	return intParse
}
func ConvertStringToFloat(data string) float64 {
	intParse, err := strconv.Atoi(data)
	if err != nil {
		return 0
	}
	return float64(intParse)
}
