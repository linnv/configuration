CFILES= $(wildcard *.c)
# COBJS  = $(patsubst %.c,%.o,$(CFILES))
# or
COBJS=$(CFILES:%.c=%.o)
${COBJS}:${CFILES}
	echo "patsubst" ${COBJS}
	echo "wildcard" ${CFILES}
all: ${COBJS}

# a.o: b.o	
# 	echo "a right part <"$<
# 	echo "a left part @"$@
# b.o: b.c
# 	echo "b right part <"$<
# 	echo "b left part @"$@
# all: b.o a.o
