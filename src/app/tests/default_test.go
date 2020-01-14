package test

import (
	"bytes"
	_"context"
	_"database/sql"
	"fmt"
	"io/ioutil"
	_"math"
	_"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
	"app/models"
	 _ "github.com/go-sql-driver/mysql" // import your required driver
)



func TestRegisterModels(t *testing.T) {
	orm.RegisterDataBase("default", "mysql", "root:root"+"@/rationing_system?charset=utf8", 30)
	orm.RegisterModel(new(models.Ration))
}

func TestCRUD(t *testing.T) {
	dORM := orm.NewOrm()
	var ration models.Ration
	ration.PacketType = "Food"
	ration.PacketContent = "MRE"
	ration.Calories = 1000
	ration.ExpiryDate = time.Now()
	_, err := dORM.Insert(&ration)
	throwFail(t, err)

	r := &models.Ration{Id: ration.Id}
	err = dORM.Read(r)
	throwFail(t, err)

	throwFail(t, AssertIs(r.PacketType, "Food"))
	throwFail(t, AssertIs(r.PacketContent, "MRE"))
	throwFail(t, AssertIs(r.Calories, 1000))

	ration.PacketContent = "Biscuit"
	num, err := dORM.Update(&ration)
	throwFail(t, err)
	throwFail(t, AssertIs(num, 1))

	r = &models.Ration{Id: ration.Id}
	err = dORM.Read(r)
	throwFailNow(t, err)
	throwFail(t, AssertIs(r.PacketContent, "Biscuit"))

	num, err = dORM.Delete(&ration)
	throwFail(t, err)
	throwFail(t, AssertIs(num, 1))
}

type argInt []int

// get int by index from int slice
func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}


// ToStr interface to string
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type argAny []interface{}

// get interface by index from interface slice
func (a argAny) Get(i int, args ...interface{}) (r interface{}) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

func ValuesCompare(is bool, a interface{}, args ...interface{}) (ok bool, err error) {
	if len(args) == 0 {
		return false, fmt.Errorf("miss args")
	}
	b := args[0]
	arg := argAny(args)

	switch v := a.(type) {
	case reflect.Kind:
		ok = reflect.ValueOf(b).Kind() == v
	case time.Time:
		if v2, vo := b.(time.Time); vo {
			if arg.Get(1) != nil {
				format := ToStr(arg.Get(1))
				a = v.Format(format)
				b = v2.Format(format)
				ok = a == b
			} else {
				err = fmt.Errorf("compare datetime miss format")
				goto wrongArg
			}
		}
	default:
		ok = ToStr(a) == ToStr(b)
	}
	ok = is && ok || !is && !ok
	if !ok {
		if is {
			err = fmt.Errorf("expected: `%v`, get `%v`", b, a)
		} else {
			err = fmt.Errorf("expected: `%v`, get `%v`", b, a)
		}
	}

wrongArg:
	if err != nil {
		return false, err
	}

	return true, nil
}

func AssertIs(a interface{}, args ...interface{}) error {
	if ok, err := ValuesCompare(true, a, args...); !ok {
		return err
	}
	return nil
}

func AssertNot(a interface{}, args ...interface{}) error {
	if ok, err := ValuesCompare(false, a, args...); !ok {
		return err
	}
	return nil
}

func getCaller(skip int) string {
	pc, file, line, _ := runtime.Caller(skip)
	fun := runtime.FuncForPC(pc)
	_, fn := filepath.Split(file)
	data, err := ioutil.ReadFile(file)
	var codes []string
	if err == nil {
		lines := bytes.Split(data, []byte{'\n'})
		n := 10
		for i := 0; i < n; i++ {
			o := line - n
			if o < 0 {
				continue
			}
			cur := o + i + 1
			flag := "  "
			if cur == line {
				flag = ">>"
			}
			code := fmt.Sprintf(" %s %5d:   %s", flag, cur, strings.Replace(string(lines[o+i]), "\t", "    ", -1))
			if code != "" {
				codes = append(codes, code)
			}
		}
	}
	funName := fun.Name()
	if i := strings.LastIndex(funName, "."); i > -1 {
		funName = funName[i+1:]
	}
	return fmt.Sprintf("%s:%s:%d: \n%s", fn, funName, line, strings.Join(codes, "\n"))
}

func throwFail(t *testing.T, err error, args ...interface{}) {
	if err != nil {
		con := fmt.Sprintf("\t\nError: %s\n%s\n", err.Error(), getCaller(2))
		if len(args) > 0 {
			parts := make([]string, 0, len(args))
			for _, arg := range args {
				parts = append(parts, fmt.Sprintf("%v", arg))
			}
			con += " " + strings.Join(parts, ", ")
		}
		t.Error(con)
		t.Fail()
	}
}

func throwFailNow(t *testing.T, err error, args ...interface{}) {
	if err != nil {
		con := fmt.Sprintf("\t\nError: %s\n%s\n", err.Error(), getCaller(2))
		if len(args) > 0 {
			parts := make([]string, 0, len(args))
			for _, arg := range args {
				parts = append(parts, fmt.Sprintf("%v", arg))
			}
			con += " " + strings.Join(parts, ", ")
		}
		t.Error(con)
		t.FailNow()
	}
}