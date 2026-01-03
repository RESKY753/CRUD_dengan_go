// package controller

// import (
// 	"net/http"
// 	"path/filepath"
// 	"text/template"
// 	"database/sql"
// )

// type Employee struct{
// 	Id string
// 	Name string
// 	NPWP string
// 	address string
// }
// func NewIndexEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		rows, err := db.Query("SELECT id, name, npwp, address FROM employee")

// 		if err != nil {
// 			w.Write([]byte(err.Error()))
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		var employees[]Employee
// 		for rows.Next() {
// 			var employee Employee

// 			err = rows.Scan(
// 				&employee.Id,
// 				&employee.Name,
// 				&employee.NPWP,
// 				&employee.address,
// 			)

// 			if err != nil {
// 			w.Write([]byte(err.Error()))
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 			}


// 			employees = append(employees, employee)
// 		}

// 		fp := filepath.Join("views", "index.html")

// 		tmpl, err := template.ParseFiles(fp)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}

// 		data := make(map[string]any)
// 		data["employees"] = employees

// 		err = tmpl.Execute(w, data)

// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 	}
// } 


package controller

import (
	"bytes"
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"
)

type Employee struct {
	Id      string
	Name    string
	NPWP    string
	Address string
}

func NewIndexEmployee(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, npwp, address FROM employee")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var employee Employee

			err = rows.Scan(
				&employee.Id,
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			employees = append(employees, employee)
		}

		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data := make(map[string]any)
		data["employees"] = employees

		// 🔑 WAJIB BUFFER di sini
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		buf.WriteTo(w)
	}
}



//