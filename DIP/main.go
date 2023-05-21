package main

import "fmt"

func main() {
	fmt.Println("\nModul tingkat tinggi tidak boleh bergantung pada modul tingkat rendah. Keduanya seharusnya bergantung pada abstraksi dan abstraksi tidak boleh bergantung pada detail. Detail seharusnya bergantung pada abstraksi")

	repo := NewAbstractionRepo()
	svc := NewAbstractionService(repo)

	CRUD(svc)
}
