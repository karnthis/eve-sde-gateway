package database

import (
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/lib/pq"
)

var Instance *sqlx.DB

func InitDB() {
	db, err := sqlx.Connect("postgres", "user=owo_user dbname=sdeyamlschema password=unsafe host=sde-pg sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	Instance = db
}

func ReadTraits(id int) []Traits {
	qry := `Select
    "attributeID",
    "valueFloat",
    eU."unitName",
    dAT."displayName",
    "highIsGood"
    from evesde."dgmTypeAttributes" dTA
        left join evesde."dgmAttributeTypes" dAT using ("attributeID")
        left join evesde."eveUnits" eU using ("unitID")
    where dTA."typeID" = $1 and "published" = true`

	result := []Traits{}
	err := Instance.Select(&result, qry, id)
	if err != nil {
		log.Println(err)
	}
	return result
}

func ReadShips() []Ship {
	qry := `Select
    "groupID",
    "groupName",
    "raceID",
    "raceName",
    "typeID",
    "typeName",
    iT."description"
    from evesde."invCategories" iC
        left join evesde."invGroups" iG using ("categoryID")
        left join evesde."invTypes" iT using ("groupID")
        left join evesde."chrRaces" cR using("raceID")
    where iC."categoryID" = 6 and iT."published" = true`

	result := []Ship{}
	err := Instance.Select(&result, qry)
	if err != nil {
		log.Println(err)
	}
	return result
}

func ReadModules() []Module {
	qry := `Select
    "groupID",
    "groupName",
    "typeID",
    "typeName",
    "effectID",
    "effectName",
    iT."description"
    from evesde."invCategories" iC
        left join evesde."invGroups" iG using ("categoryID")
        left join evesde."invTypes" iT using ("groupID")
        inner join (
            select
                "typeID",
                "effectID",
                "effectName"
            from evesde."dgmTypeEffects" dTE
            left join evesde."dgmEffects" dE using ("effectID")
            where "effectID" in (11, 12, 13) and dE."published" = true
    ) jJ using ("typeID")
    where iC."categoryID" = 7 and iT."published" = true`

	result := []Module{}
	err := Instance.Select(&result, qry)
	if err != nil {
		log.Println(err)
	}
	return result
}

func ReadCharges() []Charge {
	qry := `Select
    "groupID",
    "groupName",
    "typeID",
    "typeName",
    iT."description"
    from evesde."invCategories" iC
        left join evesde."invGroups" iG using ("categoryID")
        left join evesde."invTypes" iT using ("groupID")
    where iC."categoryID" = 8 and iT."published" = true`

	result := []Charge{}
	err := Instance.Select(&result, qry)
	if err != nil {
		log.Println(err)
	}
	return result
}
