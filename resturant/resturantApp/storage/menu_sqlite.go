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
		var food *models.SectionFoodType
		for i := range menu.Sections {
			if menu.Sections[i].ID == sectionID {
				section = &menu.Sections[i]
				foodFound := false
				for j, f := range section.Foods {
					if f.Section_id == sectionID && foodID == menu.Sections[i].Foods[j].ID && foodPrice == menu.Sections[i].Foods[j].Price {
						foodFound = true
						break
					}
				}
				if !foodFound {
					food = &models.SectionFoodType{
						Price:       foodPrice,
						Section_id:  sectionID,
						ID:          foodID,
						Name:        foodName,
						Description: foodDescription,
						Status:      foodStatus,
						Photos:      foodPhotos,
					}
					section.Foods = append(section.Foods, *food)
				}
				break
			}
		}
		if section == nil {
			section = &models.SectionViewModel{
				ID: sectionID, Title: sectionTitle,
			}
			food = &models.SectionFoodType{
				Price:       foodPrice,
				Section_id:  sectionID,
				ID:          foodID,
				Name:        foodName,
				Description: foodDescription,
				Status:      foodStatus,
				Photos:      foodPhotos,
			}
			section.Foods = append(section.Foods, *food)
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

func (repo *MenuSqliteDB) Delete(itemId string) error {
	tx, err := repo.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `DELETE FROM section_foods
	WHERE section_id IN (
	  SELECT id FROM sections
	  WHERE menu_id = ?
	);`
	stmtx, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmtx.Exec(itemId)
	if err != nil {
		return err
	}
	query = `DELETE FROM sections
	WHERE menu_id = ?;`
	stmtx, err = tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmtx.Exec(itemId)
	if err != nil {
		return err
	}
	query = `DELETE FROM menus
	WHERE id = ?;`
	stmtx, err = tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmtx.Exec(itemId)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
