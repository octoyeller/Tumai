CC = wails build
CF = -clean
o_name = Tumai

run: compile
	build/bin/$(o_name)

compile:
	$(CC) $(CF)
