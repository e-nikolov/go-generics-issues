`gotip build -gcflags=all=-G=3 main.go`


```
# command-line-arguments
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x28 pc=0x5f1f7b]

goroutine 1 [running]:
cmd/compile/internal/typecheck.(*crawler).checkGenericType(0xc00012f380, 0x0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/crawler.go:227 +0x3b
cmd/compile/internal/typecheck.(*crawler).markInlBody.func1({0xe59030, 0xc00012f380})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/crawler.go:201 +0xb2
cmd/compile/internal/ir.Visit.func1({0xe59030, 0xc00012f380})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:105 +0x30
cmd/compile/internal/ir.(*SelectorExpr).doChildren(0xc00011acc0, 0xc000120318)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/node_gen.go:1067 +0x62
cmd/compile/internal/ir.DoChildren(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:94
cmd/compile/internal/ir.Visit.func1({0xe59738, 0xc00011acc0})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:106 +0x57
cmd/compile/internal/ir.(*CallExpr).doChildren(0xc000144480, 0xc000120318)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/node_gen.go:260 +0x64
cmd/compile/internal/ir.DoChildren(...)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:94
cmd/compile/internal/ir.Visit.func1({0xe57b18, 0xc000144480})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:106 +0x57
cmd/compile/internal/ir.Visit({0xe57b18, 0xc000144480}, 0xc000120240)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:108 +0xb8
cmd/compile/internal/ir.VisitList({0xc00011c280, 0x1, 0xc000130700}, 0xc0001305b0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/ir/visit.go:114 +0x65
cmd/compile/internal/typecheck.(*crawler).markInlBody(0xc00011c290, 0xc0001305b0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/crawler.go:217 +0x22d
cmd/compile/internal/typecheck.(*crawler).markObject(0xc000132120, 0xc00012f2b0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/crawler.go:35 +0x2f
cmd/compile/internal/typecheck.crawlExports({0xc00068c000, 0x3, 0x4dfad3})
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/crawler.go:23 +0x9f
cmd/compile/internal/typecheck.WriteExports({0xe3f3e0, 0xc000114d80}, 0x1)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/typecheck/iexport.go:272 +0x65
cmd/compile/internal/noder.WriteExports(0xc00011c1e0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/noder/export.go:40 +0x7a
cmd/compile/internal/gc.dumpCompilerObj(0xc00011c1e0)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/obj.go:107 +0x28
cmd/compile/internal/gc.dumpobj1({0x7ffe58dcc016, 0x24}, 0x3)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/obj.go:63 +0x185
cmd/compile/internal/gc.dumpobj()
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/obj.go:44 +0x36
cmd/compile/internal/gc.Main(0xd13700)
        /home/enikolov/sdk/gotip/src/cmd/compile/internal/gc/main.go:318 +0x10b6
main.main()
        /home/enikolov/sdk/gotip/src/cmd/compile/main.go:55 +0xdd
```