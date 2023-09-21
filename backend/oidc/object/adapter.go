package object

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb" // db = mssql
	_ "github.com/go-sql-driver/mysql"   // db = mysql
	"github.com/xorm-io/core"
	"github.com/xorm-io/xorm"
	"runtime"
	"strconv"
)

// Adapter represents the MySQL adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	dbName         string
	Engine         *xorm.Engine
}

var adapter *Adapter

func InitAdapter() {

	host := Conf.MySQLConfig.Host
	port := strconv.Itoa(Conf.MySQLConfig.Port)
	user := Conf.MySQLConfig.User
	pwd := Conf.MySQLConfig.Password
	dbName := Conf.MySQLConfig.DB
	path := user + ":" + pwd + "@tcp" + "(" + host + ":" + port + ")" + "/"
	adapter = NewAdapter("mysql", path, dbName)
	fmt.Println(path)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "")
	adapter.Engine.SetTableMapper(tbMapper)
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(driverName string, dataSourceName string, dbName string) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName
	a.dbName = dbName

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func finalizer(a *Adapter) {
	err := a.Engine.Close()
	if err != nil {
		panic(err)
	}
}

func (a *Adapter) open() {
	dataSourceName := a.dataSourceName + a.dbName
	engine, err := xorm.NewEngine(a.driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	a.Engine = engine
}

func (a *Adapter) close() {
	_ = a.Engine.Close()
	a.Engine = nil
}
