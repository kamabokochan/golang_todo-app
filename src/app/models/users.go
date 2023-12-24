package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	// qmark style (?,?,?) 引数を渡している
	cmd := `insert into users (
        uuid,
        name,
        email,
        password,
        created_at) values (?, ?, ?, ?, ?)`

	// Dbは modelsパッケージ内のため、参照可能
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, 
		created_at from users where id = ?`

	// QueryRow: 単一行クエリ 1レコード分を取得
	// Scan: QueryRow から取得した1行の結果をスキャンして、変数に格納する
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	// セッション作成
	insertSessionCMD := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(insertSessionCMD, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	// user_idとemailが一致するセッションを取得
	getSessionCMD := `select id, uuid, email, user_id, created_at
		from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(getSessionCMD, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
		from sessions where uuid = ?`

	err = Db.QueryRow(cmd, s.UUID).Scan(
		&s.ID,
		&s.UUID,
		&s.Email,
		&s.UserID,
		&s.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if s.ID != 0 {
		valid = true
	}
	return valid, err
}

func (s *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, s.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at FROM users
	where id = ?`
	err = Db.QueryRow(cmd, s.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err
}
