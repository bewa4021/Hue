package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bcurren/go-hue"
)

// Power for json
type Power struct {
	On bool `json:"on"`
}

func main() {
	get()
	// fmt.Println(" ")
	// put()
	// bridge()
	// cycle()
}

func cycle() {
	var iter = 0
	// putLightURL := "http://192.168.1.2/api/7ewM79slMdQIuLy82dYY9oi4mjE0yTS8pabIjUy3/lights/17/state"
	var powerState = [8]bool{false, false, false, false, true, true, true, true}
	var lightNum = [8]int{17, 18, 19, 20, 17, 18, 19, 20}
	for iter < 8 {
		put(powerState[iter], lightNum[iter])
		iter++
		time.Sleep(500 * time.Millisecond)
	}

}

func put(state bool, light int) {
	putLightURL :=
		"http://192.168.1.2/api/7ewM79slMdQIuLy82dYY9oi4mjE0yTS8pabIjUy3/lights/" + strconv.Itoa(light) + "/state"
	var jsonStr = Power{On: state}
	data, err := json.Marshal(jsonStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)

	req, err := http.NewRequest(http.MethodPut, putLightURL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Fail")
		log.Fatal(err)
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Fail")
		log.Fatal(err)
		return
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fail")
		log.Fatal(err)
		return
	}
	bs := string(body)
	fmt.Println(bs)
}

func get() {
	userID := "7ewM79slMdQIuLy82dYY9oi4mjE0yTS8pabIjUy3"
	resp, err := http.Get("http://192.168.1.2/api/" + userID + "/lights/17")
	if err != nil {
		fmt.Println("Fail")
		log.Fatal(err)
		return
	}
	var jsonStr = Power{On: false}
	data, err := json.Marshal(jsonStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(data))
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fail")
		log.Fatal(err)
		return
	}
	bs := string(body)
	fmt.Println(bs)
}

func bridge() {
	bridges, err := hue.FindBridges()
	if err != nil {
		fmt.Println("Fail")
		log.Fatal(err)
		return
	}
	fmt.Println(bridges)
}
