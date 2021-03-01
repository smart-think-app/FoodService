package type_utils

import "encoding/json"

func ConvertStructToJSONString(data interface{}) (string, error) {
	e, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(e), nil
}
