#!/bin/bash
 go test -bench=. -benchtime=1000x| grep Benchmark | awk '{print $1,$3}' | sort -nk2 | head > top10.txt