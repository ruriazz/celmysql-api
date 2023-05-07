package services

import (
	"bytes"
	"log"

	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/celmysql-api/dto"
	"github.com/celmysql-api/entity"
	"github.com/celmysql-api/mapping"
	"github.com/go-playground/validator/v10"
)

// // A Response struct to map the Entire Response
// type Response struct {
// 	Name    string    `json:"name"`
// 	Pokemon []Pokemon `json:"pokemon_entries"`
// }

// // A Pokemon Struct to map every pokemon to.
// type Pokemon struct {
// 	EntryNo int            `json:"entry_number"`
// 	Species PokemonSpecies `json:"pokemon_species"`
// }

// // A struct to map our Pokemon's Species which includes it's name
// type PokemonSpecies struct {
// 	Name string `json:"name"`
// }

type RajaOngkirService struct {
	DB       *sql.DB
	Validate *validator.Validate
}

func NewRajaOngkirService(DB *sql.DB, validate *validator.Validate) IRajaOngkirService {
	return &RajaOngkirService{
		DB:       DB,
		Validate: validate,
	}
}

func (service *RajaOngkirService) Create(ctx context.Context, payload dto.CreateRajaOngkirDto) mapping.RajaOngkirVm {
	// err := service.Validate.Struct(request)
	// common.PanicIfError(err)

	// tx, err := service.DB.Begin()
	// common.PanicIfError(err)
	// defer common.CommitOrRollback(tx)

	// rajaOngkir := entity.RajaOngkir{
	// 	RajaOngkirCode:           &request.RajaOngkirCode,
	// 	RajaOngkirName:           &request.RajaOngkirName,
	// 	Description:         new(string),
	// 	OptimisticLockField: 0,
	// 	GCRecord:            0,
	// 	Deleted:             false,
	// 	UserInserted:        &request.UserInserted,
	// 	InsertedDate:        time.Now(),
	// 	LastUserId:          new(string),
	// 	LastUpdate:          time.Now(),
	// }

	// rajaOngkir = service.RajaOngkirRepository.Save(ctx, tx, rajaOngkir)

	// return mapping.ToRajaOngkirResponse(rajaOngkir)
	url := "https://jsonplaceholder.typicode.com/todos"
	client := http.Client{Timeout: 5 * time.Second}
	// todo := Todo{11, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(payload)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))

	request.Header.Set("Authorization", os.ExpandEnv("$BEARER_TOKEN"))
	request.Header.Set("Content-Type", "application/json") // => your content-type
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var rajaOngkir entity.RajaOngkir
	json.Unmarshal(bodyBytes, &rajaOngkir)
	fmt.Printf("%+v\n", rajaOngkir)
	return mapping.ToRajaOngkirResponse(rajaOngkir)

}

// func (service *RajaOngkirService) Update(ctx context.Context, request dto.UpdateRajaOngkirDto, oid string) mapping.RajaOngkirVm {
// 	err := service.Validate.Struct(request)
// 	common.PanicIfError(err)

// 	tx, err := service.DB.Begin()
// 	common.PanicIfError(err)
// 	defer common.CommitOrRollback(tx)

// 	rajaOngkir, err := service.RajaOngkirRepository.FindById(ctx, tx, oid)
// 	if err != nil {
// 		panic(common.NewNotFoundError(err.Error()))
// 	}

// 	rajaOngkir.RajaOngkirCode = &request.RajaOngkirCode
// 	rajaOngkir.RajaOngkirName = &request.RajaOngkirName
// 	rajaOngkir.LastUserId = &request.LastUserId

// 	rajaOngkir = service.RajaOngkirRepository.Update(ctx, tx, rajaOngkir)

// 	return mapping.ToRajaOngkirResponse(rajaOngkir)
// }

// func (service *RajaOngkirService) Delete(ctx context.Context, oid string) {
// 	tx, err := service.DB.Begin()
// 	common.PanicIfError(err)
// 	defer common.CommitOrRollback(tx)

// 	rajaOngkir, err := service.RajaOngkirRepository.FindById(ctx, tx, oid)
// 	if err != nil {
// 		panic(common.NewNotFoundError(err.Error()))
// 	}

// 	service.RajaOngkirRepository.Delete(ctx, tx, rajaOngkir)
// }

// func (service *RajaOngkirService) FindById(ctx context.Context, oid string) mapping.RajaOngkirVm {
// 	tx, err := service.DB.Begin()
// 	common.PanicIfError(err)
// 	defer common.CommitOrRollback(tx)

// 	rajaOngkir, err := service.RajaOngkirRepository.FindById(ctx, tx, oid)
// 	if err != nil {
// 		panic(common.NewNotFoundError(err.Error()))
// 	}

// 	return mapping.ToRajaOngkirResponse(rajaOngkir)
// }

func (service *RajaOngkirService) Find(ctx context.Context, criteria string) []mapping.RajaOngkirVm {
	url := "https://jsonplaceholder.typicode.com/todos"
	client := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Authorization", os.ExpandEnv("$BEARER_TOKEN"))
	request.Header.Set("Content-Type", "application/json") // => your content-type
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(responseData))
	var rajaOngkir []entity.RajaOngkir
	json.Unmarshal(responseData, &rajaOngkir)
	return mapping.ToRajaOngkirResponses(rajaOngkir)

}

// func (service *RajaOngkirService) Create(ctx context.Context, request dto.CreateRajaOngkirDto) Todo {
// 	url := "https://jsonplaceholder.typicode.com/todos"
// 	client := http.Client{Timeout: 5 * time.Second}
// 	todo := Todo{11, 2, "lorem ipsum dolor sit amet", true}
// 	jsonReq, err := json.Marshal(todo)
// 	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
// 	request.Header.Set("Authorization", os.ExpandEnv("$BEARER_TOKEN"))
// 	request.Header.Set("Content-Type", "application/json") // => your content-type
// 	response, err := client.Do(request)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(response.Body)

// 	var todoStruct Todo
// 	json.Unmarshal(bodyBytes, &todoStruct)
// 	fmt.Printf("%+v\n", todoStruct)

// 	// bodyString := string(bodyBytes)
// 	// fmt.Println(bodyString)

// 	// responseData, err := ioutil.ReadAll(response.Body)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// var todoStruct Todo

// 	// json.Unmarshal(responseData, &todoStruct)
// 	// fmt.Println(string(responseData))
// 	return todoStruct
// }

// func (service *RajaOngkirService) Update(ctx context.Context) Todo {
// 	fmt.Println("3. Performing Http Put...")
// 	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
// 	jsonReq, err := json.Marshal(todo)
// 	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to string
// 	bodyString := string(bodyBytes)
// 	fmt.Println(bodyString)

// 	// Convert response body to Todo struct
// 	var todoStruct Todo
// 	json.Unmarshal(bodyBytes, &todoStruct)
// 	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
// 	return todoStruct
// }

// func (service *RajaOngkirService) Delete(ctx context.Context) Todo {
// 	fmt.Println("4. Performing Http Delete...")
// 	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
// 	jsonReq, err := json.Marshal(todo)
// 	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to string
// 	bodyString := string(bodyBytes)
// 	fmt.Println(bodyString)

// 	var todoStruct Todo
// 	return todoStruct
// }

// func (service *RajaOngkirService) Find(ctx context.Context) Response {
// 	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
// 	common.PanicIfError(err)
// 	responseData, err := ioutil.ReadAll(response.Body)
// 	common.PanicIfError(err)
// 	fmt.Println(string(responseData))
// 	var responseObject Response
// 	json.Unmarshal(responseData, &responseObject)

// 	fmt.Println(responseObject.Name)
// 	fmt.Println(len(responseObject.Pokemon))

// 	for i := 0; i < len(responseObject.Pokemon); i++ {
// 		fmt.Println(responseObject.Pokemon[i].Species.Name)
// 	}

// 	return responseObject
// }
