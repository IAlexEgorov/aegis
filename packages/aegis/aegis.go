package aegis

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	log "k8s/packages/logging"
	"os/exec"
	"strings"
)

const loginPostgres string = `{"User": "postgres","Password": "iIkWD4l8bd", Dbname: "aegis", Sslmode: "disable"}`

type Aegis struct {
	Id        string
	Name      string
	Namespace string
	// Components []string {"mlfow", "airflow"} // без селдона - только батч
}

type PostgresConnection struct {
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

func (pc *PostgresConnection) Parse(data string) string {
	err := json.Unmarshal([]byte(data), &pc)
	if err != nil {
		log.ErrorLogger.Println(err)
	}
	return "user=" + pc.User + " password=" + pc.Password +
		" dbname=" + pc.Dbname + " sslmode=" + pc.Sslmode
}

func (a Aegis) CreateProject() {
	var db *sql.DB = a.ConnectToDB()
	a.InsertRow(db, a.Id, a.Name, a.Namespace)
}

func (a Aegis) ConnectToDB() (db *sql.DB) {
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.ErrorLogger.Println(err)
		}
	}(db)

	var pc PostgresConnection = PostgresConnection{}
	db, err := sql.Open("postgres", pc.Parse(loginPostgres))
	if err != nil {
		log.ErrorLogger.Println(err)
	}
	return
}

func (a Aegis) InsertRow(db *sql.DB, id, name, ns string) {
	_, err := db.Exec("insert into Aegis (Name, Namespace) values ($1, $2)", name, ns)
	if err != nil {
		log.ErrorLogger.Println(err)
	}
	log.InfoLogger.Println("INSERT INTO Aegis", name, ns)
	a.HelmCreate(name, ns)
}

func (a Aegis) HelmCreate(name, ns string) {
	s := "helm install " + name + " -f test/values.yaml test/. -n " + ns + " --create-namespace"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.ErrorLogger.Println(err)
	}
	log.InfoLogger.Println("Aegis:", name, "was created in ns:", ns)
}
