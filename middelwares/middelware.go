package middlewares

type Middleware struct {
	AuthMiddleware IAuthMiddleware
}

func Initialise() Middleware {
	return Middleware{
		AuthMiddleware: InitialiseAuthMiddleware(),
	}
}
