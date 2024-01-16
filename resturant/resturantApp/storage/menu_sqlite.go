package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"mehrangcode.ir/resturant/app/database"
	"mehrangcode.ir/resturant/app/models"
)

type MenuSqliteDB struct {
	DB *sqlx.DB
}

func NewMenuSqliteDB() *MenuSqliteDB {
	return &MenuSqliteDB{
		DB: database.Connection(),
	}
}

func (repo *MenuSqliteDB) GetAll() ([]models.MenuViewModel, error) {
	query := `
	SELECT 
		m.id AS menu_id,
		m.title AS menu_title,
		s.id AS section_id,
		s.title AS section_title,
		f.id AS food_id,
		sf.price as food_price,
		f.name AS food_name,
		f.description AS food_description,
		f.status AS food_status,
		f.photos AS food_photos
	FROM 
		menus m
	LEFT JOIN 
		sections s ON m.id = s.menu_id
	LEFT JOIN 
		section_foods sf ON s.id = sf.section_id
	LEFT JOIN 
		foods f ON sf.food_id = f.id
	`
	rows, err := repo.DB.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	menus := make(map[string]*models.MenuViewModel)
	for rows.Next() {
		var (
			menuID          string
			menuTitle       string
			sectionID       string
			sectionTitle    string
			foodID          string
			foodPrice       uint
			foodName        string
			foodDescription *string
			foodStatus      uint
			foodPhotos      *string
		)
		err := rows.Scan(&menuID, &menuTitle, &sectionID, &sectionTitle, &foodID, &foodPrice, &foodName, &foodDescription, &foodStatus, &foodPhotos)
		if err != nil {
			return nil, err
		}
		menu, exists := menus[menuID]
		if !exists {
			menu = &models.MenuViewModel{
				ID:    menuID,
				Title: menuTitle,
			}
			menus[menuID] = menu
		}
		var section *models.SectionViewModel
		for i := range menu.Sections {
			if menu.Sections[i].ID == sectionID {
				section = &menu.Sections[i]
				break
			}
		}
		if section == nil {
			section = &models.SectionViewModel{
				ID: sectionID, Title: sectionTitle,
			}
			FoodSection := models.SectionFoodType{
				Price:       foodPrice,
				Section_id:  sectionID,
				ID:          foodID,
				Name:        foodName,
				Description: foodDescription,
				Status:      foodStatus,
				Photos:      foodPhotos,
			}
			section.Foods = append(section.Foods, FoodSection)
			menu.Sections = append(menu.Sections, *section)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	var menusSlice []models.MenuViewModel
	for _, menu := range menus {
		menusSlice = append(menusSlice, *menu)
	}
	return menusSlice, nil
}

func (repo *MenuSqliteDB) Create(payload models.MenuDTO) (string, error) {
	tx, err := repo.DB.Beginx()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	query := `INSERT INTO menus (title,description,status) VALUES(?,?,?)`
	stmtx, err := tx.Prepare(query)
	if err != nil {
		return "", err
	}
	res, err := stmtx.Exec(payload.Title, payload.Description, 1)
	if err != nil {
		return "", err
	}
	menu_id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	fmt.Println("menuID: ", menu_id)
	for _, section := range payload.Sections {
		sectionQuery := `
		INSERT INTO sections 
		(title,description,status,menu_id) 
		VALUES
		(:title,:description,:status,:menu_id)`
		stmtx, err := tx.Prepare(sectionQuery)
		if err != nil {
			return "", err
		}
		res, err := stmtx.Exec(section.Title, section.Description, 1, menu_id)
		if err != nil {
			return "", err
		}
		section_id, err := res.LastInsertId()
		if err != nil {
			return "", err
		}
		fmt.Println("secctionId: ", section_id)

		for _, food := range section.Foods {
			foodQuey := `
			INSERT INTO section_foods (price,food_id,section_id) VALUES (:price,:food_id,:section_id)`
			stmtx, err := tx.Prepare(foodQuey)
			if err != nil {
				return "", err
			}
			_, err = stmtx.Exec(food.Price, food.Food_id, section_id)
			if err != nil {
				return "", err
			}
			fmt.Println("Food", food.Food_id)
		}
	}
	tx.Commit()
	return fmt.Sprint(menu_id), nil
}

// func main() {
// 	// Example data from the database
// 	data := []struct {
// 		MenuID       string
// 		MenuTitle    string
// 		SectionID    string
// 		SectionTitle string
// 		FoodID       string
// 		FoodPrice    int
// 		FoodName     string
// 		FoodStatus   uint
// 	}{
// 		{"1", "sunday", "1", "starter", "1", 200000, "Pizza", 1},
// 		{"1", "sunday", "1", "starter", "3", 200000, "Spagety", 1},
// 		{"1", "sunday", "2", "main", "1", 80000, "Pizza", 1},
// 		{"1", "sunday", "2", "main", "2", 500000, "Hot dog", 1},
// 		{"1", "sunday", "2", "main", "1", 1000000, "Pizza", 1},
// 		// Add more records as needed
// 	}

// 	// Process the data into the desired JSON structure
// 	menusMap := make(map[string]*Menu)
// 	for _, record := range data {
// 		menu, exists := menusMap[record.MenuID]
// 		if !exists {
// 			menu = &Menu{
// 				ID:    record.MenuID,
// 				Title: record.MenuTitle,
// 			}
// 			menusMap[record.MenuID] = menu
// 		}

// 		var section *Section
// 		found := false
// 		for i, sec := range menu.Sections {
// 			if sec.ID == record.SectionID {
// 				section = &menu.Sections[i]
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			section = &Section{
// 				ID:    record.SectionID,
// 				Title: record.SectionTitle,
// 			}
// 			menu.Sections = append(menu.Sections, *section)
// 		}

// 		food := Food{
// 			Price: record.FoodPrice,
// 			Name:  record.FoodName,
// 			// Description and Photos are nil in the example data
// 			Status: record.FoodStatus,
// 		}
// 		section.Foods = append(section.Foods, food)
// 	}

// 	// Convert the menus map to a slice
// 	var menus []Menu
// 	for _, menu := range menusMap {
// 		menus = append(menus, *menu)
// 	}

// 	// Marshal the slice to JSON
// 	jsonData, err := json.MarshalIndent(menus, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshaling to JSON:", err)
// 		return
// 	}

// 	// Print the JSON
// 	fmt.Println(string(jsonData))
// }
