`
MVVM is a variant of MVC that separates the application into a model, a view, and a view model. The view model is a layer between the view and the model that exposes data and functionality to the view and handles user input from the view.
`

package main

import (
	"fmt"
	"log"
)

// Model represents the data and business logic of the application.
type Model struct {
	ID   int
	Name string
}

// ViewModel is a layer between the view and the model that exposes data and functionality to the view and handles user input from the view.
type ViewModel struct {
	model *Model
}

// GetName returns the name of the model.
func (vm *ViewModel) GetName() string {
	return vm.model.Name
}

// SetName updates the name of the model.
func (vm *ViewModel) SetName(name string) {
	vm.model.Name = name
}

// View is responsible for rendering the user interface.
type View struct {
	viewModel *ViewModel
}

// Render displays the view model to the user.
func (v *View) Render() {
	fmt.Println("Name:", v.viewModel.GetName())
}

func main() {
	model := &Model{ID: 123, Name: "My Entity"}
	viewModel := &ViewModel{model: model}
	view := &View{viewModel: viewModel}

	if err := view.Render(); err != nil {
		log.Fatal(err)
	}

	if err := viewModel.SetName("New Name"); err != nil {
		log.Fatal(err)
	}

	if err := view.Render(); err != nil {
		log.Fatal(err)
	}
}
