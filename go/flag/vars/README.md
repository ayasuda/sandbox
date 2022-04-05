# A sample code of how to define flags by assinge variable

can set `--bool-flag`, `--int-flag` flag or else.

## Start to run

```
$ go run main.go \
  --bool-flag \
  --int-flag 47 \
  --int64-flag 53 \
  --uint-flag 59 \
  --uint64-flag 61 \
  --float64-flag 3.14 \
  --string-flag "foobar" \
  --duration-flag 12h34m56s
bool-flag has value true
int-flag has value 47
int64-flag has value 53
uint-flag has value 59
uint64-flag has value 61
float64-flag has value 3.140000
string-flag has value 'foobar'
duration-flag has value 45296000000000
```
