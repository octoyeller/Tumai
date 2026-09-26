CC = wails build
CF = -clean
o_name = tumai

run: compile
	build/bin/$(o_name)

compile:
	$(CC) $(CF)

release: compile
	cp /usr/lib/x86_64-linux-gnu/libwebkit2gtk-4.0.so libwebkit2gtk-4.0.so
	zip -r9 t.zip build/bin/$(o_name) libwebkit2gtk-4.0.so

cleanup:
	rm t.zip libwebkit2gtk-4.0.so
