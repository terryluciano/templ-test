package handler

import (
	"net/http"
	"strings"

	"github.com/terryluciano/templ-test/internal/service"
	"github.com/terryluciano/templ-test/internal/validation"
)

func HandleAuthSignup(w http.ResponseWriter, r *http.Request) {
	var input validation.SignupSchema

	if err := decodeJSON(r, &input); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json")
		return
	}

	input.Email = strings.ToLower(strings.TrimSpace(input.Email))
	input.FirstName = strings.TrimSpace(input.FirstName)
	input.LastName = strings.TrimSpace(input.LastName)

	if err := validation.Validate.Struct(input); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := service.AuthSignup(&input)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, user)
}
