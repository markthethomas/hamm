package hamm

import (
	"github.com/youtube/vitess/go/vt/sqlparser"
)

// const (
// 	Statement sqlparser.Statement = 0
// 	Statement sqlparser.Insert    = 0
// )

type Query struct {
	Text      string
	Statement sqlparser.Statement
}

func Parse(q string) {
	stmt, err := sqlparser.Parse(sql)

	if err != nil {
		return
	}
}

// sql := "SELECT * FROM users WHERE a = 'abc'"
// 	stmt, err := sqlparser.Parse(sql)
// 	fmt.Printf("%+v", stmt)
// 	if err != nil {
// 		// fmt.Println(err)
// 		// Do something with the err
// 	}

// 	// Otherwise do something with stmt
// 	switch stmt := stmt.(type) {
// 	case *sqlparser.Select:
// 		spew.Dump(stmt)
// 	case *sqlparser.Insert:
// 		spew.Dump(stmt)
// 	}
