package collector


import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/kr/pretty"
	"encoding/json"
	"strings"
	"errors"
)
type Table struct{
	Name string 
	Fields []string
	Values string

}
var Tables = map[string]*Table{
	"ticker":&Table{
		"24hrTicker",

		[]string{"Timestamp TIMESTAMP",
		 "Symbol TEXT",
		 "Price FLOAT", 
		 "Close FLOAT", 
		 "Open FLOAT",
		 "Price_Change FLOAT",
		 "Price_Change_Percent FLOAT",
		 "Last_Quantity FLOAT",
		 "Last_Price FLOAT",
		 "Weighted_Average_Price FLOAT",
		 "First_Trade_Price FLOAT",
		},
		"",
	},

	"miniTicker":&Table{
		"miniTicker",

		[]string{"Timestamp TIMESTAMP",
		 "Symbol TEXT",
		 "Price FLOAT", 
		 "Close FLOAT", 
		 "Open FLOAT",
		 "Price_Change FLOAT",
		 "Price_Change_Percent FLOAT",
		 "Last_Quantity FLOAT",
		 "Last_Price FLOAT",
		 "Weighted_Average_Price FLOAT",
		 "First_Trade_Price FLOAT",
		},
		"",
	},
}
func(t *Table)GetFields()string{
	fields := "("+t.Fields[0]
	for i := 1; i < len(t.Fields); i++ {
		fields += ","+t.Fields[i];
	}
	return fields+")"; 
}
func(t *Table)GetValues()string{
	if(t.Values == ""){
		t.Values = "VALUES(?";
		for _,_ = range(t.Fields[1:]){
			t.Values += ",?";
		}
		t.Values += ")"
	}
	return t.Values;
}

type database struct{
	DB *sql.DB
	Transactions map[string]*sql.Stmt
}
func newDatabaseManager()*database{
	db, err := sql.Open("sqlite3","./foo.sqlite3");
	if err != nil{
		log.Fatal(err);
	}
	return &database{db,map[string]*sql.Stmt{}}

}



func(db *database)newTable(table_name string)(err error){
	table,ok := Tables[table_name];
	if  !ok{
		return errors.New("Table "+table.Name+" Not found");
	}
	res,err := db.DB.Exec("CREATE TABLE "+table.Name+" "+table.GetFields()+"");
	if(err != nil){
		pretty.Println(res,err)
	}
	return db.PrepareTransaction(table)
}
func(db *database)PrepareTransaction(table *Table)error{
	tx, err := db.DB.Prepare("INSERT INTO table "+table.Name+" "+table.GetValues());
	if(err == nil){
		db.Transactions[table.Name] = tx
	}
	return err
}


func(db *database)Manage(data []byte){
	var stream Stream
	if err := json.Unmarshal(data,&stream);err != nil{
		panic(err)
	}
	if stream.Name[0] == '!'{stream.Name = stream.Name[1:]}
	names := strings.Split(stream.Name,"@")
	if (names[1] == "arr"){
		db.InsertArrayData(names[0],stream)
		return
	}
	//db.InsertSingleData(names[1]+"_"+names[0])

	pretty.Println(stream)
}

func(db *database)InsertArrayData(table_name string,stream Stream){
	pretty.Println(table_name);
	tx, ok := db.Transactions[table_name];
	if !ok{
		if err := db.newTable(table_name); err !=  nil{panic(err)}
	}
	n , _ := stream.Data.(*map[string]interface{})
	pretty.Println(n)
	_, err := tx.Exec("asd");
	panic(err)

}

type Stream struct{
	Name string `json:"stream"`
	Data interface{} `json:"data"`
}
