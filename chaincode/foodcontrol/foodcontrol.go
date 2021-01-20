/*
Business Blockchain Training & Consulting SpA. All Rights Reserved.
www.blockchainempresarial.com
email: ricardo@blockchainempresarial.com
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for control the food
type SmartContract struct {
	contractapi.Contract
}

//defino una estructura muy basica de un activo llamado Food que tiene 
//dos propiedades
//Food describes basic details of what makes up a food
// puede crearse una estructura que sea parte de otra estructura si las
//propiedades son muy complejas
type Food struct {
	Farmer  string `json:"farmer"`
	Variety string `json:"variety"`
}

// la funcion es parte de la struct SmartContract
func (s *SmartContract) Set(ctx contractapi.TransactionContextInterface, foodId string, farmer string, variety string) error {

	//Validaciones de sintaxis

	//validaciones de negocio

	food := Food{
		Farmer:  farmer,
		Variety: variety,
	}
	//transformo a bytes el elemento food
	foodAsBytes, err := json.Marshal(food)
	if err != nil {
		fmt.Printf("Marshal error: %s", err.Error())
		return err
	}
	//llamamos al elemento  getstub que tiene un elemento putstate que 
	//me permite guardar en el ledger
	return ctx.GetStub().PutState(foodId, foodAsBytes)
	//clave foodID y el valor es el conjunto de bytes
}
//y ya tenemos la funcion que nos permite guardar en la blockchain


func (s *SmartContract) Query(ctx contractapi.TransactionContextInterface, foodId string) (*Food, error) {

	foodAsBytes, err := ctx.GetStub().GetState(foodId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if foodAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", foodId)
	}

	food := new(Food)

	err = json.Unmarshal(foodAsBytes, food)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error. %s", err.Error())
	}

	return food, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create foodcontrol chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting foodcontrol chaincode: %s", err.Error())
	}
}
