package main

import (
	"fmt"
)

// Clase Tarea
type Task struct {
	name        string
	descripcion string
	completed   bool
}

func (t *Task) imprimirDatos() {
	fmt.Printf("Nombre de la tarea: %s\nDescripcion de la tarea: %s\nCompletado: %t\n", t.name, t.descripcion, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

// Clase ListaDeTareas
type ListaDeTareas struct {
	listaDeTareas []Task
}

func (ldt *ListaDeTareas) addTasksToList(tareas ...*Task) {

	for _, tarea := range tareas {
		ldt.listaDeTareas = append(ldt.listaDeTareas, *tarea)
	}
}

func (ldt *ListaDeTareas) imprimeLista() {
	for indice, tarea := range ldt.listaDeTareas {
		fmt.Printf("%d - %v\n", indice, tarea)
	}
}

func (ldt *ListaDeTareas) eliminarTareas(indice int) {
	ldt.listaDeTareas = append(ldt.listaDeTareas[:indice], ldt.listaDeTareas[indice+1:]...)
}

func main() {

	tarea1 := Task{
		name:        "Limpiar la casa",
		descripcion: "Limpiar arena de luneta",
	}

	tarea2 := Task{}
	tarea2.name = "Cocinar"
	tarea2.descripcion = "Comprar ingredientes para hacer molletrón"

	newLista := ListaDeTareas{}

	newLista.addTasksToList(&tarea1, &tarea2)

	newLista.imprimeLista()

	tarea3 := Task{}
	tarea3.name = "Comprar dinos pequeña xd"
	tarea3.descripcion = "Comer dinos pekeñap"
	tarea3.markCompleted()

	lista2 := ListaDeTareas{}
	fmt.Println("==================================================")
	lista2.addTasksToList(&tarea1, &tarea3)
	lista2.imprimeLista()
	fmt.Println("==================================================")
	lista2.eliminarTareas(0)
	lista2.imprimeLista()
}
