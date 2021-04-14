package repository

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shubhamsnehi/golang-api-testing/entity"
	// "gorm.io/driver/mysql"
)

var DB *gorm.DB
var err error

type mysqlRepo struct{}

func Open() error {
	//dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open("mysql", "root:root@tcp(172.30.128.1:8080)/test") //DB connection
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/test") //DB connection
	if err != nil {
		log.Println("Could not connect to db")
		return err
	}
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(2)

	db.DB().SetConnMaxLifetime(30 * time.Second)

	DB = db
	return nil
}

func NewMySQLRepository() PostRepository {
	if err = Open(); err != nil {
		log.Println("Unsucessful")
	} else {
		log.Println("DB Connected Sucessfully New")
	}
	return &mysqlRepo{}
}

func (*mysqlRepo) Save(user *entity.User) (*entity.User, error) {
	if err = Open(); err != nil {
		log.Println("Unsucessful")
	} else {
		log.Println("DB Connected Sucessfully New")
	}
	DB.Create(&user)
	return user, nil
}

func (*mysqlRepo) FindAll() ([]entity.User, error) {
	if err = Open(); err != nil {
		log.Println("Unsucessful")
	} else {
		log.Println("DB Connected Sucessfully New")
	}
	posts := []entity.User{}
	DB.Table("users").Find(&posts)
	return posts, nil
}

func (*mysqlRepo) FindByID(id string) (*entity.User, error) {
	if err = Open(); err != nil {
		log.Println("Unsucessful")
	} else {
		log.Println("DB Connected Sucessfully New")
	}
	result := entity.User{}
	DB.Find(&result, id)
	DB.Table("users").Where("id = ?", id)
	return &result, nil
}

func (*mysqlRepo) Delete(post *entity.User) error {
	if err = Open(); err != nil {
		log.Println("Unsucessful")
	} else {
		log.Println("DB Connected Sucessfully New")
	}

	DB.Delete(&post, post.ID)
	return nil
}
