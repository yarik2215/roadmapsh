package tracker

import "testing"

func TestTaskTracker_Add(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		storage TaskStorage
		// Named input parameters for target function.
		description string
		status      TaskStatus
		want        Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := NewTracker(tt.storage)
			got := ta.Add(tt.description, tt.status)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskTracker_Remove(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		storage TaskStorage
		// Named input parameters for target function.
		id      int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := NewTracker(tt.storage)
			gotErr := ta.Remove(tt.id)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Remove() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Remove() succeeded unexpectedly")
			}
		})
	}
}
