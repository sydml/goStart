package main

import "fmt"

func main() {
	var subject map[string]string
	subject = make(map[string]string)
	subject["language"] = "english"
	subject["other"] = "test"
	fmt.Println("language is:", subject["language"])
	for sub := range subject {
		fmt.Println("subject:", subject[sub])
	}
	name, ok := subject["language"]
	if ok {
		fmt.Println("language is", name)
	} else {
		fmt.Println("language is not exists")
	}

	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
