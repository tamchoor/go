#!/bin/bash

#  go test -bench=. -benchtime=1000x| grep Benchmark | awk '{print $1,$3}' | sort -nk2 | head > top.txt

go test -bench=. -benchtime=1000000x -cpuprofile cpu.out | go tool pprof ./cpu.out

# top