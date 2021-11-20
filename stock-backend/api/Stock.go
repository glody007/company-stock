// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addresse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"functionCalled\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nbrOwners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountEther\",\"type\":\"uint256\"}],\"name\":\"sendEtherTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendStockTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stockQte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"stockQteOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c0604052601460808190527f436f6d70616e792053746f636b204f7074696f6e00000000000000000000000060a090815261003e91600191906100f8565b506040805180820190915260038082526243534f60e81b6020909201918252610069916002916100f8565b50620f4240600355600060045534801561008257600080fd5b5060058054336001600160a01b03199182168117835560035460009182526006602052604082205591546007805460018101825593527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68890920180546001600160a01b03909316929091169190911790556101cc565b82805461010490610191565b90600052602060002090601f016020900481019282610126576000855561016c565b82601f1061013f57805160ff191683800117855561016c565b8280016001018555821561016c579182015b8281111561016c578251825591602001919060010190610151565b5061017892915061017c565b5090565b5b80821115610178576000815560010161017d565b600181811c908216806101a557607f821691505b602082108114156101c657634e487b7160e01b600052602260045260246000fd5b50919050565b610992806101db6000396000f3fe6080604052600436106100f35760003560e01c8063a0e67e2b1161008a578063cc60e62f11610059578063cc60e62f14610260578063cc81afc814610296578063e4fc6b6d146102ab578063f4ee80a2146102c057600080fd5b8063a0e67e2b146101f1578063a63da98114610213578063ae6e22f114610235578063c4e41b221461024b57600080fd5b80631fcbf7c3116100c65780631fcbf7c31461018957806324665e191461019c578063256c3ab4146101bc57806395d89b41146101dc57600080fd5b806306fdde03146100f857806307546172146101235780630e29df221461015b57806318160ddd14610165575b600080fd5b34801561010457600080fd5b5061010d6102d5565b60405161011a9190610753565b60405180910390f35b34801561012f57600080fd5b50600554610143906001600160a01b031681565b6040516001600160a01b03909116815260200161011a565b610163610363565b005b34801561017157600080fd5b5061017b60035481565b60405190815260200161011a565b34801561019557600080fd5b5033610143565b3480156101a857600080fd5b506101636101b73660046107bd565b610396565b3480156101c857600080fd5b506101636101d73660046107bd565b6103d1565b3480156101e857600080fd5b5061010d61052d565b3480156101fd57600080fd5b5061020661053a565b60405161011a91906107e9565b34801561021f57600080fd5b503360009081526006602052604090205461017b565b34801561024157600080fd5b5061017b60045481565b34801561025757600080fd5b5060035461017b565b34801561026c57600080fd5b5061017b61027b366004610836565b6001600160a01b031660009081526006602052604090205490565b3480156102a257600080fd5b5061010d610607565b3480156102b757600080fd5b50610163610614565b3480156102cc57600080fd5b5060075461017b565b600180546102e29061085a565b80601f016020809104026020016040519081016040528092919081815260200182805461030e9061085a565b801561035b5780601f106103305761010080835404028352916020019161035b565b820191906000526020600020905b81548152906001019060200180831161033e57829003601f168201915b505050505081565b6040805180820190915260098082526839b2b73222ba3432b960b91b6020909201918252610393916000916106ba565b50565b6040516001600160a01b0383169082156108fc029083906000818181858888f193505050501580156103cc573d6000803e3d6000fd5b505050565b3360009081526006602052604090205481111561042657336000908152600660205260409081902054905163cf47918160e01b815261041d918391600401918252602082015260400190565b60405180910390fd5b6001600160a01b03821660009081526006602052604090205461048f57600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b0319166001600160a01b0384161790555b33600090815260066020526040812080548392906104ae9084906108a5565b90915550506001600160a01b038216600090815260066020526040812080548392906104db9084906108bc565b9091555050604080513381526001600160a01b03841660208201529081018290527f3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd33459060600160405180910390a15050565b600280546102e29061085a565b60075460609060009067ffffffffffffffff81111561055b5761055b6108d4565b604051908082528060200260200182016040528015610584578160200160208202803683370190505b50905060005b60075481101561060157600781815481106105a7576105a76108ea565b9060005260206000200160009054906101000a90046001600160a01b03168282815181106105d7576105d76108ea565b6001600160a01b0390921660209283029190910190910152806105f981610900565b91505061058a565b50919050565b600080546102e29061085a565b60005b600754811015610393576106a860078281548110610637576106376108ea565b9060005260206000200160009054906101000a90046001600160a01b0316600354600660006007868154811061066f5761066f6108ea565b60009182526020808320909101546001600160a01b0316835282019290925260400190205461069e904761091b565b6101b7919061093a565b806106b281610900565b915050610617565b8280546106c69061085a565b90600052602060002090601f0160209004810192826106e8576000855561072e565b82601f1061070157805160ff191683800117855561072e565b8280016001018555821561072e579182015b8281111561072e578251825591602001919060010190610713565b5061073a92915061073e565b5090565b5b8082111561073a576000815560010161073f565b600060208083528351808285015260005b8181101561078057858101830151858201604001528201610764565b81811115610792576000604083870101525b50601f01601f1916929092016040019392505050565b6001600160a01b038116811461039357600080fd5b600080604083850312156107d057600080fd5b82356107db816107a8565b946020939093013593505050565b6020808252825182820181905260009190848201906040850190845b8181101561082a5783516001600160a01b031683529284019291840191600101610805565b50909695505050505050565b60006020828403121561084857600080fd5b8135610853816107a8565b9392505050565b600181811c9082168061086e57607f821691505b6020821081141561060157634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000828210156108b7576108b761088f565b500390565b600082198211156108cf576108cf61088f565b500190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006000198214156109145761091461088f565b5060010190565b60008160001904831182151516156109355761093561088f565b500290565b60008261095757634e487b7160e01b600052601260045260246000fd5b50049056fea26469706673582212203c3c1f9d378c446df1b9107e7c1239ececfb3510377fc72bedc4b7b9a35c82b964736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Addresse is a free data retrieval call binding the contract method 0x1fcbf7c3.
//
// Solidity: function addresse() view returns(address)
func (_Api *ApiCaller) Addresse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "addresse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresse is a free data retrieval call binding the contract method 0x1fcbf7c3.
//
// Solidity: function addresse() view returns(address)
func (_Api *ApiSession) Addresse() (common.Address, error) {
	return _Api.Contract.Addresse(&_Api.CallOpts)
}

// Addresse is a free data retrieval call binding the contract method 0x1fcbf7c3.
//
// Solidity: function addresse() view returns(address)
func (_Api *ApiCallerSession) Addresse() (common.Address, error) {
	return _Api.Contract.Addresse(&_Api.CallOpts)
}

// FunctionCalled is a free data retrieval call binding the contract method 0xcc81afc8.
//
// Solidity: function functionCalled() view returns(string)
func (_Api *ApiCaller) FunctionCalled(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "functionCalled")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FunctionCalled is a free data retrieval call binding the contract method 0xcc81afc8.
//
// Solidity: function functionCalled() view returns(string)
func (_Api *ApiSession) FunctionCalled() (string, error) {
	return _Api.Contract.FunctionCalled(&_Api.CallOpts)
}

// FunctionCalled is a free data retrieval call binding the contract method 0xcc81afc8.
//
// Solidity: function functionCalled() view returns(string)
func (_Api *ApiCallerSession) FunctionCalled() (string, error) {
	return _Api.Contract.FunctionCalled(&_Api.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Api *ApiCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Api *ApiSession) GetOwners() ([]common.Address, error) {
	return _Api.Contract.GetOwners(&_Api.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Api *ApiCallerSession) GetOwners() ([]common.Address, error) {
	return _Api.Contract.GetOwners(&_Api.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Api *ApiCaller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Api *ApiSession) GetTotalSupply() (*big.Int, error) {
	return _Api.Contract.GetTotalSupply(&_Api.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Api *ApiCallerSession) GetTotalSupply() (*big.Int, error) {
	return _Api.Contract.GetTotalSupply(&_Api.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Api *ApiCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Api *ApiSession) Minter() (common.Address, error) {
	return _Api.Contract.Minter(&_Api.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Api *ApiCallerSession) Minter() (common.Address, error) {
	return _Api.Contract.Minter(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCallerSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// NbrOwners is a free data retrieval call binding the contract method 0xf4ee80a2.
//
// Solidity: function nbrOwners() view returns(uint256)
func (_Api *ApiCaller) NbrOwners(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "nbrOwners")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NbrOwners is a free data retrieval call binding the contract method 0xf4ee80a2.
//
// Solidity: function nbrOwners() view returns(uint256)
func (_Api *ApiSession) NbrOwners() (*big.Int, error) {
	return _Api.Contract.NbrOwners(&_Api.CallOpts)
}

// NbrOwners is a free data retrieval call binding the contract method 0xf4ee80a2.
//
// Solidity: function nbrOwners() view returns(uint256)
func (_Api *ApiCallerSession) NbrOwners() (*big.Int, error) {
	return _Api.Contract.NbrOwners(&_Api.CallOpts)
}

// StockQte is a free data retrieval call binding the contract method 0xa63da981.
//
// Solidity: function stockQte() view returns(uint256)
func (_Api *ApiCaller) StockQte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "stockQte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StockQte is a free data retrieval call binding the contract method 0xa63da981.
//
// Solidity: function stockQte() view returns(uint256)
func (_Api *ApiSession) StockQte() (*big.Int, error) {
	return _Api.Contract.StockQte(&_Api.CallOpts)
}

// StockQte is a free data retrieval call binding the contract method 0xa63da981.
//
// Solidity: function stockQte() view returns(uint256)
func (_Api *ApiCallerSession) StockQte() (*big.Int, error) {
	return _Api.Contract.StockQte(&_Api.CallOpts)
}

// StockQteOf is a free data retrieval call binding the contract method 0xcc60e62f.
//
// Solidity: function stockQteOf(address account) view returns(uint256)
func (_Api *ApiCaller) StockQteOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "stockQteOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StockQteOf is a free data retrieval call binding the contract method 0xcc60e62f.
//
// Solidity: function stockQteOf(address account) view returns(uint256)
func (_Api *ApiSession) StockQteOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.StockQteOf(&_Api.CallOpts, account)
}

// StockQteOf is a free data retrieval call binding the contract method 0xcc60e62f.
//
// Solidity: function stockQteOf(address account) view returns(uint256)
func (_Api *ApiCallerSession) StockQteOf(account common.Address) (*big.Int, error) {
	return _Api.Contract.StockQteOf(&_Api.CallOpts, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCallerSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// TotalEther is a free data retrieval call binding the contract method 0xae6e22f1.
//
// Solidity: function totalEther() view returns(uint256)
func (_Api *ApiCaller) TotalEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEther is a free data retrieval call binding the contract method 0xae6e22f1.
//
// Solidity: function totalEther() view returns(uint256)
func (_Api *ApiSession) TotalEther() (*big.Int, error) {
	return _Api.Contract.TotalEther(&_Api.CallOpts)
}

// TotalEther is a free data retrieval call binding the contract method 0xae6e22f1.
//
// Solidity: function totalEther() view returns(uint256)
func (_Api *ApiCallerSession) TotalEther() (*big.Int, error) {
	return _Api.Contract.TotalEther(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Api *ApiCallerSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Api *ApiTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Api *ApiSession) Distribute() (*types.Transaction, error) {
	return _Api.Contract.Distribute(&_Api.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Api *ApiTransactorSession) Distribute() (*types.Transaction, error) {
	return _Api.Contract.Distribute(&_Api.TransactOpts)
}

// SendEther is a paid mutator transaction binding the contract method 0x0e29df22.
//
// Solidity: function sendEther() payable returns()
func (_Api *ApiTransactor) SendEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendEther")
}

// SendEther is a paid mutator transaction binding the contract method 0x0e29df22.
//
// Solidity: function sendEther() payable returns()
func (_Api *ApiSession) SendEther() (*types.Transaction, error) {
	return _Api.Contract.SendEther(&_Api.TransactOpts)
}

// SendEther is a paid mutator transaction binding the contract method 0x0e29df22.
//
// Solidity: function sendEther() payable returns()
func (_Api *ApiTransactorSession) SendEther() (*types.Transaction, error) {
	return _Api.Contract.SendEther(&_Api.TransactOpts)
}

// SendEtherTo is a paid mutator transaction binding the contract method 0x24665e19.
//
// Solidity: function sendEtherTo(address receiver, uint256 amountEther) returns()
func (_Api *ApiTransactor) SendEtherTo(opts *bind.TransactOpts, receiver common.Address, amountEther *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendEtherTo", receiver, amountEther)
}

// SendEtherTo is a paid mutator transaction binding the contract method 0x24665e19.
//
// Solidity: function sendEtherTo(address receiver, uint256 amountEther) returns()
func (_Api *ApiSession) SendEtherTo(receiver common.Address, amountEther *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendEtherTo(&_Api.TransactOpts, receiver, amountEther)
}

// SendEtherTo is a paid mutator transaction binding the contract method 0x24665e19.
//
// Solidity: function sendEtherTo(address receiver, uint256 amountEther) returns()
func (_Api *ApiTransactorSession) SendEtherTo(receiver common.Address, amountEther *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendEtherTo(&_Api.TransactOpts, receiver, amountEther)
}

// SendStockTo is a paid mutator transaction binding the contract method 0x256c3ab4.
//
// Solidity: function sendStockTo(address receiver, uint256 amount) returns()
func (_Api *ApiTransactor) SendStockTo(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendStockTo", receiver, amount)
}

// SendStockTo is a paid mutator transaction binding the contract method 0x256c3ab4.
//
// Solidity: function sendStockTo(address receiver, uint256 amount) returns()
func (_Api *ApiSession) SendStockTo(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendStockTo(&_Api.TransactOpts, receiver, amount)
}

// SendStockTo is a paid mutator transaction binding the contract method 0x256c3ab4.
//
// Solidity: function sendStockTo(address receiver, uint256 amount) returns()
func (_Api *ApiTransactorSession) SendStockTo(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendStockTo(&_Api.TransactOpts, receiver, amount)
}

// ApiSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Api contract.
type ApiSentIterator struct {
	Event *ApiSent // Event containing the contract specifics and raw log

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
func (it *ApiSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiSent)
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
		it.Event = new(ApiSent)
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
func (it *ApiSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiSent represents a Sent event raised by the Api contract.
type ApiSent struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_Api *ApiFilterer) FilterSent(opts *bind.FilterOpts) (*ApiSentIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &ApiSentIterator{contract: _Api.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_Api *ApiFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *ApiSent) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiSent)
				if err := _Api.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_Api *ApiFilterer) ParseSent(log types.Log) (*ApiSent, error) {
	event := new(ApiSent)
	if err := _Api.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
