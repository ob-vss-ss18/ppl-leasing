package leasing

import (
	"time"
)

type Leasing_contract struct{
	Leasing_id int  `json:"leasing_id"`
	Kunden_id int `json:"kunden_id"`
	Products []int `json:"products"`
	Datum time.Time `json:"datum"`
	Rabatt int `json:"rabatt"`
	Service_flat bool `json:"service_flat"`
	Testwert bool `json:"testwert"`
	Versicherung bool `json:"versicherung"`

}

func CreateLeasingContract(contract *Leasing_contract) (*Leasing_contract, error) {
	var id int
	sqlStatement := `
			INSERT INTO Leasing (Datum, Kunden_ID, Testwert, Versicherung, Serviceflat, Rabatt) 
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING Leasing_ID`
	err := db.QueryRow(sqlStatement, contract.Datum, contract.Kunden_id, contract.Testwert,
		contract.Versicherung, contract.Service_flat, contract.Rabatt).Scan(&id)
	if err != nil {
		return nil, err
	}
	contract.Leasing_id = id
	//Products speichern
	for _, product := range contract.Products {
		_, error := db.Exec(`
				INSERT INTO Leased_Products (Leasing_ID, Product_ID)
				VALUES ($1, $2)`, id, product)
		if error != nil {
			return nil, error
		}
	}
	//Daten an Billing und Mailing weiterleiten
	return contract, nil
}

func GetLeasingContractByID(id int)(*Leasing_contract, error) {
	var kunden_id int
	var datum time.Time
	var rabatt int
	var service_flat bool
	var testwert bool
	var versicherung bool
	var leasing_id int
	err := db.QueryRow("SELECT * FROM Leasing WHERE Leasing_ID=$1", id).
		Scan(&leasing_id, &datum, &kunden_id, &testwert, &versicherung, &service_flat, &rabatt)
	if err != nil {
		return nil, err
	}
	var products []int
	rows, error := db.Query(`SELECT Product_Id FROM Leased_Products WHERE Leasing_ID=$1`, id)
	if error != nil {
		return nil, error
	}
	defer rows.Close()
	for rows.Next() {
		var product_id int
		err = rows.Scan(&product_id)
		if err != nil {
			return nil, err
		}
		products = append(products, product_id)
	}
	return &Leasing_contract{
		Leasing_id: id,
		Kunden_id: kunden_id,
		Products: products,
		Datum: datum,
		Rabatt: rabatt,
		Service_flat: service_flat,
		Testwert: testwert,
		Versicherung: versicherung,
	}, nil
}
