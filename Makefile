run-migration:
	./scripts/migrate.sh

open-database:
	./scripts/open_db.sh

jet-gen:
	./scripts/jet_gen.sh

mock-gen:
	./scripts/gen_mock.sh

docker-up:
	docker-compose up -d