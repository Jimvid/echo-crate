# Echo-crate

## MakeFile

Build the application

```bash
make build
```

Run the application

```bash
make run
```

live reload the application

```bash
make watch
```

Create DB container

```bash
make docker-up
```

Shutdown DB container

```bash
make docker-down
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
make build-tailwind
```
