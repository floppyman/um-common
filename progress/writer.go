package progress

import (
	"context"
	"io"
)

// Writer is an [io.Writer] that reports the number of bytes written to it via a callback.
// and supports cancellation via a context added using [Reader.WithContext].
//
// The callback is called at most every "updateInterval" bytes.
// The updateInterval can be set using the [Writer.WithUpdateInterval] method.
//
// The following is an example of how to use [Writer] to report the progress of
// writing to a file:
//
//	file, _ := os.Create("file.txt")
//	progressWriter := progress.NewWriter(func(processedBytes int) {
//	    fmt.Printf("Processed %d bytes\n", processedBytes)
//	})
//	writerWithProgress := io.MultiWriter(file, progressWriter)
//	io.Copy(writerWithProgress, bytes.NewReader(bytes.Repeat([]byte{42}, 1024*1024)))
//
// Or with a context.
//
//  ctx := context.Background()
//	file, _ := os.Create("file.txt")
//	progressWriter := progress.NewWriter(func(processedBytes int) {
//	    fmt.Printf("Processed %d bytes\n", processedBytes)
//	}).WithContext(ctx)
//	writerWithProgress := io.MultiWriter(file, progressWriter)
//	io.Copy(writerWithProgress, bytes.NewReader(bytes.Repeat([]byte{42}, 1024*1024)))
type Writer struct {
	processedBytes int
	progressFn     func(processedBytes int)
	lastUpdate     int
	updateInterval int
	ctx            context.Context
}

func NewWriter(progressFn func(processedBytes int)) *Writer {
	return &Writer{
		progressFn:     progressFn,
		updateInterval: defaultUpdateInterval,
		ctx:            nil}
}

func (w *Writer) WithUpdateInterval(bytes int) *Writer {
	w.updateInterval = bytes
	return w
}

func (w *Writer) WithContext(ctx context.Context) *Writer {
	w.ctx = ctx
	return w
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if w.ctx != nil && w.ctx.Err() != nil {
		return 0, context.Canceled
	}
	
	w.processedBytes += len(p)
	if w.lastUpdate == 0 || w.processedBytes-w.lastUpdate > w.updateInterval {
		w.progressFn(w.processedBytes)
		w.lastUpdate = w.processedBytes
	}
	return len(p), nil
}

var _ io.Writer = (*Writer)(nil)
