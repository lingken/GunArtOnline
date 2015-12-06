package util

type Database struct {
	database map[string]interface{}
}

func NewDatabase() *Database {
	db := Database{
		database: make(map[string]interface{}),
	}
	return &db
}

func (db *Database) Put(key string, value interface{}) {
	db.database[key] = value
}

func (db *Database) GetValue(key string) (interface{}, bool) {
	val, ok := db.database[key]
	return val, ok
}
