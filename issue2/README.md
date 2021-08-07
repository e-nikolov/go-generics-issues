`gotip build -gcflags=all=-G=3 main.go`


```
# command-line-arguments
panic: runtime error: index out of range [1] with length 1

goroutine 1 [running]:
cmd/compile/internal/noder.typecheckaste(0xd0, {0xe44468, 0xc0001397c0}, 0x0, 0x8, {0xc00011ee20, 0x1, 0xc00043bef0}, 0x1)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/transform.go:511 +0x3d5
cmd/compile/internal/noder.(*irgen).stmt(0xc00015c300, {0xe448e8, 0xc0003ffcc0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stmt.go:134 +0xd0b
cmd/compile/internal/noder.(*irgen).stmts(0xc000448d20, {0xc00011ebd0, 0x1, 0xbd3578})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stmt.go:18 +0xaf
cmd/compile/internal/noder.(*irgen).funcBody(0xc00015c300, 0xc000446dc0, 0xc00043be00, 0xc000139740, 0xc000139780)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/func.go:46 +0x25e
cmd/compile/internal/noder.(*irgen).funcLit(0xc00015c300, {0xe43800, 0xc000139bc0}, 0xc0003ffca0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/expr.go:456 +0xe8
cmd/compile/internal/noder.(*irgen).expr0(0xc00015c300, {0xe43800, 0xc000139bc0}, {0xe44678, 0xc0003ffca0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/expr.go:108 +0x905
cmd/compile/internal/noder.(*irgen).expr(0xc00015c300, {0xe44678, 0xc0003ffca0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/expr.go:79 +0x5d4
cmd/compile/internal/noder.(*irgen).exprs(0x10, {0xc0004593a0, 0x1, 0x403829})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/expr.go:412 +0x8e
cmd/compile/internal/noder.(*irgen).exprList(0x0, {0xe44678, 0xc0003ffca0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/expr.go:395 +0x85
cmd/compile/internal/noder.(*irgen).stmt(0xc00015c300, {0xe44378, 0xc000139800})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stmt.go:75 +0x3ab
cmd/compile/internal/noder.(*irgen).stmts(0xc000448a10, {0xc00011ebe0, 0x1, 0x7f3ac53bf768})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/stmt.go:18 +0xaf
cmd/compile/internal/noder.(*irgen).funcBody(0xc00015c300, 0xc000446c60, 0x0, 0xc000139680, 0xc000139700)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/func.go:46 +0x25e
cmd/compile/internal/noder.(*irgen).funcDecl(0xc00015c300, 0xc000459788, 0xc000137b60)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/decl.go:108 +0x26f
cmd/compile/internal/noder.(*irgen).decls(0xc00015c300, {0xc0003ffea0, 0x2, 0xc00043bae0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/decl.go:28 +0x152
cmd/compile/internal/noder.(*irgen).generate(0xc00015c300, {0xc00011eb40, 0x2, 0x203000})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/irgen.go:232 +0x4b9
cmd/compile/internal/noder.check2({0xc00011eb40, 0x2, 0x2})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/irgen.go:92 +0x16d
cmd/compile/internal/noder.LoadPackage({0xc00013a120, 0x2, 0x0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/noder.go:90 +0x365
cmd/compile/internal/gc.Main(0xd135c0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/main.go:190 +0xaf3
main.main()
        /home/enikolov/sdk/gotip/src/cmd/compile/main.go:55 +0xdd
```