package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (repo *MenuSqliteDB) Create(payload models.MenuDTO) (string, error) {
	tx, err := repo.DB.Beginx()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	// Insert menu data
	menuID, err := insertMenu(tx, payload)
	if err != nil {
		return "", err
	}

	// Insert section data
	for _, section := range payload.Sections {
		sectionID, err := insertSection(tx, section, menuID)
		if err != nil {
			return "", err
		}

		// Insert food data
		for _, food := range section.Foods {
			err := insertFood(tx, food, sectionID)
			if err != nil {
				return "", err
			}
		}
	}

	tx.Commit()
	return fmt.Sprint(menuID), nil
}

// insertMenu inserts a menu into the menus table and returns the menu ID
func insertMenu(tx *sqlx.Tx, payload models.MenuDTO) (int64, error) {
	query := `INSERT INTO menus (title,description,status) VALUES(?,?,?)`
	hash, err := utils.HashingPassword("1234")
	if err != nil {
		return 0, err
	}
	res, err := tx.Exec(query, payload.Title, payload.Description, 1, hash)
	if err != nil {
		return 0, err
	}
	menuID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("menuID: ", menuID)
	return menuID, nil
}

// insertSection inserts a section into the sections table and returns the section ID
func insertSection(tx *sqlx.Tx, section models.SectionDTO, menuID int64) (int64, error) {
	query := `
	INSERT INTO sections 
	(title,description,status,menu_id) 
	VALUES
	(:title,:description,:status,:menu_id)`
	res, err := tx.NamedExec(query, section)
	if err != nil {
		return 0, err
	}
	sectionID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("sectionID: ", sectionID)
	return sectionID, nil
}

// insertFood inserts a food into the section_foods table
func insertFood(tx *sqlx.Tx, food models.FoodDTO, sectionID int64) error {
	query := `
	INSERT INTO section_foods (price,food_id,section_id) VALUES (:price,:food_id,:section_id)`
	_, err := tx.NamedExec(query, food)
	if err != nil {
		return err
	}
	fmt.Println("Food", food.Food_id)
	return nil
}

// GetAll returns all the menus from the database
func (repo *MenuSqliteDB) GetAll() ([]models.MenuDTO, error) {
	// Create an empty slice of MenuDTO
	menus := []models.MenuDTO{}

	// Query the menus table and join with the sections and section_foods tables
	query := `
	SELECT m.id, m.title, m.description, m.status, s.id, s.title, s.description, s.status, sf.price, sf.food_id
	FROM menus m
	LEFT JOIN sections s ON m.id = s.menu_id
	LEFT JOIN section_foods sf ON s.id = sf.section_id
	ORDER BY m.id, s.id, sf.food_id
	`
	rows, err := repo.DB.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the rows and scan the values into variables
	var (
		menuID, sectionID, foodID                      int
		menuTitle, menuDesc, sectionTitle, sectionDesc string
		menuStatus, sectionStatus                      bool
		foodPrice                                      float64
	)

	for rows.Next() {
		err := rows.Scan(&menuID, &menuTitle, &menuDesc, &menuStatus, &sectionID, &sectionTitle, &sectionDesc, &sectionStatus, &foodPrice, &foodID)
		if err != nil {
			return nil, err
		}

		// Check if the current menu is already in the slice
		menuIndex := findMenuIndex(menus, menuID)
		if menuIndex == -1 {
			// If not, create a new menu and append it to the slice
			menu := models.MenuDTO{
				ID:          menuID,
				Title:       menuTitle,
				Description: menuDesc,
				Status:      menuStatus,
			}
			menus = append(menus, menu)
			menuIndex = len(menus) - 1
		}

		// Check if the current section is already in the menu
		sectionIndex := findSectionIndex(menus[menuIndex].Sections, sectionID)
		if sectionIndex == -1 {
			// If not, create a new section and append it to the menu
			section := models.SectionDTO{
				ID:          sectionID,
				Title:       sectionTitle,
				Description: sectionDesc,
				Status:      sectionStatus,
			}
			menus[menuIndex].Sections = append(menus[menuIndex].Sections, section)
			sectionIndex = len(menus[menuIndex].Sections) - 1
		}

		// Check if the current food is already in the section
		foodIndex := findFoodIndex(menus[menuIndex].Sections[sectionIndex].Foods, foodID)
		if foodIndex == -1 {
			// If not, create a new food and append it to the section
			food := models.FoodDTO{
				Food_id: foodID,
				Price:   foodPrice,
			}
			menus[menuIndex].Sections[sectionIndex].Foods = append(menus[menuIndex].Sections[sectionIndex].Foods, food)
		}
	}

	// Return the slice of menus and nil error
	return menus, nil
}

// findMenuIndex returns the index of a menu in a slice of menus, or -1 if not found
func findMenuIndex(menus []models.MenuDTO, menuID int) int {
	for i, menu := range menus {
		if menu.ID == menuID {
			return i
		}
	}
	return -1
}

// findSectionIndex returns the index of a section in a slice of sections, or -1 if not found
func findSectionIndex(sections []models.SectionDTO, sectionID int) int {
	for i, section := range sections {
		if section.ID == sectionID {
			return i
		}
	}
	return -1
}

// findFoodIndex returns the index of a food in a slice of foods, or -1 if not found
func findFoodIndex(foods []models.FoodDTO, foodID int) int {
	for i, food := range foods {
		if food.Food_id == foodID {
			return i
		}
	}
	return -1
}
