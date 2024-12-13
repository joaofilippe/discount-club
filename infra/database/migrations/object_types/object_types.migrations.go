package objecttypesmigrations

import "github.com/joaofilippe/discount-club/infra/database"

func createObjectTypeENUM(conn *database.Connection) error {
	tx := conn.Get().MustBegin()

	_, err := tx.Exec(createUserTypeENUM)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
