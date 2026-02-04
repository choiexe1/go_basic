package miniorder

type ReceiptWriter struct {
	buffer []byte
}

func (r *ReceiptWriter) Write(p []byte) (int, error) {
	r.buffer = append(r.buffer, p...)
	return len(p), nil
}

func (r *ReceiptWriter) String() string {
	return string(r.buffer)
}
