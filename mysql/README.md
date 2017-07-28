# go-activerecord for mysql

#Usage
<pre>
import  github.com/snail007/go-activerecord/mysql
0.configure 
var dbCfg = NewDBConfig()
dbCfg.Password = "admin"
db, err := NewDB(dbCfg)
if err != nil {
		fmt.Printf("ERR:%s", err)
        return
}

Default config is below:

Charset:                  "utf8",
Collate:                  "utf8_general_ci",
Database:                 "test",
Host:                     "127.0.0.1",
Port:                     3306,
Username:                 "root",
Password:                 "",
TablePrefix:              "",
TablePrefixSqlIdentifier: "",
Timeout:                  3000,
SetMaxOpenConns:          500,
SetMaxIdleConns:          50,



1.Select
	rs, err := db.Query(db.AR().
                Select("*").
                From("log").
                Where(map[string]interface{}{
                    "id": 11,
                })
    )
    if err != nil {
		fmt.Printf("ERR:%s", err)
	} else {
		fmt.Println(rs.Rows())
	}
    //db.AR() return a new *mysql.ActiveRecord,you can use it to build you sql.
    //all of db.AR() usage to see <a href="https://github.com/snail007/go-activerecord/blob/master/mysql/mysql_test.go">mysql_test.go</a>

    rs is a ResultSet,all of ResultSet method and properties is :
    ResultSet.Len()
        how many rows of select
    ResultSet.MapRows(keyColumn string) (rowsMap map[string]map[string]string)
        get a map which key is each value of row[keyColumn]
    ResultSet.Rows() (rows []map[string]string)
        get rows of select
    ResultSet.Row() (row map[string]string)
        get first of rows
    ResultSet.Values(column string) (values []string)
        get an array contains each row[column] 
    ResultSet.MapValues(keyColumn, valueColumn string) (values map[string]string)
        get a map key is each row[column],value is row[valueColumn]
    ResultSet.Value(column string) (value string)
        get first row[column] of rows
    ResultSet.LastInsertId
        if sql type is insert , this is the last insert id
    ResultSet.RowsAffected
        if sql type is write , this is the count of rows affected

	
2.Insert & Insert Batch
    Insert:
        rs, err := db.Exec(db.AR().Insert("test", map[string]interface{}{
			"id":   "id11122",
			"name": "333",
		}))
    Insert Batch:
        rs, err := db.Exec(db.AR().InsertBatch("test", []map[string]interface{}{
            map[string]interface{}{
                "id":   "id11122",
                "name": "333",
            },
            map[string]interface{}{
                "id":   "id11122",
                "name": "4444",
            },
        }))
    lastInsertId:=rs.LastInsertId
    rowsAffected:=rs.RowsAffected
    fmt.printf("last insert id : %d,rows affected : %d",lastInsertId,rowsAffected)
	
3.Update & Update Batch
    Update:
        rs, err := db.Exec(db.AR().Update("test", map[string]interface{}{
			"id":   "id11122",
			"name": "333",
		}),map[string]interface{}{
			"pid":   223,
		}))
    Update Batch:
    rs, err := db.Exec(db.AR().UpdateBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"id":   "id11122",
			"name": "333",
		},
		map[string]interface{}{
			"id":   "id11122",
			"name": "4444",
		},
	}, "id"))
    rowsAffected:=rs.RowsAffected
    fmt.printf("rows affected : %d",rowsAffected)
4.Delete
    rs, err := db.Exec(db.AR().Delete("test", map[string]interface{}{
        "pid":   223,
    }))
    rowsAffected:=rs.RowsAffected
    fmt.printf("rows affected : %d",rowsAffected)
5.Raw SQL Query
    rs, err := db.Exec(db.AR().Raw("insert into test(id,name) values (?,?)", 555,"6666"))
    if err != nil {
        fmt.Printf("ERR:%s", err)
    } else {
        fmt.Println(rs.RowsAffected, rs.LastInsertId)
    }
</pre>