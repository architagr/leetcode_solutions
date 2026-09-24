# Workshop tooling for "It worked in dev".
#
# The repo is 17 separate Go modules - every episode is standalone so that a
# reader can clone one directory and run it - so there is no go.work and no
# single `go test ./...` that covers everything. These targets scope to the
# workshop module and leave the episodes alone.
#
#   make verify   the pre-work check. Run this before the workshop.
#   make bench    the benchmark harness: -count=10 through benchstat.
#   make clean    remove generated benchmark output.

WORKSHOP := series/it-worked-in-dev/workshop
GOBIN    := $(shell go env GOPATH)/bin
BENCHSTAT := $(GOBIN)/benchstat

# Benchmarks are run a fixed number of times rather than for a fixed duration.
# -benchtime=200x makes two runs on the same machine comparable, which is the
# entire point of running them ten times and handing the output to benchstat.
#
# It applies to `bench`, not to `verify`. The setup check measures a
# sub-nanosecond function, and 200 iterations of one of those is mostly timer
# resolution - it showed the effect at 2x instead of 3.3x and buried the point.
# verify uses the default 1s, which is the right call for a microbenchmark.
BENCHTIME := 200x
COUNT     := 10

# verify measures a sub-nanosecond function on whatever laptop an attendee
# happens to own, often a busy one. A single run there is worthless - measured
# on a loaded machine it reported the SumTo pair 1.3x apart when they are the
# same, which would teach the opposite of the lesson. Ten short runs through
# benchstat instead: it prints the variance, so noise shows up as noise.
VERIFYTIME  := 100ms
VERIFYCOUNT := 10
VERIFYOUT   := verify-bench.txt

GREEN := \033[32m
BOLD  := \033[1m
DIM   := \033[2m
OFF   := \033[0m

.PHONY: verify bench clean help

help:
	@echo "make verify   pre-work check - run this before the workshop"
	@echo "make bench    run the harness: -count=$(COUNT) through benchstat"
	@echo "make clean    remove generated benchmark output"

## verify - everything an attendee needs working, checked in about ten seconds.
##
## Deliberately more than "does go exist". It compiles the module, runs the
## tests, installs benchstat, and then runs two benchmarks whose numbers
## disagree - so the check doubles as the first lesson of Day 1.
verify:
	@printf "$(BOLD)It worked in dev - workshop setup check$(OFF)\n\n"

	@command -v go >/dev/null 2>&1 || { \
	  echo "go is not on your PATH. Install Go 1.21+ from https://go.dev/dl/"; exit 1; }
	@go version | awk '{print "  go        " $$3}'

	@go version | grep -Eq 'go1\.(2[1-9]|[3-9][0-9])' || { \
	  echo; echo "  Go 1.21 or newer is required. Yours is older."; \
	  echo "  Update from https://go.dev/dl/ and run make verify again."; exit 1; }

	@cd $(WORKSHOP) && go build ./... && echo "  build     ok"
	@cd $(WORKSHOP) && go test -count=1 ./... >/dev/null && echo "  tests     ok"

	@test -x $(BENCHSTAT) || { \
	  printf "  benchstat installing (one time, needs network)\n"; \
	  go install golang.org/x/perf/cmd/benchstat@latest; }
	@test -x $(BENCHSTAT) && echo "  benchstat ok"

	@printf "\n$(GREEN)$(BOLD)  Setup complete. You are ready for Saturday.$(OFF)\n\n"

	@printf "$(BOLD)Now the part worth looking at.$(OFF)\n"
	@printf "$(DIM)Each pair below runs the same function with the same argument, once$(OFF)\n"
	@printf "$(DIM)throwing the result away and once keeping it. Same work, same cost.$(OFF)\n"
	@printf "$(DIM)Ten runs each, through benchstat, because one run is not evidence -$(OFF)\n"
	@printf "$(DIM)takes about ten seconds.$(OFF)\n\n"
	@cd $(WORKSHOP) && go test -bench=. -benchtime=$(VERIFYTIME) -count=$(VERIFYCOUNT) -run=XXX . > $(VERIFYOUT) 2>&1
	@cd $(WORKSHOP) && $(BENCHSTAT) $(VERIFYOUT) | sed 's/^/  /'
	@rm -f $(WORKSHOP)/$(VERIFYOUT)
	@printf "\n$(DIM)  Read the two pairs against each other. Mix is about 3x apart. SumTo$(OFF)\n"
	@printf "$(DIM)  is the same on both sides once you allow for the +/- benchstat prints.$(OFF)\n\n"
	@printf "$(DIM)  The fast half of Mix never ran. A third of a nanosecond is roughly one$(OFF)\n"
	@printf "$(DIM)  clock cycle, and five operations do not happen in one cycle - the$(OFF)\n"
	@printf "$(DIM)  number is its own tell.$(OFF)\n\n"
	@printf "$(DIM)  Why one pair and not the other is hour one of Day 1. The obvious$(OFF)\n"
	@printf "$(DIM)  answer - inlining - is wrong; check with go test -gcflags=-m.$(OFF)\n"
	@printf "$(DIM)  Bring the question, not the answer.$(OFF)\n\n"
	@printf "  Add benchstat to your PATH if it is not already:\n"
	@printf '$(DIM)    export PATH="$$(go env GOPATH)/bin:$$PATH"$(OFF)\n'

## bench - the harness attendees keep.
##
## Ten runs, fixed iteration count, piped through benchstat. One run is not
## evidence: it reports variance, not just a number, and a difference it marks
## "~" is a difference you cannot claim.
bench: | $(BENCHSTAT)
	@cd $(WORKSHOP) && go test -bench=. -benchtime=$(BENCHTIME) -count=$(COUNT) -run=XXX . \
	  | tee bench.txt
	@echo
	@cd $(WORKSHOP) && $(BENCHSTAT) bench.txt

$(BENCHSTAT):
	go install golang.org/x/perf/cmd/benchstat@latest

clean:
	@rm -f $(WORKSHOP)/bench.txt $(WORKSHOP)/$(VERIFYOUT)
	@echo "removed $(WORKSHOP)/bench.txt"
