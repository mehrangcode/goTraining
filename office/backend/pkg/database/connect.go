package database

func Connect(DB_Name string) {
	switch DB_Name {
	case "postgres":
		postgresDbConnect()
	default:
		sqliteDbConnect()
	}
}
