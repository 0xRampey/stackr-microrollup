// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package settlement

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
	_ = abi.ConvertType
)

// SettlementBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type SettlementBatchHeader struct {
	PrevHash  [32]byte
	StateRoot [32]byte
	TxRoot    [32]byte
}

// SettlementTx is an auto generated low-level Go binding around an user-defined struct.
type SettlementTx struct {
	Signature []byte
}

// SettlementMetaData contains all meta data concerning the Settlement contract.
var SettlementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"}],\"name\":\"AppRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"}],\"name\":\"isAppRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredApps\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"PrevHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"StateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"TxRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structSettlement.BatchHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Signature\",\"type\":\"bytes\"}],\"internalType\":\"structSettlement.Tx[]\",\"name\":\"txList\",\"type\":\"tuple[]\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SettlementABI is the input ABI used to generate the binding from.
// Deprecated: Use SettlementMetaData.ABI instead.
var SettlementABI = SettlementMetaData.ABI

// Settlement is an auto generated Go binding around an Ethereum contract.
type Settlement struct {
	SettlementCaller     // Read-only binding to the contract
	SettlementTransactor // Write-only binding to the contract
	SettlementFilterer   // Log filterer for contract events
}

// SettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type SettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SettlementSession struct {
	Contract     *Settlement       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SettlementCallerSession struct {
	Contract *SettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SettlementTransactorSession struct {
	Contract     *SettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type SettlementRaw struct {
	Contract *Settlement // Generic contract binding to access the raw methods on
}

// SettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SettlementCallerRaw struct {
	Contract *SettlementCaller // Generic read-only contract binding to access the raw methods on
}

// SettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SettlementTransactorRaw struct {
	Contract *SettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSettlement creates a new instance of Settlement, bound to a specific deployed contract.
func NewSettlement(address common.Address, backend bind.ContractBackend) (*Settlement, error) {
	contract, err := bindSettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Settlement{SettlementCaller: SettlementCaller{contract: contract}, SettlementTransactor: SettlementTransactor{contract: contract}, SettlementFilterer: SettlementFilterer{contract: contract}}, nil
}

// NewSettlementCaller creates a new read-only instance of Settlement, bound to a specific deployed contract.
func NewSettlementCaller(address common.Address, caller bind.ContractCaller) (*SettlementCaller, error) {
	contract, err := bindSettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementCaller{contract: contract}, nil
}

// NewSettlementTransactor creates a new write-only instance of Settlement, bound to a specific deployed contract.
func NewSettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*SettlementTransactor, error) {
	contract, err := bindSettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementTransactor{contract: contract}, nil
}

// NewSettlementFilterer creates a new log filterer instance of Settlement, bound to a specific deployed contract.
func NewSettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*SettlementFilterer, error) {
	contract, err := bindSettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SettlementFilterer{contract: contract}, nil
}

// bindSettlement binds a generic wrapper to an already deployed contract.
func bindSettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SettlementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Settlement *SettlementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Settlement.Contract.SettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Settlement *SettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settlement.Contract.SettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Settlement *SettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Settlement.Contract.SettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Settlement *SettlementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Settlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Settlement *SettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Settlement *SettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Settlement.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_Settlement *SettlementCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_Settlement *SettlementSession) Aggregator() (common.Address, error) {
	return _Settlement.Contract.Aggregator(&_Settlement.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_Settlement *SettlementCallerSession) Aggregator() (common.Address, error) {
	return _Settlement.Contract.Aggregator(&_Settlement.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Settlement *SettlementCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Settlement *SettlementSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Settlement.Contract.Balances(&_Settlement.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Settlement *SettlementCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Settlement.Contract.Balances(&_Settlement.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32)
func (_Settlement *SettlementCaller) Batches(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "batches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32)
func (_Settlement *SettlementSession) Batches(arg0 *big.Int) ([32]byte, error) {
	return _Settlement.Contract.Batches(&_Settlement.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(bytes32)
func (_Settlement *SettlementCallerSession) Batches(arg0 *big.Int) ([32]byte, error) {
	return _Settlement.Contract.Batches(&_Settlement.CallOpts, arg0)
}

// IsAppRegistered is a free data retrieval call binding the contract method 0x8403be91.
//
// Solidity: function isAppRegistered(address app) view returns(bool)
func (_Settlement *SettlementCaller) IsAppRegistered(opts *bind.CallOpts, app common.Address) (bool, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "isAppRegistered", app)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAppRegistered is a free data retrieval call binding the contract method 0x8403be91.
//
// Solidity: function isAppRegistered(address app) view returns(bool)
func (_Settlement *SettlementSession) IsAppRegistered(app common.Address) (bool, error) {
	return _Settlement.Contract.IsAppRegistered(&_Settlement.CallOpts, app)
}

// IsAppRegistered is a free data retrieval call binding the contract method 0x8403be91.
//
// Solidity: function isAppRegistered(address app) view returns(bool)
func (_Settlement *SettlementCallerSession) IsAppRegistered(app common.Address) (bool, error) {
	return _Settlement.Contract.IsAppRegistered(&_Settlement.CallOpts, app)
}

// RegisteredApps is a free data retrieval call binding the contract method 0xa6c4cce9.
//
// Solidity: function registeredApps(address ) view returns(bool)
func (_Settlement *SettlementCaller) RegisteredApps(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "registeredApps", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredApps is a free data retrieval call binding the contract method 0xa6c4cce9.
//
// Solidity: function registeredApps(address ) view returns(bool)
func (_Settlement *SettlementSession) RegisteredApps(arg0 common.Address) (bool, error) {
	return _Settlement.Contract.RegisteredApps(&_Settlement.CallOpts, arg0)
}

// RegisteredApps is a free data retrieval call binding the contract method 0xa6c4cce9.
//
// Solidity: function registeredApps(address ) view returns(bool)
func (_Settlement *SettlementCallerSession) RegisteredApps(arg0 common.Address) (bool, error) {
	return _Settlement.Contract.RegisteredApps(&_Settlement.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Settlement *SettlementTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Settlement *SettlementSession) Deposit() (*types.Transaction, error) {
	return _Settlement.Contract.Deposit(&_Settlement.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Settlement *SettlementTransactorSession) Deposit() (*types.Transaction, error) {
	return _Settlement.Contract.Deposit(&_Settlement.TransactOpts)
}

// RegisterApp is a paid mutator transaction binding the contract method 0x6d6c85c3.
//
// Solidity: function registerApp() returns()
func (_Settlement *SettlementTransactor) RegisterApp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "registerApp")
}

// RegisterApp is a paid mutator transaction binding the contract method 0x6d6c85c3.
//
// Solidity: function registerApp() returns()
func (_Settlement *SettlementSession) RegisterApp() (*types.Transaction, error) {
	return _Settlement.Contract.RegisterApp(&_Settlement.TransactOpts)
}

// RegisterApp is a paid mutator transaction binding the contract method 0x6d6c85c3.
//
// Solidity: function registerApp() returns()
func (_Settlement *SettlementTransactorSession) RegisterApp() (*types.Transaction, error) {
	return _Settlement.Contract.RegisterApp(&_Settlement.TransactOpts)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x95eb88bc.
//
// Solidity: function submitBatch((bytes32,bytes32,bytes32) header, (bytes)[] txList) returns()
func (_Settlement *SettlementTransactor) SubmitBatch(opts *bind.TransactOpts, header SettlementBatchHeader, txList []SettlementTx) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "submitBatch", header, txList)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x95eb88bc.
//
// Solidity: function submitBatch((bytes32,bytes32,bytes32) header, (bytes)[] txList) returns()
func (_Settlement *SettlementSession) SubmitBatch(header SettlementBatchHeader, txList []SettlementTx) (*types.Transaction, error) {
	return _Settlement.Contract.SubmitBatch(&_Settlement.TransactOpts, header, txList)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x95eb88bc.
//
// Solidity: function submitBatch((bytes32,bytes32,bytes32) header, (bytes)[] txList) returns()
func (_Settlement *SettlementTransactorSession) SubmitBatch(header SettlementBatchHeader, txList []SettlementTx) (*types.Transaction, error) {
	return _Settlement.Contract.SubmitBatch(&_Settlement.TransactOpts, header, txList)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Settlement *SettlementTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Settlement *SettlementSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Settlement.Contract.Withdraw(&_Settlement.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Settlement *SettlementTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Settlement.Contract.Withdraw(&_Settlement.TransactOpts, amount)
}

// SettlementAppRegisteredIterator is returned from FilterAppRegistered and is used to iterate over the raw logs and unpacked data for AppRegistered events raised by the Settlement contract.
type SettlementAppRegisteredIterator struct {
	Event *SettlementAppRegistered // Event containing the contract specifics and raw log

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
func (it *SettlementAppRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementAppRegistered)
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
		it.Event = new(SettlementAppRegistered)
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
func (it *SettlementAppRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementAppRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementAppRegistered represents a AppRegistered event raised by the Settlement contract.
type SettlementAppRegistered struct {
	App common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppRegistered is a free log retrieval operation binding the contract event 0x0d540ad8f39e07d19909687352b9fa017405d93c91a6760981fbae9cf28bfef7.
//
// Solidity: event AppRegistered(address indexed app)
func (_Settlement *SettlementFilterer) FilterAppRegistered(opts *bind.FilterOpts, app []common.Address) (*SettlementAppRegisteredIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "AppRegistered", appRule)
	if err != nil {
		return nil, err
	}
	return &SettlementAppRegisteredIterator{contract: _Settlement.contract, event: "AppRegistered", logs: logs, sub: sub}, nil
}

// WatchAppRegistered is a free log subscription operation binding the contract event 0x0d540ad8f39e07d19909687352b9fa017405d93c91a6760981fbae9cf28bfef7.
//
// Solidity: event AppRegistered(address indexed app)
func (_Settlement *SettlementFilterer) WatchAppRegistered(opts *bind.WatchOpts, sink chan<- *SettlementAppRegistered, app []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "AppRegistered", appRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementAppRegistered)
				if err := _Settlement.contract.UnpackLog(event, "AppRegistered", log); err != nil {
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

// ParseAppRegistered is a log parse operation binding the contract event 0x0d540ad8f39e07d19909687352b9fa017405d93c91a6760981fbae9cf28bfef7.
//
// Solidity: event AppRegistered(address indexed app)
func (_Settlement *SettlementFilterer) ParseAppRegistered(log types.Log) (*SettlementAppRegistered, error) {
	event := new(SettlementAppRegistered)
	if err := _Settlement.contract.UnpackLog(event, "AppRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementBatchSubmittedIterator is returned from FilterBatchSubmitted and is used to iterate over the raw logs and unpacked data for BatchSubmitted events raised by the Settlement contract.
type SettlementBatchSubmittedIterator struct {
	Event *SettlementBatchSubmitted // Event containing the contract specifics and raw log

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
func (it *SettlementBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementBatchSubmitted)
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
		it.Event = new(SettlementBatchSubmitted)
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
func (it *SettlementBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementBatchSubmitted represents a BatchSubmitted event raised by the Settlement contract.
type SettlementBatchSubmitted struct {
	BatchHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchSubmitted is a free log retrieval operation binding the contract event 0xa53eb2683cef0c1e23d2e5dfdb2065ff4537fc3bbe0b632f0786c986435662f7.
//
// Solidity: event BatchSubmitted(bytes32 indexed batchHash)
func (_Settlement *SettlementFilterer) FilterBatchSubmitted(opts *bind.FilterOpts, batchHash [][32]byte) (*SettlementBatchSubmittedIterator, error) {

	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "BatchSubmitted", batchHashRule)
	if err != nil {
		return nil, err
	}
	return &SettlementBatchSubmittedIterator{contract: _Settlement.contract, event: "BatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchBatchSubmitted is a free log subscription operation binding the contract event 0xa53eb2683cef0c1e23d2e5dfdb2065ff4537fc3bbe0b632f0786c986435662f7.
//
// Solidity: event BatchSubmitted(bytes32 indexed batchHash)
func (_Settlement *SettlementFilterer) WatchBatchSubmitted(opts *bind.WatchOpts, sink chan<- *SettlementBatchSubmitted, batchHash [][32]byte) (event.Subscription, error) {

	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "BatchSubmitted", batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementBatchSubmitted)
				if err := _Settlement.contract.UnpackLog(event, "BatchSubmitted", log); err != nil {
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

// ParseBatchSubmitted is a log parse operation binding the contract event 0xa53eb2683cef0c1e23d2e5dfdb2065ff4537fc3bbe0b632f0786c986435662f7.
//
// Solidity: event BatchSubmitted(bytes32 indexed batchHash)
func (_Settlement *SettlementFilterer) ParseBatchSubmitted(log types.Log) (*SettlementBatchSubmitted, error) {
	event := new(SettlementBatchSubmitted)
	if err := _Settlement.contract.UnpackLog(event, "BatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Settlement contract.
type SettlementDepositIterator struct {
	Event *SettlementDeposit // Event containing the contract specifics and raw log

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
func (it *SettlementDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementDeposit)
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
		it.Event = new(SettlementDeposit)
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
func (it *SettlementDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementDeposit represents a Deposit event raised by the Settlement contract.
type SettlementDeposit struct {
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed depositor, uint256 amount)
func (_Settlement *SettlementFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address) (*SettlementDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return &SettlementDepositIterator{contract: _Settlement.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed depositor, uint256 amount)
func (_Settlement *SettlementFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SettlementDeposit, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementDeposit)
				if err := _Settlement.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed depositor, uint256 amount)
func (_Settlement *SettlementFilterer) ParseDeposit(log types.Log) (*SettlementDeposit, error) {
	event := new(SettlementDeposit)
	if err := _Settlement.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
