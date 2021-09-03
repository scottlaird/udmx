# udmx

`udmx` provides a Go interface for talking to cheap "udmx"
USB-[DMX](https://en.wikipedia.org/wiki/DMX512) interfaces such as
[this one](https://www.amazon.com/dp/B07GT3S6V6).

This was developed using
[markusb/uDMX-linux](https://github.com/markusb/uDMX-linux) as
documentation of the interface to udmx devices, but doesn't reuse any
of the code.

## Example

```
import (
  "github.com/google/gousb"
  "github.com/scottlaird/udmx"
)

func main() {
  ctx := gousb.NewContext()
  udmx, err := udmx.NewUDMXDevice(ctx)
  if err != nil {
    panic(err)
  }

  // Set DMX device #5 to `17`.
  udmx.Set(5, 17)
}
```

## Disclaimer

This is not an official Google project.
