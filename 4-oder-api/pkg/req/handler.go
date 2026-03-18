// package req реализует функции обработки запроса
package req

import (
	"net/http"
	"order-api/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := DecodeBody[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 400)
		return nil, err
	}
	if err := ValidateBody(body); err != nil {
		res.Json(*w, err.Error(), 400)
		return nil, err
	}
	return &body, nil
}
