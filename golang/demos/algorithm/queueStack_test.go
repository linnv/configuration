package demo

import "testing"

func Test_twoQueueOneStack_Pop(t *testing.T) {
	type fields struct {
		queueX []int
		queueY []int
	}
	tests := []struct {
		name   string
		fields fields
		wantR  int
	}{
		{"normal", fields{}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := &twoQueueOneStack{
				queueX: tt.fields.queueX,
				queueY: tt.fields.queueY,
			}
			if gotR := qs.Pop(); gotR != -1 {
				t.Errorf("twoQueueOneStack.Pop() = %v, want %v", gotR, tt.wantR)
			}
			qs.Push(1)
			if gotR := qs.Pop(); gotR != 1 {
				t.Errorf("twoQueueOneStack.Pop() = %v, want %v", gotR, tt.wantR)
			}
			if gotR := qs.Size(); gotR != 0 {
				t.Errorf("twoQueueOneStack.Pop() = %v, want %v", gotR, tt.wantR)
			}
			qs.Push(1)
			qs.Push(2)
			if gotR := qs.Pop(); gotR != 2 {
				t.Errorf("twoQueueOneStack.Pop() = %v, want %v", gotR, tt.wantR)
			}
			if gotR := qs.Size(); gotR != 1 {
				t.Errorf("twoQueueOneStack.Pop() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
