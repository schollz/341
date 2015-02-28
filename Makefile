.PHONY: all

all:
	convert 341.png -filter point -resize 1200% -quiet look-at-me.png
