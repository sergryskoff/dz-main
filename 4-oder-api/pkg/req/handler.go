// package req реализует функции обработки запроса
package req

import (
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := DecodeBody[T](r.Body)
	if err != nil {
		return nil, err
	}
	if err := ValidateBody(body); err != nil {
		return nil, err
	}
	return &body, nil
}
