package request

import (
	"go_pro_api/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		response.JSON(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		response.JSON(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	return &body, nil
}
