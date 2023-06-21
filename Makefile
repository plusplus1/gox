
.PHONY: format
format: gomod
	@for s in $(shell git status -s | grep -e '.go$$'|grep -v '^D ' |awk '{print $$NF}'); do \
		echo "go fmt $$s";\
		go fmt $$s;\
	done

.PHONY: gomod
gomod:
	go mod tidy
	go mod vendor


