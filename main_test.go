package main

import "testing"

// Funcion para testear la funcion crearPersonaje
func TestCrearPersonaje(t *testing.T) {
	// Creo el personaje
	personajeTest := crearPersonaje("Test", "golpe de espada", 100, 3)

	// Verifico que se cree correctamente
	if personajeTest == nil {
		t.Errorf("El personaje no se creo correctamente.")
	}
}

// Función para testear la función restarVida
func TestRestarVida(t *testing.T) {
	// Creo un personaje de prueba
	personajePrueba := crearPersonaje("Test", "Habilidad", 100, 1)

	// Caso 1: Daño no letal
	ataque := 20
	murio := restarVida(personajePrueba, ataque)
	if murio {
		t.Errorf("Se esperaba que el personaje no muriera, pero se informó que falleció.")
	}
	if personajePrueba.Vida != 80 {
		t.Errorf("La vida del personaje después del ataque debería ser de 80, pero es de: %d.", personajePrueba.Vida)
	}

	// Caso 2: Daño letal
	// Reinicio la vida del personaje para el segundo caso de prueba
	personajePrueba.Vida = 100
	ataque = 100
	murio = restarVida(personajePrueba, ataque)
	if !murio {
		t.Errorf("Se esperaba que el personaje muriera, pero no se informa que murió.")
	}
	if personajePrueba.Vida != 0 {
		t.Errorf("La vida del personaje después del ataque debería ser de 0, pero es de: %d.", personajePrueba.Vida)
	}
}

// Función para testear la función sumarVida
func TestSumarVida(t *testing.T) {
	// Crear un personajePrueba de prueba con vida baja y algunas pósimas
	personajePrueba := crearPersonaje("Test", "Habilidad", 50, 2)

	// Caso 1: Curación sin alcanzar la vida máxima
	cura := 30
	sumarVida(personajePrueba, cura)
	if personajePrueba.Vida != 80 {
		t.Errorf("La vida del personajePrueba después de curarse debería ser 80, pero es: %d.", personajePrueba.Vida)
	}
	if personajePrueba.Posima != 1 {
		t.Errorf("El número de pósimas después de curarse debería ser 1, pero es %d.", personajePrueba.Posima)
	}

	// Caso 2: Curación alcanzando la vida máxima
	cura = 50
	sumarVida(personajePrueba, cura)
	if personajePrueba.Vida != 100 {
		t.Errorf("La vida del personajePrueba después de curarse debería ser 100 (vida máxima), pero es %d.", personajePrueba.Vida)
	}
	if personajePrueba.Posima != 0 {
		t.Errorf("El número de pósimas después de curarse debería ser 0, pero es %d.", personajePrueba.Posima)
	}

	// Caso 3: Intento de curación sin pósimas disponibles
	cura = 10
	sumarVida(personajePrueba, cura)
	if personajePrueba.Vida != 100 {
		t.Errorf("La vida del personajePrueba no debería cambiar si no hay pósimas disponibles, pero es %d.", personajePrueba.Vida)
	}
	if personajePrueba.Posima != 0 {
		t.Errorf("El número de pósimas no debería cambiar si no hay pósimas disponibles, pero es %d.", personajePrueba.Posima)
	}
}
