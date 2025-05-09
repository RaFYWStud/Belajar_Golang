package migrations

import "database/sql"

type createAccountTable struct{}

func (m *createAccountTable) SkipProd() bool {
	return false
}

func getCreateAccountTable() migration {
	return &createAccountTable{}
}

func (m *createAccountTable) Name() string {
	return "create-account"
}

func (m *createAccountTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`CREATE TABLE account (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)

	if err != nil {
		return err
	}

	return err
}

func (m *createAccountTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec("DROP TABLE account")

	return err
}

