package modHttp

func Http_Initialize() error {
	return getSingleHttpServer().initialize()
}
func Http_Start() error {
	return getSingleHttpServer().start()
}
