package database

import "database/sql"

type Ship struct {
	GroupID     int            `db:"groupID"`
	GroupName   string         `db:"groupName"`
	RaceID      sql.NullInt64  `db:"raceID"`
	RaceName    sql.NullString `db:"raceName"`
	TypeID      int            `db:"typeID"`
	TypeName    string         `db:"typeName"`
	Description string         `db:"description"`
}

type Traits struct {
	AttributeID int            `db:"attributeID"`
	ValueFloat  float64        `db:"valueFloat"`
	UnitName    sql.NullString `db:"unitName"`
	DisplayName string         `db:"displayName"`
	HighIsGood  bool           `db:"highIsGood"`
}

type Module struct {
	GroupID     int    `db:"groupID"`
	GroupName   string `db:"groupName"`
	TypeID      int    `db:"typeID"`
	TypeName    string `db:"typeName"`
	EffectID    int    `db:"effectID"`
	EffectName  string `db:"effectName"`
	Description string `db:"description"`
}

type Charge struct {
	GroupID     int    `db:"groupID"`
	GroupName   string `db:"groupName"`
	TypeID      int    `db:"typeID"`
	TypeName    string `db:"typeName"`
	Description string `db:"description"`
}
