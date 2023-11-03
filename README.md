
# Qrcode Decoder

A simple Qrcode decoder http service

Thanks to [GoCV](https://github.com/hybridgroup/gocv)

Docker image: `zxfishhack/qrcode`

## Usage

url: /
content-type: multipart/form

|field| description             |                  |
|---|-------------------------|------------------|
|url| url of qrcode image     | optional         |
|file| content of qrcode image | optional,higher priority |


## Benchmark

Machine: MacBook Pro 14" 2021, Apple M1 Pro (10 Cores), 32G

```bash
hey  -z 1m -m POST -H "Content-Type: multipart/form-data; boundary=862d8f33b69588ae76771bcbdaf6037586418b21b042fbc9ba017ed75a9f" -T "multipart/form-data" -D req.txt http://localhost:8080/
```

```
goos: darwin
goarch: arm64
pkg: github.com/zxfishhack/qrcode
BenchmarkDecodeQrCode
BenchmarkDecodeQrCode-10    	     138	   8824558 ns/op	     360 B/op	      14 allocs/op
PASS
```