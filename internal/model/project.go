package model

import "time"

type Project struct {
	ID        int
	Title     string
	Company   string
	MainSkill string
	UnitPrice int
	RemoteType string
	TeamSize  string
	Status    string
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func StatusLabel(status string) string {
	switch status {
	case "considering":
		return "検討中"
	case "applied":
		return "応募済み"
	case "interview":
		return "面談予定"
	case "declined":
		return "見送り"
	case "accepted":
		return "参画決定"
	default:
		return status
	}
}