LAST_COMMIT := $(shell git rev-parse --short HEAD)
LAST_COMMIT_DATE := $(shell git show -s --format=%ci ${LAST_COMMIT})

# Last commit on which date
lc:
	echo $(LAST_COMMIT_DATE) $(LAST_COMMIT)