# Echo-crate

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```

## Tailwind

Download the Tailwind standalone CLI for your platform [here](https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.4.1).

```sh
# Example for macOS arm64
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tailwind
```

Then run the following command to rebuild styles on change:

```bash
make watch-tailwind
```

Build CSS for production

```bash
make watch-tailwind
```
