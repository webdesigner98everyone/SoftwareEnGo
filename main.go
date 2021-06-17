/*Creado por Luis Alberto Forero
cali-valle el 16-06-2021*/
package main

import (
	//"log"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

//Funcion conexion a la BD
func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "sistema"
	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

// esto sirve para obtener la informacion en una carpeta en este caso plantillas
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	//Solicitud para acceder a la funcion inicio
	http.HandleFunc("/", Inicio)
	//Solicitud para acceder a la funcion crear
	http.HandleFunc("/crear", Crear)
	//Solicitud para acceder a la funcion insertar
	http.HandleFunc("/insertar", Insertar)
	//Solicitud para acceder a la funcion Borrar
	http.HandleFunc("/borrar", Borrar)
	//Solicitud para acceder a la funcion Editar
	http.HandleFunc("/editar", Editar)
	//Solicitud para acceder a la funcion Actualizar
	http.HandleFunc("/actualizar", Actualizar)
	//Respuesta del servidor con puerto 8080
	fmt.Println("Servidor corriendo....")
	http.ListenAndServe(":8080", nil)
}

//estructura de la tabla empleado
type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

//funcion plantilla inicio
func Inicio(w http.ResponseWriter, r *http.Request) {
	//Cnexion a BD
	conexionEstablecida := conexionBD()
	//mostrar registros
	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")
	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	aregloEmpleado := []Empleado{}

	//recorremos los datos de la sentencia SELECT
	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
		aregloEmpleado = append(aregloEmpleado, empleado)
	}
	//Mostramos por consola todos los registros almacenados en la BD por medio de una matriz
	//fmt.Println(aregloEmpleado)

	//mostramos nuestra plantilla tmpl que creamos en la carpeta plantillas
	plantillas.ExecuteTemplate(w, "inicio", aregloEmpleado)
}

//funcion plantilla Editar
func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")

	//Conexion a BD
	conexionEstablecida := conexionBD()
	//sentencia SQL para editar el registro por medio del ID
	registro, err := conexionEstablecida.Query("SELECT * FROM empleados WHERE id=? ", idEmpleado)

	empleado := Empleado{}
	for registro.Next() {
		var id int
		var nombre, correo string
		err = registro.Scan(&id, &nombre, &correo)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
	}
	plantillas.ExecuteTemplate(w, "editar", empleado)
}

//funcion plantilla crear
func Crear(w http.ResponseWriter, r *http.Request) {
	//mostramos nuestra plantilla tmpl que creamos en la carpeta plantillas
	plantillas.ExecuteTemplate(w, "crear", nil)
}

//funcion plantilla insertar
func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//agregamos los campos que asignamos en crear.tmpl
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		//Conexion a BD
		conexionEstablecida := conexionBD()
		//Sentencia SQL Inserta Registro
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, correo)
		//Redireccionamos a la pagina inicial
		http.Redirect(w, r, "/", 301)
	}
}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//agregamos los campos que asignamos en crear.tmpl
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		//Conexion a BD
		conexionEstablecida := conexionBD()
		//Sentencia SQL Actualizar Registro
		modificarRegistros, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?, correo=? WHERE id=? ")
		if err != nil {
			panic(err.Error())
		}
		modificarRegistros.Exec(nombre, correo, id)
		//Redireccionamos a la pagina inicial
		http.Redirect(w, r, "/", 301)
	}
}

//funcion plantilla Borrar
func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	//Conexion a BD
	conexionEstablecida := conexionBD()
	//Sentencia SQL elimina Registro
	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM empleados WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	borrarRegistro.Exec(idEmpleado)
	//Redireccionamos a la pagina inicial
	http.Redirect(w, r, "/", 301)
}
