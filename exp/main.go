package main

import "fmt"

func main() {
	menus := []Menu{
		{MenuID: "1", MenuTitle: "sunday", SectionID: "1", SectionTitle: "starter", FoodID: "1", FoodPrice: 200000, FoodName: "Pizza", FoodStatus: 1},
		{MenuID: "1", MenuTitle: "sunday", SectionID: "1", SectionTitle: "starter", FoodID: "3", FoodPrice: 200000, FoodName: "Spagety", FoodStatus: 1},
		{MenuID: "1", MenuTitle: "sunday", SectionID: "2", SectionTitle: "main", FoodID: "1", FoodPrice: 80000, FoodName: "Pizza", FoodStatus: 1},
		// {MenuID: "1", MenuTitle: "sunday", SectionID: "2", SectionTitle: "main", FoodID: "2", FoodPrice: 500000, FoodName: "Hotdog", FoodStatus: 1},
		// {MenuID: "1", MenuTitle: "sunday", SectionID: "2", SectionTitle: "main", FoodID: "1", FoodPrice: 1000000, FoodName: "Pizza", FoodStatus: 1},
		// {MenuID: "1", MenuTitle: "sunday", SectionID: "1", SectionTitle: "starter", FoodID: "2", FoodPrice: 200000, FoodName: "Hotdog", FoodStatus: 1},
	}
	lists := convert(menus)
	fmt.Println(lists)
	iterate()
}
