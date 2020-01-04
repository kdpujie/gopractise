package main

func main() {

	//StartKsMobApi()
	//Start_Type_func()
	JsonStart()
	/*
		app := &baidu.App{
			AppId:      proto.String("appId_af23eadaf342af234"),
			AppVersion: &baidu.Version{Major: proto.Uint32(5), Minor: proto.Uint32(2)},
		}
		request := &baidu.MobadsRequest{
			RequestId:  proto.String("requestId-123asdfafasdf"),
			ApiVersion: &baidu.Version{Major: proto.Uint32(2), Minor: proto.Uint32(1)},
			App:        app,
		}
		//序列化
		data, err := proto.Marshal(request)
		if err != nil {
			log.Fatal("marshaling error:", err)
		}
		response, _ := http.Post("http://127.0.0.1:8080/api", "application/octet-stream", bytes.NewBuffer(data))
		data, _ = ioutil.ReadAll(response.Body)
		log.Println(string(data))
	*/
	/**
	t := &pb.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3, 4, 5},
		Optionalgroup: &pb.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}

	//序列化
	data, err := proto.Marshal(t)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	log.Println("1. 序列化后:", string(data))
	//反序列化:
	newTest := &pb.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}
	if t.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", t.GetLabel(), newTest.GetLabel())
	}
	newTest.Label = proto.String("hello,world.")
	log.Printf("2. 反序列化后:Label=%q , Type=%d,reps=%d", newTest.GetLabel(), newTest.GetType(), newTest.GetReps())
	**/
}
