package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/zarkius/go-2-1/db/dbview"
)

func init() {
	err := dbview.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
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
		users, err := getUsers()
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
		err := createUser(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// redirige al usuario a la lista de usuarios
		http.Redirect(w, r, "/usuarios", http.StatusSeeOther)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func createUser(r *http.Request) error {
	// obtiene los datos del formulario
	err := r.ParseForm()
	if err != nil {
		return err
	}

	// crea un nuevo usuario con los datos del formulario
	user := User{
		Name:  r.Form.Get("name"),
		Email: r.Form.Get("email"),
	}

	// guarda el usuario en la base de datos
	err = saveUser(user)
	if err != nil {
		return err
	}

	return nil
}

func getUsers() ([]User, error) {
	// obtiene los usuarios de la base de datos
	users, err := db.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func saveUser(user User) error {
	// guarda el usuario en la base de datos
	err := db.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}
