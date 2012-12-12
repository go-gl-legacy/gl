
GLenum goglInit() {
#ifdef __APPLE__
	return (GLenum)0;
#else
	return glewInit();
#endif
}

