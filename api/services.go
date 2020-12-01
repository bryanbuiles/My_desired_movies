package api

// Services struct que lista los diferentes servicios
// son los servicios que va a tener el programa
type Services struct {
	search MovieSearch
}

// WebServices servicios web
type WebServices struct {
	s Services
}

// NewServices Nuevo servicio
func NewServices() Services {
	return Services{
		search: &MovieService{},
	}
}

// Start comienza un nuevo servicio
func start() *WebServices { // comieza el servicio
	return &WebServices{s: NewServices()}
}
