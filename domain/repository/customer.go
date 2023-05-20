package repository

import (
	"fmt"
	"log"
	"serasa-hotel/domain/utils"
	"serasa-hotel/models"
	"strings"
)

func (repo Connection) InsertCustomer(customer models.Customer) (id int64, err error) {
	sql := `INSERT INTO customer (name, document, phone_number, email) VALUES ($1, $2, $3, $4) RETURNING id`

	err = repo.db.QueryRow(sql, customer.Name, customer.Document, customer.PhoneNumber, customer.Email).Scan(&id)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	defer repo.db.Close()
	return
}

func (repo Connection) GetCustomer(id int64) (customer models.Customer, err error) {
	row := repo.db.QueryRow(`SELECT * FROM customer WHERE id = $1`, id)
	err = row.Scan(&customer.ID, &customer.Name, &customer.Document, &customer.PhoneNumber, &customer.Email)
	if err != nil {
		log.Println(err)
		return customer, err
	}
	return
}

func (repo Connection) GetAllCustomer(customerParams utils.CustomerParams) (customers []models.Customer, err error) {
	sql := `SELECT * FROM customer`
	whereToSelect := []string{}
	if customerParams.Name != "" {
		whereToSelect = append(whereToSelect, fmt.Sprintf("name like '%%%s%%'", customerParams.Name))
	}
	if customerParams.Phone != "" {
		whereToSelect = append(whereToSelect, fmt.Sprintf("phone_number like '%%%s%%'", customerParams.Phone))
	}
	if customerParams.Document != "" {
		whereToSelect = append(whereToSelect, fmt.Sprintf("document like '%%%s%%'", customerParams.Document))
	}
	if len(whereToSelect) != 0 {
		whereClause := strings.Join(whereToSelect, " and ")
		sql = fmt.Sprintf("%s where %s", sql, whereClause)
	}
	rows, err := repo.db.Query(sql)
	if err != nil {
		log.Println(err)
		return customers, err
	}
	var customer models.Customer
	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Document, &customer.PhoneNumber, &customer.Email)
		if err != nil {
			log.Printf("Erro ao pegar clientes: %v", err)
		}
		customers = append(customers, customer)
	}
	return
}

func (repo Connection) UpdateCustomer(id int64, customer models.Customer) (int64, error) {
	res, err := repo.db.Exec(`UPDATE customer SET name=$1, document=$2, phone_number=$3, email=$4 WHERE id=$5`, customer.Name, customer.Document, customer.PhoneNumber, customer.Email, id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer repo.db.Close()
	return res.RowsAffected()
}

func (repo Connection) DeleteCustomer(id int64) (int64, error) {
	res, err := repo.db.Exec(`DELETE FROM customer WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer repo.db.Close()
	return res.RowsAffected()
}
