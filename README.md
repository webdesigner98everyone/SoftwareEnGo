# sistema
Ten encuenta los siguientes pasos para ejecutar el software desarrollado en go.....

---------------------------------------------------------------------------------------------------------------------------------
1. Tener el servidor Xampp ya que estamos utilizando una bd mysql local
2. Configura los puertos de tu xampp quedando asi con los puertos 4040 ya que la aplicacion esta corriendo con el puerto 8080 y esto puedo generar conflicto al momento de ejecutar la BD
3. una vez tengamos mysql ejecuta en nuestro xampp importa la BD llamada sistema.sql ubicada en la carpeta Bd de nuestro proyecto, esta sera la BD que estaremos manejando.
4. Valida de que las tablas esten ejecutadas dentro de nuestra BD exportada.
5. Una vez tengamos nuestra BD lista ubicaremos nuestro software llamado sistema en disco c: (C:\sistema), preferiblemente.
6. Una vez ubiquemos nuestro software desarrollado en go, lo abriremos en nuestro editor de codigo Visual-Code
7. Listos????, ya con el software en visual Code, observaremos la estructura de carpetas de nuestro software y vamos a presionar las siguientes combinancion de teclas para abrir nuestra terminal (Ctrl + ñ)
8. Ya con nuestra terminal en visual-code ejecutada daremos el siguiente comando para ejecutar nuestro servidor de nuestro software (go run main.go) <---------------Lo anterior nos ayudara a ejecutar nuestro software desarrollado en go
9. cuando ejecutemos nuestro comando en la terminal veremos en consola un msj que dice: "Servidor corriendo....", esto indicara que nuestro software no tuvo errores locales y en breve nos aparece una pantalla del firewall para asignarle los permisos como administrador para ejecutar nuestro main.go (No te preocupes por esta autorizacion de firewall, esto ocurre debido a que nuestro equipo no esta aconstumbrado a ejecutar esta clase de .exe)
10. Ya como ultimo paso una vez autoricemos nuestro firewall para ejecutra nuestro software vamos a nuestro navegador y escriberemos la siguiente ruta par ver en pantalla nuestro software (http://localhost:8080/)

Hemos Terminado!!!!!! es hora de que ingreses datos y puedas agregar, editar, eliminar registros y ten encuenta que todo lo que hagas se reflejara en nuestra BD de mysql.



HERRAMIENTAS:
1. Xampp -- BDmysql
2. Visual-Code