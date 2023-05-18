package repository

import (
	"log"
	"serasa-hotel/db"
	"serasa-hotel/models"
)

func (repo Connection) InsertCustomer(customer models.Customer) (id int64, err error) {
	sql := `INSERT INTO customer (name, document, phone_number, email) VALUES ($1, $2, $3, $4) RETURNING id`

	err = repo.db.QueryRow(sql, customer.Name, customer.Document, customer.PhoneNumber, customer.Email).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	defer repo.db.Close()
	return
}

func (repo Connection) GetCustomer(id int64) (customer models.Customer, err error) {
	row := repo.db.QueryRow(`SELECT * FROM customer where id=$1`, id)
	defer repo.db.Close()

	err = row.Scan(&customer.ID, &customer.Name, &customer.Document, &customer.PhoneNumber, &customer.Email)
	if err != nil {
		log.Println(err)
	}
	return
}

func (repo Connection) GetAllCustomer() (customers []models.Customer, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM customer`)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var customer models.Customer
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Document, &customer.PhoneNumber, &customer.Email)
		if err != nil {
			log.Println(err)
		}
		customers = append(customers, customer)
	}

	return
}

// func (repo Customer) GetByName() (name string) (customers []models.Customer, err error) {
// 	conn, err := db.OpenConnection()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer conn.Close()

// 	rows, err := conn.Query(`SELECT * FROM customer where name like '%$1'`, name)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	for rows.Next() {
// 		var customer models.Customer
// 		err = rows.Scan(&customer.ID, &customer.Name, &customer.Document, &customer.PhoneNumber, &customer.Email)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		customers = append(customers, customer)
// 	}

// 	return
// }

func (repo Connection) UpdateCustomer(id int64, customer models.Customer) (int64, error) {
	res, err := repo.db.Exec(`UPDATE customer SET name=$1, document=$2, phone_number=$3, email=$4 WHERE id=$5`, customer.Name, customer.Document, customer.PhoneNumber, customer.Email, id)
	if err != nil {
		log.Println(err)
	}
	defer repo.db.Close()
	return res.RowsAffected()
}

func (repo Connection) DeleteCustomer(id int64) (int64, error) {
	res, err := repo.db.Exec(`DELETE FROM customer WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
	}
	defer repo.db.Close()

	return res.RowsAffected()
}
