run:
	go run app/src/app/app.go

test:
	# ----------------- preparing workspace ---------------------------
	mkdir -p coverage
	# ----------------- running tests ---------------------------------
	go test -v -race -coverprofile=coverage/coverage-usecase.out -covermode=atomic ./domain/usecase/...
	
	go test -v -race -coverprofile=coverage/coverage-postgresql-adapter-raw.out -covermode=atomic \
		./infrastructure/adapters/postgresql-adapter/...
	
	go test -v -race -coverprofile=coverage/coverage-web-api-raw.out -covermode=atomic \
		./infrastructure/entry-points/web-api/...

	go test -v -race -coverprofile=coverage/coverage-validation-raw.out -covermode=atomic \
		./infrastructure/helpers/validation/...

	# ------------------ preparing coverage output --------------------
	sed 1,1d coverage/coverage-postgresql-adapter-raw.out | cat > coverage/coverage-postgresql-adapter.out
	sed 1,1d coverage/coverage-web-api-raw.out | cat > coverage/coverage-web-api.out
	sed 1,1d coverage/coverage-validation-raw.out | cat > coverage/coverage-validation.out

	cat coverage/coverage-usecase.out >> coverage.out
	cat coverage/coverage-postgresql-adapter.out >> coverage.out
	# cat coverage/coverage-web-api.out >> coverage.out
	cat coverage/coverage-validation.out >> coverage.out

	# ------------------- showing coverage ----------------------------
	go tool cover -html=coverage.out

	# ------------------ removing files -------------------------------
	rm -r coverage coverage.out



run:
	go run app/src/app/app.go