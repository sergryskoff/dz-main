// package req реализует функцию кодирования в json
package req

import (
	"encoding/json"
	"io"
)

func DecodeBody[T any](body io.ReadCloser) (T, error) {
	var payload T
	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return payload, err
	}
	return payload, nil
}
