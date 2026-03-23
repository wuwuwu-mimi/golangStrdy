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


## defer
类似于 Java中的finally

#### defer的执行顺序
先调用的最后执行。 先进后出

#### defer和return谁先执行
defer 在方法的声明周期结束之后才会触发
所以说 defer一定在return之后执行


## 数组
### 静态数组 
```golang
    // 定义数组
    var name [nums]type

    name := [nums]type{values...}
```

静态数组作为方法中的参数时，如果要调用这个方法，数组必须跟形参一致 
```golang
    func test(arr [10]int){}
    //调用时只能是 [10]int 类型的
    //且静态数组是 值传递
```

### 动态数组 silce
动态数组是 引用传递
```golang
    // 定义方式
    //1 声明一个切片，并且初始化，默认值是1,2,3 ，长度len是 3
    silce1 := []int{1, 2, 3, 4}

    //2声明一个切片，但是并不分配空间，可以使用make初始化
    var silce2 []int // 此时不可以操作数组
    silce2 = make([]int ,3) // 开辟三个空间

    //3 声明切片，并且分配空间
    var silce3 []int = make([]int,3)

    //也可以直接使用 :=
    silce4 := make([]int ,3)
    
```

#### 切片容量的追加

```golang
    // 表示 声明了一个 长度 len为3 ， 容量为cap 为5 的切片
    silce := make([]int, 3, 5)
	fmt.Printf("len = %d , cap = %d,,silce = %v\n", len(silce), cap(silce), silce) // 3 , 5 , [0,0,0]

    //可以使用 append追加元素
    silce = append(4)
    //此时 长度为 4 容量为 5
    //再追加 就是长度为 5 容量为 5
    silce = append(5)
    fmt.Printf("len = %d , cap = %d,,silce = %v\n", len(silce), cap(silce), silce) // 5,5 , [0,0,0,4,5]

    //如果 再次追加 长度会超过容量 会触发扩容，会将容量增加为原来的2倍 （1.18+）以后 扩容阈值变为 256 扩容变为平滑递增，而不是直接2倍
    silce = append(6)
    fmt.Printf("len = %d , cap = %d,,silce = %v\n", len(silce), cap(silce), silce) // 6,10 , [0,0,0,4,5,6]
```

#### 切片的截取
```golang
    // 可以使用 : 对切片进行截取

    silce := []int{1,2,3,4,5}
    [0:] 从索引0截取到最后
    [0:2] 左闭右开区间截取 0 , 1
    [:endIndex] 从头截取到endIndex - 1 位置的元素
    [:] 截取全部 

    // 截取之后产生的新切片 还是指向 原始切片，对新切片修改，还是会影响到原始切片

    // 可以使用 copy 来产生一个全新切片，也就是深拷贝
```

## map

定义map
var myMap map[type]type 此时不可用，必须使用make分配空间才可以使用
myMap := make(map[type]type) 不需要指定空间，会自动分配
可以直接添加内容
	Mymap3 := map[string]string{
		"one": "Java",
		"two": "Python",
	}
#### 使用方法

```golang
var myMap = map[string]int{
		"China": 1,
	}
	//添加
	myMap["China"] = 1
	myMap["USA"] = 2
	myMap["Japan"] = 3

	// 遍历
	for Key, value := range myMap {
		fmt.Println("key = ", Key)
		fmt.Println("value = ", value)
	}
	//修改
	myMap["USA"] = 0

	//删除
	delete(myMap, "Japan")

	fmt.Println("=====================")

	for Key, value := range myMap {
		fmt.Println("key = ", Key)
		fmt.Println("value = ", value)
	}

    // 在map当做形参时，传递给函数的是 一个指向原map的指针，在方法中对map进行修改，也会对原map修改
    func changeValue(myMap map[stirng]int){}
```
