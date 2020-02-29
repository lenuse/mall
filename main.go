package main

import (
	"mall/model"
	"mall/utils"
	"os/signal"
	"syscall"
	"time"

	//"mall/conf"
	"flag"
	"fmt"
	"os"
)

var help bool
var schema string
var httpServer string
var port string

type TableDesc struct {
	Attnum int
	Table  string
	Type   string
	Len    int
	IsNull bool
}

func init() {
	flag.BoolVar(&help, "help", false, "this help")
	flag.StringVar(&schema, "database", "", "数据库表")
	flag.StringVar(&httpServer, "server", "", "yong")
	flag.StringVar(&port, "port", ":8090", "yong")
}

func main() {
	flag.Parse()
	if help {
		if help {
			flag.Usage()
		}
	}
	ctl := make(chan os.Signal, 1)
	signal.Notify(ctl, os.Interrupt, os.Kill)
	go func() {
		for s := range ctl {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				ExitFunc()
			default:
				fmt.Println("other", s)

			}
		}
	}()
	fmt.Println("进程启动...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}

	//if "" != schema {
	//	schemaSlice := strings.Split(schema, ".")
	//	if 2 != len(schemaSlice) {
	//		fmt.Println("schema参数为点号分隔")
	//		os.Exit(1)
	//	}
	//	//schemaName := schemaSlice[0]
	//	//tableName := schemaSlice[1]
	//
	//}

	//app := gin.Default()
	//router.New(app)
	//loggerInit()
	//_ = app.Run(":8090")
}

//func loggerInit() {
//	lf, err := os.OpenFile(conf.NewConf().LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
//	if err != nil {
//		logger.Fatalf("Failed to open log file: %v", err)
//	}
//	defer lf.Close()
//	defer logger.Init("Logger", false, true, lf)
//}

func getTableStructure(schemaName, tableName string) {
	sql := `
	select
	a.attnum,
	a.attname as table,
	t.typname as type,
	a.attlen as len,
	a.attnotnull as is_null
from
	pg_class c,
	pg_attribute a,
	pg_type t,
	pg_namespace n
where
	a.attrelid = c.oid
	and a.atttypid = t.oid
	and n.oid = c.relnamespace
	and c.relname = ?
	and n.nspname = ?
	and a.attnum > 0
order by
	a.attnum;
`
	rows, err := model.NewInst().Raw(sql, tableName, schemaName).Rows()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()
	file, err := os.OpenFile("./model/"+tableName+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fileContent := `package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type %s struct {
`
	fileContent = fmt.Sprintf(fileContent, tableName)
	var structName, structType string
	for rows.Next() {
		var taleDesc TableDesc

		err = model.NewInst().ScanRows(rows, &taleDesc)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		structName = utils.StringToUpper(taleDesc.Table, '_')
		structType = TransformType(taleDesc.Type)
		fileContent += fmt.Sprintf("\n    %s    %s  `json:\"%s\"`", structName, structType, taleDesc.Table)
	}
	fileContent += "\n}"
	n, err := file.WriteString(fileContent)
	if n != len([]byte(fileContent)) || err != nil {
		fmt.Println("写入失败", err)
		os.Exit(0)
	} else {
		fmt.Println("写入成功", err)
		os.Exit(0)
	}
}

func TransformType(dataType string) (res string) {
	switch dataType {
	default:
		res = "string"
		break
	case "bool":
		res = "bool"
		break
	case "bigint", "bigserial":
		res = "int64"
		break
	case "smallint", "integer", "int4", "int8", "real", "double precision", "smallserial", "serial":
		res = "int"
		break
	case "numeric", "decimal":
		res = "float64"
		break
	}
	return
}

func ExitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
