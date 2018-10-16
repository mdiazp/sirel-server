package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mdiazp/sirel-server/api/models/models2"

	// Postgresql Driver
	_ "github.com/lib/pq"
)

var (
	// ErrResultNotFound ...
	ErrResultNotFound = orm.ErrNoRows

	//queryUserProperties
	queryUserProperties = "k_user.id, " +
		"k_user.username, " +
		"k_user.name, " +
		"k_user.email, " +
		"k_user.send_notifications_to_email, " +
		"k_user.rol, " +
		"k_user.enable"

	// QueryLocalAdmins ...
	QueryLocalAdmins = "select " +
		queryUserProperties +
		" from local_admin join k_user on user_id=k_user.id where local_id=?"
	// QueryAreaAdmins ...
	QueryAreaAdmins = "select " +
		queryUserProperties +
		" from area_admin join k_user on user_id=k_user.id where area_id=?"
)

func init() {
	orm.Debug = true

	orm.RegisterModel(
		new(KUser),
		new(Area),
		new(Local),
		new(Reservation),
		new(Notification),
	)
}

// NewModel ...
func NewModel() models2.Model {
	dbName := beego.AppConfig.String("DB_SOURCE_NAME")
	dbUser := beego.AppConfig.String("DB_USER")
	dbPassword := beego.AppConfig.String("DB_PASSWORD")

	dbDriver := "postgres"
	dbAlias := "default"

	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		dbUser, dbName, dbPassword)

	orm.RegisterDriver(dbDriver, orm.DRPostgres)
	orm.RegisterDataBase(dbAlias, dbDriver, conn)

	o := orm.NewOrm()
	m := models2.NewModel(o)
	return m
}

func GetQueryBuilder() (orm.QueryBuilder, error) {
	return orm.NewQueryBuilder("mysql")

	// var x models2.Area

	return nil, nil
}
