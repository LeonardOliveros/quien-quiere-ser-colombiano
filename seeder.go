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

	// Sample questions from the exam - you would add all 450 questions here,
	questions := []Question{
		// CULTURA - Símbolos Patrios,
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es el nombre oficial actual del país?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Nueva Granada", IsCorrect: false, Order: 1},
				{Text: "Gran Colombia", IsCorrect: false, Order: 2},
				{Text: "República de Colombia", IsCorrect: true, Order: 3},
				{Text: "Estados Unidos de Colombia", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿En qué año Colombia adoptó definitivamente su nombre actual de \"República de Colombia\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "1819", IsCorrect: false, Order: 1},
				{Text: "1863", IsCorrect: false, Order: 2},
				{Text: "1886", IsCorrect: true, Order: 3},
				{Text: "1991", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Quién dio el nombre de \"República de Colombia\" el 15 de febrero de 1819?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Francisco de Paula Santander", IsCorrect: false, Order: 1},
				{Text: "Simón Bolívar", IsCorrect: true, Order: 2},
				{Text: "Antonio Nariño", IsCorrect: false, Order: 3},
				{Text: "Cristóbal Colón", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "El nombre de Colombia es un homenaje a:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Simón Bolívar", IsCorrect: false, Order: 1},
				{Text: "Francisco Miranda", IsCorrect: false, Order: 2},
				{Text: "Cristóbal Colón", IsCorrect: true, Order: 3},
				{Text: "Rafael Núñez", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuáles son los colores de la bandera de Colombia en orden de arriba hacia abajo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Rojo, amarillo, azul", IsCorrect: false, Order: 1},
				{Text: "Azul, amarillo, rojo", IsCorrect: false, Order: 2},
				{Text: "Amarillo, azul, rojo", IsCorrect: true, Order: 3},
				{Text: "Amarillo, rojo, azul", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "En la bandera nacional, el color amarillo representa:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "El cielo", IsCorrect: false, Order: 1},
				{Text: "La sangre de los libertadores", IsCorrect: false, Order: 2},
				{Text: "La riqueza del suelo, la armonía y la justicia", IsCorrect: true, Order: 3},
				{Text: "Los océanos", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Qué representa el color azul en la bandera?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "La sangre derramada", IsCorrect: false, Order: 1},
				{Text: "El cielo, los ríos y los dos océanos", IsCorrect: true, Order: 2},
				{Text: "La riqueza del suelo", IsCorrect: false, Order: 3},
				{Text: "La vegetación", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "El color rojo de la bandera simboliza:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Los ríos", IsCorrect: false, Order: 1},
				{Text: "La riqueza", IsCorrect: false, Order: 2},
				{Text: "La sangre derramada por los libertadores", IsCorrect: true, Order: 3},
				{Text: "El cielo", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Quién propuso los colores de la bandera basándose en la teoría de los colores de Goethe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Simón Bolívar", IsCorrect: false, Order: 1},
				{Text: "Francisco Miranda", IsCorrect: true, Order: 2},
				{Text: "Rafael Núñez", IsCorrect: false, Order: 3},
				{Text: "Francisco de Paula Santander", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿En qué fecha se declaró oficialmente la bandera actual mediante la Ley 124?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "15 de febrero de 1819", IsCorrect: false, Order: 1},
				{Text: "13 de julio de 1887", IsCorrect: true, Order: 2},
				{Text: "20 de julio de 1810", IsCorrect: false, Order: 3},
				{Text: "7 de agosto de 1819", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Quién diseñó el escudo de armas de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Simón Bolívar", IsCorrect: false, Order: 1},
				{Text: "Francisco de Paula Santander", IsCorrect: true, Order: 2},
				{Text: "Rafael Núñez", IsCorrect: false, Order: 3},
				{Text: "Antonio Nariño", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿En qué año fue diseñado el escudo nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "1810", IsCorrect: false, Order: 1},
				{Text: "1819", IsCorrect: false, Order: 2},
				{Text: "1834", IsCorrect: true, Order: 3},
				{Text: "1886", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Qué ave aparece en el escudo nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Águila", IsCorrect: false, Order: 1},
				{Text: "Cóndor", IsCorrect: true, Order: 2},
				{Text: "Colibrí", IsCorrect: false, Order: 3},
				{Text: "Guacamaya", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es el lema nacional que aparece en el escudo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "\"Unidad y Libertad\"", IsCorrect: false, Order: 1},
				{Text: "\"Paz y Justicia\"", IsCorrect: false, Order: 2},
				{Text: "\"Libertad y Orden\"", IsCorrect: true, Order: 3},
				{Text: "\"Patria y Honor\"", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "En la franja superior del escudo, ¿qué fruto aparece?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Manzana de oro", IsCorrect: false, Order: 1},
				{Text: "Granada de oro", IsCorrect: true, Order: 2},
				{Text: "Piña de oro", IsCorrect: false, Order: 3},
				{Text: "Naranja de oro", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Qué simboliza el gorro frigio en el escudo nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "La riqueza", IsCorrect: false, Order: 1},
				{Text: "El comercio", IsCorrect: false, Order: 2},
				{Text: "La libertad", IsCorrect: true, Order: 3},
				{Text: "La agricultura", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "Los dos buques en la franja inferior del escudo representan:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "La marina de guerra", IsCorrect: false, Order: 1},
				{Text: "Los dos océanos (Caribe y Pacífico)", IsCorrect: true, Order: 2},
				{Text: "El comercio con España", IsCorrect: false, Order: 3},
				{Text: "La independencia naval", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Quién escribió la letra del himno nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Oreste Sindici", IsCorrect: false, Order: 1},
				{Text: "Rafael Núñez", IsCorrect: true, Order: 2},
				{Text: "José María Samper", IsCorrect: false, Order: 3},
				{Text: "Jorge Isaacs", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Quién compuso la música del himno nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Rafael Núñez", IsCorrect: false, Order: 1},
				{Text: "Francisco de Paula Santander", IsCorrect: false, Order: 2},
				{Text: "Oreste Sindici", IsCorrect: true, Order: 3},
				{Text: "José María Córdova", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿En qué año fue compuesto el himno nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "1819", IsCorrect: false, Order: 1},
				{Text: "1834", IsCorrect: false, Order: 2},
				{Text: "1887", IsCorrect: true, Order: 3},
				{Text: "1910", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es el ave nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "El águila", IsCorrect: false, Order: 1},
				{Text: "El cóndor de los Andes", IsCorrect: true, Order: 2},
				{Text: "El colibrí", IsCorrect: false, Order: 3},
				{Text: "El tucán", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es la flor nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "La rosa", IsCorrect: false, Order: 1},
				{Text: "El girasol", IsCorrect: false, Order: 2},
				{Text: "La orquídea", IsCorrect: true, Order: 3},
				{Text: "El clavel", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es el nombre científico de la flor nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Rosa colombiana", IsCorrect: false, Order: 1},
				{Text: "Cattleya trianae", IsCorrect: true, Order: 2},
				{Text: "Helianthus annus", IsCorrect: false, Order: 3},
				{Text: "Dianthus caryophyllus", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es el árbol nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "El roble", IsCorrect: false, Order: 1},
				{Text: "La ceiba", IsCorrect: false, Order: 2},
				{Text: "La palma de cera del Quindío", IsCorrect: true, Order: 3},
				{Text: "El guayacán", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es considerado el símbolo cultural de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "El café", IsCorrect: false, Order: 1},
				{Text: "La esmeralda", IsCorrect: false, Order: 2},
				{Text: "El sombrero vueltiao", IsCorrect: true, Order: 3},
				{Text: "La mochila wayúu", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Cuál es la moneda oficial de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "El bolívar", IsCorrect: false, Order: 1},
				{Text: "El peso colombiano", IsCorrect: true, Order: 2},
				{Text: "El dólar", IsCorrect: false, Order: 3},
				{Text: "El real", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA",
			Text:        "¿cuál es considerado el plato insigne de Bogotá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Bandeja paisa", IsCorrect: false, Order: 1},
				{Text: "Ajiaco", IsCorrect: true, Order: 2},
				{Text: "Sancocho", IsCorrect: false, Order: 3},
				{Text: "Tamal", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA",
			Text:        "¿Qué plato es típico de la región antioqueña?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Ajiaco", IsCorrect: false, Order: 1},
				{Text: "Mote de queso", IsCorrect: false, Order: 2},
				{Text: "Bandeja paisa", IsCorrect: true, Order: 3},
				{Text: "Lechona", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA",
			Text:        "El sancocho de gallina es un plato típico del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Valle del Cauca", IsCorrect: true, Order: 1},
				{Text: "Antioquia", IsCorrect: false, Order: 2},
				{Text: "Bogotá", IsCorrect: false, Order: 3},
				{Text: "Costa Caribe", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA",
			Text:        "La lechona es un plato tradicional del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Valle del Cauca", IsCorrect: false, Order: 1},
				{Text: "Tolima", IsCorrect: true, Order: 2},
				{Text: "Santander", IsCorrect: false, Order: 3},
				{Text: "Boyacá", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA",
			Text:        "El mote de queso es típico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Antioquia", IsCorrect: false, Order: 1},
				{Text: "La Costa Caribe", IsCorrect: true, Order: 2},
				{Text: "Bogotá", IsCorrect: false, Order: 3},
				{Text: "Los Llanos", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA",
			Text:        "¿Cuál de estos ingredientes NO se menciona como parte del ajiaco?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Pollo", IsCorrect: false, Order: 1},
				{Text: "Tres tipos de papa", IsCorrect: false, Order: 2},
				{Text: "Arroz", IsCorrect: true, Order: 3},
				{Text: "Guascas", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "¿En qué ciudad se celebra el Carnaval más famoso de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Cartagena", IsCorrect: false, Order: 1},
				{Text: "Barranquilla", IsCorrect: true, Order: 2},
				{Text: "Santa Marta", IsCorrect: false, Order: 3},
				{Text: "Bogotá", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "El Carnaval de Negros y Blancos se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Barranquilla", IsCorrect: false, Order: 1},
				{Text: "Pasto", IsCorrect: true, Order: 2},
				{Text: "Popayán", IsCorrect: false, Order: 3},
				{Text: "Cali", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "¿En qué mes se celebra el Carnaval de Barranquilla?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Enero", IsCorrect: false, Order: 1},
				{Text: "Febrero", IsCorrect: true, Order: 2},
				{Text: "Julio", IsCorrect: false, Order: 3},
				{Text: "Diciembre", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "La Feria de las Flores se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Bogotá", IsCorrect: false, Order: 1},
				{Text: "Cali", IsCorrect: false, Order: 2},
				{Text: "Medellín", IsCorrect: true, Order: 3},
				{Text: "Barranquilla", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "¿En qué mes se realiza la Feria de las Flores?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Mayo", IsCorrect: false, Order: 1},
				{Text: "Agosto", IsCorrect: true, Order: 2},
				{Text: "Octubre", IsCorrect: false, Order: 3},
				{Text: "Diciembre", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "El Festival Vallenato se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Barranquilla", IsCorrect: false, Order: 1},
				{Text: "Valledupar", IsCorrect: true, Order: 2},
				{Text: "Cartagena", IsCorrect: false, Order: 3},
				{Text: "Santa Marta", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "El Festival Internacional de Teatro se realiza en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Medellín", IsCorrect: false, Order: 1},
				{Text: "Cali", IsCorrect: false, Order: 2},
				{Text: "Bogotá", IsCorrect: true, Order: 3},
				{Text: "Manizales", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "¿Cada cuántos años se celebra el Festival Iberoamericano de Teatro de Bogotá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{Text: "Cada año", IsCorrect: false, Order: 1},
				{Text: "Cada 2 años", IsCorrect: true, Order: 2},
				{Text: "Cada 3 años", IsCorrect: false, Order: 3},
				{Text: "Cada 4 años", IsCorrect: false, Order: 4},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES Y CELEBRACIONES",
			Text:        "El Festival de la Leyenda Vallenata rinde homenaje a:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Carlos Vives",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Virgen del Rosario",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Francisco el Hombre",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Rafael Escalona",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "¿Quién es conocido por sus esculturas voluminosas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Alejandro Obregón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Fernando Botero",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Enrique Grau",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Luis Caballero",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "¿Dónde se encuentra el Museo del Oro?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "¿Cuál es el teatro más antiguo de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Teatro Colón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Teatro Heredia (Hoy Teatro Colón)",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Teatro Jorge Eliécer Gaitán",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Teatro Pablo Tobón Uribe",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "El Teatro Colón está ubicado en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "¿Quién escribió \"Cien años de soledad\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Jorge Isaacs",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Gabriel García Márquez",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Rafael Pombo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "José Eustasio Rivera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "¿En qué año Gabriel García Márquez ganó el Premio Nobel de Literatura?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1972",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1982",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1992",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "2002",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "¿Quién escribió \"María\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Gabriel García Márquez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Jorge Isaacs",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "José Eustasio Rivera",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tomás Carrasquilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTES Y CULTURA",
			Text:        "\"La Vorágine\" fue escrita por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Jorge Isaacs",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José Eustasio Rivera",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Gabriel García Márquez",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Rafael Pombo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "¿De qué región es originario el vallenato?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Costa Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Llanos Orientales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pacífico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "La cumbia es originaria de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Andes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Costa Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "El bambuco es característico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Costa Caribe",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La región Andina",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "El joropo es el baile típico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La región Andina",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Costa Caribe",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Los Llanos Orientales",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "El currulao es una danza del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Caribe",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pacífico",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Andes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "¿Cuál es el instrumento principal del vallenato?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La guitarra",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El acordeón",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La flauta",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El tambor",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "La gaita es un instrumento típico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Costa Atlántica",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Andes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES",
			Text:        "¿Cuál es el deporte más popular en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Ciclismo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Fútbol",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Boxeo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Atletismo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES",
			Text:        "¿Cuál es considerado el deporte nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Fútbol",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ciclismo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tejo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Béisbol",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES",
			Text:        "¿Quiénes son mencionados como ciclistas destacados en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "James Rodríguez y Falcao",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nairo Quintana y Egan Bernal",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Mariana Pajón y Caterine Ibargüen",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Luis Herrera y Fabio Parra",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN",
			Text:        "¿Cuál es la religión predominante en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Protestantismo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Catolicismo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Islam",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Judaísmo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN",
			Text:        "¿Qué porcentaje de la población se identifica como católica en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "60%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "70%",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "79%",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "90%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Cuál de estas ciudades es mencionada como Patrimonio de la Humanidad por la UNESCO?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Parque Nacional Tayrona está ubicado en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Guajira",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Magdalena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Bolívar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Atlántico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Dónde se encuentra el Santuario de Las Lajas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Boyacá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nariño",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cundinamarca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "Caño Cristales es conocido como:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El río más largo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El río de los cinco colores",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El río más caudaloso",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El río sagrado",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿En qué departamento se encuentra Caño Cristales?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Amazonas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Meta",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Caquetá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Guaviare",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Eje Cafetero incluye los departamentos de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia, Valle y Cauca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Quindío, Caldas y Risaralda",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tolima, Huila y Cundinamarca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Santander, Boyacá y Norte de Santander",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "Villa de Leyva está ubicada en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cundinamarca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Boyacá",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Qué ciudad es conocida como \"La Ciudad Amurallada\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Mompox",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Nevado del Ruiz se encuentra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Sierra Nevada de Santa Marta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Cordillera Central",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Cordillera Oriental",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Cordillera Occidental",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Cuál es el pico más alto de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nevado del Ruiz",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nevado del Huila",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Pico Cristóbal Colón",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Pico Simón Bolívar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "La Ciudad Perdida (Teyuna) se encuentra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Sierra Nevada de Santa Marta",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Los Andes",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Amazonía",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Chocó",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Qué isla colombiana es mencionada como destino turístico en el Caribe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Gorgona",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Malpelo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "San Andrés",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Todas las anteriores",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El desierto de La Tatacoa está en el departamento de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Guajira",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Huila",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cesar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Magdalena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "La Catedral de Sal está ubicada en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nemocón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Zipaquirá",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tunja",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Chía",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Parque Arqueológico de San Agustín está en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nariño",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Huila",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cauca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Putumayo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Cuál es la capital del departamento del Amazonas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Florencia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Mocoa",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Leticia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Mitú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Chocó tiene costas sobre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el Pacífico",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Solo el Atlántico",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Ambos océanos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Ningún océano",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Cuál es el río más importante de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Río Cauca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Río Magdalena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Río Amazonas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Río Orinoco",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Archipiélago del Rosario está cerca de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "San Andrés",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿En qué mes se celebran las principales festividades de Semana Santa en Popayán?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Marzo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Abril (varía según el año)",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Mayo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Junio",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Puente de Boyacá es importante porque:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Es el más largo del país",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Allí se libró la batalla decisiva de la independencia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Es Patrimonio Mundial",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Conecta dos océanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El Museo Nacional está ubicado en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "¿Cuál es la bebida más representativa de Colombia a nivel mundial?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El aguardiente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El café",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El guarapo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La chicha",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO",
			Text:        "El sombrero vueltiao es originario de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Córdoba y Sucre",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Boyacá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuántas regiones naturales tiene Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "4",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "5",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "6",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "7",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuáles son las seis regiones naturales de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Caribe, Pacífica, Andina, Orinoquía, Amazonía, Insular",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Norte, Sur, Este, Oeste, Centro, Islas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Costa, Sierra, Selva, Llanos, Montaña, Mar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Atlántica, Pacífica, Central, Oriental, Sur, Insular",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la región más extensa de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Andina",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Amazonía",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Orinoquía",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Caribe",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La región Andina se caracteriza por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sus playas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sus tres cordilleras",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Su selva tropical",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sus llanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Qué región es conocida como los Llanos Orientales?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Amazonía",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Orinoquía",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Andina",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Caribe",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La región Insular incluye:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo San Andrés",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Solo Providencia",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "San Andrés, Providencia y Santa Catalina",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo las islas del Pacífico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuántos kilómetros de costa tiene Colombia en el Caribe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1.200 km",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1.600 km",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "2.000 km",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "2.500 km",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuántos kilómetros de costa tiene Colombia en el Pacífico?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "900 km",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1.100 km",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1.300 km",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "1.500 km",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuántos departamentos tiene Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "30",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "31",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "32",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "33",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del departamento de Antioquia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Medellín",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pereira",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del departamento del Atlántico es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Valledupar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital de Bolívar?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Montería",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sincelejo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital de Boyacá es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Duitama",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sogamoso",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tunja",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Chiquinquirá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del departamento de Caldas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Armenia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pereira",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Manizales",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Ibagué",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Caquetá es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Leticia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Florencia",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Mocoa",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Villavicencio",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Cauca?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pasto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Popayán",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Neiva",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Cesar es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Valledupar",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Riohacha",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Montería",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del departamento de Córdoba?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sincelejo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Montería",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Valledupar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital de Cundinamarca es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Zipaquirá",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Facatativá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Girardot",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del departamento del Chocó?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Buenaventura",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tumaco",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Quibdó",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Istmina",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Huila es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pitalito",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Garzón",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Neiva",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Plata",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital de La Guajira?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Valledupar",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Riohacha",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Maicao",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Magdalena es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ciénaga",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Meta?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Yopal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arauca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Villavicencio",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Granada",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital de Nariño es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Ipiales",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tumaco",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Pasto",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Túquerres",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital de Norte de Santander?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bucaramanga",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cúcuta",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Pamplona",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ocaña",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Putumayo es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Puerto Asís",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Mocoa",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sibundoy",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Villa Garzón",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Quindío?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pereira",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Manizales",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Armenia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Calarcá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital de Risaralda es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Armenia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Manizales",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Pereira",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Dosquebradas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital de Santander?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cúcuta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Barrancabermeja",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Bucaramanga",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Girón",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital de Sucre es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Montería",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sincelejo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Corozal",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tolú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Tolima?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Neiva",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ibagué",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Espinal",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Honda",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Valle del Cauca es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Buenaventura",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Palmira",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cali",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Tuluá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Arauca?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Yopal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arauca",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Saravena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tame",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Casanare es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Villavicencio",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arauca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Yopal",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Aguazul",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Vichada?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Puerto Inírida",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Puerto Carreño",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Puerto López",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Puerto Gaitán",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Guainía es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Mitú",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Puerto Inírida",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "San José del Guaviare",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Puerto Carreño",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Guaviare?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Mitú",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Calamar",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "San José del Guaviare",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Retorno",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "La capital del Vaupés es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Leticia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Mitú",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Carurú",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Taraira",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Cuál es la capital del Amazonas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Florencia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Puerto Nariño",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Leticia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Chorrera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "San Andrés y Providencia tiene su capital en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Providencia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "San Andrés",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Santa Catalina",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Johnny Cay",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Con cuántos países limita Colombia por tierra?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "3",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "4",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "5",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "6",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "Colombia limita al este con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Ecuador y Perú",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Venezuela y Brasil",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Panamá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Océano Atlántico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "Colombia limita al sur con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Venezuela y Brasil",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ecuador y Perú",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Panamá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Bolivia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIONES NATURALES",
			Text:        "¿Con qué país limita Colombia al noroeste?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Venezuela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Costa Rica",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Panamá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Nicaragua",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODOS Y NOMBRES HISTÓRICOS",
			Text:        "¿Cuál fue el nombre anterior a \"República de Colombia\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Virreinato del Perú",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nueva Granada",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Gran Venezuela",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nueva España",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODOS Y NOMBRES HISTÓRICOS",
			Text:        "La Gran Colombia incluía los territorios actuales de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Colombia y Venezuela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Colombia, Venezuela y Ecuador",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Colombia, Venezuela, Ecuador y Panamá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo Colombia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODOS Y NOMBRES HISTÓRICOS",
			Text:        "¿En qué año se dio el nombre de \"República de Colombia\" en el Congreso de Angostura?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1810",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1819",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1821",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1830",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODOS Y NOMBRES HISTÓRICOS",
			Text:        "¿Qué otros nombres tuvo Colombia antes de 1886?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo Nueva Granada",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nueva Granada, Confederación Granadina, Estados Unidos de Colombia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo Estados Unidos de Colombia",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Virreinato y Nueva Granada",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "¿En qué fecha se celebra la independencia de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "7 de agosto de 1819",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "20 de julio de 1810",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "12 de octubre de 1492",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "15 de febrero de 1819",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "El 20 de julio de 1810 marca:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Batalla de Boyacá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Grito de Independencia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La creación de la Gran Colombia",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La muerte de Bolívar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "¿Qué evento desencadenó el Grito de Independencia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Batalla de Boyacá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Florero de Llorente",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La llegada de Bolívar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La invasión francesa a España",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "La Batalla de Boyacá ocurrió el:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "20 de julio de 1810",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "7 de agosto de 1819",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "15 de febrero de 1819",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "25 de diciembre de 1819",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "¿Quién comandó las tropas patriotas en la Batalla de Boyacá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Simón Bolívar",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "José María Córdova",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "El período conocido como \"La Patria Boba\" se caracterizó por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La unidad total",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Guerras civiles entre centralistas y federalistas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La paz absoluta",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El dominio español",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "¿Entre qué años ocurrió la Patria Boba?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1810-1815",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "1819-1830",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1830-1840",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1850-1860",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "La Reconquista española fue liderada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pablo Morillo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Francisco de Miranda",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "José María Barreiro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "¿Quién es conocido como \"El Libertador\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Simón Bolívar",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "José de San Martín",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "¿Quién es conocido como \"El Hombre de las Leyes\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Camilo Torres",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "Antonio Nariño es conocido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Ganar la Batalla de Boyacá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Traducir los Derechos del Hombre y del Ciudadano",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ser el primer presidente",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Componer el himno",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "Policarpa Salavarrieta fue:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Primera dama",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Heroína y mártir de la independencia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Escritora",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pintora",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA INDEPENDENCIA",
			Text:        "¿En qué año fue fusilada La Pola?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1810",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1817",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1819",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1821",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "¿En qué año se disolvió la Gran Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1825",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text: "1830",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1835",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1840",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "¿Quién fue el primer presidente de la Nueva Granada después de la disolución de la Gran Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Joaquín Mosquera",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "José Ignacio de Márquez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "La Constitución de Rionegro se promulgó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1853",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1858",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1863",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1886",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "La Constitución de 1863 establecía un sistema:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Centralista",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Federalista",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Monárquico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Socialista",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "¿Quién promovió la Constitución de 1886?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tomás Cipriano de Mosquera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rafael Núñez",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Miguel Antonio Caro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "José María Melo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "La Constitución de 1886 estableció un sistema:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Federalista",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Centralista",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Parlamentario",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Comunista",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "La Guerra de los Mil Días ocurrió entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1885-1888",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1899-1902",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1903-1906",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1910-1913",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX",
			Text:        "¿En qué año se separó Panamá de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1899",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1903",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1910",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1915",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "¿Quién fue Jorge Eliécer Gaitán?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Presidente de Colombia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Líder liberal asesinado en 1948",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "General conservador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Escritor",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "El Bogotazo ocurrió el:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "9 de abril de 1948",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "20 de julio de 1950",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "7 de agosto de 1946",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "13 de junio de 1953",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "El período de La Violencia se sitúa aproximadamente entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1930-1940",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1946-1958",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1960-1970",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1970-1980",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "¿Quién fue el único dictador militar del siglo XX en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Laureano Gómez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Gustavo Rojas Pinilla",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Alberto Lleras",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Mariano Ospina",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "El Frente Nacional fue un acuerdo entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Liberales y comunistas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Liberales y conservadores",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Todos los partidos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Colombia y Venezuela",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "¿Entre qué años funcionó el Frente Nacional?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1948-1958",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1958-1974",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1974-1990",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1990-2000",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "La actual Constitución de Colombia fue promulgada en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1986",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1989",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1991",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1994",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "¿Quién fue César Gaviria?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Presidente durante la Constituyente de 1991",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Magistrado de la Corte",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Líder guerrillero",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "General del ejército",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "El proceso de paz con las FARC se firmó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "2012",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "2014",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "2016",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "2018",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "¿Quién fue presidente durante la firma del acuerdo de paz con las FARC?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Álvaro Uribe",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Juan Manuel Santos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Iván Duque",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Gustavo Petro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "PRINCIPIOS CONSTITUCIONALES",
			Text:        "Según la Constitución, Colombia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un Estado federal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un Estado social de derecho",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una monarquía constitucional",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una república parlamentaria",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "PRINCIPIOS CONSTITUCIONALES",
			Text:        "La soberanía reside en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El pueblo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Corte Constitucional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "PRINCIPIOS CONSTITUCIONALES",
			Text:        "Colombia se organiza en forma de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "República unitaria y centralizada",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "República unitaria, descentralizada, con autonomía de sus entidades territoriales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Estado federal",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Confederación de estados",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "PRINCIPIOS CONSTITUCIONALES",
			Text:        "Son fines esenciales del Estado:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Servir a la comunidad",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Promover la prosperidad general",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Garantizar la efectividad de los derechos y deberes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Todas las anteriores",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "DERECHOS FUNDAMENTALES",
			Text:        "El derecho a la vida es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Renunciable",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Inviolable",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Condicionado",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Temporal",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "DERECHOS FUNDAMENTALES",
			Text:        "¿Cuál de estos es un derecho fundamental?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Derecho a la propiedad",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Derecho al libre desarrollo de la personalidad",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Derecho a la vivienda",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Derecho al trabajo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "DERECHOS FUNDAMENTALES",
			Text:        "La acción de tutela protege:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Todos los derechos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Solo derechos económicos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Derechos fundamentales",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo el derecho a la vida",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "DERECHOS FUNDAMENTALES",
			Text:        "La tutela debe ser resuelta en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "10 días",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "30 días",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "60 días",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "90 días",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "DERECHOS FUNDAMENTALES",
			Text:        "El habeas corpus protege:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La propiedad",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La libertad personal",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El trabajo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La educación",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "DERECHOS FUNDAMENTALES",
			Text:        "¿Desde qué edad se es ciudadano colombiano?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "14 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "16 años",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "18 años",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "21 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "¿Cuáles son las tres ramas del poder público?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nacional, Departamental, Municipal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ejecutiva, Legislativa, Judicial",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Política, Económica, Social",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Civil, Penal, Administrativa",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "La Rama Ejecutiva está encabezada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Corte Suprema",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Presidente de la República",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Fiscal General",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "¿Quién es el Jefe de Estado y Jefe de Gobierno?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Vicepresidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Presidente del Congreso",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Presidente de la República",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Primer Ministro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "El período presidencial es de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "3 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "4 años",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "5 años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "6 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "¿Se permite la reelección presidencial inmediata?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sí, una vez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sí, indefinidamente",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "No",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo en caso de guerra",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "El Congreso de la República está compuesto por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el Senado",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Solo la Cámara de Representantes",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Senado y Cámara de Representantes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Senado, Cámara y Consejo de Estado",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "¿Cuántos senadores componen el Senado?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "100",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "102",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "108",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "110",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "Los senadores son elegidos por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Circunscripción departamental",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Circunscripción nacional",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los gobernadores",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "El período de senadores y representantes es de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "2 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "4 años",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "6 años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "8 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMAS DEL PODER PÚBLICO",
			Text:        "La función principal del Congreso es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Juzgar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Hacer las leyes",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ejecutar las leyes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nombrar ministros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMA JUDICIAL",
			Text:        "La Corte Constitucional está compuesta por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "7 magistrados",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "9 magistrados",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "11 magistrados",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "13 magistrados",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMA JUDICIAL",
			Text:        "El período de los magistrados de la Corte Constitucional es de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "4 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "6 años",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "8 años",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Vitalicio",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMA JUDICIAL",
			Text:        "La Corte Suprema de Justicia tiene:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "20 magistrados",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "23 magistrados",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "25 magistrados",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "30 magistrados",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMA JUDICIAL",
			Text:        "El Consejo de Estado es el máximo tribunal de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Lo penal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Lo constitucional",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Lo contencioso administrativo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Lo civil",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMA JUDICIAL",
			Text:        "¿Quién dirige la Fiscalía General de la Nación?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Vicefiscal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Fiscal General",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Procurador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Defensor del Pueblo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "RAMA JUDICIAL",
			Text:        "El Fiscal General es elegido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Corte Suprema de Justicia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Voto popular",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANISMOS DE CONTROL",
			Text:        "El Ministerio Público está conformado por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo la Procuraduría",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Procuraduría y Defensoría del Pueblo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Procuraduría, Defensoría y Personerías",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo la Defensoría",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANISMOS DE CONTROL",
			Text:        "El Procurador General de la Nación es elegido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Senado",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Cámara de Representantes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Corte Constitucional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANISMOS DE CONTROL",
			Text:        "La Contraloría General de la República se encarga de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Hacer las leyes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El control fiscal",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Juzgar delitos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Defender los derechos humanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANISMOS DE CONTROL",
			Text:        "El Contralor General es elegido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Congreso en pleno",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Corte de Cuentas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Voto popular",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANISMOS DE CONTROL",
			Text:        "El Defensor del Pueblo es elegido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Cámara de Representantes",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Senado",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El pueblo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANIZACIÓN ELECTORAL",
			Text:        "La organización electoral está conformada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo la Registraduría",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Consejo Nacional Electoral y Registraduría",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo el Consejo Electoral",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tribunales electorales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANIZACIÓN ELECTORAL",
			Text:        "El Registrador Nacional del Estado Civil es nombrado por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Los presidentes de las altas cortes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Consejo Nacional Electoral",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ORGANIZACIÓN ELECTORAL",
			Text:        "El voto en Colombia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Obligatorio",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un derecho y un deber ciudadano",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo para hombres",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo para mayores de 21 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ENTIDADES TERRITORIALES",
			Text:        "Son entidades territoriales:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo los departamentos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Departamentos, municipios, distritos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Departamentos, municipios, distritos y territorios indígenas",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo municipios",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ENTIDADES TERRITORIALES",
			Text:        "Los gobernadores son elegidos por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Voto popular",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Las asambleas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ENTIDADES TERRITORIALES",
			Text:        "El período de los gobernadores es de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "2 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "3 años",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "4 años",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "5 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ENTIDADES TERRITORIALES",
			Text:        "Los alcaldes son elegidos por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los concejos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los gobernadores",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Voto popular",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ENTIDADES TERRITORIALES",
			Text:        "El período de los alcaldes es de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "2 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "3 años",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "4 años",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "5 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ENTIDADES TERRITORIALES",
			Text:        "Las Asambleas Departamentales son:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Órganos de control",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Corporaciones administrativas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tribunales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ministerios",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ENTIDADES TERRITORIALES",
			Text:        "Los Concejos Municipales:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nombran al alcalde",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Aprueban los acuerdos municipales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Eligen al gobernador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nombran jueces",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "MECANISMOS DE PARTICIPACIÓN",
			Text:        "Son mecanismos de participación ciudadana:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el voto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Voto, referendo, plebiscito, consulta popular",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo el plebiscito",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo la tutela",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "MECANISMOS DE PARTICIPACIÓN",
			Text:        "El referendo sirve para:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Elegir presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Aprobar o derogar leyes",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Nombrar ministros",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Elegir jueces",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "MECANISMOS DE PARTICIPACIÓN",
			Text:        "La revocatoria del mandato procede para:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Senadores",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Gobernadores y alcaldes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Magistrados",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "MECANISMOS DE PARTICIPACIÓN",
			Text:        "La iniciativa popular legislativa requiere el respaldo de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "5% de los ciudadanos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "0.5% del censo electoral",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "10% del censo electoral",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "30% del censo electoral",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ÓRGANOS AUTÓNOMOS",
			Text:        "El Banco de la República es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un ministerio",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Una entidad privada",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El banco central con autonomía",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Un departamento administrativo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ÓRGANOS AUTÓNOMOS",
			Text:        "La función principal del Banco de la República es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Prestar dinero al gobierno",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Mantener el poder adquisitivo de la moneda",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Financiar empresas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cobrar impuestos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ÓRGANOS AUTÓNOMOS",
			Text:        "Las Corporaciones Autónomas Regionales se encargan de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La educación",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El medio ambiente",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La salud",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La seguridad",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ÓRGANOS AUTÓNOMOS",
			Text:        "La Comisión Nacional del Servicio Civil protege:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los derechos laborales privados",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El sistema de mérito en el empleo público",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Las pensiones",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La educación",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "REFORMAS CONSTITUCIONALES",
			Text:        "La Constitución puede ser reformada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Congreso mediante acto legislativo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Solo la Corte Constitucional",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Los gobernadores",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "REFORMAS CONSTITUCIONALES",
			Text:        "Un acto legislativo requiere:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "4 debates",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "6 debates",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "8 debates",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "10 debates",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "REFORMAS CONSTITUCIONALES",
			Text:        "También se puede reformar la Constitución por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Decreto presidencial",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Referendo o Asamblea Constituyente",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Orden judicial",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Acuerdo de gobernadores",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "CONTROL CONSTITUCIONAL",
			Text:        "La acción pública de inconstitucionalidad puede ser interpuesta por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo abogados",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Solo el Procurador",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cualquier ciudadano",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo el Presidente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "CONTROL CONSTITUCIONAL",
			Text:        "Las sentencias de la Corte Constitucional:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pueden ser apeladas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Son definitivas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Son solo consultivas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Requieren aprobación del Congreso",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ESTADOS DE EXCEPCIÓN",
			Text:        "Los estados de excepción son:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo guerra exterior",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Guerra exterior, conmoción interior, emergencia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo emergencia económica",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo catástrofe natural",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ESTADOS DE EXCEPCIÓN",
			Text:        "El estado de guerra exterior lo declara:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Presidente con firma de todos los ministros",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Senado mediante declaración",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ESTADOS DE EXCEPCIÓN",
			Text:        "El estado de conmoción interior puede durar inicialmente:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "30 días",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "90 días",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "180 días",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un año",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ESTADOS DE EXCEPCIÓN",
			Text:        "El estado de emergencia puede declararse por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "15 días",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "30 días",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "60 días",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "90 días",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ESTADOS DE EXCEPCIÓN",
			Text:        "Durante los estados de excepción:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Se suspende la Constitución",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "No pueden suspenderse los derechos humanos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Se disuelve el Congreso",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Se suspenden todas las garantías",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "ESTADOS DE EXCEPCIÓN",
			Text:        "Los decretos de estados de excepción tienen:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Vigencia permanente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Fuerza de ley mientras dure el estado",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo son recomendaciones",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "No requieren control",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "¿En qué fechas específicas se celebra el Carnaval de Barranquilla?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Diciembre 24-31",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "40 días antes de la Semana Santa",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Julio 15-20",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Agosto 1-7",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "¿Qué se celebra en la Feria de las Flores además del desfile de silleteros?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el desfile",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Desfile de autos antiguos, cabalgata, fondas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo música",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo gastronomía",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "El Festival Iberoamericano de Teatro de Bogotá se celebra:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cada año",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cada dos años en años pares",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cada tres años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cada cuatro años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "¿En qué fechas se celebra el Carnaval de Negros y Blancos en Pasto?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Diciembre 28 a enero 6",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Febrero 1-10",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Marzo 15-20",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Abril 1-7",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "El Festival del Mono Núñez se realiza en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Valledupar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ginebra, Valle",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Manizales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "¿Qué tipo de música se celebra en el Festival del Mono Núñez?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Vallenato",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Música andina colombiana",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Salsa",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Reggaetón",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "La Feria de Manizales se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Enero",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Junio",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Agosto",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Diciembre",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "¿Qué festividad importante se celebra en Popayán durante la Semana Santa?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Carnaval",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Procesiones religiosas tradicionales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Festival de música",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Feria ganadera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "El Festival Internacional de Cine de Cartagena (FICCI) se fundó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1960",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text: "1970",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1980",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1990",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVIDADES Y FERIAS REGIONALES",
			Text:        "La Feria de Cali se celebra del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1 al 10 de enero",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "25 de diciembre al 3 de enero",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "15 al 25 de julio",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1 al 10 de agosto",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "El plato \"mote de queso\" incluye como ingrediente principal:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Papa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ñame",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Yuca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Plátano",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "¿Qué tipo de pescado se usa tradicionalmente en el arroz con coco caribeño?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Trucha",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pargo rojo, lebranche o bocachico",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Salmón",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Atún",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "El \"Rondón\" de San Andrés incluye un ingrediente llamado:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Carne de res",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pigtail (colita de cerdo)",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Pollo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pavo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "Los \"domplines\" del Rondón son:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pescados pequeños",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tortillas de maíz",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Vegetales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Mariscos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "¿De qué región es típico el \"sancocho de rabo\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Andina",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Pacífico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Orinoquía",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "La \"lechona tolimense\" se rellena con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo arroz",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arroz, arveja, carne de cerdo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo verduras",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pescado",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "El \"puchero santafereño\" es originario de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "¿Qué hierba característica lleva el ajiaco?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cilantro",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Guascas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Perejil",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Albahaca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "Los \"aborrajados\" del Valle del Cauca son:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sopas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Plátanos maduros rellenos de queso y apanados",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Carnes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ensaladas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL DETALLADA",
			Text:        "El \"tilote\" o \"titoté\" es parte de la preparación de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bandeja paisa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arroz con coco",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sancocho",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ajiaco",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ETNIAS Y DEMOGRAFÍA",
			Text:        "¿Cuáles son algunos de los grupos indígenas de la Amazonía colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Wayúu y Arhuacos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tikuna, Uitoto, Cubeo, Desano, Tucano",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Muiscas y Taironas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Quimbayas y Calimas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ETNIAS Y DEMOGRAFÍA",
			Text:        "Los grupos Emberá, Wounaan y Awá habitan principalmente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Amazonía",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La región Pacífica",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Sierra Nevada",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ETNIAS Y DEMOGRAFÍA",
			Text:        "¿Qué porcentaje de la población colombiana se identifica como católica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "60%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "70%",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "79%",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "90%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ETNIAS Y DEMOGRAFÍA",
			Text:        "La población afrocolombiana se concentra principalmente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el Pacífico",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Costa Pacífica y Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo el interior",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo San Andrés",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ETNIAS Y DEMOGRAFÍA",
			Text:        "Los Tikunas son una etnia característica de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Sierra Nevada",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Amazonía",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Chocó",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "Débora Arango fue reconocida principalmente como:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Escultora",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Acuarelista",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ceramista",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Fotógrafa",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "¿Quién fundó el Teatro Experimental de Cali (TEC)?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Santiago García",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Enrique Buenaventura",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Fanny Mikey",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Carlos José Reyes",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "El TEC fue fundado en el año:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1945",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1955",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1965",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1975",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "¿Qué artista colombiano murió en Mónaco en 2023?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Alejandro Obregón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Fernando Botero",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Omar Rayo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Enrique Grau",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "La \"Paloma de la Paz\" de Botero fue entregada al presidente:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Álvaro Uribe",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Juan Manuel Santos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Iván Duque",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "César Gaviria",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "¿En qué año se fundó la primera Escuela de Bellas Artes en Bogotá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1810",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1834",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1886",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1910",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "Eduardo Ramírez Villamizar es conocido por su estilo:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Barroco",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Geométrico",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Impresionista",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Realista",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "¿Quién fue Gregorio Vásquez de Arce y Ceballos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un escultor colonial",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un general",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un poeta",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "El muralismo colombiano fue influenciado por artistas como:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pablo Picasso",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Diego Rivera",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Andy Warhol",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Salvador Dalí",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARTE Y ARTISTAS ESPECÍFICOS",
			Text:        "Pedro Nel Gómez fue representante del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Arte abstracto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Muralismo y realismo socialista",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Impresionismo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Arte conceptual",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA REGIONAL",
			Text:        "El instrumento principal del currulao es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El acordeón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La marimba de chonta",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La guitarra",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El tambor",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA REGIONAL",
			Text:        "La gaita es un instrumento característico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Costa Atlántica",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Andes",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA REGIONAL",
			Text:        "¿De qué material está hecha la marimba de chonta?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Metal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Madera de palma",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Bambú",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Plástico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA REGIONAL",
			Text:        "El mapalé es una danza originaria de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Andes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Costa Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Amazonas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA REGIONAL",
			Text:        "¿Qué instrumento se usa en el joropo llanero?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Marimba",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arpa, cuatro y maracas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Acordeón",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Gaita",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "El Santuario de Las Lajas está ubicado específicamente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El centro de Pasto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ipiales, frontera con Ecuador",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Túquerres",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Cocha",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "¿En qué meses Caño Cristales muestra sus colores más vivos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Diciembre-enero",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Julio-noviembre",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Marzo-abril",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Todo el año",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "¿Cuántos colores se dice que tiene Caño Cristales?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tres",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cinco",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Siete",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Diez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "La Ciudad Perdida (Teyuna) fue construida por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Muiscas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Tayronas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Quimbayas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Incas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "¿En qué año aproximadamente se construyó Ciudad Perdida?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "500 d.C.",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "800 d.C.",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1200 d.C.",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1500 d.C.",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "El Parque Arqueológico de San Agustín es famoso por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pinturas rupestres",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Estatuas megalíticas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Pirámides",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Acueductos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "¿Cuántas hectáreas aproximadamente tiene el Parque Tayrona?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "5.000",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "15.000",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "25.000",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "35.000",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "El desierto de La Tatacoa es conocido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sus dunas de arena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sus formaciones de tierra roja y gris",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sus oasis",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sus cactus gigantes",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "¿Qué profundidad tiene la Catedral de Sal de Zipaquirá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "50 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "100 metros",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "180 metros",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "250 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "TURISMO Y SITIOS ESPECÍFICOS",
			Text:        "El Puente de Boyacá es importante porque allí:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Se firmó la independencia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Se libró la batalla del 7 de agosto de 1819",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Se fundó Bogotá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Se construyó el primer puente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA Y BILLETES",
			Text:        "¿Qué animal aparece en la moneda de mil pesos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Oso de anteojos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tortuga caguama",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cóndor",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Jaguar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA Y BILLETES",
			Text:        "La moneda de quinientos pesos tiene representada:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Una guacamaya",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Una rana de cristal",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un frailejón",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un oso",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA Y BILLETES",
			Text:        "¿Qué planta aparece en la moneda de cien pesos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Orquídea",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Palma de cera",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Frailejón",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Ceiba",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA Y BILLETES",
			Text:        "El oso de anteojos aparece en la moneda de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "20 pesos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "50 pesos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "100 pesos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "200 pesos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA Y BILLETES",
			Text:        "La guacamaya bandera está en la moneda de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "100 pesos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "200 pesos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "500 pesos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1000 pesos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "LITERATURA Y ESCRITORES",
			Text:        "¿En qué año murió Gabriel García Márquez?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "2010",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "2012",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "2014",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "2016",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "LITERATURA Y ESCRITORES",
			Text:        "\"La María\" de Jorge Isaacs es considerada una novela:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Realista",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Romántica",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "De ciencia ficción",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Policíaca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "LITERATURA Y ESCRITORES",
			Text:        "Rafael Pombo es conocido principalmente por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Novelas históricas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cuentos infantiles y fábulas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ensayos políticos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Teatro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "LITERATURA Y ESCRITORES",
			Text:        "Tomás Carrasquilla escribió sobre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La vida urbana",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Costumbres antioqueñas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Historia militar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ciencia ficción",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "LITERATURA Y ESCRITORES",
			Text:        "José Asunción Silva fue principalmente:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Novelista",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Poeta modernista",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Dramaturgo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cronista",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES ESPECÍFICOS",
			Text:        "¿Cómo se juega el tejo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Con pelotas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Lanzando discos metálicos a mechas con pólvora",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Con raquetas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Con palos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES ESPECÍFICOS",
			Text:        "Mariana Pajón es campeona en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Atletismo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "BMX",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Natación",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Gimnasia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES ESPECÍFICOS",
			Text:        "Caterine Ibargüen ganó medalla olímpica en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Maratón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Triple salto",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Lanzamiento de jabalina",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Salto con garrocha",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES ESPECÍFICOS",
			Text:        "¿Quién fue René Higuita?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Ciclista",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Portero de fútbol famoso por \"el escorpión\"",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Boxeador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Atleta",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTES ESPECÍFICOS",
			Text:        "Nairo Quintana ha ganado:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tour de Francia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Giro de Italia y Vuelta a España",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo carreras nacionales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Campeonatos de pista",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARQUITECTURA Y TEATROS",
			Text:        "El Teatro Heredia está ubicado en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARQUITECTURA Y TEATROS",
			Text:        "¿En qué año se fundó el Teatro Colón de Bogotá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1792",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1825",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1892",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1920",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARQUITECTURA Y TEATROS",
			Text:        "El Teatro Pablo Tobón Uribe está en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Medellín",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARQUITECTURA Y TEATROS",
			Text:        "¿Qué característica arquitectónica tiene el Teatro Colón?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Estilo moderno",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Estilo neoclásico italiano",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Estilo colonial",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Estilo art déco",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "ARQUITECTURA Y TEATROS",
			Text:        "El Teatro Jorge Eliécer Gaitán antes se llamaba:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Teatro Nacional",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Teatro Municipal",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Teatro Colombia",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Teatro Popular",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN Y TRADICIONES",
			Text:        "¿Qué porcentaje de protestantes hay en Colombia aproximadamente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "5%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "10%",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "13%",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "20%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN Y TRADICIONES",
			Text:        "Las procesiones de Semana Santa de Popayán datan del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Siglo XVI",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Siglo XVII",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Siglo XVIII",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Siglo XIX",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN Y TRADICIONES",
			Text:        "El Santuario de Monserrate en Bogotá está a una altura de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "2.640 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "3.152 metros",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "3.500 metros",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "4.000 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN Y TRADICIONES",
			Text:        "La Catedral Primada de Bogotá se construyó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1538",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1650",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1807-1823",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1900",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN Y TRADICIONES",
			Text:        "El Divino Niño es una devoción popular centrada en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "El piso térmico cálido en Colombia va desde:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "0 a 500 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "0 a 1.000 metros",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "500 a 1.500 metros",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1.000 a 2.000 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "¿Qué cultivos son típicos del piso térmico cálido?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Papa y trigo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Banano, cacao, caña de azúcar",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Café y flores",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cebada y quinua",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "El piso térmico templado está entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "500 y 1.000 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1.000 y 2.000 metros",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "2.000 y 3.000 metros",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "3.000 y 4.000 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "¿En qué piso térmico se cultiva principalmente el café?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cálido",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Templado",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Frío",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Páramo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "El piso térmico frío va desde:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1.000 a 2.000 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "2.000 a 3.000 metros",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "3.000 a 4.000 metros",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "4.000 a 5.000 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "¿Qué se cultiva en el piso térmico frío?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Café y plátano",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Papa, cebolla, trigo, cebada",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Caña y cacao",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Palma y arroz",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "El páramo comienza aproximadamente a:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "2.000 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "2.500 metros",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "3.000 metros",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "3.500 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PISOS TÉRMICOS Y CULTIVOS",
			Text:        "¿Qué planta es característica del páramo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Café",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Frailejón",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Palma",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Bambú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS",
			Text:        "¿Cuántos kilómetros de costa tiene la región Caribe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1.200 km",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1.600 km",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "2.000 km",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "2.400 km",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS",
			Text:        "La Ciénaga Grande de Santa Marta es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un lago",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un humedal costero",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un río",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una bahía",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS",
			Text:        "¿Qué península se encuentra en La Guajira?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Península de Barú",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Península de La Guajira",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Península de Urabá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "No hay península",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS",
			Text:        "El Golfo de Urabá está entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Guajira y Magdalena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Antioquia y Chocó",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Bolívar y Atlántico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Córdoba y Sucre",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS",
			Text:        "La Sierra Nevada de Santa Marta es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Parte de los Andes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La formación montañosa costera más alta del mundo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un volcán",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una meseta",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS",
			Text:        "¿Qué altura alcanza el Pico Cristóbal Colón?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "4.500 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "5.200 metros",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "5.775 metros",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "6.000 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS",
			Text:        "Los Montes de María están ubicados entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Guajira y Cesar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Bolívar y Sucre",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Córdoba y Antioquia",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Magdalena y Atlántico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN PACÍFICA - CARACTERÍSTICAS",
			Text:        "¿Cuántos kilómetros de costa tiene el Pacífico colombiano?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "900 km",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1.100 km",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1.300 km",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "1.500 km",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN PACÍFICA - CARACTERÍSTICAS",
			Text:        "¿Cuál es el principal puerto del Pacífico?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tumaco",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Buenaventura",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Bahía Solano",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nuquí",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN PACÍFICA - CARACTERÍSTICAS",
			Text:        "El río Atrato desemboca en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Golfo de Urabá (Caribe)",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Magdalena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Amazonas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN PACÍFICA - CARACTERÍSTICAS",
			Text:        "¿Qué característica tiene la selva del Chocó biogeográfico?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Es seca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Es una de las más lluviosas del mundo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Es templada",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Es desértica",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN PACÍFICA - CARACTERÍSTICAS",
			Text:        "Los manglares del Pacífico son importantes porque:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Producen madera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Son criaderos de especies marinas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tienen oro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Producen sal",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "¿Cuántas cordilleras atraviesan la región Andina?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Una",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Dos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tres",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cuatro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "El Valle del Magdalena está entre las cordilleras:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Oriental y Occidental",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Central y Oriental",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Central y Occidental",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "No está entre cordilleras",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "El Valle del Cauca está entre las cordilleras:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Oriental y Central",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Central y Occidental",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo en la Occidental",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo en la Central",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "¿Dónde está ubicado el Valle de Aburrá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "En Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "En Medellín",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "En Cali",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "En Bucaramanga",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "El Macizo Colombiano es importante porque:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Es el pico más alto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nacen los ríos Magdalena, Cauca, Caquetá y Patía",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tiene petróleo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Es un desierto",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "¿En qué departamento está el Cañón del Chicamocha?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Boyacá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santander",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Norte de Santander",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cundinamarca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "El Nevado del Cocuy está en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Boyacá",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tolima",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Caldas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ANDINA - DETALLES",
			Text:        "¿Cuál es el páramo más grande del mundo ubicado en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Páramo de Santurbán",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Páramo de Sumapaz",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Páramo de Chingaza",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Páramo del Cocuy",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ORINOQUÍA - CARACTERÍSTICAS",
			Text:        "La Orinoquía representa aproximadamente qué porcentaje del territorio nacional:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "15%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "20%",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "28%",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "35%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ORINOQUÍA - CARACTERÍSTICAS",
			Text:        "¿Cómo se conoce también a la Orinoquía?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La selva",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Llanos Orientales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La sabana",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La pradera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ORINOQUÍA - CARACTERÍSTICAS",
			Text:        "El río Meta es tributario del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Magdalena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Amazonas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Orinoco",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Caquetá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ORINOQUÍA - CARACTERÍSTICAS",
			Text:        "La Serranía de la Macarena es famosa por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sus minas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Caño Cristales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sus nevados",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sus desiertos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ORINOQUÍA - CARACTERÍSTICAS",
			Text:        "¿Qué actividad económica predomina en los Llanos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Minería",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ganadería extensiva",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Industria",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pesca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN ORINOQUÍA - CARACTERÍSTICAS",
			Text:        "Los \"morichales\" de los Llanos son:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Montañas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Palmeras de moriche en zonas húmedas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ríos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Lagunas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN AMAZÓNICA - DETALLES",
			Text:        "¿Qué extensión aproximada tiene la Amazonía colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "200.000 km²",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "300.000 km²",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "403.000 km²",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "500.000 km²",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN AMAZÓNICA - DETALLES",
			Text:        "¿Cuántas especies de plantas se estiman en la Amazonía colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "10.000",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "20.000",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "30.000",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "40.000",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN AMAZÓNICA - DETALLES",
			Text:        "¿Qué porcentaje de especies vegetales amazónicas son endémicas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "30%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "40%",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "50%",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "60%",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN AMAZÓNICA - DETALLES",
			Text:        "La Victoria amazónica es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un árbol",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un nenúfar gigante",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una orquídea",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una palma",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN AMAZÓNICA - DETALLES",
			Text:        "¿Cuál es la capital del Guaviare?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Mitú",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "San José del Guaviare",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Puerto Inírida",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Leticia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN AMAZÓNICA - DETALLES",
			Text:        "El río Putumayo sirve de frontera con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Brasil",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Venezuela",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Ecuador y Perú",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Panamá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN INSULAR",
			Text:        "¿A qué distancia está San Andrés de la costa colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "200 km",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "400 km",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "700 km",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "1.000 km",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN INSULAR",
			Text:        "¿Qué idiomas se hablan en San Andrés además del español?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Francés",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Inglés y creole",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Portugués",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Holandés",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN INSULAR",
			Text:        "La isla de Providencia es conocida por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Su industria",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Su barrera de coral",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sus minas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Su agricultura",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN INSULAR",
			Text:        "Malpelo es importante porque:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tiene población",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Es santuario de fauna marina",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tiene petróleo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Es turística",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "REGIÓN INSULAR",
			Text:        "Gorgona fue utilizada como:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Base militar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Prisión de alta seguridad",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Hospital",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Universidad",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "¿Cuántos kilómetros tiene el río Magdalena?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1.200 km",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1.540 km",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1.800 km",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "2.000 km",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "El río Cauca es afluente del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Orinoco",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Amazonas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Magdalena",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Atrato",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "¿Dónde nace el río Magdalena?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Páramo de Sumapaz",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Macizo Colombiano (Laguna del Magdalena)",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sierra Nevada",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nudo de los Pastos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "El río Caquetá desemboca en el:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Orinoco",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Amazonas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Pacífico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Magdalena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "La Laguna de la Cocha está en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Boyacá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nariño",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cundinamarca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "CLIMA Y BIODIVERSIDAD",
			Text:        "¿Cuántas especies de aves tiene Colombia aproximadamente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1.200",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1.500",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1.900",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "2.200",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "CLIMA Y BIODIVERSIDAD",
			Text:        "Colombia ocupa el primer lugar mundial en diversidad de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Mamíferos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Orquídeas y aves",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Reptiles",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Peces",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "CLIMA Y BIODIVERSIDAD",
			Text:        "¿Cuántas especies de orquídeas se estiman en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1.500",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "2.500",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "3.500",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "4.500",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "CLIMA Y BIODIVERSIDAD",
			Text:        "El oso de anteojos habita principalmente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La selva amazónica",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los bosques andinos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los llanos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El desierto",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "CLIMA Y BIODIVERSIDAD",
			Text:        "El delfín rosado se encuentra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El mar Caribe",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Los ríos amazónicos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El río Magdalena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PARQUES NACIONALES",
			Text:        "¿Cuántos parques nacionales naturales tiene Colombia aproximadamente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "30",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "40",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "50",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "60",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PARQUES NACIONALES",
			Text:        "El Parque Nacional Natural más grande es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tayrona",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Chiribiquete",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Nevados",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sierra Nevada",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PARQUES NACIONALES",
			Text:        "El Parque Nacional El Tuparro está en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Amazonas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Vichada",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Guainía",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Vaupés",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PARQUES NACIONALES",
			Text:        "El Parque Nacional Los Katíos es compartido con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Venezuela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Brasil",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Panamá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Ecuador",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "PARQUES NACIONALES",
			Text:        "El Parque Nacional Amacayacu protege principalmente:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Páramos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Selva amazónica",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Manglares",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Desiertos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "FRONTERAS Y LÍMITES",
			Text:        "Colombia tiene frontera terrestre con cuántos países:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "3",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "4",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "5",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "6",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "FRONTERAS Y LÍMITES",
			Text:        "La frontera más larga de Colombia es con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Brasil",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Venezuela",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Perú",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ecuador",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "FRONTERAS Y LÍMITES",
			Text:        "El río Arauca sirve de frontera con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Brasil",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Venezuela",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ecuador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Perú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "FRONTERAS Y LÍMITES",
			Text:        "La frontera con Panamá está marcada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El río Atrato",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Tapón del Darién",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La cordillera",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El mar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "FRONTERAS Y LÍMITES",
			Text:        "Colombia tiene fronteras marítimas con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "5 países",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "7 países",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "9 países",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "11 países",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "ASPECTOS ECONÓMICOS REGIONALES",
			Text:        "La principal actividad económica de la región Caribe es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Minería",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Turismo, ganadería y agroindustria",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo pesca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo industria",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "ASPECTOS ECONÓMICOS REGIONALES",
			Text:        "El Cerrejón en La Guajira es una mina de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Oro",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Carbón",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Esmeraldas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sal",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "ASPECTOS ECONÓMICOS REGIONALES",
			Text:        "Las esmeraldas colombianas provienen principalmente de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Boyacá y Cundinamarca",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Chocó",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "ASPECTOS ECONÓMICOS REGIONALES",
			Text:        "El petróleo se extrae principalmente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo Santander",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Meta, Casanare, Arauca",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo la costa",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Eje Cafetero",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "ASPECTOS ECONÓMICOS REGIONALES",
			Text:        "La zona bananera del Urabá está en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Magdalena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Antioquia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Córdoba",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Chocó",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODO PRECOLOMBINO",
			Text:        "Los Muiscas habitaban principalmente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Costa Caribe",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El altiplano cundiboyacense",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Valle del Cauca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Sierra Nevada",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODO PRECOLOMBINO",
			Text:        "Los Tayronas se ubicaban en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Andes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Sierra Nevada de Santa Marta",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Amazonía",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Llanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODO PRECOLOMBINO",
			Text:        "Los Quimbayas fueron famosos por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Su agricultura",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Su orfebrería en oro",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Su arquitectura",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Su navegación",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODO PRECOLOMBINO",
			Text:        "El Dorado era un mito relacionado con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Tayronas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Muiscas y la laguna de Guatavita",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Incas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Caribes",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "PERÍODO PRECOLOMBINO",
			Text:        "Los Zenúes desarrollaron un sistema de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Acueductos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Canales de drenaje en La Mojana",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Terrazas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Caminos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "¿En qué año llegó Rodrigo de Bastidas a la costa colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1492",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1499",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1502",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1510",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "Santa Marta fue fundada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pedro de Heredia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rodrigo de Bastidas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Gonzalo Jiménez de Quesada",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sebastián de Belalcázar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "¿En qué año fue fundada Santa Marta?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1510",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1525",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1533",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1538",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "Cartagena de Indias fue fundada en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1525",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1533",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1538",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1540",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "¿Quién fundó Cartagena?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Rodrigo de Bastidas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pedro de Heredia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Jiménez de Quesada",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nicolás de Federmán",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "Bogotá fue fundada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sebastián de Belalcázar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Gonzalo Jiménez de Quesada",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Nicolás de Federmán",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pedro de Heredia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "¿En qué fecha fue fundada Bogotá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "6 de agosto de 1537",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "6 de agosto de 1538",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "20 de julio de 1538",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "12 de octubre de 1538",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA CONQUISTA",
			Text:        "El cacique muisca que resistió a los españoles fue:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tisquesusa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sagipa",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Aquiminzaque",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Todos los anteriores",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "¿En qué año se creó la Audiencia de Santafé?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1538",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1549",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1564",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1580",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "El Virreinato de Nueva Granada se creó por primera vez en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1700",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1717",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1739",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1750",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "¿Por qué se suprimió temporalmente el Virreinato en 1723?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Falta de recursos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rebeliones",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Guerra con Inglaterra",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Problemas administrativos y económicos",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "El Virreinato se restableció definitivamente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1730",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1739",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1750",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1760",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "La economía colonial se basaba principalmente en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Agricultura y ganadería",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Explotación de oro y trata de esclavos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Comercio",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Industria",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "Las encomiendas eran:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tierras sin dueño",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sistema de trabajo indígena asignado a españoles",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ciudades",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Minas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "La mita era:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un impuesto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Trabajo forzoso indígena en las minas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una fiesta",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un cargo político",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "LA COLONIA",
			Text:        "Los resguardos indígenas eran:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Prisiones",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tierras comunales para indígenas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Hospitales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Escuelas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "Las Reformas Borbónicas fueron implementadas por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Felipe V",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Fernando VI",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Carlos III y Carlos IV",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Fernando VII",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "¿En qué siglo se implementaron las Reformas Borbónicas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "XVI",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "XVII",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "XVIII",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "XIX",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "La expulsión de los jesuitas ocurrió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1750",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1767",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1780",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1800",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "El movimiento de los Comuneros ocurrió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1770",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1775",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1781",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1790",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "Los Comuneros se levantaron en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Socorro, Santander",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Popayán",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "¿Quién lideró a los Comuneros?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José Antonio Galán",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Camilo Torres",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Francisco de Miranda",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "Las capitulaciones de Zipaquirá fueron:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un tratado de paz",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un acuerdo entre Comuneros y autoridades",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una rendición",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una declaración de guerra",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "¿Qué pasó con José Antonio Galán?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Fue perdonado",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Fue ejecutado y descuartizado",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Se exilió",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Fue nombrado gobernador",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "ILUSTRACIÓN Y EXPEDICIÓN BOTÁNICA",
			Text:        "La Expedición Botánica fue dirigida por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Alexander von Humboldt",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José Celestino Mutis",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Francisco José de Caldas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Jorge Tadeo Lozano",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "ILUSTRACIÓN Y EXPEDICIÓN BOTÁNICA",
			Text:        "La Expedición Botánica comenzó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1760",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1783",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1790",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1800",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "ILUSTRACIÓN Y EXPEDICIÓN BOTÁNICA",
			Text:        "Francisco José de Caldas era conocido como:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Precursor",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Sabio",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Libertador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Prócer",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "ILUSTRACIÓN Y EXPEDICIÓN BOTÁNICA",
			Text:        "Antonio Nariño tradujo:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Biblia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Derechos del Hombre y del Ciudadano",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Contrato Social",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Enciclopedia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "ILUSTRACIÓN Y EXPEDICIÓN BOTÁNICA",
			Text:        "¿En qué año tradujo Nariño los Derechos del Hombre?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1790",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1793",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1795",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1800",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "El Memorial de Agravios fue escrito por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Camilo Torres",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "¿En qué año se escribió el Memorial de Agravios?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1808",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1809",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1810",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1811",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "El florero de Llorente ocurrió un:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Lunes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Martes",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Viernes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Domingo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "¿Quiénes protagonizaron el incidente del florero?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Llorente y Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los hermanos Morales y Llorente",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Santander y Llorente",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nariño y Llorente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "La primera república (Patria Boba) duró de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1808-1815",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1810-1815",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1812-1816",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1815-1819",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "¿Por qué se llamó \"Patria Boba\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Por las derrotas militares",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Por las guerras civiles entre federalistas y centralistas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Por la falta de ejército",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Por la pobreza",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "El Congreso de las Provincias Unidas se reunió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tunja",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Socorro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "La Reconquista española fue liderada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pablo Morillo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Francisco de Miranda",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "José María Barreiro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "Pablo Morillo era conocido como:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Libertador",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Pacificador",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Conquistador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El General",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "El Régimen del Terror duró de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1810-1815",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1815-1819",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1819-1821",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1821-1830",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "La Campaña Libertadora de 1819 duró:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "50 días",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "77 días",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "100 días",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "150 días",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "¿Desde dónde partió Bolívar en la Campaña Libertadora?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Caracas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Llanos de Casanare",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cúcuta",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "La Batalla del Pantano de Vargas fue el:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "25 de julio de 1819",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "7 de agosto de 1819",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "10 de agosto de 1819",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "20 de julio de 1819",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "¿Quién comandó las tropas españolas en Boyacá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pablo Morillo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José María Barreiro",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sámano",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Latorre",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "El Congreso de Cúcuta se realizó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1819",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1821",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1823",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1825",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "INDEPENDENCIA",
			Text:        "La Constitución de Cúcuta estableció:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un sistema federal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un sistema centralista",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una monarquía",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una dictadura",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "La Gran Colombia existió entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1810-1830",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1819-1831",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1821-1835",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1825-1840",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "La capital de la Gran Colombia era:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Caracas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Quito",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Panamá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "¿Por qué se disolvió la Gran Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Guerra con España",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Diferencias regionales y políticas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Muerte de Bolívar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Invasión extranjera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "La Convención de Ocaña fue en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1826",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1828",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1830",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1832",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "La Noche Septembrina fue:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Una fiesta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un intento de asesinato a Bolívar",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una batalla",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una reunión política",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "¿Quién salvó a Bolívar en la Noche Septembrina?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Manuela Sáenz",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sucre",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Córdova",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "Bolívar murió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Caracas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "GRAN COLOMBIA",
			Text:        "¿En qué fecha murió Bolívar?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "17 de diciembre de 1830",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "20 de julio de 1831",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "7 de agosto de 1832",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "24 de diciembre de 1833",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "El primer presidente de la Nueva Granada fue:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santander",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Joaquín Mosquera",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "José Ignacio de Márquez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "La Guerra de los Supremos ocurrió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1830-1832",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1839-1842",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1845-1848",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1850-1853",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "La abolición definitiva de la esclavitud fue en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1821",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1835",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1851",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1863",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "¿Quién decretó la abolición de la esclavitud?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Mosquera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José Hilario López",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Murillo Toro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Núñez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "La Constitución de Rionegro era:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Centralista",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Federalista radical",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Monárquica",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Socialista",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "Los Estados Unidos de Colombia existieron entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1850-1886",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1863-1886",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1870-1900",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1880-1910",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "La Regeneración fue liderada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Mosquera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rafael Núñez",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Miguel Antonio Caro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Carlos Holguín",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "El Concordato con la Iglesia se firmó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1863",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1887",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1892",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1900",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "La Guerra de los Mil Días causó aproximadamente:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "10.000 muertos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "50.000 muertos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "100.000 muertos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "200.000 muertos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "¿Quién era presidente cuando se perdió Panamá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Rafael Reyes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José Manuel Marroquín",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Miguel Antonio Caro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Carlos Holguín",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "El tratado Herrán-Hay fue:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Aprobado por Colombia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rechazado por el Congreso colombiano",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Firmado con Panamá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un tratado de paz",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XIX - REPÚBLICA",
			Text:        "Panamá declaró su independencia con apoyo de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Inglaterra",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francia",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Estados Unidos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Venezuela",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "Rafael Reyes gobernó entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1900-1904",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1904-1909",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1910-1914",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1914-1918",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "El quinquenio de Reyes se caracterizó por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Guerras civiles",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Modernización y autoritarismo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Democracia plena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Aislamiento internacional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "La República Conservadora duró:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1886-1910",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1886-1930",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1900-1930",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1910-1940",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "La Masacre de las Bananeras ocurrió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1925",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1928",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1930",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1935",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "¿Cuántos trabajadores murieron en la Masacre de las Bananeras?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Número exacto desconocido (versiones varían)",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text: "13",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "100",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "3.000",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "La República Liberal comenzó con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "López Pumarejo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Olaya Herrera",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Eduardo Santos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Alberto Lleras",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "La \"Revolución en Marcha\" fue de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Olaya Herrera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Alfonso López Pumarejo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Eduardo Santos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Laureano Gómez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "Jorge Eliécer Gaitán fue asesinado a las:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "11:00 am",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1:15 pm",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "3:00 pm",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "5:00 pm",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "El asesino de Gaitán fue:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nunca se identificó",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Juan Roa Sierra",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un sicario",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un militar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "SIGLO XX",
			Text:        "Se estima que en El Bogotazo murieron:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "100 personas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Entre 500 y 3.000 personas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "5.000 personas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "10.000 personas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿En qué fecha se conmemora el Día de la Independencia de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "7 de agosto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "20 de julio",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "12 de octubre",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "19 de abril",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue conocido como \"El Libertador\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Simón Bolívar",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "José María Córdova",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Batalla de Boyacá ocurrió en el año:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1810",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1819",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1821",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1830",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Cuál fue el primer nombre de Colombia como república independiente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nueva Granada",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Gran Colombia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "República de Colombia",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Estados Unidos de Colombia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El período conocido como \"La Patria Boba\" se caracterizó por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Guerras civiles entre centralistas y federalistas",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "La paz total en el territorio",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El dominio español absoluto",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La alianza con Inglaterra",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue el primer presidente de la Gran Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Simón Bolívar",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Camilo Torres",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Guerra de los Mil Días ocurrió entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1899-1902",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "1810-1813",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1850-1853",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1930-1933",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Qué evento histórico ocurrió el 9 de abril de 1948?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La independencia de Panamá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El asesinato de Jorge Eliécer Gaitán (El Bogotazo)",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La firma de la Constitución",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El fin del Frente Nacional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El Frente Nacional fue un acuerdo político entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Liberales y Comunistas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Conservadores y Liberales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Todos los partidos políticos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Colombia y Estados Unidos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿En qué año se separó Panamá de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1903",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text: "1910",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1885",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1920",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue conocido como \"El Hombre de las Leyes\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Rafael Núñez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Constitución de Rionegro (1863) establecía un sistema:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Centralista",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Federalista",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Monárquico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Socialista",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue el autor del Himno Nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "José María Córdova",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rafael Núñez",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Oreste Síndici (música) y Rafael Núñez (letra)",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Jorge Isaacs",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El período de \"La Violencia\" en Colombia se sitúa principalmente entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "1920-1930",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1948-1958",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1960-1970",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1980-1990",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Cuál fue la primera constitución de Colombia independiente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Constitución de Socorro (1810)",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Constitución de Cúcuta (1821)",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Constitución de 1886",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Constitución de 1991",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El movimiento de los Comuneros ocurrió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1781",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text: "1810",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1819",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1850",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién tradujo los Derechos del Hombre y del Ciudadano al español?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Antonio Nariño",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Francisco de Miranda",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Camilo Torres",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Regeneración fue un movimiento político liderado por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Rafael Núñez",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Jorge Eliécer Gaitán",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Laureano Gómez",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Alfonso López Pumarejo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿En qué año se firmó el actual acuerdo de paz con las FARC?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "2014",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "2016",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "2018",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "2020",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El M-19 fue un grupo guerrillero que se desmovilizó en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1985",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1990",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1995",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "2000",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es la capital de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Bogotá D.C.",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "Colombia limita al norte con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Venezuela y Brasil",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Mar Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Panamá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ecuador",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuántos departamentos tiene Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "30",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "31",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "32",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "33",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El pico más alto de Colombia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pico Simón Bolívar",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Nevado del Ruiz",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Nevado del Huila",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Páramo de Sumapaz",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es el río más importante de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Río Cauca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Río Magdalena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Río Orinoco",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Río Amazonas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La región de la Orinoquía también es conocida como:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Llanos Orientales",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "La Costa Atlántica",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Pacífico",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Amazonía",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es el único departamento de Colombia que tiene costas en ambos océanos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Valle del Cauca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Chocó",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Nariño",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La Sierra Nevada de Santa Marta es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La formación montañosa costera más alta del mundo",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Un volcán activo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Una isla",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un parque temático",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es la ciudad más poblada de Colombia después de Bogotá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Medellín",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Archipiélago de San Andrés y Providencia se encuentra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Océano Pacífico",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Mar Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Golfo de México",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Atlántico Sur",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es el departamento más grande de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Meta",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Amazonas",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Vichada",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Páramo de Sumapaz es importante porque:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Es el páramo más grande del mundo",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Es el volcán más activo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Es la montaña más alta",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Es el desierto más extenso",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuántas regiones naturales tiene Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "4",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "5",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "6",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "7",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Tapón del Darién se encuentra entre:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Colombia y Venezuela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Colombia y Brasil",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Colombia y Panamá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Colombia y Ecuador",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La capital del departamento de Antioquia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Medellín",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Pereira",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Manizales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál de estas ciudades NO pertenece al Eje Cafetero?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Armenia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pereira",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Manizales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Bucaramanga",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Nevado del Ruiz se encuentra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Sierra Nevada de Santa Marta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Cordillera Central",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Cordillera Oriental",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Cordillera Occidental",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es el principal puerto sobre el Océano Pacífico?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Buenaventura",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La región Insular de Colombia incluye:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo San Andrés",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "San Andrés, Providencia y las islas del Pacífico",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo Gorgona",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo Malpelo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Cañón del Chicamocha se encuentra en el departamento de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Boyacá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santander",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cundinamarca",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Norte de Santander",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Cuál es la flor nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La rosa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El girasol",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La orquídea (Cattleya trianae)",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El clavel",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El ave nacional de Colombia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El águila",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El cóndor de los Andes",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El colibrí",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El tucán",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Cuál es el árbol nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El roble",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La ceiba",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La palma de cera del Quindío",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El guayacán",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El sombrero vueltiao es originario de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Córdoba y Sucre",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Boyacá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Quién escribió \"Cien años de soledad\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Jorge Isaacs",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rafael Pombo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Gabriel García Márquez",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "José Eustasio Rivera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El vallenato es originario de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La región Andina",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Costa Caribe",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La región Pacífica",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Llanos Orientales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La cumbia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un plato típico",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un baile y género musical",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una artesanía",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una celebración religiosa",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El Carnaval de Barranquilla se celebra:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "En diciembre",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Antes de la cuaresma",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "En julio",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "En octubre",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Cuál de estos es un plato típico de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Paella",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Bandeja paisa",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tacos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pizza",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El Festival de la Leyenda Vallenata se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Valledupar",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "Fernando Botero es famoso por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Su música",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sus esculturas y pinturas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Su literatura",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Su arquitectura",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Cuál es la bebida nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El ron",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El aguardiente",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La cerveza",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El café",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La Feria de las Flores se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bogotá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Medellín",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Pereira",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El tejo es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un baile típico",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un deporte nacional",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un instrumento musical",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una comida tradicional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Quién es conocido como \"El Joe\" en la música colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Carlos Vives",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Joe Arroyo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Juanes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Shakira",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El Festival Iberoamericano de Teatro se realiza en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Medellín",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Bogotá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La mochila wayúu es una artesanía de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Los Embera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Wayúu de La Guajira",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Muiscas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Kogui",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El bambuco es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un instrumento",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un género musical y baile andino",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Una comida",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una fiesta",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Cuál de estos escritores colombianos ganó el Premio Nobel de Literatura?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Jorge Isaacs",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José Asunción Silva",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Gabriel García Márquez",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Álvaro Mutis",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La Semana Santa en Popayán es famosa por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sus carnavales",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sus procesiones religiosas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sus conciertos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sus competencias deportivas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Qué artículo de la Constitución establece que Colombia es un Estado Social de Derecho?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Artículo 5",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Artículo 1",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Artículo 10",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Artículo 20",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "La tutela protege derechos:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Económicos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Colectivos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Fundamentales",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Culturales únicamente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Cuántos años dura el período de un alcalde en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "3 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "4 años",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "5 años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "2 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "El Defensor del Pueblo es elegido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Cámara de Representantes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Voto popular",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Qué organismo tiene la función de acusar ante el Senado al Presidente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Corte Suprema",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Fiscalía",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Cámara de Representantes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Procuraduría",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "El derecho a la doble nacionalidad en Colombia está:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Prohibido",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Permitido",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Restringido a ciertos países",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo para nacidos en Colombia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Cuál es la edad mínima para ser elegido Presidente de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "25 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "30 años",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "35 años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "40 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "La revocatoria del mandato es un mecanismo que aplica para:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Presidente y Vicepresidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Senadores y Representantes",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Gobernadores y Alcaldes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Todos los cargos de elección popular",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Cuántos magistrados tiene la Corte Suprema de Justicia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "23 magistrados",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "20 magistrados",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "25 magistrados",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "30 magistrados",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "El voto en Colombia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Obligatorio",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un derecho y un deber ciudadano",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Solo para mayores de 21 años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Exclusivo para nacionales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "La iniciativa legislativa puede ser ejercida por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el Gobierno",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Solo el Congreso",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Gobierno, Congreso y ciudadanos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Solo los ciudadanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Qué porcentaje del censo electoral se requiere para convocar un referendo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "3%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "5%",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "10%",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "15%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "El Fiscal General de la Nación es elegido por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Corte Suprema de Justicia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Congreso",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Voto popular",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "Los estados de excepción en Colombia son:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Guerra exterior, conmoción interior, emergencia",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Solo guerra y emergencia",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Solo conmoción interior",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Guerra, paz y emergencia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Cuál es el quórum mínimo para sesionar en el Congreso?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La mitad más uno",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un tercio",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Una cuarta parte",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Dos tercios",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "El Banco de la República tiene como función principal:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Prestar dinero al Gobierno",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Mantener la capacidad adquisitiva de la moneda",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Financiar empresas privadas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cobrar impuestos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Quién nombra al Procurador General de la Nación?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Presidente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Senado",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Corte Constitucional",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Consejo de Estado",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "La acción popular protege:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Derechos fundamentales",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Derechos colectivos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Derechos individuales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo el medio ambiente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "El período de los gobernadores es de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "2 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "3 años",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "4 años",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "5 años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CONSTITUCION",
			SubCategory: "",
			Text:        "¿Cuántos debates requiere un proyecto de ley para ser aprobado?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "2 debates",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "3 debates",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "4 debates",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "5 debates",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El grito de independencia del 20 de julio de 1810 ocurrió por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La invasión francesa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El incidente del florero de Llorente",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La muerte de un virrey",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Un terremoto",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue conocido como \"El Precursor\" de la independencia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Simón Bolívar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco de Miranda",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Antonio Nariño",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "José María Carbonell",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Campaña Libertadora de 1819 fue comandada por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Bolívar",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Páez",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sucre",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Qué países conformaban la Gran Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Colombia y Venezuela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Colombia, Venezuela, Ecuador, Panamá",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Colombia, Ecuador y Perú",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Solo Colombia y Panamá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Constitución de 1886 fue promovida por:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Tomás Cipriano de Mosquera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rafael Núñez",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "José Hilario López",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Manuel Murillo Toro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿En qué año fue abolida definitivamente la esclavitud en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1810",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1821",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1851",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1886",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Batalla del Pantano de Vargas ocurrió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1810",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1815",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1819",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1821",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue el primer presidente de la República de la Nueva Granada?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Joaquín Mosquera",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "José Ignacio de Márquez",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tomás Cipriano de Mosquera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El período conocido como \"La Hegemonía Conservadora\" duró aproximadamente:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "10 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "20 años",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "30 años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "45 años (1886-1930)",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La \"Masacre de las Bananeras\" ocurrió en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1928",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text: "1935",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1948",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1950",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue el presidente durante el cual Colombia perdió a Panamá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Rafael Núñez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "José Manuel Marroquín",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Miguel Antonio Caro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Rafael Reyes",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La República Liberal (1930-1946) comenzó con el presidente:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Alfonso López Pumarejo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Enrique Olaya Herrera",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Eduardo Santos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Alberto Lleras",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El Palacio de Justicia fue tomado por el M-19 en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1980",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1985",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text: "1989",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text: "1990",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue conocido como \"El Tribuno del Pueblo\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Jorge Eliécer Gaitán (José Acevedo y Gómez 1810)",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Luis Carlos Galán",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Laureano Gómez",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Rojas Pinilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La \"Séptima Papeleta\" fue un movimiento estudiantil que promovió:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La educación gratuita",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Asamblea Nacional Constituyente",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El voto femenino",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La paz con las guerrillas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿En qué año las mujeres obtuvieron el derecho al voto en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text: "1945",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text: "1954",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text: "1957",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text: "1961",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El Tratado Urrutia-Thomson de 1914 fue firmado con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Venezuela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Estados Unidos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ecuador",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Brasil",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "La Batalla de Palonegro fue parte de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La independencia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Guerra de los Mil Días",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Guerra con Perú",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Las guerras civiles del siglo XIX",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "¿Quién fue el único presidente militar del siglo XX en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Gustavo Rojas Pinilla",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Gabriel París",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Alberto Ruiz Novoa",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Rafael Reyes",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "HISTORIA",
			SubCategory: "",
			Text:        "El proceso de descentralización administrativa comenzó con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Constitución de 1886",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Frente Nacional",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Constitución de 1991",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Ley 60 de 1993",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es la extensión territorial aproximada de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "500.000 km²",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "800.000 km²",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1.141.748 km²",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "1.500.000 km²",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "Colombia tiene frontera marítima con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "5 países",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "7 países",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "9 países",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "11 países",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La capital del departamento del Meta es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Yopal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Villavicencio",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Florencia",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Arauca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es el departamento más pequeño de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Quindío",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Atlántico",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "San Andrés y Providencia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Risaralda",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Parque Nacional Natural más grande de Colombia es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Chiribiquete",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "La Macarena",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Cocuy",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tayrona",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál de estos ríos NO nace en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Magdalena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cauca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Orinoco",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Putumayo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La Península de La Guajira limita con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Panamá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Venezuela",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Brasil",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nicaragua",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuántas cordilleras atraviesan el territorio colombiano?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Una",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Dos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tres",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cuatro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Golfo de Urabá se encuentra entre los departamentos de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Chocó y Córdoba",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Antioquia y Chocó",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Córdoba y Sucre",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Atlántico y Magdalena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La capital del departamento de Boyacá es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Sogamoso",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Duitama",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tunja",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Chiquinquirá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es el punto más septentrional (norte) de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cabo de la Vela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Punta Gallinas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cabo Tiburón",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Punta Barú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Macizo Colombiano es importante porque:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Es el pico más alto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Es donde nacen los principales ríos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Es la capital ecológica",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Es el desierto más grande",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál de estas islas NO pertenece a Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Gorgona",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Malpelo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Margarita",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Rosario",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La depresión Momposina se encuentra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Costa Caribe",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Llanos Orientales",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Valle del Magdalena",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Amazonía",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La capital del departamento del Cesar es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Valledupar",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Aguachica",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Codazzi",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Bosconia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Desierto de la Tatacoa está ubicado en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La Guajira",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Huila",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cesar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Magdalena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál es el volcán más activo de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nevado del Ruiz",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Galeras",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Puracé",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Nevado del Huila",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "La Ciénaga Grande de Santa Marta es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un río",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un humedal costero",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un páramo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una bahía",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "¿Cuál de estas ciudades está más al sur?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Pasto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Leticia",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Mocoa",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Florencia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "GEOGRAFIA",
			SubCategory: "",
			Text:        "El Nudo de los Pastos o de Huaca se encuentra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Nariño",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Cauca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Putumayo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Huila",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El himno nacional de Colombia debe interpretarse en actos oficiales con:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Solo el coro",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Coro y primera estrofa",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Todas las estrofas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Coro y última estrofa",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Cuáles son los colores de la bandera de Colombia en orden?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Rojo, amarillo, azul",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Amarillo, azul, rojo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Azul, amarillo, rojo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Amarillo, rojo, azul",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El Festival Internacional de Cine de Cartagena (FICCI) se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Enero",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Marzo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Julio",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Diciembre",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Quién pintó \"La violencia\" y \"Los músicos\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Fernando Botero",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Alejandro Obregón",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Enrique Grau",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Omar Rayo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El escritor de \"La Vorágine\" fue:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "José Eustasio Rivera",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Jorge Isaacs",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tomás Carrasquilla",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Eduardo Caballero Calderón",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La champeta es un género musical originario de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sincelejo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El Festival de Teatro de Manizales se realiza cada:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Año",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Dos años",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Tres años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cinco años",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El ajiaco es un plato típico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Costa Caribe",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Bogotá y el altiplano",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Santander",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Quién compuso \"Colombia tierra querida\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Rafael Escalona",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Lucho Bermúdez",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Carlos Vives",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Alejo Durán",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El Carnaval de Negros y Blancos se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Popayán",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pasto",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ipiales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Túquerres",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La totuma es:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Un instrumento musical",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un recipiente tradicional",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un baile",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una comida",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Quién escribió \"María\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Gabriel García Márquez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Jorge Isaacs",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "José Asunción Silva",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Rafael Pombo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El Festival del Mono Núñez es de música:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Vallenata",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Andina colombiana",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Llanera",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tropical",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La Feria de Cali se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Agosto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Octubre",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Diciembre",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Febrero",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El acordeón es el instrumento principal del:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bambuco",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Vallenato",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Porro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Mapalé",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Cuál de estos es Patrimonio de la Humanidad por la UNESCO?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Carnaval de Barranquilla",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "La Feria de las Flores",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Festival Vallenato",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Feria de Cali",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El tejido del sombrero fino Sandoná es típico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Córdoba",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Nariño",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Boyacá",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tolima",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "El joropo es el baile típico de:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La región Andina",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Llanos Orientales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Costa Pacífica",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Costa Caribe",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "¿Quién es el autor de \"El olvido que seremos\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Héctor Abad Faciolince",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "William Ospina",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Juan Gabriel Vásquez",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Santiago Gamboa",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "CULTURA",
			SubCategory: "",
			Text:        "La Fiesta del Mar se celebra en:",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "San Andrés",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Buenaventura",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
		{
			Category:    "SÍMBOLOS PATRIOS",
			SubCategory: "NOMBRE OFICIAL",
			Text:        "¿Qué prócer dio el nombre de \"República de Colombia\" en el Congreso de Angostura en 1819?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El nombre fue dado por Simón Bolívar el 15 de febrero de 1819 en el Congreso de Angostura para reemplazar el nombre anterior de \"La Nueva Granada\".",
			Choices: []Choice{
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco Miranda",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Simón Bolívar",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Rafael Núñez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "SÍMBOLOS PATRIOS",
			SubCategory: "BANDERA NACIONAL",
			Text:        "¿Qué significado se le atribuye al color amarillo en la bandera nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El amarillo representa la riqueza del suelo, pero también simboliza la armonía y la justicia.",
			Choices: []Choice{
				{
					Text:      "La sangre derramada por los libertadores",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El cielo y los ríos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La riqueza del suelo, la armonía y la justicia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La soberanía sobre los océanos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "SÍMBOLOS PATRIOS",
			SubCategory: "ESCUDO NACIONAL",
			Text:        "¿Qué elemento se encuentra en la franja central del escudo de armas de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En la franja central, sobre un fondo de platino, se encuentra un gorro frigio sostenido por una lanza.",
			Choices: []Choice{
				{
					Text:      "Una granada de oro abierta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Dos buques con las velas desplegadas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Un gorro frigio sostenido por una lanza",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Dos cornucopias con monedas y frutas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "SÍMBOLOS PATRIOS",
			SubCategory: "HIMNO NACIONAL",
			Text:        "¿Quién fue el encargado de componer la música del himno nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La música fue magistralmente creada por el italiano Oreste Sindici a petición de José Domingo Torres.",
			Choices: []Choice{
				{
					Text:      "Rafael Núñez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco de Paula Santander",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Oreste Sindici",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "José Domingo Torres",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "SÍMBOLOS PATRIOS",
			SubCategory: "AVE NACIONAL",
			Text:        "¿Qué característica física distingue a los machos del cóndor andino?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los machos del cóndor destacan por tener una cresta carnosa roja en la cabeza.",
			Choices: []Choice{
				{
					Text:      "Un collar negro en la base del cuello",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Una cresta carnosa roja en la cabeza",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un plumaje completamente blanco",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una envergadura de alas de 100 km",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "SÍMBOLOS PATRIOS",
			SubCategory: "FLOR NACIONAL",
			Text:        "¿Cuál es el nombre científico de la orquídea considerada la flor nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La flor nacional es la *Cattleya Trianae*, conocida popularmente como \"flor de mayo\" o \"lirio de mayo\".",
			Choices: []Choice{
				{
					Text:      "Victoria Regia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ceroxylon Quindiuense",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cattleya Trianae",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Palma de Cera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "SÍMBOLOS PATRIOS",
			SubCategory: "ÁRBOL NACIONAL",
			Text:        "¿Hasta qué altura puede crecer la Palma de Cera del Quindío?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Palma de Cera (*Ceroxylon Quindiuense*) es una palmera singular que se erige con alturas de hasta 70 metros.",
			Choices: []Choice{
				{
					Text:      "Hasta 30 metros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Hasta 50 metros",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Hasta 70 metros",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Hasta 100 metros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLO CULTURAL",
			Text:        "¿De qué cultura indígena tiene raíces el sombrero vueltiao?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El sombrero vueltiao es una pieza artesanal característica de las sabanas del Caribe con raíces en la cultura Zenú.",
			Choices: []Choice{
				{
					Text:      "Cultura Tayrona",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cultura Wayúu",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cultura Zenú",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cultura Arhuaca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA",
			Text:        "¿Qué personaje aparece homenajeado en el billete de cien mil pesos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El billete de cien mil pesos rinde homenaje al presidente Carlos Lleras Restrepo (1966-1970).",
			Choices: []Choice{
				{
					Text:      "Gabriel García Márquez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Alfonso López Michelsen",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Carlos Lleras Restrepo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "José Asunción Silva",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA",
			Text:        "¿Qué elementos de fauna aparecen en el anverso del billete de cincuenta mil pesos junto a Gabriel García Márquez?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En el anverso del billete de cincuenta mil, junto a García Márquez, se encuentra un colibrí piquicorto y un caracol burgao.",
			Choices: []Choice{
				{
					Text:      "Una rana arborícola y una flor",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Un colibrí piquicorto y un caracol burgao",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Un abejorro y una puya",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Una tortuga y una serpiente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA",
			Text:        "¿Qué fruto acompaña la imagen del expresidente Alfonso López Michelsen en el billete de veinte mil pesos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En el anverso del billete de veinte mil se identifica la figura del mandatario junto con el fruto del anón.",
			Choices: []Choice{
				{
					Text:      "El fruto del café",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El fruto del cacao",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El fruto del anón",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El fruto de la palma de cera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA",
			Text:        "¿Quién es la protagonista representada en el billete de diez mil pesos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La antropóloga Virginia Gutiérrez de Pineda es la protagonista del billete de diez mil pesos.",
			Choices: []Choice{
				{
					Text:      "Débora Arango",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Policarpa Salavarrieta",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Virginia Gutiérrez de Pineda",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Soledad Acosta de Samper",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA",
			Text:        "¿Qué paisaje se presenta en el reverso del billete de cinco mil pesos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En el reverso del billete de cinco mil pesos se presenta el paisaje del páramo andino con los frailejones, el cóndor y el oso de anteojos.",
			Choices: []Choice{
				{
					Text:      "El Valle del Cocora",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ciudad Perdida",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El páramo andino",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Caño Cristales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA",
			Text:        "¿Qué artista paisa es homenajeadas en el billete de dos mil pesos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El billete de dos mil pesos rinde homenaje a la artista Débora Arango con un entramado de sus obras.",
			Choices: []Choice{
				{
					Text:      "Virginia Gutiérrez",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Débora Arango",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Frida Kahlo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Beatriz González",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MONEDA",
			Text:        "¿Qué animal está representado en la moneda de doscientos pesos de la nueva familia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La moneda de doscientos pesos representa a la guacamaya bandera.",
			Choices: []Choice{
				{
					Text:      "El oso de anteojos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La rana de cristal",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La guacamaya bandera",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La tortuga caguama",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "HISTORIA",
			Text:        "¿En qué año se creó el Servicio Oficial de Turismo, la primera institución pública encargada del desarrollo turístico en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Según la historia de la actividad turística descrita, aunque en 1918 se empezaron a crear leyes, fue hasta 1932 cuando se creó el Servicio Oficial de Turismo",
			Choices: []Choice{
				{
					Text:      "1918",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1931",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1932",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "1968",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "LEGISLACIÓN",
			Text:        "¿Qué entidad se creó mediante la Ley 60 de 1968 para encargarse de la planeación, regulación y financiación del turismo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El comienzo del turismo como política de Estado inició con la Ley 60 de 1968, la cual creó a la Corporación Nacional de Turismo (CNT)",
			Choices: []Choice{
				{
					Text:      "El Servicio Oficial de Turismo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Corporación Nacional de Turismo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Ministerio de Comercio",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Policía de Turismo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "LITORAL",
			Text:        "¿Cuáles son las tres ciudades costeras principales donde se concentra la mayor oferta hotelera del turismo litoral?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En el turismo litoral, la mayor oferta hotelera, restaurantes y bares se ha enfocado en las playas de Santa Marta, Barranquilla y Cartagena de Indias",
			Choices: []Choice{
				{
					Text:      "Tumaco, Guapi y Buenaventura",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Riohacha, Coveñas y Tolú",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Santa Marta, Barranquilla y Cartagena de Indias",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "San Andrés, Providencia y Santa Catalina",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "NATURALEZA",
			Text:        "¿Qué municipio del departamento de Santander es conocido por su oferta en turismo de aventura y deportes extremos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Destinos como San Gil, en el departamento de Santander, son conocidos por su oferta en turismo de aventura y actividades de montaña",
			Choices: []Choice{
				{
					Text:      "Barichara",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "San Gil",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Bucaramanga",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Socorro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "RURAL",
			Text:        "¿Qué modalidad turística ha emergido como una alternativa al turismo de sol y playa, permitiendo la interacción con espacios naturales del campo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Dentro de las nuevas opciones turísticas rurales ha emergido el agroecoturismo como una alternativa al clásico turismo de sol y playa",
			Choices: []Choice{
				{
					Text:      "El turismo religioso",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El turismo de negocios",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El agroecoturismo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El turismo urbano",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "LITORAL",
			Text:        "¿Qué ciudades del Pacífico funcionan como puertos internacionales para el transporte con países como Ecuador, Perú y Chile?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En la zona del Pacífico se destacan las ciudades de Tumaco y Buenaventura, las cuales funcionan como puertos internacionales",
			Choices: []Choice{
				{
					Text:      "Nuquí y Bahía Solano",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tumaco y Buenaventura",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Guapi y Timbiquí",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Quibdó y Ladrilleros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "FESTIVALES",
			Text:        "¿Qué festividad se celebra en la ciudad de Medellín según el texto?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Entre las festividades mencionadas que reflejan la herencia cultural se encuentra la Feria de las Flores en Medellín",
			Choices: []Choice{
				{
					Text:      "El Carnaval de Barranquilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Festival Vallenato",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Feria de las Flores",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Carnaval de Negros y Blancos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "REGIONES",
			Text:        "¿Qué característica natural define al Eje Cafetero además de las plantaciones de café?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El Eje Cafetero se destaca por la presencia del vulcanismo propio del Parque Nacional Natural Los Nevados",
			Choices: []Choice{
				{
					Text:      "Las playas de arena blanca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El desierto de la Tatacoa",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La presencia del vulcanismo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Los manglares extensos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "REGIONES",
			Text:        "¿Qué ciudades se configuran como las más importantes en términos de arquitectura urbana y centros financieros?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Bogotá, Medellín y Cali se configuran como las ciudades más importantes en términos de arquitectura urbana y centros financieros",
			Choices: []Choice{
				{
					Text:      "Cartagena, Santa Marta y Barranquilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Bogotá, Medellín y Cali",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Bucaramanga, Pereira y Manizales",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cúcuta, Ibagué y Villavicencio",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "TURISMO",
			SubCategory: "INFRAESTRUCTURA",
			Text:        "¿Cuál es la razón principal de la baja actividad turística hacia el departamento del Chocó?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Las playas del Pacífico, especialmente hacia el Chocó, no tienen un turismo muy activo debido a la baja cantidad de vías de comunicación terrestre, ya que la mayoría son fluviales",
			Choices: []Choice{
				{
					Text:      "La falta de atractivos naturales",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La baja cantidad de vías de comunicación terrestre",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La ausencia de población local",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El clima excesivamente frío",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "CONSTITUCIÓN DE 1991",
			Text:        "¿Qué sectores de la población fueron reconocidos explícitamente con derechos en la Constitución Política de 1991 tras haber sido excluidos anteriormente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La nueva Constitución de 1991 reconoció los derechos de grupos que anteriormente habían sido excluidos, específicamente los indígenas, afrocolombianos y mujeres",
			Choices: []Choice{
				{
					Text:      "Los gremios económicos y terratenientes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los indígenas, afrocolombianos y mujeres",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los miembros de la fuerza pública y el clero",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los inmigrantes europeos y asiáticos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "ECONOMÍA",
			Text:        "¿Qué fenómeno económico caracterizó a los gobiernos de Virgilio Barco y César Gaviria en la década de 1990?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Durante estos gobiernos se consolidó la apertura económica, lo que implicó la eliminación de barreras al comercio exterior, la inversión extranjera y la reducción de controles estatales sobre la empresa privada",
			Choices: []Choice{
				{
					Text:      "La nacionalización de la banca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El proteccionismo industrial",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La apertura económica",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El desarrollo agrario sostenible",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "INSTITUCIONES",
			Text:        "¿Qué entidad antecedió al Ministerio de Cultura y formaba parte del Ministerio de Educación antes de la Ley 397 de 1997?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Antes de la creación del Ministerio de Cultura, las funciones correspondían al Instituto Colombiano de Cultura, conocido como Colcultura, el cual formaba parte del Ministerio de Educación",
			Choices: []Choice{
				{
					Text:      "El Instituto Caro y Cuervo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Museo Nacional",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Colcultura",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Artesanías de Colombia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "CONFLICTO ARMADO",
			Text:        "¿Qué zona geográfica específica fue desmilitarizada durante el gobierno de Andrés Pastrana para llevar a cabo diálogos con las FARC?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Se estableció una zona desmilitarizada conocida como El Caguán, que comprendía los municipios de San Vicente del Caguán, La Uribe, Mesetas, La Macarena y Vista Hermosa",
			Choices: []Choice{
				{
					Text:      "La zona de despeje de El Caguán",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "El nudo de Paramillo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La región del Catatumbo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Montes de María",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "POLÍTICA DE SEGURIDAD",
			Text:        "¿Cuál fue el nombre de la política implementada por Álvaro Uribe Vélez destinada a combatir los grupos armados ilegales y recuperar el control territorial?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Durante sus dos mandatos, Álvaro Uribe implementó la política de Seguridad Democrática, la cual buscaba combatir grupos ilegales y mejorar la seguridad general mediante el fortalecimiento de las fuerzas armadas",
			Choices: []Choice{
				{
					Text:      "Paz Total",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Seguridad Democrática",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Alianza para el Progreso",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Plan Colombia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "SALUD PÚBLICA",
			Text:        "¿Cómo denominó el expresidente Iván Duque al aislamiento decretado durante la pandemia del COVID-19?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El expresidente Iván Duque denominó a la cuarentena como \"aislamiento preventivo obligatorio colaborativo e inteligente\"",
			Choices: []Choice{
				{
					Text:      "Confinamiento estricto total",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cuarentena generalizada",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Aislamiento preventivo obligatorio colaborativo e inteligente",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Distanciamiento social preventivo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "EL BOGOTAZO",
			Text:        "¿Qué evento desencadenó la ola de destrucción y caos conocida como el \"Bogotazo\" el 9 de abril de 1948?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El \"Bogotazo\" fue desatado por el asesinato del líder liberal Jorge Eliécer Gaitán cuando salía de su oficina en Bogotá",
			Choices: []Choice{
				{
					Text:      "La renuncia de Alfonso López Pumarejo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El fraude electoral de 1970",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El asesinato de Jorge Eliécer Gaitán",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El golpe de estado de Rojas Pinilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "GRUPOS ARMADOS",
			Text:        "¿Cuál fue el motivo principal que llevó a la fundación del Movimiento 19 de abril (M-19) en 1974?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El M-19 se fundó debido a la frustración generada por el fraude electoral en las elecciones presidenciales de 1970, donde se otorgó la victoria a Misael Pastrana sobre Gustavo Rojas Pinilla",
			Choices: []Choice{
				{
					Text:      "La muerte de Camilo Torres",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El fraude electoral de las elecciones de 1970",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La caída del Muro de Berlín",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La firma del Frente Nacional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "NARCOTRÁFICO",
			Text:        "¿Qué sucedió con la estructura del narcotráfico tras la muerte de Pablo Escobar y la captura de los líderes del cartel de Cali?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Tras la desaparición de los grandes capos, hubo una reorganización donde grupos ilegales como las guerrillas (FARC, ELN) y paramilitares (AUC) tomaron el control del negocio de producción y distribución",
			Choices: []Choice{
				{
					Text:      "El narcotráfico desapareció completamente de Colombia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El negocio pasó exclusivamente a manos de carteles extranjeros",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Guerrillas y paramilitares tomaron el control del negocio",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Estado nacionalizó la producción de coca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "PAZ",
			SubCategory: "JUSTICIA TRANSICIONAL",
			Text:        "¿Qué función cumple la Comisión de la Verdad dentro del proceso de postconflicto en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Comisión de la Verdad investiga y documenta las violaciones a los derechos humanos y crímenes del conflicto para esclarecer los hechos, reconocer a las víctimas y construir una memoria colectiva",
			Choices: []Choice{
				{
					Text:      "Juzgar penalmente a los excombatientes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Investigar y documentar hechos para construir memoria colectiva",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Otorgar amnistías inmediatas a todos los actores",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Administrar los recursos económicos de la reparación",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA COLONIAL",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "¿Cuál fue la orden religiosa expulsada tanto de España como de sus colonias en 1767 debido al fortalecimiento del Regalismo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El texto indica que el fortalecimiento del Regalismo, que defendía la autoridad del rey sobre la del Papa, permitió la expulsión de la Compañía de Jesús en 1767",
			Choices: []Choice{
				{
					Text:      "La Orden de los Agustinos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Orden de los Franciscanos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Compañía de Jesús",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Orden de los Dominicos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "INDEPENDENCIA",
			SubCategory: "MOVIMIENTOS POPULARES",
			Text:        "¿En qué centros poblados se desarrolló la revuelta de los Comuneros entre 1780 y 1781?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La revuelta de los Comuneros se desarrolló específicamente en Socorro, Charalá, Simacota y Mogotes, entre otros centros poblados",
			Choices: []Choice{
				{
					Text:      "Tunja, Villa de Leyva y Paipa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cartagena, Santa Marta y Barranquilla",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Socorro, Charalá, Simacota y Mogotes",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Popayán, Cali y Pasto",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "INDEPENDENCIA",
			SubCategory: "FIGURAS HISTÓRICAS",
			Text:        "¿Cómo se llamó el periódico en el que Antonio Nariño utilizó irónicamente el término \"Patria Boba\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Antonio Nariño utilizó el nombre \"Patria Boba\" en un periódico redactado por él mismo llamado \"Los Toros de Fucha\" para defenderse de ataques políticos",
			Choices: []Choice{
				{
					Text:      "La Bagatela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Aviso de Terremoto",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Los Toros de Fucha",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Espectador",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "ORGANIZACIÓN TERRITORIAL",
			Text:        "¿Qué territorios comprendía la República de Colombia creada por el Congreso de Angostura en 1819?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La República de Colombia creada el 17 de diciembre de 1819 comprendía los territorios que actualmente corresponden a Colombia, Ecuador, Panamá y Venezuela",
			Choices: []Choice{
				{
					Text:      "Colombia, Perú, Bolivia y Chile",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Colombia, Ecuador, Panamá y Venezuela",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Colombia, Venezuela y Brasil",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Colombia, Panamá y Costa Rica",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "POLÍTICA SIGLO XIX",
			SubCategory: "OLIMPO RADICAL",
			Text:        "¿Qué nombre recibió el país tras la promulgación de la Constitución de Rionegro en 1863?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Después de la promulgación de la Constitución de Rionegro en 1863, se le dio al país el nombre de Estados Unidos de Colombia",
			Choices: []Choice{
				{
					Text:      "República de la Nueva Granada",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Confederación Granadina",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Estados Unidos de Colombia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "República de Colombia",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "SÍMBOLOS PATRIOS",
			Text:        "¿Quién fue el encargado de musicalizar la letra del himno nacional escrita por Rafael Núñez?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El texto especifica que Rafael Núñez escribió la letra del himno nacional y el compositor italiano Oreste Síndici lo musicalizó",
			Choices: []Choice{
				{
					Text:      "José Rozo Contreras",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Francisco José de Caldas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Oreste Síndici",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Lucho Bermúdez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ECONOMÍA",
			SubCategory: "HEGEMONÍA CONSERVADORA",
			Text:        "¿En qué se invirtieron los 25 millones de dólares recibidos de Estados Unidos durante el gobierno de Pedro Nel Ospina?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El dinero recibido como indemnización por la separación de Panamá se empleó en obras de infraestructura, como la ampliación de la red ferroviaria y obras en el canal del Dique",
			Choices: []Choice{
				{
					Text:      "En el pago de la deuda externa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "En obras de infraestructura",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "En la creación de nuevos bancos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "En el fortalecimiento del ejército",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "MODERNIZACIÓN",
			SubCategory: "QUINQUENIO REYES",
			Text:        "¿Qué tratado firmado en 1905 permitió conseguir un crédito internacional para sanear la deuda externa de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En 1905 se firmó el tratado Averbury-Holguín, a través del cual se consiguió un crédito internacional que permitió sanear la deuda externa",
			Choices: []Choice{
				{
					Text:      "Tratado Urrutia-Thomson",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tratado Esguerra-Bárcenas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Tratado Averbury-Holguín",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Tratado Mallarino-Bidlack",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "POLÍTICA SIGLO XX",
			SubCategory: "REPÚBLICA LIBERAL",
			Text:        "¿Bajo qué ley se inició la construcción de la Ciudad Universitaria de la Universidad Nacional de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La reforma educativa bajo la ley 68 de 1935 inició la construcción de la Ciudad Universitaria de la Universidad Nacional de Colombia",
			Choices: []Choice{
				{
					Text:      "Ley 100 de 1936",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ley 68 de 1935",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Ley 30 de 1934",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Ley General de Educación",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS Y SOCIEDAD",
			SubCategory: "VOTO FEMENINO",
			Text:        "¿En qué fecha las mujeres votaron por primera vez en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El 1 de diciembre de 1957 se llevó a cabo el plebiscito en el que las mujeres votaron por primera vez en Colombia tras años de lucha",
			Choices: []Choice{
				{
					Text:      "13 de junio de 1953",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "10 de mayo de 1957",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1 de diciembre de 1957",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "7 de agosto de 1958",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "¿Cuáles son algunos de los ríos más importantes que caracterizan la hidrografía de la Región Pacífica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La región se caracteriza por tener una gran cantidad de ríos, entre los cuales los más importantes son el San Juan, el Atrato, el Baudó, el Mira y el Patía",
			Choices: []Choice{
				{
					Text:      "Magdalena, Cauca, Sinú, San Jorge y Nechí",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Amazonas, Orinoco, Meta, Guaviare y Vichada",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "San Juan, Atrato, Baudó, Mira y Patía",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Arauca, Casanare, Tomo, Tuparro y Bita",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Qué cadena montañosa corre paralela a la costa del Pacífico y alcanza picos superiores a los 4.000 metros sobre el nivel del mar?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Cordillera Occidental es la cadena montañosa que corre paralela a la costa del Pacífico y alcanza altitudes elevadas con picos que superan los 4.000 metros",
			Choices: []Choice{
				{
					Text:      "La Sierra Nevada de Santa Marta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Cordillera Central",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Cordillera Occidental",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Cordillera Oriental",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FLORA",
			Text:        "¿Qué nombre científico recibe la orquídea conocida comúnmente como \"flor de mayo\" presente en la región?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Dentro de las orquídeas representativas de la región, la *Cattleya trianae* es conocida como la flor de mayo",
			Choices: []Choice{
				{
					Text:      "Cattleya percivaliana",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cattleya warscewiczii",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cattleya trianae",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cattleya mossiae",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "SOCIEDAD",
			SubCategory: "POBLACIÓN",
			Text:        "¿Qué comunidades indígenas se mencionan explícitamente como habitantes de la Región Pacífica junto a la población afrocolombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En la región conviven comunidades indígenas como Emberá, Wounaan y Awá",
			Choices: []Choice{
				{
					Text:      "Wayúu, Kogui y Arhuaco",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Guambianos, Paeces y Totoroes",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Emberá, Wounaan y Awá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Tikuna, Huitoto y Bora",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "DIVISIÓN POLÍTICA",
			Text:        "¿Cuál es la capital del departamento de Nariño según la división política de la Región Pacífica mostrada en el mapa?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Según la información de departamentos y capitales de la región, la capital del departamento de Nariño es Pasto",
			Choices: []Choice{
				{
					Text:      "Quibdó",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cali",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Popayán",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pasto",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FAUNA",
			Text:        "¿Qué especies de ranas venenosas se encuentran en la Región Pacífica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En la zona se encuentran ranas venenosas, específicamente la rana dorada y la rana de cristal",
			Choices: []Choice{
				{
					Text:      "La rana toro y la rana platanera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La rana dorada y la rana de cristal",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La rana tomate y la rana de ojos rojos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La rana arlequín y la rana marsupial",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "UBICACIÓN",
			Text:        "¿Qué departamentos abarca la franja costera de la Región Pacífica de norte a sur?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Región Pacífica es una franja costera que se extiende desde el departamento de Chocó, en el norte, hasta el departamento de Nariño en el sur",
			Choices: []Choice{
				{
					Text:      "Desde La Guajira hasta el Atlántico",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Desde Antioquia hasta el Valle del Cauca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Desde el departamento de Chocó hasta el departamento de Nariño",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Desde el Meta hasta el Amazonas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y GASTRONOMÍA",
			Text:        "¿Qué elementos del folclore y la gastronomía destacan en la cultura de la Región Pacífica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La región destaca por su folclore, incluyendo el currulao y la marimba de chonta, y por su gastronomía con platos como el sancocho de pescado y el arroz con coco",
			Choices: []Choice{
				{
					Text:      "El vallenato, el acordeón, el suero costeño y la arepa de huevo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El joropo, el arpa, la ternera a la llanera y las hallacas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El currulao, la marimba de chonta, el sancocho de pescado y el arroz con coco",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El bambuco, el tiple, el ajiaco y la lechona",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FLORA",
			Text:        "¿Qué característica permite a las bromelias sobrevivir en los ambientes húmedos de la zona?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Las bromelias se caracterizan por su capacidad de almacenar agua en sus hojas, lo que les permite sobrevivir en ambientes húmedos",
			Choices: []Choice{
				{
					Text:      "Su sistema de raíces profundas que llega a las aguas subterráneas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Su capacidad de almacenar agua en sus hojas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sus espinas que evitan la pérdida de agua por transpiración",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Su simbiosis con hongos que retienen la humedad",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ECOLOGÍA",
			SubCategory: "ECOSISTEMAS",
			Text:        "¿Qué función vital cumplen los estuarios y manglares presentes en la costa del Pacífico?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los estuarios y manglares son vitales para la conservación de la biodiversidad y cumplen un papel importante en la protección de la costa contra la erosión y los fenómenos climáticos extremos",
			Choices: []Choice{
				{
					Text:      "Proveen madera para la industria de la construcción exclusivamente",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Son zonas destinadas únicamente al turismo de playa",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Filtran el agua salada para convertirla en agua dulce potable",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Conservan la biodiversidad y protegen la costa contra la erosión y fenómenos climáticos",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "GENERAL",
			Text:        "¿Qué porcentaje aproximado del territorio de Colombia abarca la llanura de la región de la Orinoquia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La región se extiende por una vasta llanura que abarca aproximadamente el 28% del territorio de Colombia",
			Choices: []Choice{
				{
					Text:      "El 15%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El 28%",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El 45%",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El 10%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "¿Cuáles son los dos ríos más importantes que atraviesan la Orinoquia y conforman su red hidrográfica principal?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Orinoquia está atravesada por dos de los ríos más importantes de Colombia: el Meta y el Orinoco, los cuales conforman una extensa red hidrográfica",
			Choices: []Choice{
				{
					Text:      "El Magdalena y el Cauca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Amazonas y el Putumayo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Meta y el Orinoco",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Atrato y el Sinú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Qué serranía de la región es conocida por albergar el famoso Cañón del Río Cristal?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Serranía de la Macarena es especialmente conocida por su belleza natural y el famoso Cañón del Río Cristal",
			Choices: []Choice{
				{
					Text:      "La Serranía de la Lindosa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Serranía del Chiribiquete",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Serranía de la Macarena",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Serranía del Perijá",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "FAUNA",
			SubCategory: "MAMÍFEROS",
			Text:        "¿Qué animal que habita en los ríos y humedales de la región es considerado el roedor más grande del mundo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El Capibara es considerado la especie de roedor más grande del mundo y habita en los ríos y humedales de la región",
			Choices: []Choice{
				{
					Text:      "El chigüiro",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La guartinaja",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El capibara",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El ñeque",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "FAUNA",
			SubCategory: "REPTILES",
			Text:        "¿Cuál es la especie de caimán endémica de la región que se encuentra en peligro de extinción?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El Caimán del Orinoco es una especie de caimán endémica de la región y se encuentra en peligro de extinción",
			Choices: []Choice{
				{
					Text:      "El caimán aguja",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El caimán del Orinoco",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El caimán negro",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La babilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA POLÍTICA",
			SubCategory: "CAPITALES",
			Text:        "¿Cuál es la capital del departamento de Vichada?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Según la información sobre los departamentos y capitales de la región, la capital de Vichada es Puerto Carreño",
			Choices: []Choice{
				{
					Text:      "Yopal",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arauca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Villavicencio",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Puerto Carreño",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DEMOGRAFÍA",
			SubCategory: "GRUPOS ÉTNICOS",
			Text:        "¿Cuáles son algunos de los grupos étnicos indígenas más representativos que habitan en la región de la Orinoquia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Entre los grupos étnicos más representativos de la región se encuentran los Sikuani, Piapoco, Puinave, Guahibo y Curripaco",
			Choices: []Choice{
				{
					Text:      "Wayuu, Nasa y Arhuacos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Embera, Zenu y Kogu",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Sikuani, Piapoco y Guahibo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Huitoto, Tikuna y Bora",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DEMOGRAFÍA",
			SubCategory: "NACIONAL",
			Text:        "¿Qué porcentaje de la población total de Colombia representan los afrocolombianos según los datos demográficos básicos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los afrocolombianos representan aproximadamente el 6.7% de la población total de Colombia, lo que corresponde a más de 3 millones de personas",
			Choices: []Choice{
				{
					Text:      "4.3%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "10.5%",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "6.7%",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "28%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DEMOGRAFÍA",
			SubCategory: "POBLACIÓN ROM",
			Text:        "¿Cuántas personas se autorreconocieron como gitanos o ROM?",
			Difficulty:  2,
			Points:      10,
			Explanation: "De acuerdo con el último Censo, un total de 2.606 personas se autorreconocieron como gitanos o ROM",
			Choices: []Choice{
				{
					Text:      "6.650 personas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "25.300 personas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "1.8 millones de personas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "2.606 personas",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ECONOMÍA",
			SubCategory: "ACTIVIDADES PRODUCTIVAS",
			Text:        "¿En qué actividades se basa principalmente la economía de la Orinoquia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La economía se basa principalmente en actividades agropecuarias, ganadería, extracción de petróleo, minería y turismo",
			Choices: []Choice{
				{
					Text:      "Industria textil y manufactura",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Pesca marítima y transporte portuario",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Actividades agropecuarias, ganadería, petróleo y minería",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Tecnología y servicios financieros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Cuál de las tres cordilleras principales de Colombia se caracteriza por ser la más alta y albergar el Nevado del Huila?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Cordillera Central es la más alta de las tres cordilleras y alberga el volcán más alto de Colombia, el Nevado del Huila",
			Choices: []Choice{
				{
					Text:      "La Cordillera Occidental",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Cordillera Oriental",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Cordillera Central",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Serranía del Baudó",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "SISTEMA OROGRÁFICO",
			Text:        "¿En qué sistema montañoso se encuentra el Pico Cristóbal Colón, considerado la montaña más alta de Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Sierra Nevada de Santa Marta se eleva abruptamente desde la costa y alcanza altitudes de más de 5.700 metros en el Pico Cristóbal Colón, lo que la convierte en la montaña más alta de Colombia",
			Choices: []Choice{
                {
                    Text:      "Sierra Nevada de Santa Marta",
                    IsCorrect: true,
                },
                {
                    Text:      "Cordillera Central",
                    IsCorrect: false,
                },
                {
                    Text:      "Cordillera Oriental",
                    IsCorrect: false,
                },
                {
                    Text:      "Serranía de la Macarena",
                    IsCorrect: false,
                },
            },
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "VALLES INTERANDINOS",
			Text:        "¿Qué cultivos predominan en el paisaje agrícola del Valle del Cauca?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El paisaje del Valle del Cauca está dominado por extensas áreas agrícolas con cultivos de caña de azúcar, café, arroz, frutas y hortalizas",
			Choices: []Choice{
				{
					Text:      "Trigo y cebada",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Caña de azúcar, café y arroz",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Palma de aceite exclusivamente",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Papa y hortalizas de clima frío",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "LLANURAS",
			Text:        "¿Qué nombre reciben las zonas de tierras bajas y pantanosas sujetas a inundaciones estacionales en la Llanura Amazónica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "A lo largo de la llanura amazónica se pueden encontrar áreas inundables conocidas como \"varzeas\" y \"igapós\", que son zonas de tierras bajas y pantanosas sujetas a inundaciones estacionales",
			Choices: []Choice{
				{
					Text:      "Sabanas y esteros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Varzeas e igapós",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Terrazas y mesetas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Deltas y estuarios",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "TERRITORIO INSULAR",
			Text:        "¿Por qué actividad de observación de fauna es mundialmente reconocida la Isla Malpelo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Isla Malpelo es uno de los lugares más importantes del mundo para la observación de tiburones y cuenta con una gran diversidad de especies marinas",
			Choices: []Choice{
                {
                    Text:      "Buceo con tiburones, especialmente el tiburón martillo",
                    IsCorrect: true,
                },
                {
                    Text:      "Avistamiento de ballenas jorobadas y sus crías",
                    IsCorrect: false,
                },
                {
                    Text:      "Observación de delfines rosados de río",
                    IsCorrect: false,
                },
                {
                    Text:      "Snorkel en la barrera de coral de los siete colores",
                    IsCorrect: false,
                },
            },
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "PISOS TÉRMICOS",
			Text:        "¿Qué tipo de vegetación es característica del piso térmico de Tierras Heladas o Páramos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los páramos se caracterizan por plantas adaptadas al clima frío, como frailejones y espeletias",
			Choices: []Choice{
                {
                    Text:      "Frailejones, musgos y líquenes",
                    IsCorrect: true,
                },
                {
                    Text:      "Manglares, palmas de coco y cacao",
                    IsCorrect: false,
                },
                {
                    Text:      "Cafetales, guadua y caña de azúcar",
                    IsCorrect: false,
                },
                {
                    Text:      "Cactus y vegetación desértica",
                    IsCorrect: false,
                },
            },
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RECURSOS NATURALES",
			Text:        "¿En qué departamento de Colombia se encuentran los principales yacimientos de esmeraldas de alta calidad?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Colombia es famosa por la producción de esmeraldas de alta calidad y sus principales yacimientos se encuentran en la región de Boyacá",
			Choices: []Choice{
				{
					Text:      "Antioquia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Chocó",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Boyacá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Córdoba",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RECURSOS MINERALES",
			Text:        "¿En qué regiones se ubican los depósitos importantes de níquel en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El país posee importantes depósitos de níquel en la región de la Sierra Nevada de Santa Marta y en el departamento de Córdoba",
			Choices: []Choice{
				{
					Text:      "Magdalena Medio y Llanos Orientales",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sierra Nevada de Santa Marta y Córdoba",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Antioquia y Tolima",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Región Andina y Costa Caribe",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA HUMANA",
			SubCategory: "GRUPOS ÉTNICOS",
			Text:        "¿Cuáles son algunos de los grupos indígenas más prominentes mencionados que conservan su cultura e idioma en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Entre los grupos indígenas más destacados se encuentran los Wayuu, los Emberá, los Nasa, los Kogui y los Arhuacos",
			Choices: []Choice{
				{
					Text:      "Los Aztecas, los Mayas y los Incas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Los Wayuu, los Emberá y los Kogui",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Los Mapuches, los Aymaras y los Quechuas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Los Zapotecas, los Mixtecas y los Toltecas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA HUMANA",
			SubCategory: "GRUPOS ÉTNICOS",
			Text:        "¿Qué comunidad específica se menciona explícitamente como parte incluida dentro del grupo poblacional de los afrocolombianos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Dentro del grupo poblacional afrocolombiano se incluye explícitamente a la comunidad de palenque",
			Choices: []Choice{
				{
					Text:      "La comunidad de palenque",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "La comunidad ROM",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La comunidad raizal",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La comunidad mestiza",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA HUMANA",
			SubCategory: "GRUPOS ÉTNICOS",
			Text:        "¿Cuál es el origen del grupo poblacional conocido como mestizos en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La población mestiza es el resultado de la mezcla entre indígenas y europeos que ocurrió durante la época colonial",
			Choices: []Choice{
				{
					Text:      "Mezcla entre africanos y europeos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Mezcla entre indígenas y asiáticos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Mezcla entre indígenas y europeos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Mezcla entre africanos e indígenas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CENTROS URBANOS",
			SubCategory: "BOGOTÁ",
			Text:        "¿Cuál es la cifra estimada de habitantes que reside dentro de la ciudad de Bogotá, sin contar su área metropolitana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Bogotá tiene una población estimada de 7'181.469 habitantes, siendo la ciudad más poblada del país",
			Choices: []Choice{
				{
					Text:      "10 millones de habitantes",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "2'372.330 habitantes",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "3.5 millones de habitantes",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "7'181.469 habitantes",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CENTROS URBANOS",
			SubCategory: "MEDELLÍN",
			Text:        "¿Qué evento cultural específico se destaca como parte de la oferta cultural de la ciudad de Medellín?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Medellín alberga diversos eventos culturales, destacándose entre ellos la Feria de las Flores",
			Choices: []Choice{
				{
					Text:      "El Carnaval de Barranquilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Feria de Cali",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Feria de las Flores",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Festival de Blancos y Negros",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CENTROS URBANOS",
			SubCategory: "CALI",
			Text:        "¿Con qué título cultural es conocida la ciudad de Cali debido a su vibrante escena musical y de baile?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Cali es conocida culturalmente como la capital de la salsa",
			Choices: []Choice{
				{
					Text:      "La capital del vallenato",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La capital de la salsa",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La ciudad de la eterna primavera",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La capital de la montaña",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ZONAS RURALES",
			SubCategory: "ECONOMÍA AGRÍCOLA",
			Text:        "¿Cuáles son los cultivos principales que sustentan la economía de las zonas rurales mencionadas en el texto?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los cultivos principales en las zonas rurales incluyen café, plátano, caña de azúcar, arroz, cacao y frutas tropicales",
			Choices: []Choice{
				{
					Text:      "Trigo, cebada, centeno y uvas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Café, plátano, caña de azúcar y cacao",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Maíz, soja, girasol y sorgo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Papa, quinua, oca y olluco",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ZONAS RURALES",
			SubCategory: "DEMOGRAFÍA RURAL",
			Text:        "¿Qué tipos de comunidades componen generalmente la demografía de las áreas rurales en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En las áreas rurales se encuentran generalmente comunidades campesinas, indígenas y afrocolombianas que han habitado estas zonas durante generaciones",
			Choices: []Choice{
				{
					Text:      "Principalmente comunidades de inmigrantes extranjeros",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Comunidades campesinas, indígenas y afrocolombianas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Exclusivamente comunidades industriales y mineras",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Mayoritariamente población urbana desplazada temporalmente",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "¿Cuál es el río más largo de Colombia que atraviesa la región Caribe y es vital para el transporte de carga y pasajeros?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El río Magdalena es descrito como el más largo de Colombia, siendo vital para el transporte de carga y pasajeros, además de ser una fuente importante de agua dulce para la región",
			Choices: []Choice{
				{
					Text:      "El río Cauca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El río Sinú",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El río Magdalena",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El río Ranchería",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Qué altura alcanza el Pico Cristóbal Colón ubicado en la Sierra Nevada de Santa Marta?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El Pico Cristóbal Colón es el pico más alto de Colombia y tiene una altitud de 5.775 metros sobre el nivel del mar",
			Choices: []Choice{
				{
					Text:      "4.500 metros sobre el nivel del mar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "5.775 metros sobre el nivel del mar",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "3.200 metros sobre el nivel del mar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "6.100 metros sobre el nivel del mar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FLORA",
			Text:        "¿Cuáles son las tres especies de manglares que se encuentran presentes en los ecosistemas costeros de la región Caribe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Entre las especies de manglares presentes en el territorio se encuentran el mangle rojo, el mangle blanco y el mangle negro",
			Choices: []Choice{
				{
					Text:      "Mangle rojo, mangle blanco y mangle negro",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Mangle azul, mangle gris y mangle verde",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Mangle dulce, mangle salado y mangle espinoso",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Mangle de río, mangle de montaña y mangle costero",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA POLÍTICA",
			SubCategory: "DEPARTAMENTOS Y CAPITALES",
			Text:        "¿Cuál es la capital del departamento de Córdoba?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Según el mapa y la lista de departamentos y capitales de la región, la capital del departamento de Córdoba (identificado con el número 6) es Montería",
			Choices: []Choice{
				{
					Text:      "Sincelejo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Riohacha",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Montería",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Valledupar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FAUNA",
			Text:        "¿Qué especies de reptiles venenosas se mencionan específicamente como parte de la fauna de la región Caribe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En la zona se encuentran serpientes venenosas, incluyendo específicamente la serpiente coral y la mapaná",
			Choices: []Choice{
				{
					Text:      "La cascabel y la anaconda",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La serpiente coral y la mapaná",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La boa constrictor y la pitón",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La víbora de pestañas y la cazadora",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DEMOGRAFÍA",
			SubCategory: "POBLACIÓN",
			Text:        "¿Qué porcentaje aproximado de la población total del país alberga la región Caribe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La región Caribe alberga cerca del 22,5% de la población total del país",
			Choices: []Choice{
				{
					Text:      "15,3%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "30,1%",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "22,5%",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "18,9%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FLORA",
			Text:        "¿Qué especies de palmas se destacan en la región por sus usos diversos como la obtención de frutos comestibles y fibras?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Las especies de palmas más destacadas en la región son la palma real, el chontaduro, la palma de vino y la palma de moriche",
			Choices: []Choice{
				{
					Text:      "Palma de cera, palma de coco y palma africana",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Palma real, chontaduro, palma de vino y palma de moriche",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Palma areca, palma botella y palma fénix",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Palma robellini, palma manila y palma yuca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA POLÍTICA",
			SubCategory: "DEPARTAMENTOS Y CAPITALES",
			Text:        "¿Cuál es la capital del departamento de Sucre?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En la lista de departamentos y capitales, Sucre (identificado con el número 5) tiene como capital a Sincelejo",
			Choices: []Choice{
				{
					Text:      "Cartagena",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Santa Marta",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Sincelejo",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Barranquilla",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿En qué ubicación geográfica específica se encuentra la Serranía de Perijá?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Serranía de Perijá se encuentra en el extremo norte de la región Caribe, específicamente en la frontera con Venezuela",
			Choices: []Choice{
				{
					Text:      "En el centro de la llanura costera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "En el extremo sur, limitando con la región andina",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "En el mar Caribe Suroccidental",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "En el extremo norte, en la frontera con Venezuela",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DEMOGRAFÍA",
			SubCategory: "GRUPOS ÉTNICOS",
			Text:        "¿Cuáles son los tres grupos indígenas mencionados explícitamente como parte de la diversidad étnica de la región Caribe?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La región es hogar de una mezcla de grupos étnicos que incluye indígenas como los wayúu, kogui y arhuaco",
			Choices: []Choice{
				{
					Text:      "Wayúu, Kogui y Arhuaco",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Emberá, Guambiano y Paez",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Ticuna, Huitoto y Nukak",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Zenú, Muisca y Pijao",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "POBLAMIENTO",
			Text:        "¿Qué penínsulas conectaba el estrecho de Beringia por el cual cruzaron los primeros pobladores humanos hacia el continente americano?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El texto especifica que el estrecho de Beringia conectaba el extremo oriental de Asia (península de Chukotka en Rusia) con el extremo occidental de América del Norte (península de Seward en Alaska).",
			Choices: []Choice{
				{
					Text:      "La península de Kamchatka y la península de Alaska",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La península de Chukotka y la península de Seward",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La península de Yucatán y la península de Florida",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La península Ibérica y la península del Labrador",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "PRECOLOMBINO",
			Text:        "Además de la lengua chibcha y caribe, ¿cuál fue la otra lengua predominante entre los grupos indígenas establecidos en el territorio?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Las lenguas predominantes entre los pueblos indígenas fueron la chibcha, arawak y caribe.",
			Choices: []Choice{
				{
					Text:      "Quechua",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Arawak",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Náhuatl",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Guaraní",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "PRECOLOMBINO",
			Text:        "¿Hace cuántos años se estima, según hallazgos arqueológicos, que el territorio colombiano ya había sido habitado por grupos nómadas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Según los hallazgos arqueológicos, se cree que el territorio fue habitado por grupos nómadas hace 16.000 años.",
			Choices: []Choice{
				{
					Text:      "Hace 5.000 años",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Hace 10.000 años",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Hace 20.000 años",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Hace 16.000 años",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "CONQUISTA",
			Text:        "¿Qué conquistador fundó la ciudad de Santa Marta en el año 1533?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La colonización se intensificó a partir de 1533 con la fundación de Santa Marta por parte de Pedro de Heredia.",
			Choices: []Choice{
				{
					Text:      "Rodrigo de Bastidas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Gonzalo Jiménez de Quesada",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Pedro de Heredia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Sebastián de Belalcázar",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "CONQUISTA",
			Text:        "¿Qué fundación marcó la culminación de la expedición de Gonzalo Jiménez de Quesada hacia el interior del territorio en 1538?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La expedición de Gonzalo Jiménez de Quesada hacia el interior culminó con la fundación de Santa Fe de Bogotá en 1538.",
			Choices: []Choice{
				{
					Text:      "Cartagena de Indias",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Popayán",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Santa Fe de Bogotá",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Villa de Leyva",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "CONQUISTA",
			Text:        "¿Qué enfermedades traídas por los españoles desencadenaron epidemias con un alto índice de mortalidad indígena?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los españoles trajeron enfermedades desconocidas para los indígenas, como el sarampión y la viruela.",
			Choices: []Choice{
				{
					Text:      "La peste negra y el cólera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El sarampión y la viruela",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La fiebre amarilla y la malaria",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La gripe española y el tifus",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "COLONIA",
			Text:        "¿Cuáles eran las dos actividades en las que se basaba la economía del Virreinato de la Nueva Granada?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El Virreinato de la Nueva Granada basaba su economía en la explotación del oro y la trata de esclavos africanos.",
			Choices: []Choice{
				{
					Text:      "La agricultura de tabaco y la ganadería",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La explotación del oro y la trata de esclavos africanos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La exportación de café y la textilería",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La pesca artesanal y el comercio de especias",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "COLONIA",
			Text:        "¿Con qué propósito utilizaba la Corona española el oro extraído del territorio del Nuevo Reino de Granada?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El oro se usaba para financiar guerras contra los luteranos en Flandes y Alemania, los ingleses anglicanos, los franceses y los turcos.",
			Choices: []Choice{
				{
					Text:      "Para construir catedrales en América",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Para pagar la deuda externa con Portugal",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Para financiar guerras en Europa contra luteranos, ingleses, franceses y turcos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Para desarrollar la industria naval en el Pacífico",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "COLONIA",
			Text:        "¿Qué entidad administrativa creada en 1549 abarcaba las provincias de Santa Marta, San Juan, Popayán, Guayana y Cartagena de Indias?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En 1549 se creó la Audiencia de Santafé, que abarcaba dichas provincias.",
			Choices: []Choice{
				{
					Text:      "El Consejo de Indias",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Real Audiencia de Santafé",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "La Capitanía General de Venezuela",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Virreinato del Perú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "HISTORIA",
			SubCategory: "REFORMAS BORBÓNICAS",
			Text:        "¿Cuál fue la intención principal de la Corona al implementar las Reformas Borbónicas un siglo más tarde?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Corona se vio obligada a tomar estas medidas con la intención de mejorar la administración de las colonias y de la península, ya que se estaban escapando de su control.",
			Choices: []Choice{
				{
					Text:      "Otorgar la independencia a los territorios de ultramar",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Mejorar la administración de las colonias y de la península",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Eliminar los impuestos a los indígenas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Promover la religión protestante en las colonias",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "UBICACIÓN Y LÍMITES",
			Text:        "¿En qué parte del país se encuentra ubicada la Región Andina y cuántas cordilleras la componen?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "En el sur del país y compuesta por una cordillera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "En el centro-occidente del país y compuesta por tres cordilleras",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "En el norte del país y compuesta por dos cordilleras",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "En el oriente del país y compuesta por cuatro cordilleras",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "¿Cuáles son los ríos importantes mencionados que nacen en la Región Andina?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Amazonas, el Orinoco y el Meta",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Magdalena, el Cauca y el Caquetá",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Sinú, el Atrato y el San Jorge",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Guaviare, el Vichada y el Tomo",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Qué fenómeno geológico dio origen a las cordilleras de la Región Andina?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "La erosión eólica y fluvial",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El vulcanismo submarino reciente",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Colisiones de las placas tectónicas",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La sedimentación de cuencas oceánicas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FAUNA",
			Text:        "¿Cuántas especies de aves se encuentran registradas en la Región Andina?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Más de 1.900 especies",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Aproximadamente 500 especies",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cerca de 3.000 especies",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Menos de 1.000 especies",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FLORA",
			Text:        "¿Qué porcentaje aproximado de las especies de plantas en la Región Andina son endémicas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Alrededor del 40%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Alrededor del 60%",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Alrededor del 25%",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Alrededor del 80%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Cuál es la función ecológica principal de los páramos según el texto?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Servir como barreras contra los vientos alisios",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Actuar como zonas de cultivo intensivo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Ser fábricas de agua y esponjas naturales",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Proveer madera para la industria",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DEMOGRAFÍA",
			SubCategory: "POBLACIÓN",
			Text:        "¿Cuántos habitantes tiene aproximadamente la Región Andina y qué porcentaje del total nacional representa?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "10 millones de habitantes, el 25% del total nacional",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "25 millones de habitantes, el 60% del total nacional",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "18 millones de habitantes, el 40,5% del total nacional",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "15 millones de habitantes, el 35% del total nacional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "DIVISION POLITICA",
			Text:        "Según el mapa de departamentos y capitales, ¿cuál es la capital del departamento de Norte de Santander?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Bucaramanga",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tunja",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cúcuta",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Ibagué",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ECONOMÍA",
			SubCategory: "AGRICULTURA",
			Text:        "¿Cuáles son algunos de los cultivos agrícolas más comunes mencionados para la Región Andina?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "Palma de aceite, arroz y algodón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Café, maíz, papa, trigo y fríjol",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Sorgo, soya y caña de azúcar exclusivamente",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Banano, plátano y cacao",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Qué cañones se mencionan como ejemplos de formaciones geológicas en la región?",
			Difficulty:  2,
			Points:      10,
			Explanation: "",
			Choices: []Choice{
				{
					Text:      "El Cañón del Colorado y el Cañón del Atuel",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Cañón del Chicamocha, el Cañón del Río Combeima y el Cañón del Río Güejar",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Cañón de la Cristalina y el Cañón del Pato",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El Cañón del Sumapaz y el Cañón del Duda",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "FRONTERAS",
			Text:        "¿Con qué países limita Colombia al oriente, según su ubicación geográfica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Al oriente, Colombia limita con Venezuela (departamentos de La Guajira, Cesar, Norte de Santander, Arauca, Vichada y Guainía) y Brasil (departamentos de Amazonas, Vaupés y Guainía).",
			Choices: []Choice{
				{
					Text:      "Perú y Ecuador",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Panamá y Venezuela",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Venezuela y Brasil",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Brasil y Perú",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "REGIÓN AMAZÓNICA",
			Text:        "¿Qué porcentaje aproximado del territorio nacional de Colombia representa la Región Amazónica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Región Amazónica de Colombia es una vasta área que representa aproximadamente el 42% del territorio nacional.",
			Choices: []Choice{
				{
					Text:      "25%",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "42%",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "60%",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "10%",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Qué nombre reciben las amplias llanuras formadas por los ríos de la región amazónica que se inundan durante la temporada de lluvias?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los ríos forman amplias llanuras inundables conocidas como \"varzeas\" e \"igapós\", las cuales se inundan en temporada de lluvias y sustentan una gran biodiversidad.",
			Choices: []Choice{
				{
					Text:      "Mesetas y colinas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Varzeas e igapós",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Serranías y valles",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cumbres andinas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FAUNA",
			Text:        "¿Cuál es considerado el pez de agua dulce más grande del mundo que habita en los ríos de la región amazónica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Entre los peces más notables de la región se incluye el pirarucú, considerado uno de los peces de agua dulce más grandes del mundo.",
			Choices: []Choice{
				{
					Text:      "La piraña",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La anguila eléctrica",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El pirarucú",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El delfín rosado",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FLORA",
			Text:        "¿Cuáles son algunas de las plantas medicinales destacadas que se encuentran en la región amazónica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Algunas de las plantas medicinales más destacadas de la zona son la ayahuasca, el sangre de grado, la uña de gato, el copaiba y la múcura.",
			Choices: []Choice{
				{
					Text:      "La manzanilla y el eucalipto",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La ayahuasca, sangre de grado y uña de gato",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El romero y la hierbabuena",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "El girasol y la orquídea",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA POLÍTICA",
			SubCategory: "DEPARTAMENTOS",
			Text:        "¿Cuál es la capital del departamento del Vaupés?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Según el mapa político y la lista de departamentos y capitales de la región amazónica, la capital del departamento del Vaupés es Mitú.",
			Choices: []Choice{
				{
					Text:      "Leticia",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Florencia",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Mitú",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Puerto Inírida",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DEMOGRAFÍA",
			SubCategory: "ETNIAS",
			Text:        "¿Cuáles son algunos de los grupos étnicos indígenas representativos que habitan la región amazónica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Algunos de los grupos étnicos más representativos de la región incluyen los Tikuna, Uitoto, Cubeo, Desano y Tucano.",
			Choices: []Choice{
				{
					Text:      "Wayuu y Emberá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Tikuna, Uitoto y Cubeo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Guambiano y Paez",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Zenú y Muisca",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "HIDROGRAFÍA",
			Text:        "¿Cuáles son los principales afluentes del río Amazonas mencionados como vitales para la región?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los principales ríos son el Amazonas y sus afluentes, como el río Putumayo, el río Caquetá (también conocido como Japurá), el río Guaviare y el río Vaupés.",
			Choices: []Choice{
				{
					Text:      "Río Magdalena y Río Cauca",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Río Orinoco y Río Meta",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Río Putumayo, Río Caquetá y Río Guaviare",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Río Sinú y Río Atrato",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "BIODIVERSIDAD",
			SubCategory: "FLORA",
			Text:        "¿Qué característica define a la vegetación de la selva ecuatorial húmeda en la Amazonía?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Se caracteriza por tener una densa vegetación compuesta por árboles altos y frondosos que forman un dosel cerrado, creando un ambiente húmedo y sombreado.",
			Choices: []Choice{
				{
					Text:      "Vegetación baja y dispersa tipo sabana",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Árboles altos y frondosos que forman un dosel cerrado",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Bosques secos con árboles espinosos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Praderas extensas sin árboles",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "GEOGRAFÍA",
			SubCategory: "RELIEVE",
			Text:        "¿Qué serranía se destaca en la región amazónica como un ejemplo de cadena montañosa más baja comparada con los Andes?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La región alberga serranías con colinas y elevaciones suaves, siendo la Serranía de La Macarena un ejemplo destacado de este tipo de formación.",
			Choices: []Choice{
				{
					Text:      "Serranía del Perijá",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Serranía de La Macarena",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Serranía del Baudó",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sierra Nevada de Santa Marta",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHO CONSTITUCIONAL",
			SubCategory: "MECANISMOS DE PROTECCIÓN",
			Text:        "¿Cuál es el mecanismo judicial idóneo para proteger de manera inmediata los derechos fundamentales cuando estos son amenazados por el Estado o un tercero?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La protección judicial de los derechos fundamentales se realiza a través de la acción de tutela, la cual es un procedimiento especial, rápido y preferente",
			Choices: []Choice{
				{
					Text:      "Acción de Cumplimiento",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Acción de Tutela",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Acción Popular",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Habeas Corpus",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHO CONSTITUCIONAL",
			SubCategory: "MECANISMOS DE PROTECCIÓN",
			Text:        "¿Qué cantidad mínima de personas se requiere para ejercer una Acción de Grupo con el fin de reparar un daño causado?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Para ejercer la acción de grupo se requiere un conjunto de mínimo 20 personas a las que se les haya causado el daño",
			Choices: []Choice{
				{
					Text:      "10 personas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "20 personas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "50 personas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "5 personas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHO CONSTITUCIONAL",
			SubCategory: "MECANISMOS DE PROTECCIÓN",
			Text:        "¿Cuál es el término establecido por la Ley 1755 de 2015 para que la administración resuelva consultas en relación con las materias a su cargo mediante el derecho de petición?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Ley 1755 de 2015 estableció un término de 30 días hábiles para consultas en relación con las materias a cargo",
			Choices: []Choice{
				{
					Text:      "10 días hábiles",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "15 días hábiles",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "20 días hábiles",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "30 días hábiles",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "PARTICIPACIÓN CIUDADANA",
			SubCategory: "MECANISMOS DE PARTICIPACIÓN",
			Text:        "¿Quién es el encargado de convocar al pueblo en un plebiscito para apoyar o rechazar una decisión del Ejecutivo?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El plebiscito es un mecanismo de participación donde el pueblo es convocado por el Presidente de la República",
			Choices: []Choice{
				{
					Text:      "El Congreso de la República",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Registrador Nacional",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Presidente de la República",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Consejo Nacional Electoral",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "PARTICIPACIÓN CIUDADANA",
			SubCategory: "MECANISMOS DE PARTICIPACIÓN",
			Text:        "¿Qué porcentaje de firmas se requiere para sustentar la solicitud de revocatoria del mandato de un gobernador o alcalde?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Se debe cumplir con el porcentaje de firmas mínimo requerido, el cual debe ser superior o igual al 40% de los votos que obtuvo el elegido",
			Choices: []Choice{
				{
					Text:      "30% de los votos obtenidos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "40% de los votos obtenidos",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "50% de los votos obtenidos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "25% de los votos obtenidos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ESTRUCTURA DEL ESTADO",
			SubCategory: "RAMA LEGISLATIVA",
			Text:        "¿Qué corporación pública representa a la Rama Legislativa y tiene la función de reformar la Constitución y hacer las leyes?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Rama Legislativa está representada por una corporación pública colegiada de elección popular denominada Congreso de la República",
			Choices: []Choice{
				{
					Text:      "El Consejo de Estado",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Corte Suprema de Justicia",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Congreso de la República",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La Asamblea Nacional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ESTRUCTURA DEL ESTADO",
			SubCategory: "ORGANISMOS DE CONTROL",
			Text:        "¿Cuál es el máximo órgano de control fiscal del Estado colombiano encargado de vigilar el buen uso de los recursos públicos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Contraloría General de la República es el máximo órgano de control fiscal del estado colombiano",
			Choices: []Choice{
				{
					Text:      "La Procuraduría General de la Nación",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La Defensoría del Pueblo",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La Fiscalía General de la Nación",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Contraloría General de la República",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ESTRUCTURA DEL ESTADO",
			SubCategory: "ÓRGANOS AUTÓNOMOS",
			Text:        "¿Qué entidad tiene como funciones básicas regular la moneda, los cambios internacionales y el crédito?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El Banco de la República es la entidad encargada de ejercer las funciones de banca central, incluyendo regular la moneda, los cambios internacionales y el crédito",
			Choices: []Choice{
				{
					Text:      "La Comisión Nacional del Servicio Civil",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Banco de la República",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Ministerio de Hacienda",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La Superintendencia Financiera",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ESTRUCTURA DEL ESTADO",
			SubCategory: "RAMA JUDICIAL",
			Text:        "¿Qué acción constitucional tutela la libertad personal cuando alguien es privado de la libertad con violación de garantías?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El habeas corpus es un derecho fundamental y una acción constitucional que tutela la libertad personal cuando alguien es privado de la libertad con violación de las garantías",
			Choices: []Choice{
				{
					Text:      "Acción de Tutela",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Habeas Corpus",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Acción de Grupo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Acción de Cumplimiento",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ESTRUCTURA DEL ESTADO",
			SubCategory: "RAMA EJECUTIVA",
			Text:        "¿Quién simboliza la unidad nacional y actúa como Jefe de Estado, Jefe de Gobierno y suprema autoridad administrativa?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Rama Ejecutiva está en cabeza del presidente de la república, quién simboliza la unidad nacional, actúa como jefe de Estado, jefe de Gobierno y suprema autoridad administrativa",
			Choices: []Choice{
				{
					Text:      "El Presidente del Senado",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Fiscal General",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Presidente de la República",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Ministro del Interior",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CONSTITUCIÓN",
			SubCategory: "PRINCIPIOS FUNDAMENTALES",
			Text:        "¿Qué significa que Colombia sea un Estado social de derecho en relación con los derechos humanos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El principio de Estado social de derecho significa que el Estado asegura la efectividad de los derechos humanos, respetándolos, protegiéndolos de intromisiones de terceros y garantizándolos a través de autoridades públicas",
			Choices: []Choice{
				{
					Text:      "Significa que el Estado asegura la efectividad de los derechos humanos, respetándolos, protegiéndolos y garantizándolos",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Significa que los derechos humanos dependen exclusivamente de la voluntad de las empresas privadas",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Significa que el Estado solo protege los derechos económicos de los ciudadanos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Significa que los derechos humanos son opcionales para las autoridades públicas",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "ORGANIZACIÓN DEL ESTADO",
			SubCategory: "DESCENTRALIZACIÓN",
			Text:        "¿En qué consiste la descentralización por colaboración según el texto?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La descentralización por colaboración se presenta cuando personas privadas ejercen funciones administrativas, como por ejemplo las Cámaras de Comercio y la Federación Nacional de Cafeteros",
			Choices: []Choice{
				{
					Text:      "En el otorgamiento de competencias administrativas a los departamentos y municipios",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "En la asignación de funciones del Estado a entidades creadas para una actividad especializada",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "En que personas privadas ejercen funciones administrativas",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "En la distribución de recursos únicamente a las grandes ciudades",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CONSTITUCIÓN",
			SubCategory: "SOBERANÍA",
			Text:        "¿A través de qué mecanismo se expresa la autoridad que reside en el pueblo según la soberanía popular?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La soberanía popular hace referencia a que la autoridad reside exclusivamente en el pueblo y se expresa a través del voto, permitiendo tomar decisiones de interés público y elegir representantes",
			Choices: []Choice{
				{
					Text:      "A través de las encuestas de opinión",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "A través de la designación directa por parte del presidente",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "A través del voto",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "A través de las juntas directivas empresariales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CONSTITUCIÓN",
			SubCategory: "JERARQUÍA NORMATIVA",
			Text:        "¿Qué ocurre en caso de incompatibilidad entre la Constitución y otras normas jurídicas?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Debido al principio de supremacía de la Constitución, que la señala como norma de normas, en caso de incompatibilidad entre ella y otras normas jurídicas, prima la aplicación de las disposiciones constitucionales",
			Choices: []Choice{
				{
					Text:      "Se aplica la norma más reciente aprobada por el Congreso",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Prima la aplicación de las disposiciones constitucionales",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Se debe realizar una consulta popular para decidir cuál aplicar",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Prima la aplicación de la ley ordinaria sobre la Constitución",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS",
			SubCategory: "DERECHOS FUNDAMENTALES",
			Text:        "¿Cuál es la posición del ordenamiento jurídico colombiano frente a la pena de muerte?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El derecho a la vida es inviolable y, por lo tanto, en el ordenamiento jurídico colombiano la pena de muerte no es permitida en el territorio nacional",
			Choices: []Choice{
				{
					Text:      "Es permitida solo en casos de traición a la patria",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Es permitida si un juez especializado la ordena",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "No es permitida en el territorio nacional",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Es permitida únicamente en tiempos de guerra internacional",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS",
			SubCategory: "IGUALDAD",
			Text:        "¿A qué hace referencia la dimensión formal del derecho a la igualdad?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La dimensión formal del derecho a la igualdad implica que la legalidad debe ser aplicada en condiciones de equidad a todos los sujetos contra quienes se dirige",
			Choices: []Choice{
				{
					Text:      "A garantizar la paridad de oportunidades económicas entre los individuos",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "A la prohibición de discriminación por razones de sexo o raza",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "A que la legalidad debe ser aplicada en condiciones de equidad a todos los sujetos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "A la distribución idéntica de recursos monetarios a toda la población",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS",
			SubCategory: "LIBERTADES INDIVIDUALES",
			Text:        "¿Qué protege el derecho al libre desarrollo de la personalidad?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Este derecho protege la capacidad y autonomía de las personas para tomar decisiones que les permitan llevar su vida de acuerdo con sus creencias y definir sus opciones vitales",
			Choices: []Choice{
				{
					Text:      "La capacidad de las personas para definir autónomamente sus opciones vitales",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "La obligación de seguir las tradiciones culturales de la mayoría",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El derecho a imponer las propias creencias a otros ciudadanos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "La facultad de actuar sin respetar las leyes del Estado",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS",
			SubCategory: "LIBERTAD DE CONCIENCIA",
			Text:        "¿Cuál es una de las garantías establecidas para la libertad de conciencia en Colombia?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Para garantizar la libertad de conciencia, se establece que nadie será obligado a actuar contra su conciencia",
			Choices: []Choice{
				{
					Text:      "Que todas las personas deben revelar sus convicciones al Estado",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Que se puede obligar a alguien a actuar contra su conciencia por el bien común",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Que nadie será obligado a actuar contra su conciencia",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Que las convicciones personales deben ser aprobadas por un juez",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS",
			SubCategory: "LIBERTAD DE LOCOMOCIÓN",
			Text:        "¿Cuándo se considera vulnerado el derecho a la libertad de locomoción?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Aunque no es absoluto, este derecho se vulnera cuando se impide el tránsito de una persona en espacios de carácter público que deben ser accesibles para todos",
			Choices: []Choice{
				{
					Text:      "Cuando se cobra peaje en una vía concesionada",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cuando se impide el tránsito de una persona en espacios de carácter público",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Cuando se prohíbe el ingreso a una propiedad privada residencial",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cuando se exige pasaporte para salir del país",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS",
			SubCategory: "DERECHO AL TRABAJO",
			Text:        "¿Cómo define el texto el derecho al trabajo en relación con la elección de la actividad?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El trabajo consiste en la realización de una actividad libremente escogida por la persona, dedicando a ella su esfuerzo intelectual o material",
			Choices: []Choice{
				{
					Text:      "Como la obligación de aceptar cualquier empleo ofrecido por el Estado",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Como la realización de una actividad asignada por las autoridades locales",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Como la realización de una actividad libremente escogida por la persona",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Como el deber de trabajar exclusivamente en sectores industriales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "DERECHOS",
			SubCategory: "ASOCIACIÓN SINDICAL",
			Text:        "¿Qué establece el artículo 39 de la Constitución respecto a la constitución de sindicatos?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El artículo 39 dispone que todos los trabajadores y empleadores tienen derecho a constituir asociaciones o sindicatos sin la intervención del Estado",
			Choices: []Choice{
				{
					Text:      "Que se requiere permiso previo del Ministerio de Trabajo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Que tienen derecho a constituir asociaciones o sindicatos sin intervención estatal",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Que solo los trabajadores del sector público pueden formar sindicatos",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Que los empleadores no tienen derecho a asociarse",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "ARTES PLÁSTICAS",
			Text:        "¿Qué material utilizaban principalmente los Quimbaya para crear figuras estilizadas con gran detalle y precisión?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los Quimbaya son conocidos por su habilidad en la orfebrería, creando figuras en oro trabajadas con técnicas como la cera perdida y la tumbaga",
			Choices: []Choice{
				{
					Text:      "Piedra volcánica",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Oro",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Arcilla roja",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Madera tallada",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "ARTES PLÁSTICAS",
			Text:        "¿Qué característica define el estilo artístico del pintor y escultor Fernando Botero?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El estilo de Fernando Botero se caracteriza por jugar con la magnificación del volumen de los elementos, dotándolos de una belleza distinta, lo cual se convirtió en su marca personal",
			Choices: []Choice{
				{
					Text:      "El uso exclusivo de figuras geométricas",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La representación abstracta del paisaje",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "La magnificación del volumen de los elementos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El uso de colores monocromáticos",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "TEATRO",
			Text:        "¿Qué institución cultural fundó Fanny Elisa Mikey en el año 1981?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Fanny Mikey logró consolidar una cultura del teatro al establecer un escenario propio en 1981 con la Fundación Teatro Nacional",
			Choices: []Choice{
				{
					Text:      "Teatro Experimental de Cali (TEC)",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Teatro Colón",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Fundación Teatro Nacional",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Teatro La Candelaria",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "CINE",
			Text:        "¿En qué año llegó el cine al territorio nacional colombiano?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El cine llegó a territorio nacional en el año de 1897, existiendo una disputa sobre si la primera función fue en Cartagena o Bucaramanga",
			Choices: []Choice{
				{
					Text:      "1920",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "1897",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "1954",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "1900",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "CINE",
			Text:        "¿Qué película dirigida por Ciro Guerra alcanzó reconocimiento internacional en 2015?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Películas como \"El abrazo de la serpiente\" (2015) alcanzaron reconocimiento internacional y consolidaron el cine colombiano en la escena global",
			Choices: []Choice{
				{
					Text:      "La estrategia del caracol",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Rodrigo D. No Futuro",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El abrazo de la serpiente",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "La gente de la universal",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "RADIO Y TELEVISIÓN",
			Text:        "¿Cuál era el objetivo principal de la emisora Radio Sutatenza dirigida a la población rural?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Radio Sutatenza se encargó de alfabetizar a distancia a gran parte de la población campesina dedicada a las labores agrícolas",
			Choices: []Choice{
				{
					Text:      "Transmitir radionovelas de entretenimiento",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Alfabetizar a distancia a la población campesina",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Emitir música folclórica exclusivamente",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Narrar eventos deportivos internacionales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "ARQUITECTURA",
			Text:        "¿Qué sitio arqueológico ubicado en Santa Marta se destaca por sus estructuras monumentales de terrazas y templos en piedra?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La Ciudad Perdida, ubicada en Santa Marta, es un famoso asentamiento indígena que cuenta con estructuras monumentales desarrolladas por culturas precolombinas",
			Choices: []Choice{
				{
					Text:      "Parque Arqueológico de San Agustín",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ciudad Perdida",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "El Pueblito Boyacense",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tierradentro",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "MÚSICA Y DANZA",
			Text:        "¿Qué ritmo musical se considera la base de las danzas de la región del Pacífico colombiano?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En el pacífico se cuenta con el currulao como base, aunque también se danzan otros ritmos como el pango y el andarele",
			Choices: []Choice{
				{
					Text:      "El Joropo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "El Bambuco",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "El Currulao",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "El Vallenato",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "LITERATURA",
			Text:        "¿Cuál es la obra literaria por la que Gabriel García Márquez recibió el premio Nobel en 1982?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Gabriel García Márquez recibió el premio Nobel en 1982 por su obra \"Cien años de Soledad\"",
			Choices: []Choice{
				{
					Text:      "El amor en los tiempos del cólera",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "La hojarasca",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Crónica de una muerte anunciada",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Cien años de Soledad",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "DEPORTE",
			Text:        "¿Bajo el mandato de qué presidente se creó el Instituto Colombiano de la Juventud y el Deporte (COLDEPORTES) en 1968?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Para 1968, el presidente Carlos Lleras Restrepo, por medio del Ministerio de Educación, creó a COLDEPORTES",
			Choices: []Choice{
				{
					Text:      "Gustavo Rojas Pinilla",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Carlos Lleras Restrepo",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Belisario Betancur",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "César Gaviria",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "RELIGIÓN",
			Text:        "Según las creencias del hinduismo mencionadas en el texto, ¿qué finalidad tiene el karma?",
			Difficulty:  2,
			Points:      10,
			Explanation: "En el hinduismo, el karma es una serie de reglas o conductas que uno tiene que seguir como parte de un proceso para forjar el carácter y ganar sabiduría",
			Choices: []Choice{
				{
					Text:      "Adorar a un único Dios verdadero",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Forjar el carácter y ganar sabiduría",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Memorizar textos sagrados",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Realizar peregrinaciones anuales",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA",
			Text:        "¿Cuáles son las tres influencias principales que se fusionaron para crear la gastronomía colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La gastronomía nacional es una fusión de diversas influencias indígenas, africanas y europeas",
			Choices: []Choice{
				{
					Text:      "Indígena, Asiática y Europea",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Indígena, Africana y Europea",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Africana, Árabe y Española",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Europea, Norteamericana e Indígena",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿De qué expresión inglesa proviene el nombre del plato típico de San Andrés, \"Rondón\"?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La palabra \"Rondón\" proviene del vocablo \"run down\" que se usaba para \"ir abajo\" en busca de las verduras o ingredientes",
			Choices: []Choice{
				{
					Text:      "Round on",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Run down",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Run on",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Roll down",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿Qué hierba local es indispensable en la preparación del Ajiaco santafereño para darle su sabor característico?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El ajiaco es una sopa espesa que combina diferentes tipos de papas, pollo, mazorcas y guascas, que es una hierba local",
			Choices: []Choice{
				{
					Text:      "Cilantro",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Laurel",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Guascas",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Orégano",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "En la preparación de la Mamona o ternera a la llanera, ¿cuál es el único condimento que se utiliza tradicionalmente?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Este plato consta del sacrificio de una becerra criolla a la cual no se le pone condimentos, sino sal espolvoreada y candela",
			Choices: []Choice{
				{
					Text:      "Pimienta y comino",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Ajo y cebolla",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cerveza y hierbas",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Sal",
					IsCorrect: true,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿Cómo se le llama al caramelo hecho de leche de coco y azúcar que da color y sabor al arroz con coco?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Para hacer el titote o caramelo que le da sabor y color al arroz se hace con la leche de coco rallado con el dulce del azúcar o de la panela",
			Choices: []Choice{
				{
					Text:      "Melado",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Titote",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Guarapo",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Tostado",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿Cuáles son los ingredientes principales que componen los Chicharrones de Pirarucú de la región Amazónica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Los chicharrones se preparan con filete de pirarucú condimentado y frito, y se acompañan de tacacho de plátano, fariña y casabe",
			Choices: []Choice{
				{
					Text:      "Filete de pirarucú, tacacho, fariña y casabe",
					IsCorrect: true,
					Order:     1,
				},
				{
					Text:      "Bagre frito, arroz con coco y patacón",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Trucha, papa francesa y ensalada",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Bocachico, yuca cocida y suero",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿Qué ingredientes introducidos por los españoles en el siglo XVI se incorporaron a la cocina colombiana?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Con la llegada de los españoles se introdujeron ingredientes como el trigo, el arroz, la carne de cerdo, el pollo y productos lácteos",
			Choices: []Choice{
				{
					Text:      "Maíz, yuca y frijoles",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Plátano, coco y mariscos",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Trigo, arroz, carne de cerdo y lácteos",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Papa, aguacate y ají",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿Cuáles son las dos influencias culinarias que dieron origen al Puchero Santafereño?",
			Difficulty:  2,
			Points:      10,
			Explanation: "El puchero santafereño tiene su origen milenario con la cocina inca y también en Andalucía (España)",
			Choices: []Choice{
				{
					Text:      "Cocina Maya y Francesa",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Cocina Azteca y Portuguesa",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Cocina Inca y Andalucía (España)",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Cocina Chibcha y Africana",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿Qué técnica de cocción se utiliza para preparar la Chuleta Valluna típica de la región Pacífica?",
			Difficulty:  2,
			Points:      10,
			Explanation: "La chuleta valluna se prepara apanando el lomo de cerdo con harina de trigo, huevos batidos y miga de pan, para luego freírlo",
			Choices: []Choice{
				{
					Text:      "Asado al carbón",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Sudado en leche de coco",
					IsCorrect: false,
					Order:     2,
				},
				{
					Text:      "Apanado y frito",
					IsCorrect: true,
					Order:     3,
				},
				{
					Text:      "Ahumado en hoja de plátano",
					IsCorrect: false,
					Order:     4,
				},
			},
		},
	
		{
			Category:    "CULTURA",
			SubCategory: "GASTRONOMÍA REGIONAL",
			Text:        "¿Qué ingredientes componen el \"Mute\", plato símbolo de los santanderes?",
			Difficulty:  2,
			Points:      10,
			Explanation: "Sus ingredientes incluyen maíz peto blanco, costillas de res y cerdo, patas de res, chorizo, habas, garbanzos, papas, entre otros",
			Choices: []Choice{
				{
					Text:      "Arroz, frijoles, carne molida y huevo",
					IsCorrect: false,
					Order:     1,
				},
				{
					Text:      "Maíz peto, carnes variadas (res, cerdo), garbanzos y habas",
					IsCorrect: true,
					Order:     2,
				},
				{
					Text:      "Pescado, leche de coco y ñame",
					IsCorrect: false,
					Order:     3,
				},
				{
					Text:      "Pollo, tres tipos de papa y guascas",
					IsCorrect: false,
					Order:     4,
				},
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
