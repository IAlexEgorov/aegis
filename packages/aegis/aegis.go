package aegis

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	log "k8s/packages/logging"
	"os/exec"
	"strconv"
	"strings"
)

const loginPostgres string = `{"User": "postgres","Password": "iIkWD4l8bd", Dbname: "aegis", Sslmode: "disable"}`

type Aegis struct {
	Id        string
	Name      string
	Namespace string
	// Components []string {"mlfow", "airflow"} // без селдона - только батч
}

func (a Aegis) CreateProject(dbCon *sql.DB) {
	var db *sql.DB = dbCon
	a.InsertRow(db)
}

func (a Aegis) ConnectToDB(user, password, dbname string, port uint16) (db *sql.DB) {
	//"root:root@tcp(127.0.0.1:8889)/mysql"
	db, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:"+strconv.Itoa(int(port))+")/"+dbname)
	if err != nil {
		panic(err)
	}
	log.InfoLogger.Println("DB Connected")
	return
}

func (a Aegis) InsertRow(db *sql.DB) {
	parsQ := "INSERT INTO `Aegis`(`Id`, `Name`, `Namespace`) VALUES ('" + a.Id + "','" + a.Name + "','" + a.Namespace + "')"
	insert, err := db.Query(parsQ)
	if err != nil {
		log.ErrorLogger.Println(err)
		panic(err)
	}
	log.InfoLogger.Println("Query was created")
	defer insert.Close()

	//a.HelmCreate(name, ns)
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
