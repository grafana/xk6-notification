import http from 'k6/http';
import notification from 'k6/x/notification';

const url = 'slack://xoxb:123456789012-1234567890123-4mt0t4l1YL3g1T5L4cK70k3N@C001CH4NN3L'

export function setup() {
  notification.send(url,"Starting test");
}

export default function () {
  http.get('http://test.k6.io');
}

export function teardown(data) {
  notification.send(url,"Finishing test");
}