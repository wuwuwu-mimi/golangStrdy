# go


## 变量声明方式

1.  var 变量名 数据类型  有数据类型的默认值
2.  var 变量名 数据类型 := 初始化值
3.  var 变量名 := 初始化值  由编译器 判断
4.  变量名 := 值 可以直接省略 var 和 类型  但是只能使用在局部变量

声明多个变量
var v1,v2 = 值，值
var(
     v1 类型 = 值
     v2 类型 = 值
     v3 类型 = 值
)

## const

可以使用 const 关键字定义常量
```golang
    const length int = 1000
```

可以使用 const 定义枚举
```golang
    const(
        BEIJING = 10
        SHANGHAI = 20
        SHENZHEN = 30
    )

    //在枚举中可以使用 iota 每行的iota都会累加1 ，第一个iota默认是0
    // 可以使用一个 常数d 对 iota操作 例如 2* iota
    const(
        BEIJING = iota //0
        SHANGHAI        //1
        SHENZHEN        //2
    )
```

## 函数

**定义函数**

```golang
    // 单返回值
    func funcName(v1,v2) type {

    }

    //多返回值 匿名
    func funcName(v1,v2) (type,type,...) {

    }
    // 多返回值 有形参名称
    func funcName(v1,v2) (name1 type ,name2 type,...){
        reutrn name1,name2
    }
```

### 使用自定义包下的函数
自定义包下的对外函数必须大写开头
```golang
    import(
        _"包名" // 表示匿名，无法使用这个包中的方法，但是会自动执行包下的init方法
        别名 "包名" // 给包起了一个别名 可以用别名.方法来调用方法
        . "包名" // 可以直接使用包下的方法，不需要 包名.方法
    )
```


## 指针
 和c一样

