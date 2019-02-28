// Package aurora provides client access to a aurora server, allowing an
// application to post transactions and lookup ledger information.
//
// Create an instance of `Client` to customize the server used, or alternatively
// use `DefaultTestNetClient` or `DefaultPublicNetClient` to access the SDF run
// aurora servers.
package aurora

import (
	"context"
	"net/http"
	"net/url"
	"sync"

	"github.com/hcnet/go/build"
	"github.com/hcnet/go/support/errors"
	"github.com/hcnet/go/xdr"
)

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{
	URL:  "https://aurora-testnet.hcnet.org",
	HTTP: http.DefaultClient,
}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{
	URL:  "https://aurora.hcnet.org",
	HTTP: http.DefaultClient,
}

// At is a paging parameter that can be used to override the URL loaded in a
// remote method call to aurora.
type At string

// Cursor represents `cursor` param in queries
type Cursor string

// Limit represents `limit` param in queries
type Limit uint

// Order represents `order` param in queries
type Order string

// StartTime is an integer values of timestamp
type StartTime int64

// EndTime is an integer values of timestamp
type EndTime int64

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

var (
	// ErrResultCodesNotPopulated is the error returned from a call to
	// ResultCodes() against a `Problem` value that doesn't have the
	// "result_codes" extra field populated when it is expected to be.
	ErrResultCodesNotPopulated = errors.New("result_codes not populated")

	// ErrEnvelopeNotPopulated is the error returned from a call to
	// Envelope() against a `Problem` value that doesn't have the
	// "envelope_xdr" extra field populated when it is expected to be.
	ErrEnvelopeNotPopulated = errors.New("envelope_xdr not populated")

	// ErrResultNotPopulated is the error returned from a call to
	// Result() against a `Problem` value that doesn't have the
	// "result_xdr" extra field populated when it is expected to be.
	ErrResultNotPopulated = errors.New("result_xdr not populated")
)

// Client struct contains data required to connect to Aurora instance
// It is okay to call methods on Client concurrently.
// A Client must not be copied after first use.
type Client struct {
	// URL of Aurora server to connect
	URL string

	// HTTP client to make requests with
	HTTP HTTP

	fixURLOnce sync.Once
}

type ClientInterface interface {
	Root() (Root, error)
	HomeDomainForAccount(aid string) (string, error)
	LoadAccount(accountID string) (Account, error)
	LoadAccountOffers(accountID string, params ...interface{}) (offers OffersPage, err error)
	LoadTradeAggregations(
		baseAsset Asset,
		counterAsset Asset,
		resolution int64,
		params ...interface{},
	) (tradeAggrs TradeAggregationsPage, err error)
	LoadTrades(
		baseAsset Asset,
		counterAsset Asset,
		offerID int64,
		resolution int64,
		params ...interface{},
	) (tradesPage TradesPage, err error)
	LoadAccountMergeAmount(p *Payment) error
	LoadMemo(p *Payment) error
	LoadOperation(operationID string) (payment Payment, err error)
	LoadOrderBook(selling Asset, buying Asset, params ...interface{}) (orderBook OrderBookSummary, err error)
	LoadTransaction(transactionID string) (transaction Transaction, err error)
	SequenceForAccount(accountID string) (xdr.SequenceNumber, error)
	StreamLedgers(ctx context.Context, cursor *Cursor, handler LedgerHandler) error
	StreamPayments(ctx context.Context, accountID string, cursor *Cursor, handler PaymentHandler) error
	StreamTransactions(ctx context.Context, accountID string, cursor *Cursor, handler TransactionHandler) error
	SubmitTransaction(txeBase64 string) (TransactionSuccess, error)
}

// Error struct contains the problem returned by Aurora
type Error struct {
	Response *http.Response
	Problem  Problem
}

// HTTP represents the HTTP client that a aurora client uses to communicate
type HTTP interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// LedgerHandler is a function that is called when a new ledger is received
type LedgerHandler func(Ledger)

// PaymentHandler is a function that is called when a new payment is received
type PaymentHandler func(Payment)

// TransactionHandler is a function that is called when a new transaction is received
type TransactionHandler func(Transaction)

// ensure that the aurora client can be used as a SequenceProvider
var _ build.SequenceProvider = &Client{}

// ensure that the aurora client implements ClientInterface
var _ ClientInterface = &Client{}
