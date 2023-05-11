# go-fitstatusb

A golang wrapper module to interact with [fit-statUSB](https://fit-iot.com/web/product/fit-statusb/) device over serial. This is experimental at best, not ready for production.

## Background

Found `fit-statUSB` on [Amazon](https://www.amazon.com/dp/B07CKFLQ5V) on a whim. I am originally a fan of [ThingM's Blink(1)](https://www.amazon.com/ThingM-Blink-USB-RGB-BLINK1MK3/dp/B07Q8944QK), but wanted to try an alternative.

## Documentation

Technical documentation for `fit-statUSB` is located at: <http://fit-pc.com/wiki/index.php/Fit-statUSB>

## Troubleshooting

I'm getting  a `permission denied` when running `go test`?

Try adding your `$USER` to the `dialout` group.

    `sudo adduser $USER dialout`

Logout of `$USER` account for changes to be applied.

## Caveats

* Right now this only works on Linux, would love to support Windows, and macOS.


## Bugs and Contribution

Please feel free to reach out. Issues and PR's are always welcome!
