<!--
 Copyright (c) 2022 aiocat
 
 This software is released under the MIT License.
 https://opensource.org/licenses/MIT
-->

# Cystem
Go binding for C `int system(const char* cmd)` command.

## Install
- You must have `gcc` to use this library.
- `go get -u github.com/aiocat/cystem`

## Usage
for `string`:
```go
cystem.RunString("echo Hello World")
```

for `[]string`:
```go
cystem.RunStringSlice([]string{
    "echo",
    "Hello",
    "World",
})
```

## License
Cystem is distributed under MIT license. for more information:
- https://raw.githubusercontent.com/aiocat/cystem/main/LICENSE