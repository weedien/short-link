package logging

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
	"time"
)

type AsyncW struct {
	io       *lumberjack.Logger
	dataChan chan []byte
	done     chan struct{}
	wg       sync.WaitGroup
}

// AsyncLumberjack initializes an AsyncW logger with the specified lumberjack.Logger.
func AsyncLumberjack(io *lumberjack.Logger) *AsyncW {
	r := &AsyncW{
		io:       io,
		dataChan: make(chan []byte, 1024),
		done:     make(chan struct{}),
	}

	r.wg.Add(1)
	go r.run()

	return r
}

func (r *AsyncW) run() {
	defer r.wg.Done()
	ticker := time.NewTicker(time.Millisecond * 333)
	defer ticker.Stop()

	var buffer []byte

	for {
		select {
		case val := <-r.dataChan:
			buffer = append(buffer, val...)
			if len(buffer) >= 100 {
				_, _ = r.io.Write(buffer)
				buffer = nil
			}
		case <-ticker.C:
			if len(buffer) > 0 {
				_, _ = r.io.Write(buffer)
				buffer = nil
			}
		case <-r.done:
			if len(buffer) > 0 {
				_, _ = r.io.Write(buffer)
			}
			return
		}
	}
}

// Write asynchronously writes the data to the logger.
func (r *AsyncW) Write(p []byte) (n int, err error) {
	select {
	case r.dataChan <- p:
		return len(p), nil
	default:
		// Data channel is full, handle the error as needed
		return 0, nil // or return an error
	}
}

// Stop closes the logger and ensures all pending data is written.
func (r *AsyncW) Stop() {
	close(r.done)
	r.wg.Wait()
	close(r.dataChan)
	err := r.io.Close()
	if err != nil {
		return
	}
}
