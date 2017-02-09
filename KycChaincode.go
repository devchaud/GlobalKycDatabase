package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Region Chaincode implementation
type KycChaincode struct {
}

var KycIndexTxStr = "_KycIndexTxStr"

type KycData struct {
	USER_NAME           string `json:"USER_NAME"`
	USER_ID             string `json:"USER_ID"`
	KYC_BANK_NAME       string `json:"KYC_BANK_NAME"`
	KYC_CREATE_DATE     string `json:"KYC_CREATE_DATE"`
	KYC_VALID_TILL_DATE string `json:"KYC_VALID_TILL_DATE"`
	KYC_DOC_BLOB        string `json:"KYC_DOC_BLOB"`
}

func (t *KycChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error
	// Initialize the chaincode

	var EmptyKYC KycData
	jsonAsBytes, _ := json.Marshal(EmptyKYC)
	err = stub.PutState(KycIndexTxStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Deployment of KYC is completed\n")
	return nil, nil
}

// Add user KYC data in Blockchain
func (t *KycChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "InsertKycDetails" {
		// Insert User's KYC data in blockchain
		return t.InsertKycDetails(stub, args)
	}
	/*else if function == "UpdateKycDetails" {
		// Update User's KYC data in blockchain
		return t.UpdateKycDetails(stub, args)
	}*/

	return nil, errors.New("Received unknown function invocation")
}

func (t *KycChaincode) InsertKycDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var KYCDetails KycData

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Need 3 arguments")
	}

	// Initialize the chaincode
	KYCDetails.USER_NAME = args[0]
	KYCDetails.USER_ID = args[1]
	KYCDetails.KYC_BANK_NAME = args[2]
	KYCDetails.KYC_DOC_BLOB = args[3]

	jsonAsBytes, _ := json.Marshal(KYCDetails)
	stub.PutState(args[1], jsonAsBytes)

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *KycChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	return nil, nil
}

func main() {
	err := shim.Start(new(KycChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
