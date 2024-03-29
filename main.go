package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {

	fmt.Println("hello. table to struct .. ")

	// 初始化
	t2s := converter.NewTable2Struct()
	// 个性化配置
	t2s.Config(&converter.T2tConfig{
		StructNameToHump: true,  // 结构体名称是否转为驼峰式，默认为false
		RmTagIfUcFirsted: false, // 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		TagToLower:       true,  // tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		JsonTagToHump:    false, // json tag是否转为驼峰，默认为false，不转换
		UcFirstOnly:      true,  // 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		SeperatFile:      false, // 每个struct放入单独的文件,默认false,放入同一个文件 (暂未提供)
	})
	// 开始迁移转换
	err := t2s.
		// 指定某个表,如果不指定,则默认全部表都迁移
		//Table("user").
		// 表前缀
		//Prefix("prefix_").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("model").
		// tag字段的key值,默认是orm
		TagKey("gorm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		// 生成的结构体保存路径
		SavePath("./model/model.go").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn("root:root@tcp(localhost:3306)/mydb?charset=utf8").
		// 执行
		Run()

	if err != nil {
		fmt.Println("generate struct fault. err:", err.Error())
		return
	}
	fmt.Println("generate struct success ..")

}
