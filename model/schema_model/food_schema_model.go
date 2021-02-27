package schema_model

import "time"

type FoodSchemaModel struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	TypeFood    int       `db:"type_food"`
	Status      int       `db:"status"`
	Description string    `db:"description"`
	UpdatedDate time.Time `db:"updated_date"`
}
