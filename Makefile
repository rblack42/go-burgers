# go-burger Makefile
PROJDIR := $(PWD)
PROJNAME := $(notdir $(PROJDIR))	
MK := mk
.DEFAULT_GOAL := all

include $(MK)/os-detect.mk
include $(MK)/golang.mk
