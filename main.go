package main

import (
	"dictionnaire/dictionary"
	"log"
	"net/http"
	"sync"
	"fmt"
)

func main() {
	dico := dictionary.NewDictionary()

	var wg sync.WaitGroup
	serverDone := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		dico.AddWord("sarah", "Female")
		dico.AddWord("paul", "Male")
		dico.AddWord("romain", "Male")
		
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		dico.RemoveWord("romain")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		dico.GetDefinition("sarah")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := dico.SaveToFile("dictionary.json")
		if err != nil {
			log.Println("Error saving dictionary to file:", err)
			return
		}
		fmt.Println("Dictionary saved successfully!")
	}()

	go func() {
		http.Handle("/", dico)
		fmt.Println("Dictionnaire Server starting...")
		log.Fatal(http.ListenAndServe(":8082", nil))
		fmt.Println("Dictionnaire Server started successfully !")
		close(serverDone) // Fermez le canal pour signaler que le serveur a terminé
	}()


	wg.Wait()

	// Attendre que le serveur ait terminé avant de quitter
	<-serverDone
}

