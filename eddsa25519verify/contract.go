// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package eddsa25519verify

import (
	"errors"
	"fmt"
	"math/big"
	"crypto/ed25519"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/vmerrs"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	VerifySignatureGasCost uint64 = 1 /* SET A GAS COST HERE */
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = abi.JSON
	_ = errors.New
	_ = big.NewInt
	_ = vmerrs.ErrOutOfGas
	_ = common.Big0
)

// Singleton StatefulPrecompiledContract and signatures.
var (

	// Eddsa25519VerifyRawABI contains the raw ABI of Eddsa25519Verify contract.
	//go:embed contract.abi
	Eddsa25519VerifyRawABI string

	Eddsa25519VerifyABI = contract.ParseABI(Eddsa25519VerifyRawABI)

	Eddsa25519VerifyPrecompile = createEddsa25519VerifyPrecompile()
)

type VerifySignatureInput struct {
	PublicKey []byte
	Message   string
	Signature []byte
}

// UnpackVerifySignatureInput attempts to unpack [input] as VerifySignatureInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackVerifySignatureInput(input []byte) (VerifySignatureInput, error) {
	inputStruct := VerifySignatureInput{}
	err := Eddsa25519VerifyABI.UnpackInputIntoInterface(&inputStruct, "verifySignature", input)

	return inputStruct, err
}

// PackVerifySignature packs [inputStruct] of type VerifySignatureInput into the appropriate arguments for verifySignature.
func PackVerifySignature(inputStruct VerifySignatureInput) ([]byte, error) {
	return Eddsa25519VerifyABI.Pack("verifySignature", inputStruct.PublicKey, inputStruct.Message, inputStruct.Signature)
}

// PackVerifySignatureOutput attempts to pack given isValid of type bool
// to conform the ABI outputs.
func PackVerifySignatureOutput(isValid bool) ([]byte, error) {
	return Eddsa25519VerifyABI.PackOutput("verifySignature", isValid)
}

// UnpackVerifySignatureOutput attempts to unpack given [output] into the bool type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackVerifySignatureOutput(output []byte) (bool, error) {
	res, err := Eddsa25519VerifyABI.Unpack("verifySignature", output)
	if err != nil {
		return false, err
	}
	unpacked := *abi.ConvertType(res[0], new(bool)).(*bool)
	return unpacked, nil
}

func verifySignature(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, VerifySignatureGasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the VerifySignatureInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackVerifySignatureInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	_ = inputStruct // CUSTOM CODE OPERATES ON INPUT

	var output bool // CUSTOM CODE FOR AN OUTPUT
	output = ed25519.Verify(inputStruct.PublicKey, []byte(inputStruct.Message), inputStruct.Signature)
	packedOutput, err := PackVerifySignatureOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createEddsa25519VerifyPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createEddsa25519VerifyPrecompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"verifySignature": verifySignature,
	}

	for name, function := range abiFunctionMap {
		method, ok := Eddsa25519VerifyABI.Methods[name]
		if !ok {
			panic(fmt.Errorf("given method (%s) does not exist in the ABI", name))
		}
		functions = append(functions, contract.NewStatefulPrecompileFunction(method.ID, function))
	}
	// Construct the contract with no fallback function.
	statefulContract, err := contract.NewStatefulPrecompileContract(nil, functions)
	if err != nil {
		panic(err)
	}
	return statefulContract
}
