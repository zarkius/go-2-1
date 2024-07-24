package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL
)

// User representa un usuario
type User struct {
	Name  string
	Email string
}

// PageData representa los datos de la página
type PageData struct {
	Title   string
	Message string
}

func init() {
	// Configuración de la cadena de conexión
	host := "localhost"
	port := "5432"
	user := "zarkius"
	password := "1234"
	dbname := "go"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/usuarios", usuariosHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// crea una instancia de PageData
	data := PageData{
		Title:   "Mi Página",
		Message: "¡Hola, mundo!",
	}

	// parsea el archivo HTML
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// renderiza la plantilla con los datos
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func usuariosHandler(w http.ResponseWriter, r *http.Request) {
	// parsea el archivo HTML
	tmpl, err := template.ParseFiles("usuarios.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "GET":
		// obtiene los usuarios de la base de datos
		users, err := getUsers() //TODO
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// renderiza la plantilla con los datos
		err = tmpl.Execute(w, users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case "POST":
		// crea un nuevo usuario
		//err := createUser(r) // TODO
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		// redirige al usuario a la lista de usuarios
		http.Redirect(w, r, "/usuarios", http.StatusSeeOther)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func getUsers() ([]User, error) {
	print("Obteniendo usuarios...")
	return []User{}, nil
}
