package main

import (
	"fmt"
	"os"

	"log"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"gopkg.in/yaml.v2"
)

func main() {
	// Read the YAML file content
	yamlData, err := os.ReadFile("./deployment.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Unmarshal the YAML data into a map[string]interface{}
	var deployment1 map[string]interface{}
	if err := yaml.Unmarshal(yamlData, &deployment1); err != nil {
		log.Fatalf("Error unmarshalling YAML data: %v", err)
	}

	// Create a CEL environment
	env, err := cel.NewEnv(
		cel.Variable("spec", types.NewMapType(cel.StringType, cel.AnyType)),
		cel.Variable("status", types.NewMapType(cel.StringType, cel.AnyType)),
	)
	if err != nil {
		log.Fatalf("Environment creation error: %v\n", err)
	}

	// Define the CEL expression to check if spec.replicas is equal to status.readyReplicas
	expression := "spec.replicas == status.readyReplicas"

	// Compile the expression
	ast, iss := env.Compile(expression)
	// Check for compilation errors
	if iss.Err() != nil {
		log.Fatalln(iss.Err())
	}

	// Create a program from the compiled expression
	prg, err := env.Program(ast)
	if err != nil {
		log.Fatalln(err)
	}

	// Evaluate the expression with the deployment variables
	out, _, err := prg.Eval(map[string]interface{}{
		"spec":   deployment1["spec"],
		"status": deployment1["status"],
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Print the result
	fmt.Println(out)
}
