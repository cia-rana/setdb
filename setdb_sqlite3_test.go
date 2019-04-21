package setbd

import (
	"fmt"
	"os"
	"testing"
)

func TestDBWithSQLite3(t *testing.T) {
	const testDataSourceName = "test.db"

	db, err := OpenWithSQLite3(testDataSourceName)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	fmt.Println(db.Insert("1"))
	fmt.Println(db.Contain("1"))
	fmt.Println(db.Erase("1"))
	fmt.Println(db.Contain("1"))

	os.Remove(testDataSourceName)
}
