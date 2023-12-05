package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Record struct {
	ID         int64  `json:"-" sql.field:"id"`
	Name       string `json:"name" sql.field:"name"`
	LastName   string `json:"last_name" sql.field:"last_name"`
	MiddleName string `json:"middle_name" sql.field:"middle_name"`
	Address    string `json:"address" sql.field:"address"`
	Phone      string `json:"phone" sql.field:"phone"`
}

func connentToServer() {
	for {
		///// мб здесь сделать проверку доступности сайта
		var command int
		fmt.Print("Выберите, что хотите сделать [1 - Создать запись, 2 - Обновить запись, 3 - Найти запись, 4 - Удалить запись]: ")
		fmt.Scanln(&command)
		if command == 1 {
			confirm := "yes"
			fmt.Print("Эта функция создаёт запись в адресной книге. Хотите продолжить? [yes]: ")
			fmt.Scanln(&confirm)
			if confirm != "yes" {
				continue
			}
			rec := &Record{}
			fmt.Print("Введите имя: ")
			fmt.Scanln(&rec.Name)
			fmt.Print("Введите фамилию: ")
			fmt.Scanln(&rec.MiddleName)
			fmt.Print("Введите отчество (при наличии): ")
			fmt.Scanln(&rec.LastName)
			fmt.Print("Введите адрес: ")
			fmt.Scanln(&rec.Address)
			fmt.Print("Введите номер телефона: ")
			fmt.Scanln(&rec.Phone)
			createRecord(rec)
			continue
		}
		if command == 2 {
			confirm := "yes"
			fmt.Print("Эта функция обновяет запись в адресной книге. Хотите продолжить? [yes]: ")
			fmt.Scanln(&confirm)
			if confirm != "yes" {
				continue
			}
			rec := &Record{}
			fmt.Print("Введите номер телефона, для которого обновляете данные: ")
			fmt.Scanln(&rec.Phone)
			fmt.Println("Введите те данные, которые хотите обновить. Остальные пропускайте")
			fmt.Print("Введите имя: ")
			fmt.Scanln(&rec.Name)
			fmt.Print("Введите фамилию: ")
			fmt.Scanln(&rec.MiddleName)
			fmt.Print("Введите отчество: ")
			fmt.Scanln(&rec.LastName)
			fmt.Print("Введите адрес: ")
			fmt.Scanln(&rec.Address)
			updateRecord(rec)
			continue
		}
		if command == 3 {
			confirm := "yes"
			fmt.Print("Эта функция поиска записи в адресной книге по известным данным. Хотите продолжить? [yes]: ")
			fmt.Scanln(&confirm)
			if confirm != "yes" {
				continue
			}
			rec := &Record{}
			fmt.Println("Введите те данные, по которым хотите найти запись")
			fmt.Print("Имя: ")
			fmt.Scanln(&rec.Name)
			fmt.Print("Фамилия: ")
			fmt.Scanln(&rec.MiddleName)
			fmt.Print("Отчество: ")
			fmt.Scanln(&rec.LastName)
			fmt.Print("Адрес: ")
			fmt.Scanln(&rec.Address)
			fmt.Print("Номер телефона: ")
			fmt.Scanln(&rec.Phone)
			getRecords(rec)
			continue
		}
		if command == 4 {
			confirm := "yes"
			fmt.Print("Эта функция удаления записи из адресной книги. Хотите продолжить? [yes]: ")
			fmt.Scanln(&confirm)
			if confirm != "yes" {
				continue
			}
			var phone []byte
			fmt.Print("Введите номер телефона, который хотите удалить из адресной книги: ")
			fmt.Scanln(&phone)
			deleteRecord(phone)
			continue
		}
		fmt.Println("Пожалуйста, введите корректную команду")
	}
}

func createRecord(rec *Record) {
	// recordData := map[string]interface{}{
	// 	"name":        "John",
	// 	"last_name":   "Doe",
	// 	"middle_name": "M.",
	// 	"address":     "123 Main St",
	// 	"phone":       "+791356640274",
	// }

	jsonData, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func updateRecord(rec *Record) {
	// updatedRecordData := map[string]interface{}{
	// 	"name":        "UpdatedName",
	// 	"last_name":   "UpdatedLastName",
	// 	"middle_name": "UpdatedMiddleName",
	// 	"address":     "UpdatedAddress",
	// 	"phone":       "79135640274",
	// }

	jsonData, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/update", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	// url := fmt.Sprintf("http://localhost:8080/update")
	// req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	fmt.Println("Error creating PUT request:", err)
	// 	return
	// }

	// client := http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending PUT request:", err)
	// 	return
	// }
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func deleteRecord(phone []byte) {
	resp, err := http.Post("http://localhost:8080/delete", "application/text", bytes.NewBuffer(phone))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	// url := fmt.Sprintf("http://localhost:8080/delete")
	// phone := "79135640274"
	// req, err := http.NewRequest("DELETE", url, bytes.NewBuffer([]byte(phone)))
	// if err != nil {
	// 	fmt.Println("Error creating DELETE request:", err)
	// 	return
	// }

	// client := http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending DELETE request:", err)
	// 	return
	// }
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func getRecords(rec *Record) {

	// jsonData := `{
	// 	"name":        "UpdatedName",
	// 	"last_name":   "UpdatedLastName",
	// 	"middle_name": "",
	// 	"address":     "",
	// 	"phone":       ""
	// }`

	jsonData, err := json.Marshal(rec)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	url := "http://localhost:8080/get"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	var recordsData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&recordsData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Records:", recordsData) ////////////////// красивый вывод
}

func main() {
	connentToServer()
}
