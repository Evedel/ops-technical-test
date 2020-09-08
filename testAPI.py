#!/usr/bin/env python

import time
import requests

addr  = 'http://127.0.0.1:8080'
tests = [
  ('/', 'hello world\n'),
  ('/healthcheck', ''),
  ('/metadata', '{"myapplication":[{"version":"1.0","description":"pre-interview technical test","lastcommitsha":"abc57858585"}]}\n')
]

time.sleep(1)
for test in tests:
  path = test[0]
  want = test[1]
  response = requests.get(addr+path)
  if response.text != want:
    print("Request: " + addr+path + " want: " + want + " have: " + response.text)
    exit(1)

print('OK')
exit(0)