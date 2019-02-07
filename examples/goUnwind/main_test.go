package main

import "testing"

func TestConstructUnwindStatment(t *testing.T) {
	expected := "UNWIND $map as m CREATE (a:User {name: m.name, position: m.position})"

	us := UnwindStatement{
		StatementType: "CREATE",
		NodeType:      "User",
		Map: []map[string]interface{}{
			map[string]interface{}{"name": "Steve", "position": "Developer"},
			map[string]interface{}{"name": "Mark", "position": "Developer"},
			map[string]interface{}{"name": "Jimmy", "position": "Product Manager"},
		},
	}

	statement := constructUnwindStatment(us)

	if statement != expected {
		t.Logf("Expected statement to be '%s' but got '%s'", expected, statement)
	}
}
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
