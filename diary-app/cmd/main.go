// cmd/main.go
package main

import (
    "diary-app/internal/repository"
    "diary-app/internal/service"
    "fmt"
    "log"
)

func main() {
    // Initialize dependencies
    repo := repository.NewDiaryRepository()
    service := service.NewDiaryService(repo)
    
    // Example usage
    created, err := service.CreateDiary(dto.CreateDiaryRequest{
        Username: "alice",
        Password: "password123",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created diary: %+v\n", created)
    
    unlocked, err := service.UnlockDiary(dto.UnlockDiaryRequest{
        Username: "alice",
        Password: "password123",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Unlocked diary: %+v\n", unlocked)
}