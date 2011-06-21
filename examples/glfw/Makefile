
include $(GOROOT)/src/Make.inc

TARG:=glfw
CGOFILES:=glfw.go
CGO_LDFLAGS:=-lglfw

include $(GOROOT)/src/Make.pkg

main: install main.go
	$(GC) main.go
	$(LD) -o $@ main.$O
