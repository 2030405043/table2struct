

    使用方法:
        
        1.使用git工具clone项目
          git clone https://github.com/2030405043/table2struct.git
        
        2.在 main() 函数中修改数据库地址、端口、数据库  启动 main() 函数 即在model/model.go文件中生成所有表映射的结构体

    
    官方教程

    一、获取
        go get -u github.com/gohouse/converter
        
     
    二、 GitHub 地址：https://github.com/gohouse/converter
    
        
    三、示例表结构
      
      CREATE TABLE `prefix_user` (
        `Id` int(11) NOT NULL AUTO_INCREMENT,
        `Email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
        `Password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
        `CreatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (`Id`)
      ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表'   
    
    
    四、用法
    
    简单代码示例
    package main
    import (
    	"fmt"
    	"github.com/gohouse/converter"
    )
    func main() {
    	err := converter.NewTable2Struct().
    		SavePath("/home/go/project/model/model.go").
    		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
    		Run()
    	fmt.Println(err)
    }
    
    详细示例
    package main
    
    import (
    	"fmt"
    	"github.com/gohouse/converter"
    )
    
    func main() {
    	// 初始化
    	t2t := converter.NewTable2Struct()
    	// 个性化配置
    	t2t.Config(&converter.T2tConfig{
    		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
    		RmTagIfUcFirsted: false,
    		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
    		TagToLower: false,
    		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
    		UcFirstOnly: false,
    		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
    		//SeperatFile: false,
    	})
    	// 开始迁移转换
    	err := t2t.
    		// 指定某个表,如果不指定,则默认全部表都迁移
    		Table("user").
    		// 表前缀
    		Prefix("prefix_").
    		// 是否添加json tag
    		EnableJsonTag(true).
    		// 生成struct的包名(默认为空的话, 则取名为: package model)
    		PackageName("model").
    		// tag字段的key值,默认是orm
    		TagKey("orm").
    		// 是否添加结构体方法获取表名
    		RealNameMethod("TableName").
    		// 生成的结构体保存路径
    		SavePath("/Users/fizz/go/src/github.com/gohouse/gupiao/model/model.go").
    		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
    		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
    		// 执行
    		Run()
    	
    	fmt.Println(err)
    }
    
    
    五、result
    
    package model
    
    import "time"
    
    type User struct {
    	Id         int     `json:"Id" orm:"Id"`
    	Email      string  `json:"Email" orm:"Email"`           // 邮箱
    	Password   string  `json:"Password" orm:"Password"`     // 密码
    	CreatedAt  string  `json:"CreatedAt" orm:"CreatedAt"`
    }
    
    func (*User) TableName() string {
    	return "user"
    }
    
