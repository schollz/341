.PHONY: all

all:
	convert src/341.png -filter point -resize 1200% -quiet preview.png
