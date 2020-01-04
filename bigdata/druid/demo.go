/**
@description	kafka 简单的消费者(consumer)，支持一个消费者消费多个topic，consumer-group和consumer负载均衡
	sarama-cluster，kafka go语言客户端库。支持kafka0.8及以上版本。
@github https://github.com/bsm/sarama-clus
@author pujie
@data	2018-01-17
**/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/shunfei/godruid"
	"log"
)

type LineIterm struct {
	CarId      string  `json:"car_id"`
	SpecId     string  `json:"spec_id"`
	CarAge     string  `json:"car_age"`
	Mileage    string  `json:"mileage"`
	SalePrice  float64 `json:"sale_price"`
	DataSource string  `json:"data_source"`
	Ts         string  `json:"ts"`
}

type ResultQuery struct {
	Line []LineIterm
}

type SqlQuery struct {
	Query     string                 `json:"query"`
	Intervals []string               `json:"intervals"`
	Context   map[string]interface{} `json:"context"`

	ResultQuery []LineIterm
}

func (q *SqlQuery) setup() { q.Query = "query" }

func (q *SqlQuery) onResponse(content []byte) error {
	res := new([]LineIterm)
	err := json.Unmarshal(content, res)
	if err != nil {
		return err
	}
	q.ResultQuery = *res
	return nil
}

func main() {
	url := "http://druid-api.taoche.com"
	client := godruid.Client{Url: url, EndPoint: "/druid/v2/sql"}
	queryRow(client)
}

func queryRow(c godruid.Client) {
	sql := "SELECT car_id,spec_id, car_age,mileage,sale_price,sale_date,publish_date,data_source, __time as ts " +
		"FROM alg_car_price_detail " +
		"WHERE __time >= CURRENT_TIMESTAMP - INTERVAL '3' DAY and data_source='guazi' and sale_status =0 and spec_id=114536"
	context := map[string]interface{}{}
	context["sqlTimeZone"] = "Asia/Shanghai"
	qr := SqlQuery{Query: sql, Context: context}
	var reqJson []byte
	var err error
	reqJson, err = json.Marshal(qr)
	if err != nil {
		return
	}
	result, err := c.QueryRaw(reqJson)
	if err != nil {
		fmt.Printf("query err: %s \n", err.Error())
	} else {
		fmt.Printf("%s \n", string(result))
		res := new([]LineIterm)
		err := json.Unmarshal(result, res)
		if err != nil {
			fmt.Printf("Unmarshal err: %s \n", err.Error())
		}
		for _, line := range *res {
			fmt.Printf("%s\t%s\t%s\t%s\t%f\t%s\t%s \n", line.CarId, line.SpecId, line.CarAge,
				line.Mileage, line.SalePrice, line.DataSource, line.Ts)
		}
	}
}

func querySearch(client godruid.Client) {
	q := &godruid.QuerySearch{
		QueryType:        "scan",
		Intervals:        []string{"2019-02-26T00:00:00.000/2019-02-28T00:00:00.000"},
		DataSource:       "alg_car_price_detail",
		SearchDimensions: []string{"spec_id", "mileage"},
	}
	err := client.Query(q)
	if err != nil {
		log.Printf("druid query err: %s \n", err.Error())
	} else {
		log.Printf("大小：%d \n", len(q.QueryResult))
		r := q.QueryResult
		for _, iterm := range r {
			fmt.Printf("数据条数:%d \n", len(iterm.Result))
			for _, d := range iterm.Result {
				fmt.Printf("%s:%s; ", d.Dimension, d.Value)
			}
			fmt.Println()
		}
	}
}
