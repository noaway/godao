package godao

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	defaultExchangeDB *gorm.DB
)

type Student struct {
	Id        int
	Name      string
	Age       int
	CreatedAt time.Time
	ABA       *string
	WANG      string `gorm:"-"`
}

func (stu *Student) TableName() string {
	return "test_table_" + fmt.Sprintf("%v", stu.Id%16)
}

func TestSQLInjection(t *testing.T) {
	initDb()

	eds := []interface{}{
		&Student{Id: 1, Name: ";and 1=(select IS_SRVROLEMEMBER('serveradmin'));--", Age: 28},
		&Student{Id: 1, Name: ";and 1=(select IS_SRVROLEMEMBER('setupadmin'));-- ", Age: 28},
		&Student{Id: 1, Name: "' or 1=1 --", Age: 28},
		&Student{Id: 1, Name: "&^$#%^&*(*(*(<>+_####''''''''", Age: 2142334343},
		&Student{Id: 1, Name: `';create table temp(id nvarchar(255),num1 nvarchar(255),num2 nvarchar(255),num3 nvarchar`, Age: 28},
		&Student{Id: 1, Name: "+and+1::int=1–", Age: 28},
		&Student{Id: 1, Name: "and 1=cast(user||123 as int)", Age: 28},
		&Student{Id: 1, Name: "order by", Age: 28},
		&Student{Id: 1, Name: "union select null,null,null", Age: 28},
		&Student{Id: 1, Name: "union+select+null,current_database(),null", Age: 28},
		&Student{Id: 1, Name: "union+select+null,relname,null from pg_stat_user_tables", Age: 28},
		&Student{Id: 1, Name: "create table shell(shell text not null);", Age: 28},
		&Student{Id: 1, Name: "select relname from pg_stat_user_tables", Age: 28},
		&Student{Id: 1, Name: "select current_database()", Age: 28},
		&Student{Id: 1, Name: "''''''''''''", Age: 28},
		&Student{Id: 1, Name: "------------", Age: 28},
		&Student{Id: 1, Name: "##########", Age: 28},
		&Student{Id: 1, Name: "select", Age: 28},
		&Student{Id: 1, Name: "insert", Age: 28},
		&Student{Id: 1, Name: "##########", Age: 28},
		&Student{Id: 1, Name: "##########", Age: 28},
		&Student{Id: 1, Name: "##########", Age: 28},
		&Student{Id: 1, Name: "##########", Age: 28},
		&Student{Id: 1, Name: "##########", Age: 28},
		&Student{Id: 1, Name: "##########", Age: 28},
	}
	if err := BatchInsert(defaultExchangeDB, eds); err != nil {
		t.Error(err)
		return
	}
}

func initDb() {
	cnf := PostgreSQLConfig{
		Host:              "127.0.0.1",
		Port:              5432,
		User:              "root",
		Password:          "111111",
		Database:          "postgres",
		SSLMode:           "disable",
		ShowSQL:           true,
		MaxIdleConnection: 1,
		MaxOpenConnection: 1,
	}

	if err := InitExchangeDB(cnf); err != nil {
		panic(err)
	}
}

func InitExchangeDB(c DBConfig) error {
	db, err := Open(c)
	if err != nil {
		return err
	}
	defaultExchangeDB = db
	return nil
}

func TestGenerateSQL(t *testing.T) {
	initDb()
	a := ""

	count := 4
	var tt []interface{}
	tslice := []int{3, 5, 2, 4}

	for i := 0; i < count; i++ {
		stu := &Student{Id: tslice[i], Name: "1", Age: 0, ABA: &a}
		stu1 := &Student{Id: tslice[i], Name: "12", Age: 0, ABA: &a}
		stu2 := &Student{Id: tslice[i], Name: "7072547", Age: 0, ABA: &a}
		tt = append(tt, stu, stu1, stu2)
	}
	t.Log(generateSQL(defaultExchangeDB, tt))

	t.Log("--------------")
	tt = []interface{}{}
	tslice = []int{2, 4, 5, 3}
	for i := 0; i < count; i++ {
		stu := &Student{Id: tslice[i], Name: "1", Age: 0, ABA: &a}
		tt = append(tt, stu)
	}
	t.Log(generateSQL(defaultExchangeDB, tt))
}

func BenchmarkGenerateSql(b *testing.B) {
	initDb()

	n := 500000
	values := make([]interface{}, n, n)
	for i := 0; i < n; i++ {
		values[i] = &Student{
			Id:   1,
			Name: "test test test",
			Age:  10,
		}
	}

	generateSQL(defaultExchangeDB, values)
}
