package domain

import (
	_ "github.com/mattn/go-sqlite3"
)

/*func hashMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func verifyPassword(db *sql.DB, username, password string) bool {
	var hashedPassword string
	query := `SELECT contraseña FROM usuarios WHERE nombre_usuario_admin = ?`
	err := db.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return hashedPassword == hashMD5(password)
}

func main() {
	db, err := sql.Open("sqlite3", "./cursos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	username := "ArnonNahmias"
	password := "Arnon123"
	if verifyPassword(db, username, password) {
		fmt.Println("Contraseña correcta")
	} else {
		fmt.Println("Contraseña incorrecta")
	}
*/

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
