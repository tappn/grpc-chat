.PHONY: gen
gen:
	protoc --proto_path=proto \
		--go_out=gen/go \
		--go_opt=paths=source_relative \
		--doc_out=doc \
		--doc_opt=markdown,docs.md,source_relative \
		${PROTO}
