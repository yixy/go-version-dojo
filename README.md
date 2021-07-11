# go-version-dojo #

## version >=2 的项目结构测试 ##

```
// valid dependencies
require github.com/yixy/go-version-dojo/v2 v1.0.0
require github.com/yixy/go-version-dojo/v2 v1.0.1
require github.com/yixy/go-version-dojo/v2 v1.0.2
require github.com/yixy/go-version-dojo/v2 v2.0.0

// invalid dependency ( note: go.mod not define in subdir)
require github.com/yixy/go-version-dojo/v2 v3.0.0
```

## different major version is not interaction ##

go.mod定义：

```
module test

go 1.16

require (
	github.com/yixy/go-version-dojo/hello v1.0.0
	github.com/yixy/go-version-dojo/v2 v2.0.0
)

```

module graph：

```
$ go mod graph
test github.com/yixy/go-version-dojo/hello@v1.0.0
test github.com/yixy/go-version-dojo/v2@v2.0.0
github.com/yixy/go-version-dojo/hello@v1.0.0 github.com/yixy/go-version-dojo@v1.0.1

```

hello-v1.0.0的依赖仍为v1.0.1，并没有被MVS替换为v2.0.0。

```
$ go run .
hello-v1.0.0 v1.0.1
v2.0.0
```

## same major version is interaction ##

###### main module is lower ######

```
module test

go 1.16

require (
	github.com/yixy/go-version-dojo v1.0.0
	github.com/yixy/go-version-dojo/hello v1.0.0
)
```

注意，go mod tidy 后 主module的直接依赖版本会自动调整

```
$ go mod graph
test github.com/yixy/go-version-dojo@v1.0.1
test github.com/yixy/go-version-dojo/hello@v1.0.0
github.com/yixy/go-version-dojo/hello@v1.0.0 github.com/yixy/go-version-dojo@v1.0.1
```

```
$ go run .
hello-v1.0.0 v1.0.1
v1.0.1
```

###### main module is higher ######

```
module test

go 1.16

require (
	github.com/yixy/go-version-dojo v1.0.2
	github.com/yixy/go-version-dojo/hello v1.0.0
)
```


```
$ go mod graph
test github.com/yixy/go-version-dojo@v1.0.2
test github.com/yixy/go-version-dojo/hello@v1.0.0
github.com/yixy/go-version-dojo/hello@v1.0.0 github.com/yixy/go-version-dojo@v1.0.1

```

```
$ go run .
hello-v1.0.0 v1.0.2
v1.0.2
```


## repalce directive ##


replace/exclude directives only apply in the main module's go.mod file and are ignored in other modules.

main module的例子：

```
module test

go 1.16

require (
	github.com/yixy/go-version-dojo/hello v1.0.0
	github.com/yixy/go-version-dojo/v2 v2.0.0
)

replace github.com/yixy/go-version-dojo v1.0.1 => github.com/yixy/go-version-dojo v1.0.2
```

```
$ go mod graph
test github.com/yixy/go-version-dojo/hello@v1.0.0
test github.com/yixy/go-version-dojo/v2@v2.0.0
github.com/yixy/go-version-dojo/hello@v1.0.0 github.com/yixy/go-version-dojo@v1.0.1
```

```
$ go run .
hello-v1.0.0 v1.0.2
v2.0.0
```

other module的例子：

```
module test

go 1.16

require (
	github.com/yixy/go-version-dojo v1.0.0
	github.com/yixy/go-version-dojo/hello v1.0.1
)
```

```
$ go mod graph
test github.com/yixy/go-version-dojo@v1.0.1
test github.com/yixy/go-version-dojo/hello@v1.0.1
github.com/yixy/go-version-dojo/hello@v1.0.1 github.com/yixy/go-version-dojo@v1.0.1
```

注意，go mod tidy 后 主module的直接依赖版本会自动调整

```
$ go run .
hello-v1.0.1 v1.0.1
v1.0.1
```