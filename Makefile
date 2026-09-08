CC = wails build
CF = -clean
o_name = tumai

run: compile
	build/bin/$(o_name)

compile:
	$(CC) $(CF)
