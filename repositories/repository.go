package repositories

type Repository struct {
	db     Postgres
	client Redis
}

type ReposOption func(*Repository)

func NewRepository(options ...ReposOption) *Repository {
	repo := &Repository{}
	for _, option := range options {
		option(repo)
	}
	return repo
}

func WithPostgresRepository(db Postgres) ReposOption {
	return func(r *Repository) {
		r.db = db
	}
}

func WithRedis(client Redis) ReposOption {
	return func(r *Repository) {
		r.client = client
	}
}
