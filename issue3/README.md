`gotip build -gcflags=all=-G=3 main.go`


```
# command-line-arguments
./main.go:9:6: internal compiler error: mayCall 
.   DYNAMICDOTTYPE tc(1) .shape.string # main.go:9:6 .shape.string
.   .   NAME-main.x esc(no) tc(1) Class:PAUTO Offset:0 OnStack INTER-interface {} # main.go:8:2
.   DYNAMICDOTTYPE-T
.   .   INDEX tc(1) Bounded PTR-*uint8 # main.go:9:6 PTR-*uint8
.   .   .   DEREF tc(1) ARRAY-[3]uintptr # main.go:9:6 ARRAY-[3]uintptr
.   .   .   .   CONVNOP tc(1) PTR-*[3]uintptr # main.go:9:6 PTR-*[3]uintptr
.   .   .   .   .   CONVNOP tc(1) UNSAFEPTR-unsafe.Pointer # main.go:9:6 UNSAFEPTR-unsafe.Pointer
.   .   .   .   .   .   NAME-main..dict esc(no) tc(1) Class:PPARAM Offset:0 OnStack uintptr # main.go:7:6
.   .   .   LITERAL-0 tc(1) uintptr # main.go:7:6

goroutine 1 [running]:
runtime/debug.Stack()
        /home/enikolov/sdk/gotip/src/runtime/debug/stack.go:24 +0x65
cmd/compile/internal/base.FatalfAt({0x486ea0, 0xc0}, {0xce767d, 0xb}, {0xc0000c1628, 0x1, 0x1})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/base/print.go:227 +0x154
cmd/compile/internal/walk.mayCall.func2({0xe58220, 0xc000486ea0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/walk.go:308 +0x3c9
cmd/compile/internal/ir.Any.func1({0xe58220, 0xc000486ea0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:130 +0x30
cmd/compile/internal/ir.Any({0xe58220, 0xc000486ea0}, 0xc00007ebe0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:132 +0xb8
cmd/compile/internal/walk.mayCall({0xe58220, 0xc000486ea0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/walk.go:299 +0x73
cmd/compile/internal/walk.walkCall1(0xc0004b8510, 0xc0000c1aa0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/expr.go:586 +0x2f8
cmd/compile/internal/walk.walkCall(0xc0004b8510, 0xe57730)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/expr.go:556 +0x6b2
cmd/compile/internal/walk.walkExpr1({0xe57b18, 0xc0004b8510}, 0xc0004b8510)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/expr.go:176 +0x485
cmd/compile/internal/walk.walkExpr({0xe57b18, 0xc0004b8510}, 0xc0000c1aa0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/expr.go:55 +0x3a5
cmd/compile/internal/walk.walkStmt({0xe57b18, 0xc0004b8510})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/stmt.go:57 +0x59d
cmd/compile/internal/walk.walkStmtList({0xc000434880, 0x5, 0xc00048a6e0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/stmt.go:167 +0x5b
cmd/compile/internal/walk.Walk(0xc00048a6e0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/walk/walk.go:43 +0xef
cmd/compile/internal/gc.prepareFunc(0xc00048a6e0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/compile.go:92 +0x6d
cmd/compile/internal/gc.enqueueFunc(0xc00048a6e0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/compile.go:66 +0x2f7
cmd/compile/internal/gc.Main(0xd13700)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/main.go:292 +0xe5f
main.main()
        /home/enikolov/sdk/gotip/src/cmd/compile/main.go:55 +0xdd

```