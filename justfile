alias t := test
alias c := check

default:
  just --list

submit year day part:
  aoc submit --year {{year}} --day {{day}} {{part}} $(just run-{{year}}-{{day}}-{{part}})

check: fmt lint test

fmt:
  gosimports -local github.com/maolonglong -w .
  gofumpt -extra -w .

lint:
  go vet ./...

test:
  #!/usr/bin/env bash
  set -euo pipefail
  tasks=$(grep -E '^run-.*?:' ./justfile | awk -F ':' '{print $1}')
  for task in $tasks; do
    want=$(grep "$task" ./justfile | awk -F ' ## ' '{print $2}')
    got=$(just "$task")
    if [[ "$got" != "$want" ]]; then
      echo "$task = $got, want $want"
      exit 1
    fi
  done

run-2023-1-1: (_go "2023" "1" "1") ## 54634
run-2023-1-2: (_go "2023" "1" "2") ## 53855

[private]
_go year day part:
  cd ./{{year}}/{{day}}/{{part}} && go run main.go
