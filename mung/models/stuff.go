package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Stuff struct {
	Model
	Gender int `json:"gender"`
}


func str(){
	//string 常用函数
	//len 字符串长度 一个中文等于3个字节 所以是9
	str1 := len("string北")
	fmt.Println("str1 = ",str1)

	str2 := "string北京"
	// 用[]rune转换输出中文,扩大每个字节的容量了,输出完整的中文
	str2r := []rune(str2)
	for i:=0; i < len(str2r); i++ {
		fmt.Printf("str2 = %c \n", str2r[i])
	}

	//字符串转[]byte ：var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("成功结果：%v 数据类型：%T \n",bytes,bytes)

	//[]byte 转 字符串：str = string([]byte{97,98,99})
	bytes1 := string([]byte{97,98,99})
	fmt.Printf("成功结果：%v 数据类型：%T \n",bytes1,bytes1)

	//10进制 转换对应的 2，8，16进制：str := strconv.FormatInt(123,2) 返回对应的字符串
	str4 := strconv.FormatInt(123,2)
	fmt.Printf("成功结果：%v 数据类型：%T \n",str4,str4)

	//查找子串是否在指定的字符串中：strings.Contains("seafood","foo")
	str5 := strings.Contains("seafood","foo")
	fmt.Printf("seafood 中 是否包含 foo: %v \n",str5)

	//统计一个字符串有几个指定的字符 strings.Count("ceheese","e")
	str6 := strings.Count("ceheese","e")
	fmt.Printf("ceheese 中有%v个e字符\n",str6)

	//不区分大小写的字符串比较(==是区分字符大小写的)
	str7 := strings.EqualFold("Abc","abc")
	fmt.Println("abc 和 Abc 的比较是",str7)

	//返回字符串第一次出现的index值，如果没有返回-1
	str8 := strings.Index("hello world","l")
	fmt.Println(str8)

	//返回子串在字符串最后一次出现的index，如没有返回-1
	str9 := strings.LastIndex("hello world", "o")
	fmt.Println(str9)

	//将指定的子串替换成另外一个子串 n可以指定你希望替换第几个，如果n=-1表示全部替换
	str10 := strings.Replace("hello world", "wo", "ha",1)
	fmt.Println(str10)

	//按照指定的某个字符，为分割表示，将一个字符串拆分成 字符串数组  explode ,
	str11 := strings.Split("hell,lo,wor,ld",",")
	fmt.Println(str11)

	//将字符串的字母进行大小写的转换
	str12 := strings.ToUpper("abc")
	str13 := strings.ToLower("Abc")
	fmt.Println(str12,str13)

	//将字符串左右两边的空格去掉
	str14 := strings.TrimSpace(" h o l lo,word ")
	fmt.Println(str14)

	//将字符串左右两边指定的字符去掉
	str15 := strings.Trim("!hello world +!"," !")
	fmt.Println(str15)

	//将字符串左边的字符去掉strings.TrimLeft("! hello!"," !")
	str16 := strings.TrimLeft("! hello!"," !")
	fmt.Println(str16)

	//判断字符串是否以指定的字符串开头
	str18 := strings.HasPrefix("http://baiu.com","http")
	fmt.Println(str18)

	//判断字符串是否以指定的字符串结尾
	str19 := strings.HasSuffix("NLT_adc.jpg","jpg")
	fmt.Println(str19)

}



func conv() {
	// int转string
	s1 := strconv.Itoa(10)
	fmt.Printf("1=: %T %s \n", s1, s1)
	// int64转string
	s2 := strconv.FormatInt(int64(10), 10)
	fmt.Printf("1=: %T %s \n", s2, s2)

	// string转int
	i1, err := strconv.Atoi("12")
	fmt.Printf("2=: %T %v  %T %v \n", i1, i1, err, err)
	// string转int64
	i2, _ := strconv.ParseInt("12",10, 64)
	fmt.Printf("2=: %T %v \n", i2, i2)

	// float转string
	v := 3.1415926535
	f1 := strconv.FormatFloat(v, 'E', -1, 32) //32
	f2 := strconv.FormatFloat(v, 'E', -1, 64) //64
	fmt.Printf("3=: %T %v \n", f1, f1)
	fmt.Printf("3=: %T %v \n", f2, f2)

	// string转float
	sf := "3.1415926535"
	f4, err := strconv.ParseFloat(sf, 64)
	fmt.Printf("4=: %T %+v \n", f4, f4)

	// 当前日期时间
	t1 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("time=: %T %v \n", t1, t1)

	// 随机数
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("rand=: %T %v \n", r1.Intn(100), r1.Intn(100))

	JsonToMap()
	MapToJson()

	JsonToStruct()
	StructToJson()
}

// JsonToMap json转map
func JsonToMap() {
	jsonStr := `{"name": "zhangsan", "age": 10,"city":"hunan"}`
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Printf("json2map: %T %v \n", m, m)
}
// MapToJson Map转json
func MapToJson(){
	m := map[string]interface{}{"name": "zhangsan", "age": 10, "city": "湖南"}
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("map2json: %T %v \n", string(jsonStr), string(jsonStr))
}

// JsonToStruct json转struct
func JsonToStruct(){
	jsonStr := `{"name": "zhTan", "phone": "13684578541"}`
	var usr User
	json.Unmarshal([]byte(jsonStr), &usr)
	fmt.Printf("json2struct:  %#v \n", usr)
}
// StructToJson struct转json
func StructToJson(){
	src := &User{Nickname: "zsx",Phone: "1369897989",Account: "wx931293129"}
	b, err := json.Marshal(src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("struct2json:  %#v \n", string(b))
}

