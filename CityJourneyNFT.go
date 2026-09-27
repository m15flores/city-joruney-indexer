// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = abi.ConvertType
	_ = time.Tick
	_ = context.Background
)

// CityJourneyNFTMetaData contains all meta data concerning the CityJourneyNFT contract.
var CityJourneyNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseUri_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseUri\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cityData\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"latitude\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"longitude\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"fromDate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toDate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentTokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"cityName_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"latitude_\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"longitude_\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"fromDate_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toDate_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintNFT\",\"inputs\":[{\"name\":\"userAddress_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId_\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cityName_\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// CityJourneyNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CityJourneyNFTMetaData.ABI instead.
var CityJourneyNFTABI = CityJourneyNFTMetaData.ABI

// CityJourneyNFT is an auto generated Go binding around an Ethereum contract.
type CityJourneyNFT struct {
	CityJourneyNFTCaller     // Read-only binding to the contract
	CityJourneyNFTTransactor // Write-only binding to the contract
	CityJourneyNFTFilterer   // Log filterer for contract events
}

// CityJourneyNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CityJourneyNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityJourneyNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CityJourneyNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityJourneyNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CityJourneyNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CityJourneyNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CityJourneyNFTSession struct {
	Contract     *CityJourneyNFT   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CityJourneyNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CityJourneyNFTCallerSession struct {
	Contract *CityJourneyNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CityJourneyNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CityJourneyNFTTransactorSession struct {
	Contract     *CityJourneyNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CityJourneyNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CityJourneyNFTRaw struct {
	Contract *CityJourneyNFT // Generic contract binding to access the raw methods on
}

// CityJourneyNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CityJourneyNFTCallerRaw struct {
	Contract *CityJourneyNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CityJourneyNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CityJourneyNFTTransactorRaw struct {
	Contract *CityJourneyNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCityJourneyNFT creates a new instance of CityJourneyNFT, bound to a specific deployed contract.
func NewCityJourneyNFT(address common.Address, backend bind.ContractBackend) (*CityJourneyNFT, error) {
	contract, err := bindCityJourneyNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFT{CityJourneyNFTCaller: CityJourneyNFTCaller{contract: contract}, CityJourneyNFTTransactor: CityJourneyNFTTransactor{contract: contract}, CityJourneyNFTFilterer: CityJourneyNFTFilterer{contract: contract}}, nil
}

// NewCityJourneyNFTCaller creates a new read-only instance of CityJourneyNFT, bound to a specific deployed contract.
func NewCityJourneyNFTCaller(address common.Address, caller bind.ContractCaller) (*CityJourneyNFTCaller, error) {
	contract, err := bindCityJourneyNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTCaller{contract: contract}, nil
}

// NewCityJourneyNFTTransactor creates a new write-only instance of CityJourneyNFT, bound to a specific deployed contract.
func NewCityJourneyNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CityJourneyNFTTransactor, error) {
	contract, err := bindCityJourneyNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTTransactor{contract: contract}, nil
}

// NewCityJourneyNFTFilterer creates a new log filterer instance of CityJourneyNFT, bound to a specific deployed contract.
func NewCityJourneyNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CityJourneyNFTFilterer, error) {
	contract, err := bindCityJourneyNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTFilterer{contract: contract}, nil
}

// bindCityJourneyNFT binds a generic wrapper to an already deployed contract.
func bindCityJourneyNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CityJourneyNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CityJourneyNFT *CityJourneyNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CityJourneyNFT.Contract.CityJourneyNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CityJourneyNFT *CityJourneyNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.CityJourneyNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CityJourneyNFT *CityJourneyNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.CityJourneyNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CityJourneyNFT *CityJourneyNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CityJourneyNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CityJourneyNFT *CityJourneyNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CityJourneyNFT *CityJourneyNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CityJourneyNFT.Contract.BalanceOf(&_CityJourneyNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CityJourneyNFT.Contract.BalanceOf(&_CityJourneyNFT.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCaller) BaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTSession) BaseUri() (string, error) {
	return _CityJourneyNFT.Contract.BaseUri(&_CityJourneyNFT.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) BaseUri() (string, error) {
	return _CityJourneyNFT.Contract.BaseUri(&_CityJourneyNFT.CallOpts)
}

// CityData is a free data retrieval call binding the contract method 0x42afa385.
//
// Solidity: function cityData(uint256 ) view returns(string cityName, int256 latitude, int256 longitude, uint256 fromDate, uint256 toDate)
func (_CityJourneyNFT *CityJourneyNFTCaller) CityData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CityName  string
	Latitude  *big.Int
	Longitude *big.Int
	FromDate  *big.Int
	ToDate    *big.Int
}, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "cityData", arg0)

	outstruct := new(struct {
		CityName  string
		Latitude  *big.Int
		Longitude *big.Int
		FromDate  *big.Int
		ToDate    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CityName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Latitude = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Longitude = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FromDate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ToDate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CityData is a free data retrieval call binding the contract method 0x42afa385.
//
// Solidity: function cityData(uint256 ) view returns(string cityName, int256 latitude, int256 longitude, uint256 fromDate, uint256 toDate)
func (_CityJourneyNFT *CityJourneyNFTSession) CityData(arg0 *big.Int) (struct {
	CityName  string
	Latitude  *big.Int
	Longitude *big.Int
	FromDate  *big.Int
	ToDate    *big.Int
}, error) {
	return _CityJourneyNFT.Contract.CityData(&_CityJourneyNFT.CallOpts, arg0)
}

// CityData is a free data retrieval call binding the contract method 0x42afa385.
//
// Solidity: function cityData(uint256 ) view returns(string cityName, int256 latitude, int256 longitude, uint256 fromDate, uint256 toDate)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) CityData(arg0 *big.Int) (struct {
	CityName  string
	Latitude  *big.Int
	Longitude *big.Int
	FromDate  *big.Int
	ToDate    *big.Int
}, error) {
	return _CityJourneyNFT.Contract.CityData(&_CityJourneyNFT.CallOpts, arg0)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTCaller) CurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTSession) CurrentTokenId() (*big.Int, error) {
	return _CityJourneyNFT.Contract.CurrentTokenId(&_CityJourneyNFT.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) CurrentTokenId() (*big.Int, error) {
	return _CityJourneyNFT.Contract.CurrentTokenId(&_CityJourneyNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CityJourneyNFT *CityJourneyNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CityJourneyNFT *CityJourneyNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CityJourneyNFT.Contract.GetApproved(&_CityJourneyNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CityJourneyNFT.Contract.GetApproved(&_CityJourneyNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CityJourneyNFT *CityJourneyNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CityJourneyNFT *CityJourneyNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CityJourneyNFT.Contract.IsApprovedForAll(&_CityJourneyNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CityJourneyNFT.Contract.IsApprovedForAll(&_CityJourneyNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTSession) Name() (string, error) {
	return _CityJourneyNFT.Contract.Name(&_CityJourneyNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) Name() (string, error) {
	return _CityJourneyNFT.Contract.Name(&_CityJourneyNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CityJourneyNFT *CityJourneyNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CityJourneyNFT *CityJourneyNFTSession) Owner() (common.Address, error) {
	return _CityJourneyNFT.Contract.Owner(&_CityJourneyNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) Owner() (common.Address, error) {
	return _CityJourneyNFT.Contract.Owner(&_CityJourneyNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CityJourneyNFT *CityJourneyNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CityJourneyNFT *CityJourneyNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CityJourneyNFT.Contract.OwnerOf(&_CityJourneyNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CityJourneyNFT.Contract.OwnerOf(&_CityJourneyNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CityJourneyNFT *CityJourneyNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CityJourneyNFT *CityJourneyNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CityJourneyNFT.Contract.SupportsInterface(&_CityJourneyNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CityJourneyNFT.Contract.SupportsInterface(&_CityJourneyNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTSession) Symbol() (string, error) {
	return _CityJourneyNFT.Contract.Symbol(&_CityJourneyNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) Symbol() (string, error) {
	return _CityJourneyNFT.Contract.Symbol(&_CityJourneyNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CityJourneyNFT *CityJourneyNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CityJourneyNFT.Contract.TokenURI(&_CityJourneyNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CityJourneyNFT.Contract.TokenURI(&_CityJourneyNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CityJourneyNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTSession) TotalSupply() (*big.Int, error) {
	return _CityJourneyNFT.Contract.TotalSupply(&_CityJourneyNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CityJourneyNFT *CityJourneyNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _CityJourneyNFT.Contract.TotalSupply(&_CityJourneyNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.Approve(&_CityJourneyNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.Approve(&_CityJourneyNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x1ed632ca.
//
// Solidity: function mint(string cityName_, int256 latitude_, int256 longitude_, uint256 fromDate_, uint256 toDate_) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) Mint(opts *bind.TransactOpts, cityName_ string, latitude_ *big.Int, longitude_ *big.Int, fromDate_ *big.Int, toDate_ *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "mint", cityName_, latitude_, longitude_, fromDate_, toDate_)
}

// Mint is a paid mutator transaction binding the contract method 0x1ed632ca.
//
// Solidity: function mint(string cityName_, int256 latitude_, int256 longitude_, uint256 fromDate_, uint256 toDate_) returns()
func (_CityJourneyNFT *CityJourneyNFTSession) Mint(cityName_ string, latitude_ *big.Int, longitude_ *big.Int, fromDate_ *big.Int, toDate_ *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.Mint(&_CityJourneyNFT.TransactOpts, cityName_, latitude_, longitude_, fromDate_, toDate_)
}

// Mint is a paid mutator transaction binding the contract method 0x1ed632ca.
//
// Solidity: function mint(string cityName_, int256 latitude_, int256 longitude_, uint256 fromDate_, uint256 toDate_) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) Mint(cityName_ string, latitude_ *big.Int, longitude_ *big.Int, fromDate_ *big.Int, toDate_ *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.Mint(&_CityJourneyNFT.TransactOpts, cityName_, latitude_, longitude_, fromDate_, toDate_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CityJourneyNFT *CityJourneyNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.RenounceOwnership(&_CityJourneyNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.RenounceOwnership(&_CityJourneyNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.SafeTransferFrom(&_CityJourneyNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.SafeTransferFrom(&_CityJourneyNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CityJourneyNFT *CityJourneyNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.SafeTransferFrom0(&_CityJourneyNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.SafeTransferFrom0(&_CityJourneyNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CityJourneyNFT *CityJourneyNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.SetApprovalForAll(&_CityJourneyNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.SetApprovalForAll(&_CityJourneyNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.TransferFrom(&_CityJourneyNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.TransferFrom(&_CityJourneyNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CityJourneyNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CityJourneyNFT *CityJourneyNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.TransferOwnership(&_CityJourneyNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CityJourneyNFT *CityJourneyNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CityJourneyNFT.Contract.TransferOwnership(&_CityJourneyNFT.TransactOpts, newOwner)
}

// CityJourneyNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CityJourneyNFT contract.
type CityJourneyNFTApprovalIterator struct {
	Event *CityJourneyNFTApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CityJourneyNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityJourneyNFTApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CityJourneyNFTApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CityJourneyNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityJourneyNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityJourneyNFTApproval represents a Approval event raised by the CityJourneyNFT contract.
type CityJourneyNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CityJourneyNFT *CityJourneyNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CityJourneyNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTApprovalIterator{contract: _CityJourneyNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CityJourneyNFT *CityJourneyNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CityJourneyNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityJourneyNFTApproval)
				if err := _CityJourneyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CityJourneyNFT *CityJourneyNFTFilterer) ParseApproval(log types.Log) (*CityJourneyNFTApproval, error) {
	event := new(CityJourneyNFTApproval)
	if err := _CityJourneyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityJourneyNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CityJourneyNFT contract.
type CityJourneyNFTApprovalForAllIterator struct {
	Event *CityJourneyNFTApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CityJourneyNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityJourneyNFTApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CityJourneyNFTApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CityJourneyNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityJourneyNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityJourneyNFTApprovalForAll represents a ApprovalForAll event raised by the CityJourneyNFT contract.
type CityJourneyNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CityJourneyNFT *CityJourneyNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CityJourneyNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTApprovalForAllIterator{contract: _CityJourneyNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CityJourneyNFT *CityJourneyNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CityJourneyNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityJourneyNFTApprovalForAll)
				if err := _CityJourneyNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CityJourneyNFT *CityJourneyNFTFilterer) ParseApprovalForAll(log types.Log) (*CityJourneyNFTApprovalForAll, error) {
	event := new(CityJourneyNFTApprovalForAll)
	if err := _CityJourneyNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityJourneyNFTMintNFTIterator is returned from FilterMintNFT and is used to iterate over the raw logs and unpacked data for MintNFT events raised by the CityJourneyNFT contract.
type CityJourneyNFTMintNFTIterator struct {
	Event *CityJourneyNFTMintNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CityJourneyNFTMintNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityJourneyNFTMintNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CityJourneyNFTMintNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CityJourneyNFTMintNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityJourneyNFTMintNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityJourneyNFTMintNFT represents a MintNFT event raised by the CityJourneyNFT contract.
type CityJourneyNFTMintNFT struct {
	UserAddress common.Address
	TokenId     *big.Int
	CityName    string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintNFT is a free log retrieval operation binding the contract event 0x86094769a8d8b54a575d56650097d483bd4ecdeb3c8dd7ac9ac20df1e0d5ae03.
//
// Solidity: event MintNFT(address userAddress_, uint256 tokenId_, string cityName_)
func (_CityJourneyNFT *CityJourneyNFTFilterer) FilterMintNFT(opts *bind.FilterOpts) (*CityJourneyNFTMintNFTIterator, error) {

	logs, sub, err := _CityJourneyNFT.contract.FilterLogs(opts, "MintNFT")
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTMintNFTIterator{contract: _CityJourneyNFT.contract, event: "MintNFT", logs: logs, sub: sub}, nil
}

// WatchMintNFT is a free log subscription operation binding the contract event 0x86094769a8d8b54a575d56650097d483bd4ecdeb3c8dd7ac9ac20df1e0d5ae03.
//
// Solidity: event MintNFT(address userAddress_, uint256 tokenId_, string cityName_)
func (_CityJourneyNFT *CityJourneyNFTFilterer) WatchMintNFT(opts *bind.WatchOpts, sink chan<- *CityJourneyNFTMintNFT) (event.Subscription, error) {

	logs, sub, err := _CityJourneyNFT.contract.WatchLogs(opts, "MintNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityJourneyNFTMintNFT)
				if err := _CityJourneyNFT.contract.UnpackLog(event, "MintNFT", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintNFT is a log parse operation binding the contract event 0x86094769a8d8b54a575d56650097d483bd4ecdeb3c8dd7ac9ac20df1e0d5ae03.
//
// Solidity: event MintNFT(address userAddress_, uint256 tokenId_, string cityName_)
func (_CityJourneyNFT *CityJourneyNFTFilterer) ParseMintNFT(log types.Log) (*CityJourneyNFTMintNFT, error) {
	event := new(CityJourneyNFTMintNFT)
	if err := _CityJourneyNFT.contract.UnpackLog(event, "MintNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityJourneyNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CityJourneyNFT contract.
type CityJourneyNFTOwnershipTransferredIterator struct {
	Event *CityJourneyNFTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CityJourneyNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityJourneyNFTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CityJourneyNFTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CityJourneyNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityJourneyNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityJourneyNFTOwnershipTransferred represents a OwnershipTransferred event raised by the CityJourneyNFT contract.
type CityJourneyNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CityJourneyNFT *CityJourneyNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CityJourneyNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTOwnershipTransferredIterator{contract: _CityJourneyNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CityJourneyNFT *CityJourneyNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CityJourneyNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityJourneyNFTOwnershipTransferred)
				if err := _CityJourneyNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CityJourneyNFT *CityJourneyNFTFilterer) ParseOwnershipTransferred(log types.Log) (*CityJourneyNFTOwnershipTransferred, error) {
	event := new(CityJourneyNFTOwnershipTransferred)
	if err := _CityJourneyNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CityJourneyNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CityJourneyNFT contract.
type CityJourneyNFTTransferIterator struct {
	Event *CityJourneyNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CityJourneyNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CityJourneyNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CityJourneyNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CityJourneyNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CityJourneyNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CityJourneyNFTTransfer represents a Transfer event raised by the CityJourneyNFT contract.
type CityJourneyNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CityJourneyNFT *CityJourneyNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CityJourneyNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CityJourneyNFTTransferIterator{contract: _CityJourneyNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CityJourneyNFT *CityJourneyNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CityJourneyNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CityJourneyNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CityJourneyNFTTransfer)
				if err := _CityJourneyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CityJourneyNFT *CityJourneyNFTFilterer) ParseTransfer(log types.Log) (*CityJourneyNFTTransfer, error) {
	event := new(CityJourneyNFTTransfer)
	if err := _CityJourneyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
