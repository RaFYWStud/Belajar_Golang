package migrations

import "database/sql"

type createToDoTable struct{}

func (m *createToDoTable) SkipProd() bool {
    return false
}

func getCreateToDoTable() migration {
    return &createToDoTable{}
}

func (m *createToDoTable) Name() string {
    return "create-todo"
}

func (m *createToDoTable) Up(conn *sql.Tx) error {
    _, err := conn.Exec(`
        CREATE TABLE to_do_list (
            Id INT PRIMARY KEY,
            Nama VARCHAR(255),
            Hari VARCHAR(50),
            ToDo TEXT
        )
    `)
    return err
}

func (m *createToDoTable) Down(conn *sql.Tx) error {
    _, err := conn.Exec("DROP TABLE to_do_list")
    return err
}