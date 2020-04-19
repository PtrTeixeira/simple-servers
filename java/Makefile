SHELL := bash

build/main.jar: build/Main.class manifest.txt
	cp manifest.txt build/
	cd build/ && jar -cvfm main.jar manifest.txt Main.class

build/Main.class: Main.java
	javac -d build/ $<

clean: 
	rm -fr build
