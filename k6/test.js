import { Trend } from 'k6/metrics';
import http from 'k6/http';
import { check,sleep } from 'k6';

export const options = {
  vus: 1, 
  duration: '30s', 
};

const trend1 = new Trend('check', true);


export default function () {

  const postUrl = 'http://localhost:8888/check';
  let payload = {
    cg: "cg1",
    cr: "cr14",
    action: "READ",
    resource: "UserProfile"
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(postUrl, JSON.stringify(payload), params);

  check(res, {
    'post status was 200': (r) => r.status === 200,
  });
  
  if (res.status == 200) {
    trend1.add(res.timings.duration);
  }
  

}