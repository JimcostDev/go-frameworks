package frameworks

import (
	"fmt"
	"net/http"
)

func SinFramework() {
	// Configuración del handler para la ruta "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Go!")
	})

	fmt.Println("Servidor corriendo en :3000")
	// Inicia el servidor en el puerto 3000
	http.ListenAndServe(":3000", nil)
}
