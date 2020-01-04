/**
修改excel；调用估价接口
**/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"ksyun.com/commons/util"
	"log"
	"net/http"
	"strconv"
)

var client *http.Client = util.DefaultClient

func main() {
	file := "/Users/pujie/Desktop/gzqd.xlsx"
	readExcel(file)
	//doEvaluate()
}

func readExcel(path string) {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}
	sheet := file.Sheets[0]
	index := 0
	for i, row := range sheet.Rows {
		if i == 0 {
			continue
		}
		if index == 10 {
			break
		}
		vId := row.Cells[5].String()
		regDate, _ := row.Cells[10].GetTime(false)
		regDateStr := regDate.Format("2006-01-02")
		pId, _ := row.Cells[3].Int()
		cId, _ := row.Cells[1].Int()
		mileage, _ := row.Cells[11].Float()
		local := NewParameter(vId, regDateStr, pId, cId, mileage)
		qingDao := NewParameter(vId, regDateStr, 370000, 370200, mileage)
		req := NewRequest()
		req.addParam(local)
		req.addParam(qingDao)
		res, err := doEvaluate(req)
		if err != nil {

		} else {
			if res.Status == 0 {
				for i, d := range res.Data {
					if d.Status == 0 {
						if i == 0 {
							writeRow(row, 13, strconv.FormatFloat(d.Data.Tc.SalePrice, 'f', 2, 64))
							writeRow(row, 14, strconv.FormatFloat(d.Data.Che300.B2cPrice, 'f', 2, 64))
						} else {
							writeRow(row, 15, strconv.FormatFloat(d.Data.Tc.SalePrice, 'f', 2, 64))
							writeRow(row, 16, strconv.FormatFloat(d.Data.Che300.B2cPrice, 'f', 2, 64))
						}
					} else {
						fmt.Printf("requestId=%s,车源：%s,mes=%s,code=%d,data=%v \n", d.RequestId, row.Cells[0].String(), d.Message, d.Status, d.Data)
					}
				}
			}
		}
		/*		for j, cell := range row.Cells {
				if j == 10 {
					t, _ := cell.GetTime(false)
					fmt.Printf("%v\t : %v", t, cell.Type())
				} else {
					fmt.Printf("%s : %v ", cell.String(), cell.Type())
				}
				fmt.Printf("%s : %v",cell)
				fmt.Println()
			}*/

		index++
	}
	err = file.Save(path)
	if err != nil {
		panic(err)
	}
}

func writeRow(row *xlsx.Row, index int, value string) {
	if index < len(row.Cells) {
		row.Cells[index].Value = value
	} else {
		row.AddCell().Value = value
	}
}

func doEvaluate(req *Request) (*BatchResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	//fmt.Printf("request:%s \n", string(data))
	httReq, err := http.NewRequest("POST", "http://uat.eval.taoche.com/api/eval/bp", bytes.NewBuffer(data))
	httReq.Header.Add("Content-Type", "application/json")
	httRes, err := client.Do(httReq)
	log.Println("1.status code = ", httRes.StatusCode)
	data, err = ioutil.ReadAll(httRes.Body)
	if err != nil {
		log.Fatal("response err:", err)
		return nil, err
	}
	//fmt.Printf("Response:%v \n", string(data))
	httpRes := &BatchResponse{}
	err = json.Unmarshal(data, &httpRes)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
		return nil, err
	}
	//fmt.Printf("json:%v \n", httpRes)
	return httpRes, nil
}

type Request struct {
	Channel  string
	UserName string
	Params   []*Parameter
}

func NewRequest() *Request {
	return &Request{Channel: "self", UserName: "pujie"}
}

func (r *Request) addParam(p *Parameter) {
	r.Params = append(r.Params, p)
}

type Parameter struct {
	VehicleModelId string
	ProvinceId     int
	CityId         int
	CarRegDate     string
	MileAge        float64
}

func NewParameter(vId, regDate string, pId, cId int, mileage float64) *Parameter {
	return &Parameter{VehicleModelId: vId, ProvinceId: pId, CityId: cId, CarRegDate: regDate, MileAge: mileage}
}

type BatchResponse struct {
	Status  int
	Message string
	Data    []Response
}

type Response struct {
	Status    int
	Message   string
	Data      Data
	RequestId string
}
type Data struct {
	Tc     TC
	Che300 Che300
}

type TC struct {
	SalePrice float64
}

type Che300 struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	B2cPrice  float64
}
