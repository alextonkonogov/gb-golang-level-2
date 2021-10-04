build-mac:
	rm -f main || true
	env GOOS=darwin GOARCH=amd64 go build main.go
#file main outs:
	#main: Mach-O 64-bit executable x86_64

build-linux:
	rm -f main || true
	env GOOS=linux GOARCH=386 go build main.go
#file main outs:
	#main: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), statically linked, Go BuildID=MFfp7PmQyEGtsR-sl_k1/99bH2yvh5cs74sYLb4wm/1yJVswEvP1qd_AtJx5mU/kEuoxW6IUi70Qf6xRN-r, not stripped

build-windows:
	rm -f main.exe || true
	env GOOS=windows GOARCH=amd64 go build main.go
#file main.exe outs:
	#main.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
