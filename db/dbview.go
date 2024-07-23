package dbview

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "zarkius"
	password = "1234"
	dbname   = "go"
)

func main() {
	// Configuración de la cadena de conexión
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Conexión a la base de datos
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado exitosamente a la base de datos")
}
