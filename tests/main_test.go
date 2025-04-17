package tests

import (
	"fullsteak/internal/database"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var TestDB *pgxpool.Pool

func TestMain(m *testing.M) {
	TestDB = database.ConnectTest()
	database.InitTables(TestDB, "../tables.sql")

	// Run tests
	code := m.Run()
	database.DropTables(TestDB)
	TestDB.Close()
	os.Exit(code)
}
