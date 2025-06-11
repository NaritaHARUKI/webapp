package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID      int    `db:"id"`
	Mime    string `db:"mime"`
	Imgdata []byte `db:"imgdata"`
}

func mimeToExt(mime string) string {
	switch mime {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	default:
		return ""
	}
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		getEnv("ISUCONP_DB_USER", "root"),
		getEnv("ISUCONP_DB_PASSWORD", "root"),
		getEnv("ISUCONP_DB_HOST", "localhost"),
		getEnv("ISUCONP_DB_PORT", "3306"),
		getEnv("ISUCONP_DB_NAME", "isuconp"),
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}
	defer db.Close()

	var posts []Post
	err = db.Select(&posts, "SELECT id, mime, imgdata FROM posts")
	if err != nil {
		log.Fatalf("DBクエリ失敗: %v", err)
	}

	outputDir := "/home/public/images"
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("ディレクトリ作成失敗: %v", err)
	}

	count := 0
	for _, p := range posts {
		ext := mimeToExt(p.Mime)
		if ext == "" {
			log.Printf("対応していない MIME タイプ: %s", p.Mime)
			continue
		}

		filename := fmt.Sprintf("%d%s", p.ID, ext)
		path := filepath.Join(outputDir, filename)
		err = ioutil.WriteFile(path, p.Imgdata, 0644)
		if err != nil {
			log.Printf("画像書き込み失敗 %s: %v", path, err)
			continue
		}
		count++
	}

	log.Printf("✅ 完了: %d 件の画像を %s に保存しました", count, outputDir)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
