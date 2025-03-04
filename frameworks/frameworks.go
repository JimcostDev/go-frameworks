package frameworks

import (
	"fmt"
	"sync"
)

// Ejecuta todos los servidores en goroutines
func RunAll() {
	var wg sync.WaitGroup
	wg.Add(3) // Cantidad de frameworks

	go func() {
		defer wg.Done()
		fmt.Println("Starting Gin...")
		Gin()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Starting Echo...")
		Echo()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Starting Fiber...")
		Fiber()
	}()

	wg.Wait()
}
