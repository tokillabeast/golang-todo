package database

type InsertRespose struct {
	Deleted int
	Errors int
	Generated_keys []string
	Inserted int
	Replaced int
	Skipped int
	Unchanged int
}
