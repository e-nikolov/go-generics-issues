`gotip build -gcflags=all=-G=3 main.go`


```
# command-line-arguments
./main.go:12:13: internal compiler error: assertion failed

goroutine 1 [running]:
runtime/debug.Stack()
        /home/enikolov/sdk/gotip/src/runtime/debug/stack.go:24 +0x65
cmd/compile/internal/base.FatalfAt({0x4968c0, 0xc0}, {0xcec5f4, 0x10}, {0x0, 0x0, 0x0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/base/print.go:227 +0x154
cmd/compile/internal/base.Fatalf(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/base/print.go:196
cmd/compile/internal/base.Assert(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/base/print.go:239
cmd/compile/internal/noder.assert(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:27
cmd/compile/internal/noder.(*subster).node.func1({0xe59f08, 0xc0004868a0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:1122 +0x8ad
cmd/compile/internal/ir.(*AssignStmt).editChildren(0xc0004916d0, 0xc00048c270)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/node_gen.go:163 +0xb5
cmd/compile/internal/ir.EditChildren(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:185
cmd/compile/internal/noder.(*subster).node.func1({0xe57730, 0xc0004905f0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:898 +0x416
cmd/compile/internal/noder.(*subster).node(0xc0004ac080, {0xe57730, 0xc0004905f0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:1147 +0xa5
cmd/compile/internal/noder.(*subster).list(0x13b0040, {0xc000488220, 0x1, 0xc00013e5f0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:1230 +0x8e
cmd/compile/internal/noder.(*irgen).genericSubst(0xc0004be000, 0xc000491540, 0xc00049a9c0, {0xc000484178, 0x1, 0x1}, 0x0, 0xc0004968c0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:711 +0xcc6
cmd/compile/internal/noder.(*irgen).getInstantiation(0xc0004be000, 0xc00049a9c0, {0xc000484170, 0x1, 0x1}, 0x0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:597 +0x33d
cmd/compile/internal/noder.(*irgen).stencil.func1({0xe57b18, 0xc0004c21b0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:109 +0x2df
cmd/compile/internal/ir.Visit.func1({0xe57b18, 0xc0004c21b0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:105 +0x30
cmd/compile/internal/ir.doNodes({0xc000488270, 0x1, 0x0}, 0xc00048c210)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/node_gen.go:1484 +0x67
cmd/compile/internal/ir.(*Func).doChildren(0xe583b0, 0xc00048a580)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/func.go:152 +0x2e
cmd/compile/internal/ir.DoChildren(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:94
cmd/compile/internal/ir.Visit.func1({0xe583b0, 0xc00048a580})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:106 +0x57
cmd/compile/internal/ir.Visit({0xe583b0, 0xc00048a580}, 0xc000496880)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:108 +0xb8
cmd/compile/internal/noder.(*irgen).stencil(0xc0004be000)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:91 +0x265
cmd/compile/internal/noder.(*irgen).generate(0xc0004be000, {0xc00011eb40, 0x2, 0x203000})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/irgen.go:261 +0x259
cmd/compile/internal/noder.check2({0xc00011eb40, 0x2, 0x2})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/irgen.go:92 +0x16d
cmd/compile/internal/noder.LoadPackage({0xc00013a110, 0x2, 0x0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/noder.go:90 +0x365
cmd/compile/internal/gc.Main(0xd13700)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/main.go:190 +0xaf3
main.main()
        /home/enikolov/sdk/gotip/src/cmd/compile/main.go:55 +0xdd
```