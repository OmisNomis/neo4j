package db

import "testing"

func TestHasValue(t *testing.T) {
	arr := []interface{}{"yes", "no", 1, 5}

	output := hasValue(arr, "yes")
	if !output {
		t.Errorf("Expected output to be %t, but it was %t", true, output)
	}

	output = hasValue(arr, 5)
	if !output {
		t.Errorf("Expected output to be %t, but it was %t", true, output)
	}

	output = hasValue(arr, "maybe")
	if output {
		t.Errorf("Expected output to be %t, but it was %t", false, output)
	}

	output = hasValue(arr, 3)
	if output {
		t.Errorf("Expected output to be %t, but it was %t", false, output)
	}
}

func TestConstructUnwind(t *testing.T) {
	expected := "UNWIND $map as m CREATE (a:User {name: m.name, position: m.position})"

	us := unwindStatement{
		statementType: "CREATE",
		nodeType:      "User",
		valuesMap: []map[string]interface{}{
			map[string]interface{}{"name": "Steve", "position": "Developer"},
			map[string]interface{}{"name": "Mark", "position": "Developer"},
			map[string]interface{}{"name": "Jimmy", "position": "Product Manager"},
		},
	}

	statement := us.constructUnwind()

	if statement != expected {
		t.Errorf("Expected statement to be '%s' but got '%s'", expected, statement)
	}
}

func TestGetUniqueProps(t *testing.T) {
	us := unwindStatement{
		statementType: "CREATE",
		nodeType:      "User",
		valuesMap: []map[string]interface{}{
			map[string]interface{}{"name": "Steve", "age": 24},
			map[string]interface{}{"name": "Mark", "position": "Developer"},
			map[string]interface{}{"name": "Jimmy", "salary": 65000},
		},
	}

	up := us.getUniqueProps()

	if len(up) != 4 {
		t.Errorf("Expected length to be '%d' but got '%d'", 4, len(up))
	}

	if !hasValue(up, "name") {
		t.Errorf("Expected returned slice to have '%s', but it did not", "name")
	}

	if !hasValue(up, "age") {
		t.Errorf("Expected returned slice to have '%s', but it did not", "age")
	}

	if !hasValue(up, "position") {
		t.Errorf("Expected returned slice to have '%s', but it did not", "position")
	}

	if !hasValue(up, "salary") {
		t.Errorf("Expected returned slice to have '%s', but it did not", "salary")
	}

}
