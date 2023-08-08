package worker

import (
	"cyberok/internal/config"
	"cyberok/internal/lib/api/logger/sl"
	"cyberok/internal/service"
	"golang.org/x/exp/slog"
	"sync"
	"time"
)

type UpdateWorker struct {
	timer   time.Duration
	done    chan bool
	service *service.Service
	log     *slog.Logger
}

func NewUpdateWorker(cfg *config.DBConfig, log *slog.Logger, service *service.Service) *UpdateWorker {
	return &UpdateWorker{
		timer:   cfg.UpdateTimer,
		done:    make(chan bool),
		service: service,
		log:     log,
	}
}

func (w *UpdateWorker) Start() {
	const op = "worker.update"
	w.log = w.log.With(
		slog.String("op", op),
	)

	ticker := time.NewTicker(w.timer)
	go func() {
		for {
			select {
			case <-w.done:
				ticker.Stop()
				return
			case <-ticker.C:
				w.update()
			}
		}
	}()
}

func (w *UpdateWorker) Close() {
	w.done <- true
	w.log.Info("UpdateWorker has been stopped")

}

func (w *UpdateWorker) update() {

	fqdns, _ := w.service.GetAll()
	if len(fqdns) == 0 {
		w.log.Info("nothing tp update.")
		return
	}
	err := w.service.TruncateIp()
	if err != nil {
		w.log.Error("truncate err", sl.Err(err))
	}
	wg := sync.WaitGroup{}
	for _, fqdn := range fqdns {
		wg.Add(1)
		go func(fqdn string) {
			defer wg.Done()
			ips, _ := w.service.FqdnResolver.LookupIp(fqdn)
			_, err := w.service.UpdateFqdn(fqdn, ips)
			if err != nil {
				w.log.Error("can't update fqdn", sl.Err(err))
			}
		}(fqdn.Name)
	}
	wg.Wait()

	w.log.Info("IP addresses have been updated")
}
