currentBranch = $(shell git symbolic-ref HEAD | sed -e 's,.*/\(.*\),\1,')

.PHONY: gen
gen:
	docker run --rm -v "$(shell pwd):/work" uber/prototool prototool all

.PHONY: lint
lint:
	docker run --rm -v "$(shell pwd):/work" uber/prototool prototool lint

.PHONY: install
install:
	@echo "当前git分支 $(currentBranch)"
	rm -rf .subsplit
	git subsplit init .
	git subsplit publish  --no-tags --heads=$(currentBranch) gen/go:git@git.github.com/donech/proto-go.git
	rm -rf .subsplit