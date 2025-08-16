package database

import "database/sql"

// User representa a entidade do banco
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// UserRepository encapsula o DB
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository cria um repositório a partir da conexão
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create insere um novo usuário no banco
func (r *UserRepository) Create(user *User) error {
	result, err := r.db.Exec("INSERT INTO users(name) VALUES(?)", user.Name)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	return err
}

// FindAll retorna todos os usuários
func (r *UserRepository) FindAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// FindByID retorna um usuário pelo ID
func (r *UserRepository) FindByID(id int64) (*User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)
	var u User
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

// Delete remove um usuário pelo ID
func (r *UserRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
