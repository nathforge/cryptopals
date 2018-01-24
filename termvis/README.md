# termvis

Just enough terminal emulator to GIFerise my Cryptopals solutions.

Saves a sequentially-numbered frame before every screen clear.

## Example

```
$ go run $GOPATH/src/github.com/nathforge/termvis/main.go \
    --font-filename=/System/Library/Fonts/Menlo.ttc \
    --output-path=frames \
    --term-width=62 \
    --term-height=8
```
