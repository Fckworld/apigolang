import *gorm.DB
func Init() *gorm.DB{
	dbURL = "postgres://postgres:admin@127.0.0.1:5432/crud"

	db, err = gorm.Open(postgres.Open(dbURL),&gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}

	db.Automigrate(&guitarra{})

	return db
}