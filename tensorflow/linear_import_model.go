package main

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {

	modelDir := "/Users/pujie/temp/test_linear"
	modelFile := filepath.Join(modelDir, "linear_model.pb")
	model, err := ioutil.ReadFile(modelFile)
	if err != nil {
		log.Fatal(err)
	}
	// 创建图 加载模型
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		log.Fatal("加载模型异常：", err.Error())
	}
	// Create a session for inference over graph.
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	inputData1 := [1][3]float32{{3.3, 4.4, 5.5}}
	//inputData2 := [1][2]float32{{34.2}}
	tensor1, err := tf.NewTensor(inputData1)
	//tensor2, err := tf.NewTensor(inputData2)
	if err != nil {
		fmt.Printf("Error NewTensor: err: %s", err.Error())
		return
	}

	result, err := session.Run(
		map[tf.Output]*tf.Tensor{
			graph.Operation("Placeholder_2").Output(0): tensor1,
			//graph.Operation("Placeholder_1").Output(0): tensor2,
		},
		[]tf.Output{
			graph.Operation("Add").Output(0),
		},
		nil)
	if err != nil {
		fmt.Printf("Error running the session with input, err: %s  ", err.Error())
		return
	}
	// 输出结果，interface{}格式
	fmt.Printf("Result value: %v \n", result[0].Value())
}
