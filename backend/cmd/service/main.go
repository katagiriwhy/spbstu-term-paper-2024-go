package main

import (
	"log"

	"github.com/backend/database"
	"github.com/backend/internal/config"
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()

	createProjectTable := "CREATE TABLE IF NOT EXISTS Project (id    serial PRIMARY KEY,    key TEXT,    title TEXT,    name text, url text)"
	_, err = db.Db.Query(createProjectTable)
	if err != nil {
		log.Fatal(err)
	}
	createAuthorTable := "CREATE TABLE IF NOT EXISTS Author(id   serial PRIMARY KEY,    name TEXT)"
	_, err = db.Db.Query(createAuthorTable)
	if err != nil {
		log.Fatal(err)
	}
	createIssuesTable := "CREATE TABLE IF NOT EXISTS Issues(id          serial PRIMARY KEY,    project_id   INT NOT NULL,    FOREIGN KEY (project_id) REFERENCES Project (id) ON DELETE CASCADE ON UPDATE CASCADE,    author_id    INT NOT NULL,    FOREIGN KEY (author_id) REFERENCES Author (id) ON DELETE CASCADE ON UPDATE CASCADE,    assignee_id  INT NOT NULL,    FOREIGN KEY (assignee_id) REFERENCES Author (id) ON DELETE CASCADE ON UPDATE CASCADE,    key         TEXT,    summary     TEXT,    description TEXT,    type        TEXT,    priority    TEXT,    status      TEXT,    created_time TIMESTAMP WITHOUT TIME ZONE,    updated_time TIMESTAMP WITHOUT TIME ZONE)"
	_, err = db.Db.Query(createIssuesTable)
	if err != nil {
		log.Fatal(err)
	}
	createStatusChangeTable := "CREATE TABLE IF NOT EXISTS statusChange (id         serial primary key,    issuesId    INT NOT NULL,    FOREIGN KEY (issuesId) REFERENCES Issues (id) ON DELETE CASCADE ON UPDATE CASCADE,    authorId   INT NOT NULL,    FOREIGN KEY (authorId) REFERENCES Author (id) ON DELETE CASCADE ON UPDATE CASCADE,    changeTime TIMESTAMP WITHOUT TIME ZONE,    fromStatus TEXT,   toStatus   TEXT)"
	_, err = db.Db.Query(createStatusChangeTable)
	if err != nil {
		log.Fatal(err)
	}
	createUser := "DROP ROLE IF EXISTS pguser; CREATE USER pguser WITH ENCRYPTED PASSWORD 'pgpwd';"
	_, err = db.Db.Query(createUser)
	if err != nil {
		log.Fatal(err)
	}
	givePrivileges := "GRANT ALL ON TABLE Project, Author, Issues, statusChange TO pguser; GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO pguser;"
	_, err = db.Db.Query(givePrivileges)
	if err != nil {
		log.Fatal(err)
	}

}

//IN SCHEMA public

/*DROP OWNED BY pguser;
DROP USER pguser;*/
