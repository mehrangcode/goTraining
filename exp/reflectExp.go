package main

import "fmt"

// Define the Menu struct
type Menu struct {
	MenuID       string
	MenuTitle    string
	SectionID    string
	SectionTitle string
	FoodID       string
	FoodPrice    uint
	FoodName     string
	FoodStatus   uint
}

// Define the nested struct
type List struct {
	ID       string
	Title    string
	Sections []struct {
		ID    string
		Title string
		Foods []struct {
			SetctionID string
			ID         string
			Name       string
			Price      uint
			Status     uint
		}
	}
}

// Define a function to convert a slice of Menu to a slice of List
func convert(menus []Menu) []List {
	// Create a map to group the menus by their ID and Title
	menusMap := make(map[string]*List)
	for _, menu := range menus {
		// Get or create the list
		list, exists := menusMap[menu.MenuID]
		fmt.Println("List: ", list, exists, menu.MenuID)
		if !exists {
			list = &List{
				ID:    menu.MenuID,
				Title: menu.MenuTitle,
			}
			menusMap[menu.MenuID] = list
		}

		var section *struct {
			ID    string
			Title string
			Foods []struct {
				SetctionID string
				ID         string
				Name       string
				Price      uint
				Status     uint
			}
		}
		found := false
		for i, sec := range list.Sections {
			if sec.ID == menu.SectionID {
				section = &list.Sections[i]
				var food *struct {
					SetctionID string
					ID         string
					Name       string
					Price      uint
					Status     uint
				}
				foodFound := false
				for i, f := range section.Foods {
					if f.SetctionID == menu.SectionID && f.ID == menu.FoodID {
						food = &section.Foods[i]
						foodFound = true
						break
					}
				}
				if !foodFound {
					food = &struct {
						SetctionID string
						ID         string
						Name       string
						Price      uint
						Status     uint
					}{
						SetctionID: menu.SectionID,
						ID:         menu.FoodID,
						Name:       menu.FoodName,
						Price:      menu.FoodPrice,
						Status:     menu.FoodStatus,
					}
				}
				section.Foods = append(section.Foods, *food)
				found = true
				break
			}
		}
		if !found {
			section = &struct {
				ID    string
				Title string
				Foods []struct {
					SetctionID string
					ID         string
					Name       string
					Price      uint
					Status     uint
				}
			}{
				ID:    menu.SectionID,
				Title: menu.SectionTitle,
			}

			var food *struct {
				SetctionID string
				ID         string
				Name       string
				Price      uint
				Status     uint
			}
			foodFound := false
			for i, f := range section.Foods {
				if f.SetctionID == menu.SectionID && f.ID == menu.FoodID {
					food = &section.Foods[i]
					foodFound = true
					break
				}
			}
			if !foodFound {
				food = &struct {
					SetctionID string
					ID         string
					Name       string
					Price      uint
					Status     uint
				}{
					SetctionID: menu.SectionID,
					ID:         menu.FoodID,
					Name:       menu.FoodName,
					Price:      menu.FoodPrice,
					Status:     menu.FoodStatus,
				}
			}
			section.Foods = append(section.Foods, *food)
			list.Sections = append(list.Sections, *section)
		}

	}

	// Convert the map to a slice
	var lists []List
	for _, list := range menusMap {
		lists = append(lists, *list)
	}

	return lists
}
