package dependencies

import (
	"act2/src/core"
	"act2/src/esp32Temperature/application"

	"act2/src/esp32Temperature/infraestructure"
	
	"act2/src/esp32Temperature/infraestructure/controllers"
	"database/sql"
	"fmt"
)

var (
	mySQL infraestructure.MySQL
	db    *sql.DB
)

func Init() {
	db, err := core.ConnectToDB()

	if err != nil {
		fmt.Println("server error")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)
	

}
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}

func GetPostSensorController()*controllers.PostSensorController{
	caseCreatePayment := application.NewPostSensor(&mySQL)
	return controllers.NewPostSensorController(caseCreatePayment)
}
