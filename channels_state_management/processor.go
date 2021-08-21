package channels_state_management

import "context"

func Processor(ctx context.Context, in chan *WorkRequest, out chan *WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-in:
			out <- Process(wr)
		}
	}
}
