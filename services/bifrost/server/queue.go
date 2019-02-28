package server

import (
	"time"
)

// poolTransactionsQueue pools transactions queue which contains only processed and
// validated transactions and sends it to HcnetAccountConfigurator for account configuration.
func (s *Server) poolTransactionsQueue() {
	s.log.Info("Started pooling transactions queue")

	for {
		transaction, err := s.TransactionsQueue.QueuePool()
		if err != nil {
			s.log.WithField("err", err).Error("Error pooling transactions queue")
			time.Sleep(time.Second)
			continue
		}

		if transaction == nil {
			time.Sleep(time.Second)
			continue
		}

		s.log.WithField("transaction", transaction).Info("Received transaction from transactions queue")
		go s.HcnetAccountConfigurator.ConfigureAccount(
			transaction.HcnetPublicKey,
			string(transaction.AssetCode),
			transaction.Amount,
		)
	}
}
