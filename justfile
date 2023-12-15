alias t := test
alias c := check

default:
  just --list

submit year day part:
  aoc submit --year {{year}} --day {{day}} {{part}} $(just run-{{year}}-{{day}}-{{part}})

check: fmt lint test

fmt: _janet-format
  fd --type=file --extension go --exec gosimports -local github.com/maolonglong -w
  fd --type=file --extension go --exec gofumpt -extra -w
  fd --type=file --extension janet --exec janet-format

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
run-2023-2-1: (_janet "2023" "2" "1") ## 2551
run-2023-2-2: (_janet "2023" "2" "2") ## 62811
run-2023-3-1: (_go "2023" "3" "1") ## 550064
run-2023-3-2: (_go "2023" "3" "2") ## 85010461
run-2023-4-1: (_go "2023" "4" "1") ## 24733
run-2023-4-2: (_go "2023" "4" "2") ## 5422730
run-2023-5-1: (_janet "2023" "5" "1") ## 510109797
run-2023-5-2: (_janet "2023" "5" "2") ## 9622622
run-2023-6-1: (_go "2023" "6" "1") ## 1108800
run-2023-6-2: (_go "2023" "6" "2") ## 36919753
run-2023-7-1: (_go "2023" "7" "1") ## 249204891
run-2023-7-2: (_go "2023" "7" "2") ## 249666369
run-2023-8-1: (_go "2023" "8" "1") ## 18113
run-2023-8-2: (_go "2023" "8" "2") ## 12315788159977
run-2023-9-1: (_go "2023" "9" "1") ## 1930746032
run-2023-9-2: (_go "2023" "9" "2") ## 1154
run-2023-10-1: (_go "2023" "10" "1") ## 6690
run-2023-10-2: (_go "2023" "10" "2") ## 525
run-2023-11-1: (_go "2023" "11" "1") ## 10313550
run-2023-11-2: (_go "2023" "11" "2") ## 611998089572
run-2023-12-1: (_janet "2023" "12" "1") ## 7379
run-2023-12-2: (_janet "2023" "12" "2") ## 7732028747925
run-2023-13-1: (_go "2023" "13" "1") ## 31739
run-2023-13-2: (_go "2023" "13" "2") ## 31539
run-2023-14-1: (_go "2023" "14" "1") ## 109345
run-2023-14-2: (_go "2023" "14" "2") ## 112452
run-2023-15-1: (_go "2023" "15" "1") ## 505459
run-2023-15-2: (_go "2023" "15" "2") ## 228508

[private]
_go year day part:
  cd ./{{year}}/{{day}}/{{part}} && go run main.go

[private]
_janet year day part:
  cd ./{{year}}/{{day}}/{{part}} && janet main.janet

[private]
_janet-format:
  [ -f ".jpm_tree/bin/janet-format" ] || jpm install spork
