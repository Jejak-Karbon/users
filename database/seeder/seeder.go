package seeder

import "github.com/born2ngopi/alterra/basic-echo-mvc/database"

func Seed() {

	conn := database.GetConnection()

	districtTableSeeder(conn)
	cityTableSeeder(conn)
	provinceTableSeeder(conn)
}
