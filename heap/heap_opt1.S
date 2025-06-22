main.foo STEXT size=261 args=0x8 locals=0x48 funcid=0x0 align=0x0
	0x0000 00000 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	TEXT	main.foo(SB), ABIInternal, $72-8
	0x0000 00000 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	CMPQ	SP, 16(R14)
	0x0004 00004 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	PCDATA	$0, $-2
	0x0004 00004 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	JLS	240
	0x000a 00010 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	PCDATA	$0, $-1
	0x000a 00010 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	PUSHQ	BP
	0x000b 00011 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	MOVQ	SP, BP
	0x000e 00014 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	SUBQ	$64, SP
	0x0012 00018 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	FUNCDATA	$0, gclocals·ISb46fRPFoZ9pIfykFK/kQ==(SB)
	0x0012 00018 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	FUNCDATA	$1, gclocals·jCgrU8XAg0ifiSJZPFgpKw==(SB)
	0x0012 00018 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	FUNCDATA	$5, main.foo.arginfo1(SB)
	0x0012 00018 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	MOVQ	AX, main.arg_val+80(SP)
	0x0017 00023 (D:/learn/golang/Test-Go/heap/heap_opt1.go:5)	MOVQ	$11, main.foo_val1+48(SP)
	0x0020 00032 (D:/learn/golang/Test-Go/heap/heap_opt1.go:6)	MOVQ	$12, main.foo_val2+40(SP)
	0x0029 00041 (D:/learn/golang/Test-Go/heap/heap_opt1.go:7)	LEAQ	type:int(SB), AX
	0x0030 00048 (D:/learn/golang/Test-Go/heap/heap_opt1.go:7)	PCDATA	$1, $0
	0x0030 00048 (D:/learn/golang/Test-Go/heap/heap_opt1.go:7)	CALL	runtime.newobject(SB)
	0x0035 00053 (D:/learn/golang/Test-Go/heap/heap_opt1.go:7)	MOVQ	AX, main.&foo_val3+56(SP)
	0x003a 00058 (D:/learn/golang/Test-Go/heap/heap_opt1.go:7)	MOVQ	$13, (AX)
	0x0041 00065 (D:/learn/golang/Test-Go/heap/heap_opt1.go:8)	MOVQ	$14, main.foo_val4+32(SP)
	0x004a 00074 (D:/learn/golang/Test-Go/heap/heap_opt1.go:9)	MOVQ	$15, main.foo_val5+24(SP)
	0x0053 00083 (D:/learn/golang/Test-Go/heap/heap_opt1.go:9)	XORL	CX, CX
	0x0055 00085 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	JMP	224
	0x005a 00090 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	MOVQ	CX, main.i+16(SP)
	0x005f 00095 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	PCDATA	$1, $1
	0x005f 00095 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	NOP
	0x0060 00096 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printlock(SB)
	0x0065 00101 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.arg_val+80(SP), AX
	0x006a 00106 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x006f 00111 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x0074 00116 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val1+48(SP), AX
	0x0079 00121 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x007e 00126 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	NOP
	0x0080 00128 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x0085 00133 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val2+40(SP), AX
	0x008a 00138 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x008f 00143 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x0094 00148 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	MOVQ	main.&foo_val3+56(SP), AX
	0x0099 00153 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x009e 00158 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	NOP
	0x00a0 00160 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x00a5 00165 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val4+32(SP), AX
	0x00aa 00170 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x00af 00175 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x00b4 00180 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val5+24(SP), AX
	0x00b9 00185 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x00be 00190 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	NOP
	0x00c0 00192 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printnl(SB)
	0x00c5 00197 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printunlock(SB)
	0x00ca 00202 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	MOVQ	main.i+16(SP), CX
	0x00cf 00207 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	INCQ	CX
	0x00d2 00210 (D:/learn/golang/Test-Go/heap/heap_opt1.go:19)	MOVQ	main.&foo_val3+56(SP), AX
	0x00d7 00215 (D:/learn/golang/Test-Go/heap/heap_opt1.go:19)	NOP
	0x00e0 00224 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	CMPQ	CX, $5
	0x00e4 00228 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	JLT	90
	0x00ea 00234 (D:/learn/golang/Test-Go/heap/heap_opt1.go:19)	ADDQ	$64, SP
	0x00ee 00238 (D:/learn/golang/Test-Go/heap/heap_opt1.go:19)	POPQ	BP
	0x00ef 00239 (D:/learn/golang/Test-Go/heap/heap_opt1.go:19)	RET
	0x00f0 00240 (D:/learn/golang/Test-Go/heap/heap_opt1.go:19)	NOP
	0x00f0 00240 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	PCDATA	$1, $-1
	0x00f0 00240 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	PCDATA	$0, $-2
	0x00f0 00240 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	MOVQ	AX, 8(SP)
	0x00f5 00245 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x00fa 00250 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	PCDATA	$0, $-1
	0x00fa 00250 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	MOVQ	8(SP), AX
	0x00ff 00255 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	NOP
	0x0100 00256 (D:/learn/golang/Test-Go/heap/heap_opt1.go:3)	JMP	0
	0x0000 49 3b 66 10 0f 86 e6 00 00 00 55 48 89 e5 48 83  I;f.......UH..H.
	0x0010 ec 40 48 89 44 24 50 48 c7 44 24 30 0b 00 00 00  .@H.D$PH.D$0....
	0x0020 48 c7 44 24 28 0c 00 00 00 48 8d 05 00 00 00 00  H.D$(....H......
	0x0030 e8 00 00 00 00 48 89 44 24 38 48 c7 00 0d 00 00  .....H.D$8H.....
	0x0040 00 48 c7 44 24 20 0e 00 00 00 48 c7 44 24 18 0f  .H.D$ ....H.D$..
	0x0050 00 00 00 31 c9 e9 86 00 00 00 48 89 4c 24 10 90  ...1......H.L$..
	0x0060 e8 00 00 00 00 48 8d 44 24 50 e8 00 00 00 00 e8  .....H.D$P......
	0x0070 00 00 00 00 48 8d 44 24 30 e8 00 00 00 00 66 90  ....H.D$0.....f.
	0x0080 e8 00 00 00 00 48 8d 44 24 28 e8 00 00 00 00 e8  .....H.D$(......
	0x0090 00 00 00 00 48 8b 44 24 38 e8 00 00 00 00 66 90  ....H.D$8.....f.
	0x00a0 e8 00 00 00 00 48 8d 44 24 20 e8 00 00 00 00 e8  .....H.D$ ......
	0x00b0 00 00 00 00 48 8d 44 24 18 e8 00 00 00 00 66 90  ....H.D$......f.
	0x00c0 e8 00 00 00 00 e8 00 00 00 00 48 8b 4c 24 10 48  ..........H.L$.H
	0x00d0 ff c1 48 8b 44 24 38 66 0f 1f 84 00 00 00 00 00  ..H.D$8f........
	0x00e0 48 83 f9 05 0f 8c 70 ff ff ff 48 83 c4 40 5d c3  H.....p...H..@].
	0x00f0 48 89 44 24 08 e8 00 00 00 00 48 8b 44 24 08 90  H.D$......H.D$..
	0x0100 e9 fb fe ff ff                                   .....
	rel 44+4 t=R_PCREL type:int+0
	rel 49+4 t=R_CALL runtime.newobject+0
	rel 97+4 t=R_CALL runtime.printlock+0
	rel 107+4 t=R_CALL runtime.printpointer+0
	rel 112+4 t=R_CALL runtime.printsp+0
	rel 122+4 t=R_CALL runtime.printpointer+0
	rel 129+4 t=R_CALL runtime.printsp+0
	rel 139+4 t=R_CALL runtime.printpointer+0
	rel 144+4 t=R_CALL runtime.printsp+0
	rel 154+4 t=R_CALL runtime.printpointer+0
	rel 161+4 t=R_CALL runtime.printsp+0
	rel 171+4 t=R_CALL runtime.printpointer+0
	rel 176+4 t=R_CALL runtime.printsp+0
	rel 186+4 t=R_CALL runtime.printpointer+0
	rel 193+4 t=R_CALL runtime.printnl+0
	rel 198+4 t=R_CALL runtime.printunlock+0
	rel 246+4 t=R_CALL runtime.morestack_noctxt+0
main.main STEXT size=270 args=0x0 locals=0x50 funcid=0x0 align=0x0
	0x0000 00000 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	TEXT	main.main(SB), ABIInternal, $80-0
	0x0000 00000 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	CMPQ	SP, 16(R14)
	0x0004 00004 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	PCDATA	$0, $-2
	0x0004 00004 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	JLS	260
	0x000a 00010 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	PCDATA	$0, $-1
	0x000a 00010 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	PUSHQ	BP
	0x000b 00011 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	MOVQ	SP, BP
	0x000e 00014 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	SUBQ	$72, SP
	0x0012 00018 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	FUNCDATA	$0, gclocals·FzY36IO2mY0y4dZ1+Izd/w==(SB)
	0x0012 00018 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	FUNCDATA	$1, gclocals·FzY36IO2mY0y4dZ1+Izd/w==(SB)
	0x0012 00018 (D:/learn/golang/Test-Go/heap/heap_opt1.go:23)	MOVQ	$666, main.arg_val+56(SP)
	0x001b 00027 (<unknown line number>)	NOP
	0x001b 00027 (D:/learn/golang/Test-Go/heap/heap_opt1.go:5)	MOVQ	$11, main.foo_val1+48(SP)
	0x0024 00036 (D:/learn/golang/Test-Go/heap/heap_opt1.go:6)	MOVQ	$12, main.foo_val2+40(SP)
	0x002d 00045 (D:/learn/golang/Test-Go/heap/heap_opt1.go:7)	MOVQ	$13, main.foo_val3+32(SP)
	0x0036 00054 (D:/learn/golang/Test-Go/heap/heap_opt1.go:8)	MOVQ	$14, main.foo_val4+24(SP)
	0x003f 00063 (D:/learn/golang/Test-Go/heap/heap_opt1.go:9)	MOVQ	$15, main.foo_val5+16(SP)
	0x0048 00072 (D:/learn/golang/Test-Go/heap/heap_opt1.go:9)	XORL	AX, AX
	0x004a 00074 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	JMP	198
	0x004c 00076 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	MOVQ	AX, main.i+8(SP)
	0x0051 00081 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	PCDATA	$1, $0
	0x0051 00081 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printlock(SB)
	0x0056 00086 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.arg_val+56(SP), AX
	0x005b 00091 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	NOP
	0x0060 00096 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x0065 00101 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x006a 00106 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val1+48(SP), AX
	0x006f 00111 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x0074 00116 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x0079 00121 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val2+40(SP), AX
	0x007e 00126 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	NOP
	0x0080 00128 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x0085 00133 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x008a 00138 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val3+32(SP), AX
	0x008f 00143 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x0094 00148 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x0099 00153 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val4+24(SP), AX
	0x009e 00158 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	NOP
	0x00a0 00160 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x00a5 00165 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printsp(SB)
	0x00aa 00170 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	LEAQ	main.foo_val5+16(SP), AX
	0x00af 00175 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printpointer(SB)
	0x00b4 00180 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printnl(SB)
	0x00b9 00185 (D:/learn/golang/Test-Go/heap/heap_opt1.go:15)	CALL	runtime.printunlock(SB)
	0x00be 00190 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	MOVQ	main.i+8(SP), AX
	0x00c3 00195 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	INCQ	AX
	0x00c6 00198 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	CMPQ	AX, $5
	0x00ca 00202 (D:/learn/golang/Test-Go/heap/heap_opt1.go:14)	JLT	76
	0x00cc 00204 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	MOVQ	main.foo_val3+32(SP), AX
	0x00d1 00209 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	MOVQ	AX, main..autotmp_16+64(SP)
	0x00d6 00214 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	CALL	runtime.printlock(SB)
	0x00db 00219 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	MOVQ	main..autotmp_16+64(SP), AX
	0x00e0 00224 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	CALL	runtime.printint(SB)
	0x00e5 00229 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	CALL	runtime.printsp(SB)
	0x00ea 00234 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	LEAQ	main.foo_val3+32(SP), AX
	0x00ef 00239 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	CALL	runtime.printpointer(SB)
	0x00f4 00244 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	CALL	runtime.printnl(SB)
	0x00f9 00249 (D:/learn/golang/Test-Go/heap/heap_opt1.go:25)	CALL	runtime.printunlock(SB)
	0x00fe 00254 (D:/learn/golang/Test-Go/heap/heap_opt1.go:26)	ADDQ	$72, SP
	0x0102 00258 (D:/learn/golang/Test-Go/heap/heap_opt1.go:26)	POPQ	BP
	0x0103 00259 (D:/learn/golang/Test-Go/heap/heap_opt1.go:26)	RET
	0x0104 00260 (D:/learn/golang/Test-Go/heap/heap_opt1.go:26)	NOP
	0x0104 00260 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	PCDATA	$1, $-1
	0x0104 00260 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	PCDATA	$0, $-2
	0x0104 00260 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	CALL	runtime.morestack_noctxt(SB)
	0x0109 00265 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	PCDATA	$0, $-1
	0x0109 00265 (D:/learn/golang/Test-Go/heap/heap_opt1.go:22)	JMP	0
	0x0000 49 3b 66 10 0f 86 fa 00 00 00 55 48 89 e5 48 83  I;f.......UH..H.
	0x0010 ec 48 48 c7 44 24 38 9a 02 00 00 48 c7 44 24 30  .HH.D$8....H.D$0
	0x0020 0b 00 00 00 48 c7 44 24 28 0c 00 00 00 48 c7 44  ....H.D$(....H.D
	0x0030 24 20 0d 00 00 00 48 c7 44 24 18 0e 00 00 00 48  $ ....H.D$.....H
	0x0040 c7 44 24 10 0f 00 00 00 31 c0 eb 7a 48 89 44 24  .D$.....1..zH.D$
	0x0050 08 e8 00 00 00 00 48 8d 44 24 38 0f 1f 44 00 00  ......H.D$8..D..
	0x0060 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 30 e8  ..........H.D$0.
	0x0070 00 00 00 00 e8 00 00 00 00 48 8d 44 24 28 66 90  .........H.D$(f.
	0x0080 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 20 e8  ..........H.D$ .
	0x0090 00 00 00 00 e8 00 00 00 00 48 8d 44 24 18 66 90  .........H.D$.f.
	0x00a0 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 10 e8  ..........H.D$..
	0x00b0 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8b  ..............H.
	0x00c0 44 24 08 48 ff c0 48 83 f8 05 7c 80 48 8b 44 24  D$.H..H...|.H.D$
	0x00d0 20 48 89 44 24 40 e8 00 00 00 00 48 8b 44 24 40   H.D$@.....H.D$@
	0x00e0 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 20 e8  ..........H.D$ .
	0x00f0 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48 83  ..............H.
	0x0100 c4 48 5d c3 e8 00 00 00 00 e9 f2 fe ff ff        .H]...........
	rel 82+4 t=R_CALL runtime.printlock+0
	rel 97+4 t=R_CALL runtime.printpointer+0
	rel 102+4 t=R_CALL runtime.printsp+0
	rel 112+4 t=R_CALL runtime.printpointer+0
	rel 117+4 t=R_CALL runtime.printsp+0
	rel 129+4 t=R_CALL runtime.printpointer+0
	rel 134+4 t=R_CALL runtime.printsp+0
	rel 144+4 t=R_CALL runtime.printpointer+0
	rel 149+4 t=R_CALL runtime.printsp+0
	rel 161+4 t=R_CALL runtime.printpointer+0
	rel 166+4 t=R_CALL runtime.printsp+0
	rel 176+4 t=R_CALL runtime.printpointer+0
	rel 181+4 t=R_CALL runtime.printnl+0
	rel 186+4 t=R_CALL runtime.printunlock+0
	rel 215+4 t=R_CALL runtime.printlock+0
	rel 225+4 t=R_CALL runtime.printint+0
	rel 230+4 t=R_CALL runtime.printsp+0
	rel 240+4 t=R_CALL runtime.printpointer+0
	rel 245+4 t=R_CALL runtime.printnl+0
	rel 250+4 t=R_CALL runtime.printunlock+0
	rel 261+4 t=R_CALL runtime.morestack_noctxt+0
go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go:info.main.foo$abstract SDWARFABSFCN dupok size=111
	0x0000 05 6d 61 69 6e 2e 66 6f 6f 00 01 03 01 22 61 72  .main.foo...."ar
	0x0010 67 5f 76 61 6c 00 00 00 00 00 00 21 66 6f 6f 5f  g_val......!foo_
	0x0020 76 61 6c 31 00 05 00 00 00 00 21 66 6f 6f 5f 76  val1......!foo_v
	0x0030 61 6c 32 00 06 00 00 00 00 21 66 6f 6f 5f 76 61  al2......!foo_va
	0x0040 6c 33 00 07 00 00 00 00 21 66 6f 6f 5f 76 61 6c  l3......!foo_val
	0x0050 34 00 08 00 00 00 00 21 66 6f 6f 5f 76 61 6c 35  4......!foo_val5
	0x0060 00 09 00 00 00 00 21 69 00 0e 00 00 00 00 00     ......!i.......
	rel 23+4 t=R_DWARFSECREF go:info.int+0
	rel 38+4 t=R_DWARFSECREF go:info.int+0
	rel 53+4 t=R_DWARFSECREF go:info.int+0
	rel 68+4 t=R_DWARFSECREF go:info.int+0
	rel 83+4 t=R_DWARFSECREF go:info.int+0
	rel 98+4 t=R_DWARFSECREF go:info.int+0
	rel 106+4 t=R_DWARFSECREF go:info.int+0
main..inittask SNOPTRDATA size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
gclocals·ISb46fRPFoZ9pIfykFK/kQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·jCgrU8XAg0ifiSJZPFgpKw== SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
main.foo.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
gclocals·FzY36IO2mY0y4dZ1+Izd/w== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
go:sehuw.12.AQ4CBQ4DC1AAAAAA SSEHUNWINDINFO dupok size=12
	0x0000 01 0e 02 05 0e 03 0b 50 00 00 00 00              .......P....
