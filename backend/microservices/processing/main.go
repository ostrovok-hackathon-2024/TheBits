package processing

import (
	"context"
	"errors"
	"log"

	"golang.org/x/sync/errgroup"
)

type Config struct {
}

type Processing struct {
	config        *Config
	workerChannel chan []byte
}

func NewProcessing(cfg *Config) *Processing {
	return &Processing{
		config:        cfg,
		workerChannel: make(chan []byte),
	}
}

func (p *Processing) Run(ctx context.Context) error {
	defer close(p.workerChannel)

	group, ctx := errgroup.WithContext(ctx)

	for {
		select {
		case <-ctx.Done():
			if err := group.Wait(); err != nil {
				return err
			}
			return errors.New("processing end by context")

		case d := <-p.workerChannel:
			group.Go(func() error {
				if err := p.process(d); err != nil {
					return err
				}
				return nil
			})

		}
	}

	return nil
}

func (p *Processing) process(data []byte) error {
	// @TODO сохранить данные, найти аналоги, перенаправить в ML
	// @TODO ответ сохранить, проанализировать на различие и сформировать сообщение в очередь для алертов
	log.Println("processing data", data)
	res, err := p.sendML(data)
	log.Println("return response from ML", res)
	return err
}

func (p *Processing) sendML(data []byte) ([]byte, error) {
	// @TODO отправить данные в ML
	log.Println("send data to ML", data)
	return []byte{}, nil
}

func (p *Processing) checkAnomaly(data []byte) error {
	// @TODO проверить на аномалию
	log.Println("check anomaly", data)
	return nil

}

func (p *Processing) sendAlert(data []byte) error {
	return nil
}
