package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

// https://docs.mongodb.com/manual/tutorial/query-documents/
// https://sqlite.org/datatype3.html
// ../../../20200527/example_tinyorm/02sqlx/main.go

type Inventory struct {
	ID      int64   `db:"id,primarykey,autoincrement"`
	Item    string  `db:"item,size:255"`
	Qty     int64   `db:"qty"`
	SizeW   float64 `db:"sizeW"`
	SizeH   float64 `db:"sizeH"`
	SizeUOM string  `db:"sizeUOM,size:15"`
	Status  string  `db:"status,size:15"` // enum?
}

func (i Inventory) Format(s fmt.State, verb rune) {
	// 手抜き
	json.NewEncoder(s).Encode(i)
	// fmt.Fprintf(s, `@@%d`, i.ID)
}

var schema = `
CREATE TABLE inventory (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  item TEXT,
  qty INTEGER,
  sizeW REAL,
  sizeH REAL,
  sizeUOM TEXT,
  status TEXT
);
`

func initDb(ctx context.Context) (*sqlx.DB, error) {
	var db *sql.DB
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open failed")
	}

	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
	)
	// populate log pre-fields here before set to
	db = sqldblogger.OpenDriver(
		":memory:",
		db.Driver(),
		zerologadapter.New(logger),
		// optional config...
	)

	sqlxdb := sqlx.NewDb(db, "sqlite3")
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "ping failed")

	}

	// create table
	sqlxdb.MustExecContext(ctx, schema)
	return sqlxdb, nil
}

func run() error {
	ctx := context.Background()

	db, err := initDb(ctx)
	if err != nil {
		return err
	}

	inventories := []Inventory{
		{Item: "journal", Qty: 25, SizeH: 14, SizeW: 21, SizeUOM: "cm", Status: "A"},
		{Item: "notebook", Qty: 50, SizeH: 8.5, SizeW: 11, SizeUOM: "in", Status: "A"},
		{Item: "paper", Qty: 100, SizeH: 8.5, SizeW: 11, SizeUOM: "in", Status: "D"},
		{Item: "planner", Qty: 75, SizeH: 22.85, SizeW: 30, SizeUOM: "cm", Status: "D"},
		{Item: "postcard", Qty: 45, SizeH: 10, SizeW: 15.25, SizeUOM: "cm", Status: "A"},
		{Item: "unknown", Qty: 99, SizeH: 9, SizeW: 9, SizeUOM: "cm", Status: "X"},
	}

	// insert
	{
		r, err := db.NamedExecContext(
			ctx,
			`INSERT INTO inventory(item, qty, sizeH, sizeW, sizeUOM, status) VALUES (:item, :qty, :sizeH, :sizeW, :sizeUOM, :status)`,
			inventories,
		)

		// insert rows - auto increment PKs will be set properly after the insert
		if err != nil {
			return errors.Wrap(err, "Insert failed")
		}
		id, err := r.LastInsertId()
		if err != nil {
			return errors.Wrap(err, "Insert failed")
		}
		fmt.Println("@", id)
	}

	// select
	{
		// SELECT * FROM inventory
		log.Println(`db.inventory.find( {} )`)
		var inventories []Inventory
		err := db.SelectContext(ctx, &inventories, "select * from inventory;")
		if err != nil {
			return errors.Wrap(err, "Select failed")
		}
		fmt.Println(inventories)
		// spew.Dump(inventories)
	}

	{
		// SELECT * FROM inventory WHERE status in ("A", "D")
		log.Println(`db.inventory.find( { status: { $in: [ "A", "D" ] } } )`)
		var inventories []Inventory

		// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
		query, args, err := sqlx.In("select * from inventory where status in (?);", []string{"A", "D"})
		if err != nil {
			return errors.Wrap(err, "sqlx.In failed")
		}

		err = db.SelectContext(ctx, &inventories, query, args...)
		if err != nil {
			return errors.Wrap(err, "Select failed")
		}
		fmt.Println(inventories)
	}

	{
		// SELECT * FROM inventory WHERE status = "A" AND qty < 30
		log.Println(`db.inventory.find( { status: "A", qty: { $lt: 30 } } )`)
		var inventories []Inventory
		err := db.SelectContext(ctx, &inventories,
			"select * from inventory where status = ? AND qty < ?;",
			"A", 30,
		)
		if err != nil {
			return errors.Wrap(err, "Select failed")
		}
		fmt.Println(inventories)
	}

	{
		// SELECT * FROM inventory WHERE status = "A" OR qty < 30
		log.Println(`db.inventory.find( { $or: [ { status: "A" }, { qty: { $lt: 30 } } ] } )`)
		var inventories []Inventory
		err := db.SelectContext(ctx, &inventories,
			"select * from inventory where status = ? OR qty < ?;",
			"A", 30,
		)
		if err != nil {
			return errors.Wrap(err, "Select failed")
		}
		fmt.Println(inventories)
	}

	{
		// SELECT * FROM inventory WHERE status = "A" AND ( qty < 30 OR item LIKE "p%")
		log.Println(`db.inventory.find( {
     status: "A",
     $or: [ { qty: { $lt: 30 } }, { item: /^p/ } ]
} )
	}`)
		var inventories []Inventory
		err := db.SelectContext(ctx, &inventories,
			"select * from inventory where status = ? AND ( qty < ? OR item LIKE ? );",
			"A", 30, "p%",
		)
		if err != nil {
			return errors.Wrap(err, "Select failed")
		}
		fmt.Println(inventories)
	}

	// TODO: join?, limit
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("!!+%v", err)
	}
}
