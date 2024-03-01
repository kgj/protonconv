package utils

import "testing"

func TestEmptyParams(t *testing.T) {
	params := Params{
		JsonFilename: "",
		DBFileName:   "",
		Password:     "",
	}

	err := Convert(params)

	want := "Please select a valid JSON file"
	if err.Error() != want {
		t.Fatal(err)
	}
}

func TestJsonFIleParams(t *testing.T) {
	params := Params{
		JsonFilename: "filename",
		DBFileName:   "",
		Password:     "",
	}

	err := Convert(params)

	want := "Please Enter a valid file name for the new keepass db"
	if err.Error() != want {
		t.Fatal(err)
	}
}

func TestPasswordEmptyParams(t *testing.T) {
	params := Params{
		JsonFilename: "filename",
		DBFileName:   "filename",
		Password:     "",
	}

	err := Convert(params)

	want := "Please enter a password for the new keepass db"
	if err.Error() != want {
		t.Fatal(err)
	}
}
