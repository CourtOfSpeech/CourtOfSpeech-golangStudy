package main

import "fmt"

func main() {

	m := map[string]string{
		"name": "jiangxin",
		"addr": "where",
		"call": "183xxxxxxxx",
		"age":  "18",
	}

	//m2 == empty map
	m2 := make(map[string]int)
	//m3 == nil
	var m3 map[string]int

	//%v相应值的默认格式
	fmt.Printf("m: %v, m2: %v, m3: %v\n", m, m2, m3)
	fmt.Println()
	//%#v相应值的Go语法表示
	fmt.Printf("%#v, %#v, %#v\n", m, m2, m3)

	fmt.Println("Traversing map m")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name := m["name"]
	fmt.Println(`m["name"] = `, name)

	//判断map中是否有该键, ok 是bool值
	if addr, ok := m["addr"]; ok {
		fmt.Println(addr)
	} else {
		fmt.Println(`key "addr" does not exist`)
	}

	fmt.Println(`Deleting values`)
	if age, ok := m["age"]; ok {
		fmt.Println(age)
		//删除key
		delete(m, "age")
		if age1, ok1 := m["age"]; ok1 {
			fmt.Println(age1)
			fmt.Println(`not deleted`)
		} else {
			fmt.Println(`key "age" does not exist,deleted`)
		}

	} else {
		fmt.Println(`key "age" does not exist,but not deleted`)
	}

}
