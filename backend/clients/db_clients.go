package clients

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"backend/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"errors"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	log.Println("Initializing database...")
	//dsn := "root:admin@tcp(127.0.0.1:3306)/proyecto?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	log.Println("Database connected successfully")

	Migrate()
	SeedDB()
}

func Migrate() {
	log.Println("Migrating database...")
	DB.AutoMigrate(&dao.Usuario{}, &dao.Course{}, &dao.Subscription{}, &dao.Chat{})
}

func hashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SeedDB() {
	log.Println("Seeding database...")

	// Hashear las contraseñas
	adminPassword := hashPassword("admin")
	userPassword := hashPassword("user")

	admin := dao.Usuario{NombreUsuario: "admin", Contrasena: adminPassword, Tipo: "admin"}
	user := dao.Usuario{NombreUsuario: "user", Contrasena: userPassword, Tipo: "normal"}

	DB.FirstOrCreate(&admin, dao.Usuario{NombreUsuario: "admin"})
	DB.FirstOrCreate(&user, dao.Usuario{NombreUsuario: "user"})

	log.Println("Database seeded successfully")

	cursos := []dao.Course{
		{Nombre: "Ingles B2", Categoria: "Idiomas", Dificultad: "Medio", Precio: 45, Descripcion: "Curso de ingles avanzado para aprobar tus examenes internacionales", ImageURL: "https://diarium.usal.es/ireneigls/files/2018/09/b2-de-ingles.jpg"},
		{Nombre: "Hacking Etico", Categoria: "Programación", Dificultad: "Dificil", Precio: 60, Descripcion: "Curso de como ser hacker sin ser hacker", ImageURL: "https://www.pmg-ssi.com/wp-content/uploads/2023/08/he.jpg"},
		{Nombre: "Rainbow six", Categoria: "Videojuegos", Dificultad: "Muy Dificil", Precio: 100, Descripcion: "Curso de como aprender a jugar al mejor juego de Ubisoft", ImageURL: "https://media.tycsports.com/files/2021/08/13/319515/rainbow-six-siege-gratis_1440x810_wmk.jpg"},
		{Nombre: "Padel principiante", Categoria: "Deportes", Dificultad: "Facil", Precio: 20, Descripcion: "Curso de como aprender padel en 4 meses", ImageURL: "https://upload.wikimedia.org/wikipedia/commons/6/6c/Agustin-tapia.jpg"},
		{Nombre: "Simulador iracing", Categoria: "Videojuegos", Dificultad: "Muy difil", Precio: 30, Descripcion: "Curso para parecerse a Max Verstappen", ImageURL: "https://i.blogs.es/4b09e7/captura-de-pantalla-2022-05-25-a-las-13.48.10/1366_2000.jpeg"},
		{Nombre: "Fotografia", Categoria: "Estilo de vida", Dificultad: "Medio", Precio: 62, Descripcion: "Curso de como hacer fotografias para diarios y revistas", ImageURL: "https://pqs.pe/wp-content/uploads/2015/10/pqs-fotografia.jpg"},
		{Nombre: "Cocteleria", Categoria: "Estilo de vida", Dificultad: "Facil", Precio: 18, Descripcion: "Curso pricipiante de como hacer los coctele mas faciles de argentina", ImageURL: "https://media.revistagq.com/photos/5df397b1b1280d00088d986e/16:9/w_2560%2Cc_limit/iStock-680866296.jpg"},
		{Nombre: "Cocina", Categoria: "Estilo de vida", Dificultad: "Medio", Precio: 30, Descripcion: "Curso de prepareacion para MASTERCHEF", ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTqbKsTAP_QuCXJ3kfxrwVSvSy2VgW-QJY1kA&s"},
		{Nombre: "Ceramica", Categoria: "Manualidades", Dificultad: "dificil", Precio: 23, Descripcion: "Curso provincial de ceramica familiar", ImageURL: "https://www.conasi.eu/blog/wp-content/uploads/2020/09/ceramica-y-porcelana-para-cocinar-1111-1.jpg"},
		{Nombre: "FIFA", Categoria: "Videojuegos", Dificultad: "Facil", Precio: 10, Descripcion: "Curso de como jugar al FIFA de manera profesional", ImageURL: "https://cloudfront-us-east-1.images.arcpublishing.com/infobae/SWFBFT2L5JDJTD4KKOYIU6FDZY.jpg"},
	}

	for _, curso := range cursos {
		DB.Create(&curso)
	}
	log.Println("Database seeded successfully")

}

func CreateUser(nombreUsuario, contrasena, tipo string) error {
	user := dao.Usuario{
		NombreUsuario: nombreUsuario,
		Contrasena:    contrasena,
		Tipo:          tipo,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	result := DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SearchUser busca un usuario en la base de datos por nombre de usuario.
func SearchUser(nombreUsuario string) error {
	var user dao.Usuario
	result := DB.Where("nombre_usuario = ?", nombreUsuario).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}
	return nil
}
