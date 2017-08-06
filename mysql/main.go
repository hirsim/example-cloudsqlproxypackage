package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cloudsqlproxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	"github.com/go-sql-driver/mysql"
	goauth "golang.org/x/oauth2/google"
)

func main() {
	var (
		dbAddr         = flag.String("db_addr", "", "Fully-qualified Cloud SQL Instance (in the form of 'project:region:instance-name')")
		dbName         = flag.String("db_name", "", "Name of database")
		dbUser         = flag.String("db_user", "", "Name of user")
		dbPassword     = flag.String("db_pass", "", "Password for user")
		credentialFile = flag.String("credential_file", "", "Path to Service Account credentials file")
	)
	flag.Parse()

	if *dbAddr == "" {
		log.Fatal("Must set -db_addr")
	}
	if *dbName == "" {
		log.Fatal("Must set -db_name")
	}
	if *dbUser == "" {
		log.Fatal("Must set -db_user")
	}

	ctx := context.Background()

	if *credentialFile != "" {
		client, err := clientFromCredentials(ctx, *credentialFile)
		if err != nil {
			log.Fatal(err)
		}
		proxy.Init(client, nil, nil)
	}

	if err := showRecords(ctx, *dbAddr, *dbName, *dbUser, *dbPassword); err != nil {
		log.Fatal(err)
	}
}

func showRecords(ctx context.Context, dbAddress, dbName, dbUser, dbPassword string) error {
	db, err := cloudsqlproxy.DialCfg(&mysql.Config{
		Addr:      dbAddress,
		DBName:    dbName,
		User:      dbUser,
		Passwd:    dbPassword,
		Net:       "cloudsql", // Cloud SQL Proxy で接続する場合は cloudsql 固定です
		ParseTime: true,       // DATE/DATETIME 型を time.Time へパースする
		TLSConfig: "",         // TLSConfig は空文字を設定しなければなりません
	})
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "SELECT * FROM guestbook")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var guestName, content string
		var date time.Time
		var entryID int64
		if err := rows.Scan(&guestName, &content, &date, &entryID); err != nil {
			return err
		}
		fmt.Printf("%s\t%s\t%s\t%d\n", guestName, content, date.Format(time.RFC3339), entryID)
	}

	return nil
}

// 参考: https://github.com/GoogleCloudPlatform/cloudsql-proxy/blob/master/tests/dialers_test.go#L89
func clientFromCredentials(ctx context.Context, file string) (*http.Client, error) {
	const SQLScope = "https://www.googleapis.com/auth/sqlservice.admin"
	var client *http.Client

	all, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cfg, err := goauth.JWTConfigFromJSON(all, SQLScope)
	if err != nil {
		return nil, err
	}

	client = cfg.Client(ctx)

	return client, nil
}
