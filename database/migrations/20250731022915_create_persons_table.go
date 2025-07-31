package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250731022915CreatePersonsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250731022915CreatePersonsTable) Signature() string {
	return "20250731022915_create_persons_table"
}

// Up Run the migrations.
func (r *M20250731022915CreatePersonsTable) Up() error {
	if !facades.Schema().HasTable("persons") {
		return facades.Schema().Create("persons", func(table schema.Blueprint) {
			table.ID()
			table.String("name")
			table.String("email")
			table.String("phone").Nullable()
			table.Integer("age").Nullable()
			table.Text("address").Nullable()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250731022915CreatePersonsTable) Down() error {
	return facades.Schema().DropIfExists("persons")
}
