package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(w http.ResponseWriter, r *http.Request, target interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return err
	}

	if err := validate.Struct(target); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, fieldErr := range validationErrors {
				errors[fieldErr.Field()] = fieldErr.Tag()
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": errors,
			})
		} else {
			http.Error(w, "Validation error", http.StatusBadRequest)
		}
		return err
	}

	return nil
}