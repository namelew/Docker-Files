package database

import (
	"commerce-api/models"
	"commerce-api/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var database *sql.DB

func OpenConnection() {
	sourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", utils.GetDotEnvVar("DBHOST"), utils.GetDotEnvVar("DBPORT"), utils.GetDotEnvVar("DBUSER"), utils.GetDotEnvVar("DBPW"), utils.GetDotEnvVar("DBNAME"))
	db, err := sql.Open("postgres", sourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products(id serial primary key, name varchar(50), description varchar(255), price numeric)")
	if err != nil {
		log.Fatal(err.Error())
	}

	database = db
}

func Insert(p models.Product) error {
	insert, err := database.Prepare("insert into products(name,description,price) values($1,$2,$3)")
	if err != nil {
		return err
	}

	_, err = insert.Exec(p.Name, p.Description, p.Price)

	if err != nil {
		return err
	}

	return nil
}

func Remove(id int) error {
	remove, err := database.Prepare("delete from products where id=$1")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = remove.Exec(id)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func Update(p models.Product) error {
	update, err := database.Prepare("update products set name=$1,description=$2,price=$3 where id=$4")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = update.Exec(p.Name, p.Description, p.Price, p.Id)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func List() ([]models.Product, error) {
	var products []models.Product

	ret, err := database.Query("select * from products")

	if err != nil {
		return nil, err
	}

	for ret.Next() {
		var p models.Product

		err = ret.Scan(&p.Id, &p.Name, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, err
}
