package main

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
	driver  neo4j.Driver
	session neo4j.Session
)

// UnwindStatement comment
type UnwindStatement struct {
	StatementType string
	NodeType      string
	Map           []map[string]interface{}
}

func hasValue(arr []interface{}, val interface{}) bool {
	for _, x := range arr {
		if x == val {
			return true
		}
	}
	return false
}

func constructUnwindStatment(us UnwindStatement) (statement string) {
	const identifier = "m"

	var props []interface{}

	for _, m := range us.Map {
		for prop := range m {
			if !hasValue(props, prop) {
				props = append(props, prop)
			}
		}
	}

	nodeProps := "{"
	for ix, prop := range props {
		nodeProps += fmt.Sprintf("%s: %s.%s", prop, identifier, prop)

		// Only add a comma if it's not the last item
		if ix < len(props)-1 {
			nodeProps += ", "
		}
	}
	nodeProps += "}"

	statement = fmt.Sprintf(
		"UNWIND $map as %s %s (a:%s %s)",
		identifier, us.StatementType, us.NodeType, nodeProps,
	)

	return
}

// CreateMultiple comment
func CreateMultiple(us UnwindStatement) error {
	statement := constructUnwindStatment(us)

	result, err := session.Run(statement, map[string]interface{}{
		"map": us.Map,
	})

	if err != nil {
		return err
	}

	for result.Next() {
		// fmt.Printf("Created Item with Userame = '%s'\n", result.Record().GetByIndex(0).(string))
	}
	if err = result.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	var err error

	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
	}

	defer driver.Close()

	session, err = driver.Session(neo4j.AccessModeWrite)
	if err != nil {
	}

	defer session.Close()

	users := UnwindStatement{
		StatementType: "CREATE",
		NodeType:      "User",
		Map: []map[string]interface{}{
			map[string]interface{}{"name": "Steve", "position": "Developer", "age": 29},
			map[string]interface{}{"name": "Mark", "position": "Developer"},
			map[string]interface{}{"name": "Jimmy", "position": "Product Manager"},
		},
	}

	groups := UnwindStatement{
		StatementType: "CREATE",
		NodeType:      "Group",
		Map: []map[string]interface{}{
			map[string]interface{}{"name": "Admin", "visibility": "hidden"},
			map[string]interface{}{"name": "Viewer"},
			map[string]interface{}{"name": "Writer"},
		},
	}

	err = CreateMultiple(users)
	if err != nil {
		log.Println("Error creating users")
	}

	err = CreateMultiple(groups)
	if err != nil {
		log.Println("Error creating groups")
	}
}
