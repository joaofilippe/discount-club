package objecttypesmigrations

import "github.com/joaofilippe/discount-club/infra/database"

func CreateObjectTypeENUM(conn *database.Connection) error {
	tx := conn.Get().MustBegin()

	_, err := tx.Exec(createUserTypeENUM)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func DropObjectTypeENUM(conn *database.Connection) error {
	tx := conn.Get().MustBegin()

	_, err := tx.Exec(dropUserTypeENUM)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
