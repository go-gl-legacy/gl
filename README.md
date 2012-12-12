gl | OpenGL Bindings for golang
===============================

This fork of [go-gl/gl](http://github.com/go-gl/gl) enables OpenGL 3.2
contexts on OSX (this is a core context with no backwards
compatability). All the deprecated functions have been removed. I
started this fork because GLEW (which gogl/gl uses) does not appear to
play nicely with Apple's OpenGL 3 implementation. If anyone knows
better, please let me know.

You should only need this if you are targetting a strict OpenGL 3+
context and need OSX compatibility (I have also tested it on Ubuntu
12.04, but for any other OS you're on your own!)

Patches should, in general, be directed to the original repo. Hopefully
some day GLEW and Apple will sort it out and I can delete this
fork.
