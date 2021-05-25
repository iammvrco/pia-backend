package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Declaración de modelos de las entidades
//Utilizamos gorm column para especificar el nombre de la columna
//ya que GORM genera un nombre cambiando espacios por guiones.
type (
	Users struct {
		UserName string `gorm:"column:UserName" json:"UserName"`
		Password string `json:"Password"`
	}

	PersonalInfo struct {
		CURP      string `json:"CURP"`
		RFC       string `json:"RFC"`
		NSS       string `json:"NSS"`
		Names     string `json:"Names"`
		LastNames string `gorm:"column:LastNames" json:"LastNames"`
		BirthDate string `gorm:"column:BirthDate" json:"BirthDate"`
	}

	Addresses struct {
		ID           string `json:"ID"`
		Name         string `json:"Name"`
		Street       string `json:"Street"`
		Number       string `json:"Number"`
		Neighborhood string `json:"Neighborhood"`
		ZipCode      string `gorm:"column:ZipCode" json:"ZipCode"`
		City         string `json:"City"`
		State        string `json:"State"`
		Country      string `json:"Country"`
		CURP         string `json:"CURP"`
	}

	Emails struct {
		Email string `json:"Email"`
		CURP  string `json:"CURP"`
	}

	Phones struct {
		PhoneNumber string `gorm:"column:PhoneNumber" json:"PhoneNumber"`
		Place       string `json:"Place"`
		CURP        string `json:"CURP"`
	}
)

//En esta sección utilizamos estas funciones para especificar
//el nombre de las entidades de la base de datos, ya que GORM
//pluraliza los nombres y separa los espacios por guiones
func (PersonalInfo) TableName() string {
	return "PersonalInfo"
}
func (Addresses) TableName() string {
	return "Addresses"
}
func (Emails) TableName() string {
	return "Emails"
}
func (Phones) TableName() string {
	return "Phones"
}

//Servicios

//PUBLICOS
func signUp(c echo.Context) error {
	//Esta función crea un nuevo registro de la entidad Users en la base de datos
	//para iniciar sesión en la aplicación, generar un token y realizar peticiones
	//a la base de datos

	//Lee los valores para la entidad Users desde la url
	user := new(Users)
	user.UserName = c.FormValue("UserName")
	user.Password = c.FormValue("Password")
	//NO ES REQUISITO USAR HASHING

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Crea el registro en la base de datos
	db.Create(&user)

	//Regresa el status y el objeto creado
	return c.JSON(http.StatusCreated, user)
}

func login(c echo.Context) error {
	//Esta función loggea al usuario en la aplicación y genera
	//un token para poder realizar peticiones a la base de datos
	//Este token se utiliza en las peticiones, por lo que es
	//necesario copiarlo en Authorized Bearer Token de la petición

	//Lee los valores de la entidad Users desde la url
	userName := c.FormValue("UserName")
	password := c.FormValue("Password")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var user Users

	//Busca los valores ingresados en la base de datos
	db.First(&user, "UserName = ?", userName)

	//Arroja un error si no existe el usuario en la base de datos
	if user.UserName != userName || user.Password != password {
		return echo.ErrUnauthorized
	}

	// Crea el token
	token := jwt.New(jwt.SigningMethodHS256)

	//Genera los claims del token
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.UserName
	claims["type"] = "Admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//Genera el token codificado
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	//Regresa el bearer token que debemos guardar para utilizar
	//en las peticiones
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

//PERSONALINFO
func getAllPersonalInfo(c echo.Context) error {
	//Esta función regresa todos los registros de la entidad
	//PersonalInfo en la base de datos

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var res []PersonalInfo

	//Busca y regresa todos los registros de tipo PersonalInfo
	//en la base de datos
	db.Find(&res)

	//Regresa el status y la lista de objetos encontrados
	return c.JSON(http.StatusOK, res)
}

func createPersonalInfo(c echo.Context) error {
	//Esta función crea un nuevo registro de la entidad PersonalInfo
	//en la base de datos

	//Lee los valores para la entidad PersonalInfo desde la url
	personalInfo := new(PersonalInfo)
	personalInfo.CURP = c.FormValue("CURP")
	personalInfo.RFC = c.FormValue("RFC")
	personalInfo.NSS = c.FormValue("NSS")
	personalInfo.Names = c.FormValue("Names")
	personalInfo.LastNames = c.FormValue("LastNames")
	personalInfo.BirthDate = c.FormValue("BirthDate")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Crea el registro en la base de datos
	db.Create(&personalInfo)

	//Regresa el status y el objeto creado
	return c.JSON(http.StatusCreated, personalInfo)
}

func updatePersonalInfo(c echo.Context) error {
	//Esta función modifica un registro de la entidad PersonalInfo
	//en la base de datos

	//Lee la llave primaria del registro que se va a modificar
	//desde la url
	CURP := c.FormValue("CURP")

	//Lee los valores para la entidad desde la url
	newPersonalInfo := new(PersonalInfo)
	newPersonalInfo.RFC = c.FormValue("RFC")
	newPersonalInfo.NSS = c.FormValue("NSS")
	newPersonalInfo.Names = c.FormValue("Names")
	newPersonalInfo.LastNames = c.FormValue("LastNames")
	newPersonalInfo.BirthDate = c.FormValue("BirthDate")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var personalInfo PersonalInfo

	//Mantenemos el mismo CURP para el objeto modificado
	personalInfo.CURP = CURP

	//Modifica el registro que cumpla la condición WHERE omitiendo el campo CURP
	//e introduciendo los datos modificados (también datos que no serán modificados y,
	//por lo tanto, siguen siendo los mismos)
	db.Model(&personalInfo).Where("CURP = ?", CURP).Omit("CURP").Updates(PersonalInfo{RFC: newPersonalInfo.RFC, NSS: newPersonalInfo.NSS, Names: newPersonalInfo.Names, LastNames: newPersonalInfo.LastNames, BirthDate: newPersonalInfo.BirthDate})

	//Regresa el status y el objeto modificado
	return c.JSON(http.StatusOK, personalInfo)
}

func deletePersonalInfo(c echo.Context) error {
	//Esta función elimina el registro de la entidad PersonalInfo
	//en la base de datos

	//Lee la llave primaria del registro que será eliminado desde la url
	CURP := c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var personalInfo PersonalInfo

	//Elimina el registro que cumpla con la condición WHERE
	db.Where("CURP = ?", CURP).Delete(&personalInfo)

	//Regresa el status NoContent
	return c.NoContent(http.StatusNoContent)
}

//ADDRESSES
func getAllAddresses(c echo.Context) error {
	//Esta función regresa todos los registros de la entidad
	//Addresses en la base de datos

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var res []Addresses

	//Busca y regresa todos los registros de tipo Addresses
	//en la base de datos
	db.Find(&res)

	//Regresa el status y la lista de objetos encontrados
	return c.JSON(http.StatusOK, res)
}

func createAddress(c echo.Context) error {
	//Esta función crea un nuevo registro de la entidad Addresses
	//en la base de datos

	//Lee los valores para la entidad Addresses desde la url
	address := new(Addresses)
	address.Name = c.FormValue("Name")
	address.Street = c.FormValue("Street")
	address.Number = c.FormValue("Number")
	address.Neighborhood = c.FormValue("Neighborhood")
	address.ZipCode = c.FormValue("ZipCode")
	address.City = c.FormValue("City")
	address.State = c.FormValue("State")
	address.Country = c.FormValue("Country")
	address.CURP = c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Crea el registro en la base de datos
	//omitiendo el ID ya que es autoincrementable en la
	//base de datos
	db.Omit("ID").Create(&address)

	//Regresa el status y el objeto creado
	return c.JSON(http.StatusCreated, address)
}

func updateAddress(c echo.Context) error {
	//Esta función modifica un registro de la entidad Addresses
	//en la base de datos

	//Lee la llave primaria del registro que se va a modificar
	//desde la url
	ID := c.FormValue("ID")

	//Lee los valores para la entidad desde la url
	newAddress := new(Addresses)
	newAddress.Name = c.FormValue("Name")
	newAddress.Street = c.FormValue("Street")
	newAddress.Number = c.FormValue("Number")
	newAddress.Neighborhood = c.FormValue("Neighborhood")
	newAddress.ZipCode = c.FormValue("ZipCode")
	newAddress.City = c.FormValue("City")
	newAddress.State = c.FormValue("State")
	newAddress.Country = c.FormValue("Country")
	newAddress.CURP = c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var address Addresses

	//Mantenemos el mismo CURP para el objeto modificado
	address.CURP = newAddress.CURP
	//Mantenemos el mismo ID para el objeto modificado
	address.ID = ID

	//Modifica el registro que cumpla la condición WHERE omitiendo el campo ID
	//y CURP e introduciendo los datos modificados (también datos que no serán
	//modificados y, por lo tanto, siguen siendo los mismos)
	db.Model(&address).Where("ID = ?", ID).Omit("ID", "CURP").Updates(Addresses{Name: newAddress.Name, Street: newAddress.Street, Number: newAddress.Number, Neighborhood: newAddress.Neighborhood, ZipCode: newAddress.ZipCode, City: newAddress.City, State: newAddress.State, Country: newAddress.Country})

	//Regresa el status y el objeto modificado
	return c.JSON(http.StatusOK, address)
}

func deleteAddress(c echo.Context) error {
	//Esta función elimina el registro de la entidad Addresses
	//en la base de datos

	//Lee la llave primaria del registro que será eliminado desde la url
	ID := c.FormValue("ID")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var address Addresses

	//Elimina el registro que cumpla con la condición WHERE
	db.Where("ID = ?", ID).Delete(&address)

	//Regresa el status NoContent
	return c.NoContent(http.StatusNoContent)
}

//EMAILS
func getAllEmails(c echo.Context) error {
	//Esta función regresa todos los registros de la entidad
	//Emails en la base de datos

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var res []Emails

	//Busca y regresa todos los registros de tipo Emails
	//en la base de datos
	db.Find(&res)

	//Regresa el status y la lista de objetos encontrados
	return c.JSON(http.StatusOK, res)
}

func createEmail(c echo.Context) error {
	//Esta función crea un nuevo registro de la entidad Emails
	//en la base de datos

	//Lee los valores para la entidad Emails desde la url
	email := new(Emails)
	email.Email = c.FormValue("Email")
	email.CURP = c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Crea el registro en la base de datos
	db.Create(&email)

	//Regresa el status y el objeto creado
	return c.JSON(http.StatusCreated, email)
}

func updateEmail(c echo.Context) error {
	//Esta función modifica un registro de la entidad Emails
	//en la base de datos

	//Lee la llave primaria del registro que se va a modificar
	//desde la url
	Email := c.FormValue("Email")

	//Lee los valores para la entidad desde la url
	newEmail := new(Emails)
	newEmail.Email = c.FormValue("NewEmail")
	newEmail.CURP = c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var email Emails

	//Mantenemos el mismo CURP para el objeto modificado
	email.CURP = newEmail.CURP

	//Modifica el registro que cumpla la condición WHERE e introduce
	//los datos modificados (también datos que no serán modificados y,
	//por lo tanto, siguen siendo los mismos)
	db.Model(&email).Where("Email = ?", Email).Update("Email", newEmail.Email)

	return c.JSON(http.StatusOK, email)
}

func deleteEmail(c echo.Context) error {
	//Esta función elimina el registro de la entidad Emails
	//en la base de datos

	//Lee la llave primaria del registro que será eliminado desde la url
	Email := c.FormValue("Email")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var email Emails

	//Elimina el registro que cumpla con la condición WHERE
	db.Where("Email = ?", Email).Delete(&email)

	//Regresa el status NoContent
	return c.NoContent(http.StatusNoContent)
}

//PHONES
func getAllPhones(c echo.Context) error {
	//Esta función regresa todos los registros de la entidad
	//Phones en la base de datos

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var res []Phones

	//Busca y regresa todos los registros de tipo Phones
	//en la base de datos
	db.Find(&res)

	//Regresa el status y la lista de objetos encontrados
	return c.JSON(http.StatusOK, res)
}

func createPhone(c echo.Context) error {
	//Esta función crea un nuevo registro de la entidad Phones
	//en la base de datos

	//Lee los valores para la entidad Phones desde la url
	phone := new(Phones)
	phone.PhoneNumber = c.FormValue("PhoneNumber")
	phone.Place = c.FormValue("Place")
	phone.CURP = c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Crea el registro en la base de datos
	db.Create(&phone)

	//Regresa el status y el objeto creado
	return c.JSON(http.StatusCreated, phone)
}

func updatePhone(c echo.Context) error {
	//Esta función modifica un registro de la entidad Phones
	//en la base de datos

	//Lee la llave primaria del registro que se va a modificar
	//desde la url
	PhoneNumber := c.FormValue("PhoneNumber")

	//Lee los valores para la entidad desde la url
	newPhone := new(Phones)
	newPhone.PhoneNumber = c.FormValue("NewPhone")
	newPhone.Place = c.FormValue("NewPlace")
	newPhone.CURP = c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var phone Phones

	//Mantenemos el mismo CURP para el objeto modificado
	phone.CURP = newPhone.CURP

	//Modifica el registro que cumpla la condición WHERE e introduce
	//los datos modificados (también datos que no serán modificados y,
	//por lo tanto, siguen siendo los mismos)
	db.Model(&phone).Where(db.Where("PhoneNumber = ?", PhoneNumber), db.Where("CURP = ?", newPhone.CURP)).Select("PhoneNumber", "Place").Updates(Phones{PhoneNumber: newPhone.PhoneNumber, Place: newPhone.Place})

	return c.JSON(http.StatusOK, phone)
}

func deletePhone(c echo.Context) error {
	//Esta función elimina el registro de la entidad Phones
	//en la base de datos

	//Lee la llave primaria del registro que será eliminado desde la url
	PhoneNumber := c.FormValue("PhoneNumber")
	CURP := c.FormValue("CURP")

	//Intenta conectarse a la base de datos para realizar la petición
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("FALLÓ CONEXIÓN A BDD", err)
	}

	//Declaración de objeto de tipo entidad que se utilizará en la petición
	var phone Phones

	//Elimina el registro que cumpla con las condiciones WHERE
	db.Where(db.Where("PhoneNumber = ?", PhoneNumber), db.Where("CURP = ?", CURP)).Delete(&phone)

	//Regresa el status NoContent
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	//Crea un grupo para las rutas de peticiones
	jwtGroup := e.Group("/jwt")

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//Configuración de CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200", "http://localhost:8080", "http://localhost:1433"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//Implementamos el bearer token para las peticiones del grupo jwtGroup
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
	}))

	//Definimos las rutas
	//Públicas
	e.POST("/login", login)
	e.POST("/signUp", signUp)

	//Reestringidas (Bearer token necesario)
	jwtGroup.GET("/personalInfo", getAllPersonalInfo)
	jwtGroup.POST("/personalInfo", createPersonalInfo)
	jwtGroup.PUT("/personalInfo", updatePersonalInfo)
	jwtGroup.DELETE("/personalInfo", deletePersonalInfo)

	jwtGroup.GET("/addresses", getAllAddresses)
	jwtGroup.POST("/addresses", createAddress)
	jwtGroup.PUT("/addresses", updateAddress)
	jwtGroup.DELETE("/addresses", deleteAddress)

	jwtGroup.GET("/emails", getAllEmails)
	jwtGroup.POST("/emails", createEmail)
	jwtGroup.PUT("/emails", updateEmail)
	jwtGroup.DELETE("/emails", deleteEmail)

	jwtGroup.GET("/phones", getAllPhones)
	jwtGroup.POST("/phones", createPhone)
	jwtGroup.PUT("/phones", updatePhone)
	jwtGroup.DELETE("/phones", deletePhone)

	//Ejecución del servidor en el puerto 1323
	e.Logger.Fatal(e.Start(":1323"))
}
