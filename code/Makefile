LIMIT ?= 15


all: bindata.go

bindata.go: dump.txt
	go-bindata -pkg moulcode dump.txt repos-list.txt

dump.txt: dumps
	cat tmp/*/dump.txt > $@
	wc dump.txt
	ls -l dump.txt


.PHONY: dumps
dumps: clones
	cd tmp; for repo in *; do ( \
	  cd $$repo; \
	  git grep . 2>/dev/null | grep -vE '(vendor/|bower_components/)' | grep -v '^Binary file ' | grep -vE '(LICENSE|\.godir|\.gitignore|\.playground|COPYING|Readme|dsp\.js|tgmath\.h|\.config):' | grep -vE '(\.md|\.yml|\.json|\.txt):' > dump.txt; \
	  wc dump.txt; \
	); done
	cat tmp/*/dump.txt | cut -d: -f1 | uniq -c | sort -n | tail -n 20


.PHONY: clones
clones: repos-list.txt
	mkdir -p tmp
	for url in $(shell cat $< | head -n $(LIMIT)); do (cd tmp; git clone $$url || true); done


repos-list.txt:
	( \
	  curl https://api.github.com/users/moul/repos\?sort\=updated\&type\=owner\&page=1 | jq -r '.[].clone_url'; \
	  curl https://api.github.com/users/moul/repos\?sort\=updated\&type\=owner\&page=2 | jq -r '.[].clone_url'; \
	  curl https://api.github.com/users/moul/repos\?sort\=updated\&type\=owner\&page=3 | jq -r '.[].clone_url'; \
	) > $@


clean:
	rm -f repos-list.txt
	rm -rf tmp
