package json

import (
	"encoding/json"
)

func UnmarshalJSON[T any](data []byte) (result *T, err error) {
	return result, json.Unmarshal(data, &result)
}
