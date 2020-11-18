# k6-plugin-notification

This is a [k6](https://github.com/loadimpact/k6) plugin using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Clone `xk6`:
  ```shell
  git clone https://github.com/k6io/xk6.git
  cd xk6
  ```

2. Build the binary:
  ```shell
  CGO_ENABLED=1 go run ./cmd/xk6/main.go build master \
    --with github.com/dgzlopes/k6-plugin-notification
  ```

## Example

```javascript
import http from 'k6/http';
import notification from 'k6/x/notification';

const url = "slack://token-a/token-b/token-c"

export function setup() {
  notification.send(url,"Starting test");
}

export default function () {
  http.get('http://test.k6.io');
}

export function teardown(data) {
  notification.send(url,"Finishing test");
}
```
