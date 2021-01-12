package tests_test

import (
	"math/rand"
	"testing"
	"time"

	. "gorm.io/gorm/utils/tests"
)

func TestDryMigrate(t *testing.T) {
	var (
		err       error
		query     string
		allModels = []interface{}{&User{}, &Account{}, &Pet{}, &Company{}, &Toy{}, &Language{}}
	)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	_ = DB.Migrator().DropTable("user_speaks", "user_friends")

	if err = DB.Migrator().DropTable(allModels...); err != nil {
		t.Fatalf("Failed to drop table, got error %v", err)
	}

	if query, err = DB.DryMigrate(allModels...); err != nil {
		t.Fatalf("Failed to auto migrate, but got error %v", err)
	}

	t.Log(query)

	if query == "" {
		t.Fatalf("Failed to generate SQL with dry-migrate command, please check your driver support")
	}
}
