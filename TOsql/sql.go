package TOsql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	DB *sql.DB
)
var ip string
var port string
var password string
var dbname string

func init() {
	Initconfig()
	DB = Connect()
}
func Initconfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	ip = viper.GetString("mysql.ip")
	port = viper.GetString("mysql.port")
	password = viper.GetString("mysql.password")
	dbname = viper.GetString("mysql.dbname")
	fmt.Println("读取配置：", "ip:", ip, "port:", port, "password:", password, "dbname:", dbname)
}
func Connect() *sql.DB {
	s := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", password, ip, port, dbname)
	db, err := sql.Open("mysql", s)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("连接数据库成功")
	}
	return db
}
