package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// テーブル名は自動的にtodosになる.
type Todo struct {
	gorm.Model // Gormの標準モデル.id,created_at,updated_at,deleted_atが含まれる.
	Text string
	Status string
}

//DBマイグレート
func dbInit() {
	// gorm.Open(使用するDBの種類, ファイル名)
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		// panic() - 処理を中断して終了する
		panic("データベース開けない(dbInit)")
	}
	// migrate実行 ファイルがなければ生成,あれば何もしない.
	db.AutoMigrate(&Todo{})
	// defer - 終了時に実行する処理を定義できる.
	defer db.Close()
}

//追加
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けない(dbInsert)")
	}
	// 構造体で指定してレコード追加
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けない(dbUpdate)")
	}
	// todoの型の構造体で,IDがidのレコードを取得
	// todoには見つかったレコードが格納される
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// 削除
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けない(dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//全件取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けない(dbGetAll)")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けない(dbGetOne)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
