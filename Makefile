vet:
	go vet ./...

fmt: vet
	go vet ./...

test: fmt
	richgo test ./...

templ:
	templ generate

build: templ
	go build -o ./bin/todo-app cmd/app/main.go

run: build
	./bin/todo-app

live/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v --proxybind="0.0.0.0" --proxyport="7331"

live/server:
	air \
	--build.cmd "go build -o bin/todo-app cmd/app/main.go" \
	--build.bin "bin/todo-app" \
	--build.delay "100" \
	--build.exclude_dir "node_modules,internal/assets/static" \
	--build.include_dir "cmd,internal" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
live/tailwind:
	npx --yes tailwindcss -i ./input.css -o ./assets/styles.css --minify --watch

# run esbuild to generate the index.js bundle in watch mode.
live/esbuild:
	npx --yes esbuild js/index.ts --bundle --outdir=assets/ --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
live/sync_assets:
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin ":" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "internal/assets/static" \
	--build.include_ext "js,css"

# start all 5 watch processes in parallel.
live: 
	# make -j5 live/templ live/server live/tailwind live/esbuild live/sync_assets
	make -j3 live/templ live/server live/sync_assets
