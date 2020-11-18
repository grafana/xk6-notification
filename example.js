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