#
# Makefile for wasm-example
#
.PHONY: usage edit build clean git
#----------------------------------------------------------------------------------
usage:
	@echo "make [edit|build]"
#----------------------------------------------------------------------------------
edit e:
	@echo "make (edit:e) [history]"
edit-go eg:
	vi main.go
edit-history eh:
	vi HISTORY.md
#----------------------------------------------------------------------------------
build b:
#----------------------------------------------------------------------------------
clean:
	rm -f bin/*
	docker system prune
#----------------------------------------------------------------------------------
run r:
	@echo "make (run:r) [test]"

run-test rt:
	go run main.go
#----------------------------------------------------------------------------------
git g:
	@echo "make (git:g) [update|store]"
git-update gu:
	git add .gitignore *.md Makefile simple/ chat/
	git commit -m "Initial commit to start coding"
	git push
git-store gs:
	git config credential.helper store
#----------------------------------------------------------------------------------

