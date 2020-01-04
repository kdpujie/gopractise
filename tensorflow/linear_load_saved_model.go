package main

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	modelDir := "/Users/pujie/temp/test_linear"
	model, err := tf.LoadSavedModel(modelDir, []string{"add"}, nil)
	inputData := [2][1]float32{{1}, {2}}
	tensor, err := tf.NewTensor(inputData)
	if err != nil {
		fmt.Printf("Error NewTensor: err: %s", err.Error())
		return
	}

	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			model.Graph.Operation("input").Output(0): tensor,
		},
		[]tf.Output{
			model.Graph.Operation("output").Output(0),
		},
		nil)
	if err != nil {
		fmt.Printf("Error running the session with input, err: %s  ", err.Error())
		return
	}
	// 输出结果，interface{}格式
	fmt.Printf("Result value: %v", result[0].Value())
}
