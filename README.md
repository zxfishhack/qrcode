
# Qrcode Decoder

A simple Qrcode decoder http service

Thanks to [GoCV](https://github.com/hybridgroup/gocv)

Docker image: `zxfishhack/qrcode`

## Benchmark

Machine: MacBook Pro 14" 2021, Apple M1 Pro (10 Cores), 32G

```bash
hey  -z 1m -m POST -H "Content-Type: multipart/form-data; boundary=862d8f33b69588ae76771bcbdaf6037586418b21b042fbc9ba017ed75a9f" -T "multipart/form-data" -D req.txt http://localhost:8080/
```

```
Summary:
  Total:	60.7083 secs
  Slowest:	1.8738 secs
  Fastest:	0.5008 secs
  Average:	1.2418 secs
  Requests/sec:	40.0604

  Total data:	97280 bytes
  Size/request:	40 bytes

Response time histogram:
  0.501 [1]	|
  0.638 [1]	|
  0.775 [16]	|■
  0.913 [42]	|■■
  1.050 [179]	|■■■■■■■■
  1.187 [596]	|■■■■■■■■■■■■■■■■■■■■■■■■■■
  1.325 [909]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  1.462 [504]	|■■■■■■■■■■■■■■■■■■■■■■
  1.599 [154]	|■■■■■■■
  1.736 [27]	|■
  1.874 [3]	|


Latency distribution:
  10% in 1.0526 secs
  25% in 1.1469 secs
  50% in 1.2449 secs
  75% in 1.3387 secs
  90% in 1.4303 secs
  95% in 1.5024 secs
  99% in 1.6093 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0002 secs, 0.5008 secs, 1.8738 secs
  DNS-lookup:	0.0001 secs, 0.0000 secs, 0.0040 secs
  req write:	0.0010 secs, 0.0002 secs, 0.0193 secs
  resp wait:	1.2405 secs, 0.5004 secs, 1.8725 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0041 secs

Status code distribution:
  [200]	2432 responses
```