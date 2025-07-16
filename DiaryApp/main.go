package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var allDiaries = &Diaries{}
var currentUserDiary *Diary

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Diary Application ---")
		fmt.Println("1. Create Diary")
		fmt.Println("2. Unlock Diary")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			createDiary(reader)
		case "2":
			unlockDiary(reader)
		case "3":
			fmt.Println("Exiting application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func createDiary(reader *bufio.Reader) {
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if allDiaries.FindByUsername(username) != nil {
		fmt.Println("Diary with this username already exists.")
		return
	}

	diary := NewDiary(username, password)
	allDiaries.Add(diary)
	fmt.Println("Diary created successfully!")
}

func unlockDiary(reader *bufio.Reader) {
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	diary := allDiaries.FindByUsername(username)
	if diary == nil {
		fmt.Println("Diary not found.")
		return
	}

	diary.UnlockDiary(password)
	if diary.IsLocked() {
		fmt.Println("Incorrect password.")
	} else {
		currentUserDiary = diary
		fmt.Println("Diary unlocked successfully!")
		diaryMenu(reader)
	}
}

func diaryMenu(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Diary Menu ---")
		fmt.Println("1. Add Entry")
		fmt.Println("2. Find Entry by ID")
		fmt.Println("3. Update Entry")
		fmt.Println("4. Delete Entry")
		fmt.Println("5. Lock Diary")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addEntry(reader)
		case "2":
			findEntryByID(reader)
		case "3":
			updateEntry(reader)
		case "4":
			deleteEntry(reader)
		case "5":
			currentUserDiary.LockDiary()
			currentUserDiary = nil
			fmt.Println("Diary locked.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addEntry(reader *bufio.Reader) {
	if currentUserDiary == nil || currentUserDiary.IsLocked() {
		fmt.Println("Please unlock a diary first.")
		return
	}

	fmt.Print("Enter entry title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter entry body: ")
	body, _ := reader.ReadString('\n')
	body = strings.TrimSpace(body)

	currentUserDiary.CreateEntry(title, body)
	fmt.Println("Entry added successfully!")
}

func findEntryByID(reader *bufio.Reader) {
	if currentUserDiary == nil || currentUserDiary.IsLocked() {
		fmt.Println("Please unlock a diary first.")
		return
	}

	fmt.Print("Enter entry ID: ")
	inputID, _ := reader.ReadString('\n')
	entryID, err := strconv.Atoi(strings.TrimSpace(inputID))
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	entry := currentUserDiary.FindEntryByID(entryID)
	if entry == nil {
		fmt.Println("Entry not found.")
	} else {
		fmt.Printf("\nEntry Found:\nID: %d\nTitle: %s\nBody: %s\nDate Created: %s\n",
			entry.ID, entry.Title, entry.Body, entry.DateCreated.Format("2006-01-02 15:04:05"))
	}
}

func updateEntry(reader *bufio.Reader) {
	if currentUserDiary == nil || currentUserDiary.IsLocked() {
		fmt.Println("Please unlock a diary first.")
		return
	}

	fmt.Print("Enter entry ID to update: ")
	inputID, _ := reader.ReadString('\n')
	entryID, err := strconv.Atoi(strings.TrimSpace(inputID))
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	entry := currentUserDiary.FindEntryByID(entryID)
	if entry == nil {
		fmt.Println("Entry not found.")
		return
	}

	fmt.Printf("Current Title: %s\nEnter new title (leave blank to keep current): ", entry.Title)
	newTitle, _ := reader.ReadString('\n')
	newTitle = strings.TrimSpace(newTitle)
	if newTitle == "" {
		newTitle = entry.Title
	}

	fmt.Printf("Current Body: %s\nEnter new body (leave blank to keep current): ", entry.Body)
	newBody, _ := reader.ReadString('\n')
	newBody = strings.TrimSpace(newBody)
	if newBody == "" {
		newBody = entry.Body
	}

	currentUserDiary.UpdateEntry(entryID, newTitle, newBody)
	fmt.Println("Entry updated successfully!")
}

func deleteEntry(reader *bufio.Reader) {
	if currentUserDiary == nil || currentUserDiary.IsLocked() {
		fmt.Println("Please unlock a diary first.")
		return
	}

	fmt.Print("Enter entry ID to delete: ")
	inputID, _ := reader.ReadString('\n')
	entryID, err := strconv.Atoi(strings.TrimSpace(inputID))
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	entry := currentUserDiary.FindEntryByID(entryID)
	if entry == nil {
		fmt.Println("Entry not found.")
		return
	}

	currentUserDiary.DeleteEntry(entryID)
	fmt.Println("Entry deleted successfully!")
}
