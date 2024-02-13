package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

const (
	TEMPLATE_FILENAME = "template.tmpl"
	MODEL_DIR         = "src/model"
	SERVICE_DIR       = "src/service"
	CONTROLLER_DIR    = "src/controller"
)

func main() {
	var newCrudname string
	var newCrudnameLower string
	switch len(os.Args) {
	case 2:
		newCrudname = os.Args[1]
		newCrudnameLower = strings.ToLower(newCrudname)
	default:
		fmt.Println("Example:\n cli/1-crud.sh TestCrud")
	}

	//* -------------------------------- GEN MODEL ------------------------------- */
	model, err := os.ReadFile(fmt.Sprintf("%s/%s", MODEL_DIR, TEMPLATE_FILENAME))
	if err != nil {
		log.Println(err)
		return
	}
	model = bytes.ReplaceAll(model, []byte("Template"), []byte(newCrudname))
	if err = os.WriteFile(fmt.Sprintf("%s/%s.model.go", MODEL_DIR, newCrudnameLower), model, os.ModePerm); err != nil {
		log.Println(err)
		return
	}
	//* -------------------------------- GEN SERVICE ------------------------------- */
	service, err := os.ReadFile(fmt.Sprintf("%s/%s", SERVICE_DIR, TEMPLATE_FILENAME))
	if err != nil {
		log.Println(err)
		return
	}
	service = bytes.ReplaceAll(service, []byte("Template"), []byte(newCrudname))
	service = bytes.ReplaceAll(service, []byte("template"), []byte(newCrudnameLower))
	if err = os.WriteFile(fmt.Sprintf("%s/%s.service.go", SERVICE_DIR, newCrudnameLower), service, os.ModePerm); err != nil {
		log.Println(err)
		return
	}
	//* -------------------------------- GEN CONTROLLER ------------------------------- */
	controller, err := os.ReadFile(fmt.Sprintf("%s/%s", CONTROLLER_DIR, TEMPLATE_FILENAME))
	if err != nil {
		log.Println(err)
		return
	}
	controller = bytes.ReplaceAll(controller, []byte("Template"), []byte(newCrudname))
	controller = bytes.ReplaceAll(controller, []byte("template"), []byte(newCrudnameLower))
	if err := os.WriteFile(fmt.Sprintf("%s/%s.controller.go", CONTROLLER_DIR, newCrudnameLower), controller, os.ModePerm); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Done, Now Add \"MongoCollection_%s\" in src/model/enum/mongo.go package\n", newCrudname)
	fmt.Printf("      And Add this to main.go\n%s\n", aurora.Green(fmt.Sprintf("controller.New%sController(apiV1)", newCrudname)))
}
