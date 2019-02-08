package db

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type unwindStatement struct {
	statementType string
	nodeType      string
	valuesMap     []map[string]interface{}
}

var (
	permissionsDBAddress  = "bolt://localhost:7687"
	permissionsDBUsername = "neo4j"
	permissionsDBPassword = "password"

	driver neo4j.Driver
)

const (
	identifier = "m"
)

func init() {
	var err error

	driver, err = neo4j.NewDriver(permissionsDBAddress, neo4j.BasicAuth(permissionsDBUsername, permissionsDBPassword, ""), func(config *neo4j.Config) {
		config.MaxConnectionPoolSize = 10
	})

	if err != nil {
		log.Fatalf("Error connecting to Permissions Database: %+v", err)
	}
}

// AddUsers takes an array of User objects and creates
// each as a User node
func AddUsers(users []User) error {
	var vm []map[string]interface{}

	for _, u := range users {
		m := map[string]interface{}{"name": u.Name, "mobile": u.Mobile}
		vm = append(vm, m)
	}

	us := unwindStatement{
		statementType: "CREATE",
		nodeType:      "User",
		valuesMap:     vm,
	}

	err := us.executeUnwind()
	if err != nil {
		return err
	}

	return nil
}

// AddGroups takes an array of Group objects and creates
// each as a Group node
func AddGroups(groups []Group) error {
	var vm []map[string]interface{}

	for _, g := range groups {
		m := map[string]interface{}{"name": g.Name}
		vm = append(vm, m)
	}

	us := unwindStatement{
		statementType: "CREATE",
		nodeType:      "Group",
		valuesMap:     vm,
	}

	err := us.executeUnwind()
	if err != nil {
		return err
	}

	return nil
}

// AddAlerts takes an array of Alert objects and creates
// each as a Alert node
func AddAlerts(alerts []Alert) error {
	var vm []map[string]interface{}

	for _, a := range alerts {
		m := map[string]interface{}{"id": a.ID}
		vm = append(vm, m)
	}

	us := unwindStatement{
		statementType: "CREATE",
		nodeType:      "Alert",
		valuesMap:     vm,
	}

	err := us.executeUnwind()
	if err != nil {
		return err
	}

	return nil
}

func hasValue(arr []interface{}, val interface{}) bool {
	for _, x := range arr {
		if x == val {
			return true
		}
	}
	return false
}

func (us *unwindStatement) getUniqueProps() (uniqueProps []interface{}) {
	for _, m := range us.valuesMap {
		for prop := range m {
			if !hasValue(uniqueProps, prop) {
				uniqueProps = append(uniqueProps, prop)
			}
		}
	}

	return
}

func (us *unwindStatement) constructUnwind() (statement string) {
	uniqueProps := us.getUniqueProps()

	props := "{"
	for ix, prop := range uniqueProps {
		props += fmt.Sprintf("%s: %s.%s", prop, identifier, prop)

		// Only add a comma if it's not the last item
		if ix < len(uniqueProps)-1 {
			props += ", "
		}
	}
	props += "}"

	statement = fmt.Sprintf(
		"UNWIND $map as %s %s (a:%s %s)",
		identifier, us.statementType, us.nodeType, props,
	)

	return
}

func (us *unwindStatement) executeUnwind() error {
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}

	defer session.Close()

	result, err := session.Run(us.constructUnwind(), map[string]interface{}{
		"map": us.valuesMap,
	})
	if err != nil {
		return err
	}

	for result.Next() {
		// I don't know why yet but without this block
		// it doesn't work
	}

	if err = result.Err(); err != nil {
		return err
	}

	return nil
}
