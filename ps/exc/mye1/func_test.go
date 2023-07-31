package main

import (
	"fmt"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	// Make new variable wiht slice
	ns := MakeSlice()

	// This weid syntax is to remove the last element
	NewSlice := append(ns[:0], ns[0+1:]...)

	// slice = append(slice[:i], slice[i+1:]...)

	if len(NewSlice) != 10 {
		// Was trying to print what the new one was here
		fmt.Println(NewSlice)
		t.Errorf("slice length is %d", len(NewSlice))
	}
}
