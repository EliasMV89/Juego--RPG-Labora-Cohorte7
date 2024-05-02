package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Estructura para representar personajes
type Personaje struct {
	Nombre    string
	Habilidad string
	Vida      int // La cantidad de vida con la que se inicia la partida
	Posima    int // Cantidad de curas disponibles
}

// Función para crear personajes
func crearPersonaje(nombre, habilidad string, vida, posima int) *Personaje {
	nuevoPersonaje := &Personaje{
		Nombre:    nombre,
		Habilidad: habilidad,
		Vida:      vida,
		Posima:    posima,
	}
	return nuevoPersonaje
}

// Función para actualizar la cantidad de vida después de un ataque
func restarVida(personaje *Personaje, ataque int) bool {
	personaje.Vida -= ataque
	if personaje.Vida <= 0 {
		fmt.Printf("'%s' ¡Ha muerto!\n", personaje.Nombre)
		return true // Indica que el personaje ha muerto
	} else {
		fmt.Printf("'%s' Ha recibido un daño de %d y su vida ahora es de: %d.\n ", personaje.Nombre, ataque, personaje.Vida)
		return false
	}
}

// Funcion para actualizar la cantidad vida después de curarse
func sumarVida(personaje *Personaje, cura int) {
	if personaje.Posima > 0 {
		if personaje.Vida < 100 && personaje.Vida+cura <= 100 {
			personaje.Vida += cura
			personaje.Posima -= 1
			fmt.Printf("'%s' Ha recibido una cura de '%d' y su vida ahora es de: %d.\n", personaje.Nombre, cura, personaje.Vida)
		} else {
			personaje.Vida = 100
			personaje.Posima -= 1
			fmt.Printf("'%s' ha recibido una cura de '%d' y su vida ahora es de: %d.\n", personaje.Nombre, cura, personaje.Vida)
		}
	} else {
		fmt.Printf("No tienes posimas disponibles para curarte.\n")
	}
}

// Función para mostrar la información del personaje
func mostrarPersonaje(personaje *Personaje) {
	fmt.Printf("******************************************************\n")
	fmt.Printf(" Nombre: %s\n Habilidad: %s\n Vida: %d\n Cantidad de posimas: %d\n", personaje.Nombre, personaje.Habilidad, personaje.Vida, personaje.Posima)
	fmt.Printf("******************************************************\n")
}

// Estructura para representar mundos
type Mundo struct {
	Nombre    string
	Ubicacion []string
}

// Función para crear mundos
func crearMundo(Nombre string, Ubicacion []string) *Mundo {
	nuevoMundo := &Mundo{
		Nombre:    Nombre,
		Ubicacion: Ubicacion,
	}
	return nuevoMundo
}

// Función para mostrar menú dentro de las ubicaciones
func menuUbicacion(personaje, enemigo *Personaje, ubicacion string) {
	if ubicacion == "" {
		personaje = personaje
		enemigo = enemigo
	}
	for {
		fmt.Printf("Te encuentras en %s, frente a amenazante %s.\n ", ubicacion, enemigo.Nombre)
		fmt.Printf("Debes seleccionar una opción: \n")
		fmt.Println("1. Atacar")
		fmt.Println("2. Curarte")
		fmt.Println("3. Mostrar información del personaje")
		fmt.Println("4. Regresar al menú principal")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		// Procesar la opción seleccionada
		switch choice {
		case 1:
			var ataque = randomizador()
			if restarVida(enemigo, ataque) {
				return // Salir al bucle principal si el personaje ha muerto
			}
			var contraataque = randomizador()
			restarVida(personaje, contraataque)
		case 2:
			var cura = randomizador()
			sumarVida(personaje, cura)
			var contraataque = randomizador()
			restarVida(personaje, contraataque)
		case 3:
			mostrarPersonaje(personaje)
		case 4:
			return
		default:
			fmt.Println("Opcióninvalida")
		}
	}
}

// Función para crear el valor aleatorio para ataque y cura
func randomizador() int {
	valores := []int{10, 20, 30}     // Lista de valores posibles
	rand.Seed(time.Now().UnixNano()) // Inicializa el generador de valores aleatorios
	return valores[rand.Intn(len(valores))]
}

func main() {

	// Creación del personaje principal y los enemigos
	personajePrincipal := crearPersonaje("John Nieve", "golpe de espada", 100, 3)
	enemigo1 := crearPersonaje("Hombre del hierro", "golpe de espada", 110, 0)
	enemigo2 := crearPersonaje("Lobo guargo", "mordida", 110, 0)
	enemigo3 := crearPersonaje("Caminante blanco", "golpe de espada de hielo", 180, 0)

	// Creación del mundo
	mundo1 := crearMundo("Poniente", []string{"Las Islas del Hierro", "El Bosque de Invernalia", "El Norte del Muro"})

	for {
		fmt.Printf("Bienvenido! %s al mundo de %s ❄ \n", personajePrincipal.Nombre, mundo1.Nombre)
		fmt.Printf("******************************************************\n")
		fmt.Printf("Debes elegir una ubicación dentro de este mundo.\n")

		fmt.Printf("******************************************************\n")
		fmt.Println("1. Islas del Hierro")
		fmt.Println("2. Bosque de Invernalia")
		fmt.Println("3. Norte del Muro")
		fmt.Println("4. Para salir de la partida")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			menuUbicacion(personajePrincipal, enemigo1, "Las Islas del Hierro")
		case 2:
			menuUbicacion(personajePrincipal, enemigo2, "El Bosque de Invernalia")
		case 3:
			menuUbicacion(personajePrincipal, enemigo3, "El Norte del Muro")
		case 4:
			fmt.Println("salir")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}
