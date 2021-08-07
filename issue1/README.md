`gotip build -gcflags=all=-G=3 main.go`


```
# command-line-arguments
main.go:5:3: internal compiler error: cannot export = (22) node
        ==> please file an issue and assign to gri@

goroutine 1 [running]:
runtime/debug.Stack()
        /home/enikolov/sdk/gotip/src/runtime/debug/stack.go:24 +0x65
cmd/compile/internal/base.FatalfAt({0x4913b0, 0xc0}, {0xd0be24, 0x47}, {0xc000455418, 0x2, 0x2})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/base/print.go:227 +0x154
cmd/compile/internal/base.Fatalf(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/base/print.go:196
cmd/compile/internal/typecheck.(*exportWriter).expr(0xc00049d7a0, {0xe575f0, 0xc0004913b0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:2056 +0x19c5
cmd/compile/internal/typecheck.(*exportWriter).exprList(0xc00049d7a0, {0xc000488340, 0x1, 0x604639})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1618 +0x76
cmd/compile/internal/typecheck.(*exportWriter).stmt(0xc00049d7a0, {0xe59468, 0xc000490460})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1512 +0x83f
cmd/compile/internal/typecheck.(*exportWriter).node(0xc000455546, {0xe59468, 0xc000490460})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1438 +0x65
cmd/compile/internal/typecheck.(*exportWriter).stmtList(0xc00049d7a0, {0xc000488310, 0x1, 0xc000455596})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1431 +0x76
cmd/compile/internal/typecheck.(*exportWriter).expr(0xc00049d7a0, {0xe57c30, 0xc000490280})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1732 +0x1efa
cmd/compile/internal/typecheck.(*exportWriter).stmt(0xc00049d7a0, {0xe575f0, 0xc0004904b0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1484 +0x22d
cmd/compile/internal/typecheck.(*exportWriter).node(0x30, {0xe575f0, 0xc0004904b0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1438 +0x65
cmd/compile/internal/typecheck.(*exportWriter).stmtList(0xc00049d7a0, {0xc0004881a0, 0x1, 0xcb7140})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1431 +0x76
cmd/compile/internal/typecheck.(*exportWriter).funcBody(0xc00049d7a0, 0xc00048a2c0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1426 +0x5b
cmd/compile/internal/typecheck.(*iexporter).doInline(0xc0003ac370, 0xc00049a5b0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:619 +0xc5
cmd/compile/internal/typecheck.(*exportWriter).funcExt(0xc00049cd20, 0xc00049a5b0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:1363 +0x1b6
cmd/compile/internal/typecheck.(*iexporter).doDecl(0xc0003ac370, 0xc00049a5b0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:503 +0x1d7
cmd/compile/internal/typecheck.WriteExports({0xe3f160, 0xc0000a7530}, 0x1)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:300 +0x2eb
cmd/compile/internal/noder.WriteExports(0xc00007e330)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/export.go:40 +0x7a
cmd/compile/internal/gc.dumpCompilerObj(0xc00007e330)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/obj.go:107 +0x28
cmd/compile/internal/gc.dumpobj1({0x7ffd056a0007, 0x24}, 0x3)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/obj.go:63 +0x185
cmd/compile/internal/gc.dumpobj()
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/obj.go:44 +0x36
cmd/compile/internal/gc.Main(0xd13510)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/main.go:321 +0x10d6
main.main()
        /home/enikolov/sdk/gotip/src/cmd/compile/main.go:55 +0xdd
```