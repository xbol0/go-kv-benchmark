# pure go kv database benchmark

ref: https://github.com/recoilme/pogreb-bench

- pogred
- pudge
- buntdb

## Usage

```
# Clone this repo and install dependencies
go run .
```

## Result

Read: pogreb > buntdb > pudge
Write: buntdb > pudge > pogreb
Disk usage: pudge > buntdb > pogreb

```
Now test 10000 count:
pudge: Total usage: 103ms, Init usage: 0ms, Read usage: 16ms, Write usage: 87ms, File size: 0M
buntdb: Total usage: 39ms, Init usage: 1ms, Read usage: 4ms, Write usage: 34ms, File size: 0M
pogreb: Total usage: 67ms, Init usage: 1ms, Read usage: 2ms, Write usage: 64ms, File size: 0M
Now test 100000 count:
pudge: Total usage: 676ms, Init usage: 0ms, Read usage: 151ms, Write usage: 525ms, File size: 6M
buntdb: Total usage: 356ms, Init usage: 0ms, Read usage: 46ms, Write usage: 310ms, File size: 7M
pogreb: Total usage: 601ms, Init usage: 1ms, Read usage: 22ms, Write usage: 578ms, File size: 8M
Now test 1000000 count:
pudge: Total usage: 7017ms, Init usage: 0ms, Read usage: 1612ms, Write usage: 5405ms, File size: 64M
buntdb: Total usage: 5472ms, Init usage: 0ms, Read usage: 1410ms, Write usage: 4062ms, File size: 76M
pogreb: Total usage: 6581ms, Init usage: 0ms, Read usage: 390ms, Write usage: 6191ms, File size: 81M
Now test 5000000 count:
pudge: Total usage: 46345ms, Init usage: 0ms, Read usage: 14259ms, Write usage: 32086ms, File size: 324M
buntdb: Total usage: 36822ms, Init usage: 0ms, Read usage: 9942ms, Write usage: 26880ms, File size: 360M
pogreb: Total usage: 37205ms, Init usage: 1ms, Read usage: 2331ms, Write usage: 34873ms, File size: 407M
```

1000op/s on 1m(large is greater):
|db|read|write|
|:-:|:-:|:-:|
|pudge|620|185|
|buntdb|709|246|
|pogreb|2564|161|

1000op/s on 5m(large is greater):
|db|read|write|
|:-:|:-:|:-:|
|pudge|350|156|
|buntdb|502|186|
|pogreb|2145|143|