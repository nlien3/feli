package clients

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"backend/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"errors"
	"time"
)

var DB *gorm.DB

func InitDB() {
	log.Println("Initializing database...")
	dsn := "root:@tcp(127.0.0.1:3306)/proyecto4?charset=utf8mb4&parseTime=True&loc=Local"
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

	// Crear un nuevo usuario sin hashear la contraseña
	plainPassword := "queso"
	quesoUser := dao.Usuario{NombreUsuario: "queso", Contrasena: plainPassword, Tipo: "normal"}

	DB.FirstOrCreate(&quesoUser, dao.Usuario{NombreUsuario: "queso"})

	log.Println("Database seeded successfully")
/*
	cursos := []dao.Course{
		{Nombre: "Ingles B2", Categoria: "Idiomas", Dificultad: "Medio", Precio: 45, Descripcion: "Curso de ingles avanzado para aprobar tus examenes internacionales", ImageURL: "https://diarium.usal.es/ireneigls/files/2018/09/b2-de-ingles.jpg"},
		{Nombre: "Hacking Etico", Categoria: "Programación", Dificultad: "Dificil", Precio: 60, Descripcion: "Curso de como ser hacker sin ser hacker", ImageURL: "https://www.pmg-ssi.com/wp-content/uploads/2023/08/he.jpg"},
	}
	for _, curso := range cursos {
		DB.Create(&curso)
	}
	log.Println("Database seeded successfully")
	*/
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
