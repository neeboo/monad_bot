// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frontrunner

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FrontrunnerParticipantData is an auto generated low-level Go binding around an user-defined struct.
type FrontrunnerParticipantData struct {
	Address common.Address
	Wins    *big.Int
	Losses  *big.Int
}

// FrontrunnerMetaData contains all meta data concerning the Frontrunner contract.
var FrontrunnerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"frontrun\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getScore\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFrontrunner.ParticipantData\",\"components\":[{\"name\":\"Address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"Wins\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Losses\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getScores\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structFrontrunner.ParticipantData[]\",\"components\":[{\"name\":\"Address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"Wins\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Losses\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"}]",
}

// FrontrunnerABI is the input ABI used to generate the binding from.
// Deprecated: Use FrontrunnerMetaData.ABI instead.
var FrontrunnerABI = FrontrunnerMetaData.ABI

// Frontrunner is an auto generated Go binding around an Ethereum contract.
type Frontrunner struct {
	FrontrunnerCaller     // Read-only binding to the contract
	FrontrunnerTransactor // Write-only binding to the contract
	FrontrunnerFilterer   // Log filterer for contract events
}

// FrontrunnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FrontrunnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrontrunnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FrontrunnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrontrunnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FrontrunnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrontrunnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FrontrunnerSession struct {
	Contract     *Frontrunner      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FrontrunnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FrontrunnerCallerSession struct {
	Contract *FrontrunnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FrontrunnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FrontrunnerTransactorSession struct {
	Contract     *FrontrunnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FrontrunnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FrontrunnerRaw struct {
	Contract *Frontrunner // Generic contract binding to access the raw methods on
}

// FrontrunnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FrontrunnerCallerRaw struct {
	Contract *FrontrunnerCaller // Generic read-only contract binding to access the raw methods on
}

// FrontrunnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FrontrunnerTransactorRaw struct {
	Contract *FrontrunnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFrontrunner creates a new instance of Frontrunner, bound to a specific deployed contract.
func NewFrontrunner(address common.Address, backend bind.ContractBackend) (*Frontrunner, error) {
	contract, err := bindFrontrunner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Frontrunner{FrontrunnerCaller: FrontrunnerCaller{contract: contract}, FrontrunnerTransactor: FrontrunnerTransactor{contract: contract}, FrontrunnerFilterer: FrontrunnerFilterer{contract: contract}}, nil
}

// NewFrontrunnerCaller creates a new read-only instance of Frontrunner, bound to a specific deployed contract.
func NewFrontrunnerCaller(address common.Address, caller bind.ContractCaller) (*FrontrunnerCaller, error) {
	contract, err := bindFrontrunner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FrontrunnerCaller{contract: contract}, nil
}

// NewFrontrunnerTransactor creates a new write-only instance of Frontrunner, bound to a specific deployed contract.
func NewFrontrunnerTransactor(address common.Address, transactor bind.ContractTransactor) (*FrontrunnerTransactor, error) {
	contract, err := bindFrontrunner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FrontrunnerTransactor{contract: contract}, nil
}

// NewFrontrunnerFilterer creates a new log filterer instance of Frontrunner, bound to a specific deployed contract.
func NewFrontrunnerFilterer(address common.Address, filterer bind.ContractFilterer) (*FrontrunnerFilterer, error) {
	contract, err := bindFrontrunner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FrontrunnerFilterer{contract: contract}, nil
}

// bindFrontrunner binds a generic wrapper to an already deployed contract.
func bindFrontrunner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FrontrunnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Frontrunner *FrontrunnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Frontrunner.Contract.FrontrunnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Frontrunner *FrontrunnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Frontrunner.Contract.FrontrunnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Frontrunner *FrontrunnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Frontrunner.Contract.FrontrunnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Frontrunner *FrontrunnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Frontrunner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Frontrunner *FrontrunnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Frontrunner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Frontrunner *FrontrunnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Frontrunner.Contract.contract.Transact(opts, method, params...)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address _address) view returns((address,uint256,uint256))
func (_Frontrunner *FrontrunnerCaller) GetScore(opts *bind.CallOpts, _address common.Address) (FrontrunnerParticipantData, error) {
	var out []interface{}
	err := _Frontrunner.contract.Call(opts, &out, "getScore", _address)

	if err != nil {
		return *new(FrontrunnerParticipantData), err
	}

	out0 := *abi.ConvertType(out[0], new(FrontrunnerParticipantData)).(*FrontrunnerParticipantData)

	return out0, err

}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address _address) view returns((address,uint256,uint256))
func (_Frontrunner *FrontrunnerSession) GetScore(_address common.Address) (FrontrunnerParticipantData, error) {
	return _Frontrunner.Contract.GetScore(&_Frontrunner.CallOpts, _address)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address _address) view returns((address,uint256,uint256))
func (_Frontrunner *FrontrunnerCallerSession) GetScore(_address common.Address) (FrontrunnerParticipantData, error) {
	return _Frontrunner.Contract.GetScore(&_Frontrunner.CallOpts, _address)
}

// GetScores is a free data retrieval call binding the contract method 0x1ea1380c.
//
// Solidity: function getScores() view returns((address,uint256,uint256)[])
func (_Frontrunner *FrontrunnerCaller) GetScores(opts *bind.CallOpts) ([]FrontrunnerParticipantData, error) {
	var out []interface{}
	err := _Frontrunner.contract.Call(opts, &out, "getScores")

	if err != nil {
		return *new([]FrontrunnerParticipantData), err
	}

	out0 := *abi.ConvertType(out[0], new([]FrontrunnerParticipantData)).(*[]FrontrunnerParticipantData)

	return out0, err

}

// GetScores is a free data retrieval call binding the contract method 0x1ea1380c.
//
// Solidity: function getScores() view returns((address,uint256,uint256)[])
func (_Frontrunner *FrontrunnerSession) GetScores() ([]FrontrunnerParticipantData, error) {
	return _Frontrunner.Contract.GetScores(&_Frontrunner.CallOpts)
}

// GetScores is a free data retrieval call binding the contract method 0x1ea1380c.
//
// Solidity: function getScores() view returns((address,uint256,uint256)[])
func (_Frontrunner *FrontrunnerCallerSession) GetScores() ([]FrontrunnerParticipantData, error) {
	return _Frontrunner.Contract.GetScores(&_Frontrunner.CallOpts)
}

// Frontrun is a paid mutator transaction binding the contract method 0x0c60e091.
//
// Solidity: function frontrun() returns()
func (_Frontrunner *FrontrunnerTransactor) Frontrun(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Frontrunner.contract.Transact(opts, "frontrun")
}

// Frontrun is a paid mutator transaction binding the contract method 0x0c60e091.
//
// Solidity: function frontrun() returns()
func (_Frontrunner *FrontrunnerSession) Frontrun() (*types.Transaction, error) {
	return _Frontrunner.Contract.Frontrun(&_Frontrunner.TransactOpts)
}

// Frontrun is a paid mutator transaction binding the contract method 0x0c60e091.
//
// Solidity: function frontrun() returns()
func (_Frontrunner *FrontrunnerTransactorSession) Frontrun() (*types.Transaction, error) {
	return _Frontrunner.Contract.Frontrun(&_Frontrunner.TransactOpts)
}
