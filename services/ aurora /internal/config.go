package aurora

import (
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/throttled/throttled"
)

// Config is the configuration for aurora.  It gets populated by the
// app's main function and is provided to NewApp.
type Config struct {
	DatabaseURL            string
	HcnetCoreDatabaseURL string
	HcnetCoreURL         string
	Port                   uint
	MaxDBConnections       int
	SSEUpdateFrequency     time.Duration
	ConnectionTimeout      time.Duration
	RateLimit              *throttled.RateQuota
	RateLimitRedisKey      string
	RedisURL               string
	FriendbotURL           *url.URL
	LogLevel               logrus.Level
	LogFile                string
	// MaxPathLength is the maximum length of the path returned by `/paths` endpoint.
	MaxPathLength     uint
	NetworkPassphrase string
	SentryDSN         string
	LogglyToken       string
	LogglyTag         string
	// TLSCert is a path to a certificate file to use for aurora's TLS config
	TLSCert string
	// TLSKey is the path to a private key file to use for aurora's TLS config
	TLSKey string
	// Ingest toggles whether this aurora instance should run the data ingestion subsystem.
	Ingest bool
	// IngestFailedTransactions toggles whether to ingest failed transactions
	IngestFailedTransactions bool
	// HistoryRetentionCount represents the minimum number of ledgers worth of
	// history data to retain in the aurora database. For the purposes of
	// determining a "retention duration", each ledger roughly corresponds to 10
	// seconds of real time.
	HistoryRetentionCount uint
	// StaleThreshold represents the number of ledgers a history database may be
	// out-of-date by before aurora begins to respond with an error to history
	// requests.
	StaleThreshold uint
	// SkipCursorUpdate causes the ingestor to skip reporting the "last imported
	// ledger" state to hcnet-core.
	SkipCursorUpdate bool
	// EnableAssetStats is a feature flag that determines whether to calculate
	// asset stats during the ingestion and expose `/assets` endpoint.
	// Enabling it has a negative impact on CPU when ingesting ledgers full of
	// many different assets related operations.
	EnableAssetStats bool
}
