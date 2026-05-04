package repository

import (
	"database/sql"

	"caselog/internal/model"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) FindAll() ([]model.Project, error) {
	rows, err := r.DB.Query(`
		SELECT id, title, company, main_skill, unit_price, remote_type, team_size, status, notes, created_at, updated_at
		FROM projects
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project

	for rows.Next() {
		var p model.Project
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Company,
			&p.MainSkill,
			&p.UnitPrice,
			&p.RemoteType,
			&p.TeamSize,
			&p.Status,
			&p.Notes,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}

		projects = append(projects, p)
	}

	return projects, rows.Err()
}

func (r *ProjectRepository) Create(p model.Project) error {
	_, err := r.DB.Exec(`
		INSERT INTO projects (
			title,
			company,
			main_skill,
			unit_price,
			remote_type,
			team_size,
			status,
			notes
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		p.Title,
		p.Company,
		p.MainSkill,
		p.UnitPrice,
		p.RemoteType,
		p.TeamSize,
		p.Status,
		p.Notes,
	)

	return err
}

func (r *ProjectRepository) FindByID(id int) (model.Project, error) {
	row := r.DB.QueryRow(`
		SELECT id, title, company, main_skill, unit_price, remote_type, team_size, status, notes, created_at, updated_at
		FROM projects
		WHERE id = ?
	`, id)

	var p model.Project
	err := row.Scan(
		&p.ID,
		&p.Title,
		&p.Company,
		&p.MainSkill,
		&p.UnitPrice,
		&p.RemoteType,
		&p.TeamSize,
		&p.Status,
		&p.Notes,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProjectRepository) Update(project model.Project) error {
	query := `
		UPDATE projects
		SET
			title = ?,
			company = ?,
			main_skill = ?,
			unit_price = ?,
			remote_type = ?,
			team_size = ?,
			status = ?,
			notes = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	_, err := r.DB.Exec(
		query,
		project.Title,
		project.Company,
		project.MainSkill,
		project.UnitPrice,
		project.RemoteType,
		project.TeamSize,
		project.Status,
		project.Notes,
		project.ID,
	)

	return err
}