import { Trend } from 'k6/metrics';
import http from 'k6/http';
import { check,sleep } from 'k6';

export const options = {
  vus: 1, 
  duration: '30s', 
};

const trend1 = new Trend('check_cg_for_access_to_cr', true);


export default function () {

  //  random number from 1 - 10
  let cgId = Math.floor(Math.random() * 10 - 1 + 1) + 1
  // random number from 1 to 1000
  let crId = Math.floor(Math.random() * 1000 - 1 + 1) + 1

// console.log("cg" + cgId)
// console.log("cr" + cgId + "::" + crId)
  

  const postUrl = 'http://localhost:8888/check';
  let payload = {
    cg: "cg" +  cgId,
    cr: "cr" + cgId + "::" + crId,
    action: "READ",
    resource: "UserProfile"
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(postUrl, JSON.stringify(payload), params);
  // console.log(res.body)

  check(res, {
    'post status was 200 and check is true': (r) => r.status === 200 && res.body.includes("true"),
  });
  
  if (res.status == 200) {
    trend1.add(res.timings.duration);
  }
  

}