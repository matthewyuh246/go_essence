package profile_test

import (
	"encoding/json"
	"os"
)

type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateProfile(filename string) (bool, error) {
	profile := Profile{
		Name: "Default User",
		Age:  25,
	}

	data, err := json.MarshalIndent(profile, "", " ")
	if err != nil {
		return false, err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
