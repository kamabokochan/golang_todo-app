package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/src/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3" // アンダースコアをつけることで未使用による自動削除を回避
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// Sprintf: 任意の型と文字列をまとめて文字列化する

	// sql文
	// CREATE TABLE IF NOT EXISTS: テーブルがなければ作成する
	// AUTOINCREMENT: 自動増幅
	// UNIQUE: 重複を禁止
	cmdU := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        uuid STRING NOT NULL UNIQUE,
        name STRING,
        email STRING,
        password STRING,
        created_at DATETIME)`, tableNameUser)

	// コマンド実行
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRIGN,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// ハッシュ化
// 不可逆なためセキュリティ向上
// また、パスワードが複数ユーザで同一でもハッシュ化により異なるハッシュ値が生成でき、攻撃者の同時解析を困難にする
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
