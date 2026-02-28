import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 100 },  // Ramp up to 100 users
    { duration: '1m', target: 500 },   // Stay at 500 users (Heavy Load)
    { duration: '30s', target: 0 },    // Ramp down
  ],
};

export default function () {
  // Use a short code you know exists in your DB
  const res = http.get('http://localhost:8080/b',{redirect:0});
  
  // Verify we are getting a 302 Redirect
  check(res, {
    'is status 302': (r) => r.status === 302,
  });

  sleep(0.01); // Simulate very fast users
}