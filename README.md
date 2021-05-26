### API Rest en Golang.

**Developers:**

**Judith Alejandra Candelaria Sánchez 1723541**

**Gretel Gabrielle González Rodríguez 1805315**

**Marco Antonio Vázquez Rivera 1678576**

### ___________________________________________________________________________________________________

Para hace un uso correcto de esta API hay que cumplir con los siguientes requerimientos:

•	Instalar SQL Server Management Studio para la base de datos.

•	Instalar Visual Studio Code para el proyecto de Go de la API.

•	Instalar Git para instalar librerías necesarias para el código del proyecto de Go de la API.

•	Instalar Golang para el proyecto de Go de la API.

•	Descargar o clonar el proyecto desde el repositorio.

•	Ejecutar el script de la base de datos (carpeta database) en SQL Server Management Studio.

•	Ejecutar el proyecto de Go de la API (carpeta backend) en Visual Studio Code.

•	Instalar Postman para realizar pruebas de las peticiones.

•	Importar la colección de peticiones en Postman (carpeta request collection postman) para realizar pruebas.


Es necesario que correr el servidor de la base de datos y el servidor de la API al mismo tiempo. A continuación, se explicará cómo hacerlo.



## Procedimiento levantar el servidor de la base de datos y ejecutar el script de la base de datos.

Una vez que hayamos descargado el proyecto desde el repositorio procederemos a abrir SQL Server Management Studio.
Aparecerá esta ventana de login, hay que asegurarnos que el nombre del servidor sea localhost.
 
Procedemos a dar clic en Connect para levantar el servidor de la base de datos

Una vez que hayamos abierto el script, procederemos a ejecutarlo de la siguiente manera.
Seleccionaremos únicamente la primera línea del script y daremos clic en Execute.
 
Después de ejecutar la primera línea, seleccionaremos el resto del código, desde la segunda línea hasta la línea final del script y daremos clic en Execut.
 

Una vez ejecutado el script correctamente la base de datos y sus entidades aparecerán en el Object Explorer.
 
A partir de este momento, el servidor ya está corriendo, la base de datos ya está creada y ya tenemos registros dentro de ella (el script ya nos proporciona registros con la sentencia INSERT INTO). El servidor se encuentra corriendo el puerto 1433 y el nombre de la base de datos es Manhattam.

Al momento de cerrar SQL Server Management Studio, el servidor se desconecta, por lo que para volver a levantar el servidor solo hay que volver a abrir SQL Server Management Studio e iniciar sesión como el primer paso.
 
No es necesario realizar todo el procedimiento anterior, ya que la base de datos ya ha sido creada.



## Procedimiento conectar la API con la base de datos, ejecutar el código y levantar el servidor de la API.

Una vez que hayamos levantado el servidor de la base de datos y ejecutado el script de la base de datos, procederemos a abrir Visual Studio Code para ejecutar el proyecto de Go de la API y levantar el servidor de la API.
 
El primer paso es realizar la conexión al servidor de la base de datos y a la base de datos. Procederemos a abrir el archivo datacontext.go y localizamos la línea 13 de código.
 
Proporcionaremos en esta línea de código el puerto donde está ejecutándose la base de datos (por default es el puerto 1433) y el nombre de la base de datos (Manhattam).
Una vez que hayamos proporcionado el puerto donde está ejecutándose la base de datos y el nombre de la base de datos, procederemos a ejecutar el proyecto, abriendo la terminal de Visual Studio Code con la combinación de teclas Ctrl+Ñ y escribiendo el comando ***go run .*** como se muestra en la siguiente captura.
 
 
Damos clic en Permitir acceso, y a partir de este momento la API ya está ejecutándose en el puerto 1323. 



## Peticiones HTTP 

A continuación, mostraremos cómo realizar peticiones a la API de manera correcta, proporcionando una documentación de cada petición http y endpoints disponibles.

### Públicas

***POST login***

Loguea en la API y genera un Token.

•	URL: localhost:1323/login?UserName=&Password=

Esta es una petición pública.

Esta peticion POST login permite iniciar sesión con un usuario y contraseña existentes en la base de datos y generar un token para la autenticación bearer token.

Este token se requiere para realizar las peticiones http a la base de datos que se mostrarán a continuación. Es necesario guardar este token y agregarlo en las peticiones en la parte de Autenticación de tipo Bearer token.

Request Params:

o	UserName: nombre de usuario registrado en la base de datos.

o	Password: contraseña del usuario registrado en la base de datos.

***POST signUp***

Sign up (Crea usuario).

•	URL: localhost:1323/signUp?UserName=&Password=

Esta petición es pública.

Esta petición POST signUp permite crear un nuevo usuario en la base de datos, por lo que el usuario que sea creado ya estará registrado y podrá generar un token para realizar peticiones a la base de datos.

Request Params:

o	UserName: nombre de usuario para registrar en la base de datos.

o	Password: contraseña del usuario para registrar en la base de datos.

### Todas las funciones que requieren Authorization Bearer Token necesitan colocar en la petición el token generado al momento de loguearse.

### personalInfo

***GET personalInfo***

Trae todos los registros en PersonalInfo.

•	URL: localhost:1323/jwt/personalInfo

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición GET personalInfo permite consultar todos los registros de la entidad PersonalInfo de la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

***POST personalInfo***

Crea un registro en PersonalInfo.

•	URL:localhost:1323/jwt/personalInfo?CURP=&RFC=&NSS=&Names=&LastNames=&BirthDate=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición POST personalInfo permite crear un registro en la entidad PersonalInfo en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	CURP: CURP a registrar.

o	RFC: RFC a registrar.

o	NSS: NSS a registrar.

o	Names: Nombres a registrar.

o	LastNames: Apellidos a registrar.

o	BirthDate: Fecha de nacimiento a registrar.

***PUT personalInfo***

Modifica un registro en PersonalInfo.

•	URL:localhost:1323/jwt/personalInfo?CURP=&RFC=&NSS=&Names=&LastNames=&BirthDate=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición PUT personalInfo permite modificar un registro en la entidad PersonalInfo en la base de datos.

Los parámetros requeridos son los datos que serán modificados y los datos que no serán modificados también.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	CURP: CURP de la persona de datos a modificar.

o	RFC: RFC modificada.

o	NSS: NSS modificado.

o	Names: Nombres modificados.

o	LastNames: Apellidos modificados

o	BirthDate: Fecha de nacimiento modificada.

***DEL personalInfo***

Elimina un registro en PersonalInfo.

•	URL: localhost:1323/jwt/personalInfo?CURP=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición DELETE personalInfo permite eliminar un registro en la entidad PersonalInfo en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	CURP: CURP de registro a eliminar

### addresses
 
***GET addresses***

Trae todos los registros en Addresses.

•	URL: localhost:1323/jwt/addresses

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición GET addresses permite consultar todos los registros de la entidad Addresses de la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.


***POST addresses***

Crea un registro en Addresses.

•	URL:localhost:1323/jwt/addresses?Name=&Street=&Number=&Neighborhood=&ZipCode=&City=&State=&Country=&CURP=

•	Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición POST addresses permite crear un registro en la entidad Addresses en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	Name: Nombre del lugar a registrar.

o	Street: Calle de la dirección a registrar.

o	Number: Número de la dirección a registrar.

o	Neighborhood: Colonia de la dirección a registrar.

o	ZipCode: Código postal de la dirección a registrar.

o	City: Ciudad de la dirección a registrar.

o	State: Estado de la dirección a registrar.

o	Country: País de la dirección a registrar.

o	CURP: CURP de la persona a quien pertenece la dirección. 

 
***PUT addresses***

Modificar un registro en Addresses.

•	URL:localhost:1323/jwt/addresses?ID=&Name=&Street=&Number=&Neighborhood=&ZipCode=&City=&State=&Country=&CURP=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición PUT addresses permite modificar un registro en la entidad Addresses en la base de datos.

Los parámetros requeridos son los datos que serán modificados y los datos que no serán modificados también.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	ID: ID de dirección a modificar.

o	Name: Nombre de la dirección modificada.

o	Street: Calle modificada.

o	Number: Número modificado.

o	Neighborhood: Colonia modificada.

o	ZipCode: Código postal modificado.

o	City: Ciudad modificada.

o	State: estado modificado.

o	Country: País modificado.

o	CURP: CURP de la persona a quien pertenece la dirección.


***DEL addresses***

Elimina un registro en Addresses.

•	URL: localhost:1323/jwt/addresses?ID=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición DELETE addresses permite eliminar un registro en la entidad Addresses en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token>

Request Params:

o	ID: ID de dirección a eliminar.

### emails

***GET emails***

Trae todos los registros en Emails.

•	URL: localhost:1323/jwt/emails

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición GET emails permite consultar todos los registros de la entidad Emails de la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.


***POST emails***

Crea un registro en Emails.

•	URL: localhost:1323/jwt/emails?Email=&CURP=

•	Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición POST emails permite crear un registro en la entidad Emails en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	Email: Email a registrar.

o	CURP: CURP de la persona a quien pertenece el Email.


***PUT emails***

Modificar un registro en Emails.

•	URL: localhost:1323/jwt/emails?Email=&NewEmail=&CURP=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición PUT emails permite modificar un registro en la entidad Emails en la base de datos.

Los parámetros requeridos son los datos que serán modificados y los datos que no serán modificados.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	Email: Email a modificar.

o	NewEmail: Email modificado

o	CURP: CURP de la persona a quien pertenece el Email.


***DEL emails***

Elimina un registro en Emials.

•	URL: localhost:1323/jwt/emails?Email=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición DELETE emails permite eliminar un registro en la entidad Emails en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	Email: Email a eliminar.

### phones

***GET phones***

Trae todos los registros en Phones.

•	URL: localhost:1323/jwt/phones

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición GET phones permite consultar todos los registros de la entidad Phones de la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.


***POST phones***

Crea un registro en Phones.

•	URL: localhost:1323/jwt/phones?PhoneNumber=&Place=&CURP=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición POST phones permite crear un registro en la entidad Phones en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params

o	PhoneNumber: Número de teléfono a registrar.

o	Place: Lugar del número a registrar.

o	CURP: CURP de la persona a quien pertenece el número.


***PUT phones***

Modificar un registro en Phones.

•	URL: localhost:1323/jwt/phones?PhoneNumber=&NewPhone=&NewPlace=&CURP=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición PUT phones permite modificar un registro en la entidad Phones en la base de datos.

Los parámetros requeridos son los datos que serán modificados y los datos que no serán modificados.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	PhoneNumber: Número de teléfono a modificar.

o	NewPhone: Número de teléfono modificado.

o	NewPlace: Lugar del número modificado.

o	CURP: CURP de la persona a quien pertenece el número.


***DEL phones***

Elimina un registro en Phones.

•	URL: localhost:1323/jwt/phones?PhoneNumber=&CURP=

Esta petición requiere Bearer Token. (Iniciar sesión, generar token y utilizarlo en las peticiones).

Esta petición DELETE phones permite eliminar un registro en la entidad Phones en la base de datos.

El bearer token generado debe ser colocado en Authorization Bearer Token en la petición.

•	Authorization: Bearer Token

•	Token: <token> generado al logearse.

Request Params:

o	PhoneNumber: Número de teléfono a eliminar.

o	CURP: CURP de la persona a quien pertenece el número.
 


 

