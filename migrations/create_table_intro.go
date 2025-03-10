package migrations

import "database/sql"

type createIntroTable struct{}

func (m *createIntroTable) SkipProd() bool {
	return false
}

func getCreateIntroTable() migration {
	return &createIntroTable{}
}

func (m *createIntroTable) Name() string {
	return "create-intro"
}

func (m *createIntroTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE intro (
			id SERIAL PRIMARY KEY,
			nama VARCHAR(255) NOT NULL,
			nama_panggilan VARCHAR(255),
			fun_fact VARCHAR(255),
			keinginan_BE VARCHAR(255),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)

	if err != nil {
		return err
	}

	return err
}

func (m *createIntroTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec("DROP TABLE intro")

	return err
}
