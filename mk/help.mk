# Makefile help system

.PHONY: help
help:
	## display project help messages
	@python3 $(MK)/pyhelp.py $(MAKEFILE_LIST)

