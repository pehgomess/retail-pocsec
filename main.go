package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var db *sql.DB

type productInfo struct {
	Id       string  `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Category string  `json:"category,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(localhost:3306)/sqlinjection?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//func getProducts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	db = getMySQLDB()
//	defer db.Close()
//	ss := []studentInfo{}
//	s := studentInfo{}
//	rows, err := db.Query("select * from products")
//	if err != nil {
//		fmt.Fprintf(w, ""+err.Error())
//	} else {
//		for rows.Next() {
//			rows.Scan(&s.Id, &s.Name, &s.Category, &s.Price)
//			ss = append(ss, s)
//		}
//		json.NewEncoder(w).Encode(ss)
//	}
//}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = getMySQLDB()
	defer db.Close()
	ss := []productInfo{}
	s := productInfo{}
	categorys, ok := r.URL.Query()["category"]
	if !ok || len(categorys[0]) < 1 {
		rows, err := db.Query("select * from products")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			for rows.Next() {
				rows.Scan(&s.Id, &s.Name, &s.Category, &s.Price)
				ss = append(ss, s)
			}
			json.NewEncoder(w).Encode(ss)
		}
	} else {
		category := categorys[0]
		rows, err := db.Query("SELECT id, name, category, price FROM products WHERE category=?", category)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			for rows.Next() {
				rows.Scan(&s.Id, &s.Name, &s.Category, &s.Price)
				ss = append(ss, s)
			}
			json.NewEncoder(w).Encode(ss)
		}
	}

}

//func addProducts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	db = getMySQLDB()
//	defer db.Close()
//	s := productInfo{}
//	json.NewDecoder(r.Body).Decode(&s)
//	id, _ := strconv.Atoi(s.Id)
//	name := strings.ToLower(s.Name)
//	category := strings.ToLower(s.Category)
//	result, err := db.Exec("insert into products(id, name, category, price) values(?,?,?,?)", id, name, category, s.Price)
//	if err != nil {
//		fmt.Fprintf(w, ""+err.Error())
//	} else {
//		_, err := result.LastInsertId()
//		if err != nil {
//			json.NewEncoder(w).Encode("{error:Record not inserted}")
//		} else {
//			json.NewEncoder(w).Encode(s)
//		}
//	}
//}

func addProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = getMySQLDB()
	defer db.Close()
	s := productInfo{}
	json.NewDecoder(r.Body).Decode(&s)
	id, _ := strconv.Atoi(s.Id)
	name := strings.ToLower(s.Name)
	category := strings.ToLower(s.Category)
	result, err := db.Exec("insert into products(id, name, category, price) values(?,?,?,?)", id, name, category, s.Price)
	if err != nil {
		fmt.Fprintf(w, ""+err.Error())
	} else {
		_, err := result.LastInsertId()
		if err != nil {
			json.NewEncoder(w).Encode("{error:Record not inserted}")
		} else {
			json.NewEncoder(w).Encode(s)
		}
	}
}

func updateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = getMySQLDB()
	defer db.Close()
	s := productInfo{}
	json.NewDecoder(r.Body).Decode(&s)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	name := strings.ToLower(s.Name)
	category := strings.ToLower(s.Category)
	result, err := db.Exec("UPDATE products SET id=?, name=?, category=?, price=? WHERE ``.products.id = ?", id, name, category, s.Price, id) //UPDATE products SET id=1, name="armario", category="moveis", price=3000.00 WHERE ``.products.id = 1;
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		_, err := result.RowsAffected()
		if err != nil {
			json.NewEncoder(w).Encode("{error:Record is not updated}")
		} else {
			json.NewEncoder(w).Encode(s)
		}
	}
}

func deleteProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = getMySQLDB()
	defer db.Close()
	//s := studentInfo{}
	//json.NewDecoder(r.Body).Decode(&s)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	result, err := db.Exec("delete from products where id=?", id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		_, err := result.RowsAffected()
		if err != nil {
			json.NewEncoder(w).Encode("{error:Record is not deleted}")
		} else {
			json.NewEncoder(w).Encode("{Record is deleted}")
		}
	}
}

func getOneProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = getMySQLDB()
	defer db.Close()
	//	queryParams := r.URL.Query().Get("name")
	ss := []productInfo{}
	s := productInfo{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	rows, err := db.Query("select id, name, category, price from products where id=?", id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&s.Id, &s.Name, &s.Category, &s.Price)
			ss = append(ss, s)
		}
		json.NewEncoder(w).Encode(ss)
	}
}

func sqlgetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db = getMySQLDB()
	defer db.Close()
	ss := []productInfo{}
	s := productInfo{}
	json.NewDecoder(r.Body).Decode(&s)
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM products WHERE id = %s", s.Id)) // sql injection test
	//rows, err := db.Query("SELECT * FROM products WHERE id = ?", s.Id)
	if err != nil {
		fmt.Fprintf(w, ""+err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&s.Id, &s.Name, &s.Category, &s.Price)
			ss = append(ss, s)
		}
		json.NewEncoder(w).Encode(ss)
	}
}

func getXSS(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, r.URL.Query().Get("KEY"))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/products", getProducts).Methods("GET")
	r.HandleFunc("/api/v1/products", addProducts).Methods("POST")
	r.HandleFunc("/api/v1/products/{id}", updateProducts).Methods("PUT")
	r.HandleFunc("/api/v1/products/{id}", deleteProducts).Methods("DELETE")
	r.HandleFunc("/api/v1/products/{id}", getOneProducts).Methods("GET")
	r.HandleFunc("/api/v1/getproducts", sqlgetProducts).Methods("POST") // sqlinjection test
	r.HandleFunc("/api/v1/xss", getXSS).Methods("GET")                  // xss test
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
