package service

import (
	"fmt"
	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hsintegration/dbaccess"
	u "hsintegration/utility"
	"log"
	"os"
	"testing"
)

var (
	fixtures *testfixtures.Loader
)

const (
	dbDriver = "postgres"
	dbHost   = "localhost"
	dbPort   = 5432

	dbName = "test_db"
	dbUser = "postgres"
	dbPass = "root"

	sslMode = "disable"

	fixtureDir = "../test/fixtures"
)

func TestMain(m *testing.M) {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPass, dbName)
	config = fmt.Sprintf("%s sslmode=%s", config, sslMode)

	db, err := gorm.Open(postgres.Open(config))
	if err != nil {
		u.Log.Error(err.Error())
	}

	dbaccess.SetupService(db, nil)

	if err != nil {
		log.Fatalf("Error Connecting to db:  %s", err)
	}

	dbb, _ := db.DB()
	fixtures, err = testfixtures.New(
		testfixtures.Database(dbb),
		testfixtures.Dialect(dbDriver),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func PrepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("error occurred in db", err)
	}
}

func TestVerifyAPI(t *testing.T) {
	PrepareTestDatabase()
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "api is not available in database",
			args: args{
				key: "oiuppooupiu",
			},
			want: false,
		},
		{
			name: "for valid api",
			args: args{
				key: "1a213976-e94b-4906-9ddd-6bef32c22728",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyAPI(tt.args.key); got != tt.want {
				t.Errorf("VerifyAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
