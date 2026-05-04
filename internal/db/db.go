package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./caselog.db")
	if err != nil {
		return nil, err
	}

	if err := createProjectsTable(database); err != nil {
		return nil, err
	}

	if err := seedProjects(database); err != nil {
		return nil, err
	}

	return database, nil
}

func createProjectsTable(database *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		company TEXT,
		main_skill TEXT,
		unit_price INTEGER,
		remote_type TEXT,
		team_size TEXT,
		status TEXT NOT NULL DEFAULT 'considering',
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := database.Exec(query)
	return err
}

func seedProjects(database *sql.DB) error {
	var count int

	if err := database.QueryRow("SELECT COUNT(*) FROM projects").Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	query := `
	INSERT INTO projects (
		title,
		company,
		main_skill,
		unit_price,
		remote_type,
		team_size,
		status,
		notes
	) VALUES
		(
			'Webサービス開発支援',
			'サンプル株式会社',
			'Go / Gin / SQLite',
			70,
			'フルリモート',
			'3〜5名',
			'considering',
			'Goを用いたAPI開発とWeb画面の実装を行う想定のサンプル案件。案件条件、使用技術、選考状況を整理するためのデータです。'
		),
		(
			'SaaS管理画面開発支援',
			'非公開',
			'TypeScript / React / Go',
			75,
			'週4リモート',
			'4名',
			'interview',
			'Reactを用いた管理画面開発と、GoによるAPI改修を行う想定。少人数チームで設計から実装、改善まで関われる案件として登録しています。'
		),
		(
			'業務システム改善プロジェクト',
			'非公開',
			'C# / ASP.NET Core / React',
			68,
			'一部リモート',
			'5〜7名',
			'applied',
			'C#経験を活かしながら、ReactやAPI開発にも関われる想定の案件。保守中心ではなく、Web開発寄りの経験を積めるかを確認したい案件です。'
		);
	`

	_, err := database.Exec(query)
	return err
}