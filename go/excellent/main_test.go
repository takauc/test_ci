﻿package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
    result := EvenOdd(10)
        if result != "even" {
            t.Errorf("expected: even, actual: %s". result)
    }
}