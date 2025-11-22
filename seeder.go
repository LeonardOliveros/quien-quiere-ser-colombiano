package main

import (
	"log"
)

func seedQuestions() {
	// Check if questions already exist
	var count int64
	db.Model(&Question{}).Count(&count)
	if count > 0 {
		log.Printf("Database already contains %d questions", count)
		return
	}

	log.Println("Seeding database with initial questions...")

	// Sample questions from the exam - you would add all 450 questions here
	questions := []Question{
		// CULTURA - Símbolos Patrios
		{
			Category:    "CULTURA",
			SubCategory: "Símbolos Patrios",
			Text:        "¿Cuál es el nombre oficial actual del país?",
			Difficulty:  1,
			Points:      10,
			Explanation: "Colombia adoptó definitivamente su nombre actual 'República de Colombia' en 1886, después de haber tenido otros nombres como Nueva Granada, Gran Colombia, Confederación Granadina y Estados Unidos de Colombia.",
			Choices: []Choice{
				{Text: "Nueva Granada", IsCorrect: false, Order: 1},
				{Text: "Gran Colombia", IsCorrect: false, Order: 2},
				{Text: "República de Colombia", IsCorrect: true, Order: 3},
				{Text: "Estados Unidos de Colombia", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "Símbolos Patrios",
			Text:        "¿Quién dio el nombre de 'República de Colombia' el 15 de febrero de 1819?",
			Difficulty:  2,
			Points:      10,
			Hint:        "Fue conocido como El Libertador",
			Explanation: "Simón Bolívar propuso el nombre en el Congreso de Angostura como homenaje a Cristóbal Colón.",
			Choices: []Choice{
				{Text: "Francisco de Paula Santander", IsCorrect: false, Order: 1},
				{Text: "Simón Bolívar", IsCorrect: true, Order: 2},
				{Text: "Antonio Nariño", IsCorrect: false, Order: 3},
				{Text: "Cristóbal Colón", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "Símbolos Patrios",
			Text:        "¿Cuáles son los colores de la bandera de Colombia en orden de arriba hacia abajo?",
			Difficulty:  1,
			Points:      10,
			Explanation: "La bandera tiene tres franjas horizontales: amarillo (ocupando la mitad superior), azul y rojo (cada uno un cuarto).",
			Choices: []Choice{
				{Text: "Rojo, amarillo, azul", IsCorrect: false, Order: 1},
				{Text: "Azul, amarillo, rojo", IsCorrect: false, Order: 2},
				{Text: "Amarillo, azul, rojo", IsCorrect: true, Order: 3},
				{Text: "Amarillo, rojo, azul", IsCorrect: false, Order: 4},
			},
		},
		// GEOGRAFÍA
		{
			Category:    "GEOGRAFIA",
			SubCategory: "Departamentos",
			Text:        "¿Cuál es la capital de Colombia?",
			Difficulty:  1,
			Points:      10,
			Explanation: "Bogotá D.C. es la capital y ciudad más grande de Colombia, ubicada en la cordillera Oriental de los Andes.",
			Choices: []Choice{
				{Text: "Medellín", IsCorrect: false, Order: 1},
				{Text: "Cali", IsCorrect: false, Order: 2},
				{Text: "Bogotá D.C.", IsCorrect: true, Order: 3},
				{Text: "Barranquilla", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "Departamentos",
			Text:        "¿Cuántos departamentos tiene Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Colombia está dividida en 32 departamentos y un distrito capital (Bogotá).",
			Choices: []Choice{
				{Text: "30", IsCorrect: false, Order: 1},
				{Text: "31", IsCorrect: false, Order: 2},
				{Text: "32", IsCorrect: true, Order: 3},
				{Text: "33", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "Regiones",
			Text:        "¿Cuántas regiones naturales tiene Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Colombia tiene 6 regiones naturales: Caribe, Pacífica, Andina, Orinoquía, Amazonía e Insular.",
			Choices: []Choice{
				{Text: "4", IsCorrect: false, Order: 1},
				{Text: "5", IsCorrect: false, Order: 2},
				{Text: "6", IsCorrect: true, Order: 3},
				{Text: "7", IsCorrect: false, Order: 4},
			},
		},
		// HISTORIA
		{
			Category:    "HISTORIA",
			SubCategory: "Independencia",
			Text:        "¿En qué fecha se conmemora el Día de la Independencia de Colombia?",
			Difficulty:  1,
			Points:      10,
			Explanation: "El 20 de julio de 1810 marca el inicio del proceso de independencia con el Grito de Independencia en Bogotá.",
			Choices: []Choice{
				{Text: "7 de agosto", IsCorrect: false, Order: 1},
				{Text: "20 de julio", IsCorrect: true, Order: 2},
				{Text: "12 de octubre", IsCorrect: false, Order: 3},
				{Text: "19 de abril", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "Independencia",
			Text:        "¿Quién fue conocido como 'El Libertador'?",
			Difficulty:  1,
			Points:      10,
			Explanation: "Simón Bolívar es conocido como El Libertador por su papel fundamental en la independencia de varios países sudamericanos.",
			Choices: []Choice{
				{Text: "Antonio Nariño", IsCorrect: false, Order: 1},
				{Text: "Francisco de Paula Santander", IsCorrect: false, Order: 2},
				{Text: "Simón Bolívar", IsCorrect: true, Order: 3},
				{Text: "José María Córdova", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "Independencia",
			Text:        "La Batalla de Boyacá ocurrió en el año:",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Batalla de Boyacá del 7 de agosto de 1819 fue decisiva para la independencia de Nueva Granada.",
			Choices: []Choice{
				{Text: "1810", IsCorrect: false, Order: 1},
				{Text: "1819", IsCorrect: true, Order: 2},
				{Text: "1821", IsCorrect: false, Order: 3},
				{Text: "1830", IsCorrect: false, Order: 4},
			},
		},
		// CONSTITUCIÓN
		{
			Category:    "CONSTITUCION",
			SubCategory: "Principios",
			Text:        "Según la Constitución, Colombia es un Estado:",
			Difficulty:  2,
			Points:      10,
			Explanation: "El artículo 1 de la Constitución define a Colombia como un Estado social de derecho, organizado en forma de República unitaria, descentralizada.",
			Choices: []Choice{
				{Text: "Federal, democrático y participativo", IsCorrect: false, Order: 1},
				{Text: "Unitario, centralizado y democrático", IsCorrect: false, Order: 2},
				{Text: "Social de derecho, unitario, descentralizado", IsCorrect: true, Order: 3},
				{Text: "Monárquico constitucional", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "Estructura del Estado",
			Text:        "¿Cuáles son las tres ramas del poder público en Colombia?",
			Difficulty:  1,
			Points:      10,
			Explanation: "La Constitución establece la separación de poderes en tres ramas: Ejecutiva, Legislativa y Judicial.",
			Choices: []Choice{
				{Text: "Ejecutiva, Legislativa y Judicial", IsCorrect: true, Order: 1},
				{Text: "Presidencial, Senatorial y Judicial", IsCorrect: false, Order: 2},
				{Text: "Nacional, Departamental y Municipal", IsCorrect: false, Order: 3},
				{Text: "Civil, Penal y Administrativa", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "Derechos",
			Text:        "¿Desde qué edad se adquiere la ciudadanía en Colombia?",
			Difficulty:  1,
			Points:      10,
			Explanation: "Según la Constitución, se es ciudadano colombiano al cumplir 18 años de edad.",
			Choices: []Choice{
				{Text: "16 años", IsCorrect: false, Order: 1},
				{Text: "18 años", IsCorrect: true, Order: 2},
				{Text: "21 años", IsCorrect: false, Order: 3},
				{Text: "14 años", IsCorrect: false, Order: 4},
			},
		},
		// Add more questions for variety in each category
		{
			Category:    "CULTURA",
			SubCategory: "Festividades",
			Text:        "¿En qué ciudad se celebra el Carnaval más famoso de Colombia?",
			Difficulty:  1,
			Points:      10,
			Explanation: "El Carnaval de Barranquilla es Patrimonio Cultural Inmaterial de la Humanidad.",
			Choices: []Choice{
				{Text: "Cartagena", IsCorrect: false, Order: 1},
				{Text: "Barranquilla", IsCorrect: true, Order: 2},
				{Text: "Santa Marta", IsCorrect: false, Order: 3},
				{Text: "Bogotá", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "Límites",
			Text:        "Colombia limita al este con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "Al este, Colombia limita con Venezuela y Brasil.",
			Choices: []Choice{
				{Text: "Ecuador y Perú", IsCorrect: false, Order: 1},
				{Text: "Venezuela y Brasil", IsCorrect: true, Order: 2},
				{Text: "Panamá", IsCorrect: false, Order: 3},
				{Text: "El Océano Atlántico", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "Siglo XX",
			Text:        "¿En qué año se promulgó la actual Constitución de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Constitución Política de Colombia fue promulgada el 4 de julio de 1991.",
			Choices: []Choice{
				{Text: "1986", IsCorrect: false, Order: 1},
				{Text: "1989", IsCorrect: false, Order: 2},
				{Text: "1991", IsCorrect: true, Order: 3},
				{Text: "1994", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "Instituciones",
			Text:        "El Congreso de la República está conformado por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "El Congreso es bicameral, compuesto por el Senado y la Cámara de Representantes.",
			Choices: []Choice{
				{Text: "Senado y Asamblea Nacional", IsCorrect: false, Order: 1},
				{Text: "Cámara Alta y Cámara Baja", IsCorrect: false, Order: 2},
				{Text: "Senado y Cámara de Representantes", IsCorrect: true, Order: 3},
				{Text: "Parlamento y Senado", IsCorrect: false, Order: 4},
			},
		},
	}

	// Insert questions with their choices
	for _, q := range questions {
		if err := db.Create(&q).Error; err != nil {
			log.Printf("Error seeding question: %v", err)
		}
	}

	log.Printf("Successfully seeded %d questions", len(questions))
}
