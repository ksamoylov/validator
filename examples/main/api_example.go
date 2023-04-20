package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ksamoylov/validator"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)

	fmt.Println("Starting server...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
		return
	}

	var data validator.Fields

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprintf(w, "Ok")
}

func validate(data validator.Fields) error {
	fRules := validator.NewFieldRules(10, 2)
	rule := validator.NewRule([]string{"name", "surname"}, "string", *fRules)
	model := validator.NewModel(data, []validator.Rule{*rule})
	_, validationErrors := model.Validate()

	if len(validationErrors) > 1 {
		return errors.New(validator.ToStringErrors(validationErrors))
	}

	return nil
}
