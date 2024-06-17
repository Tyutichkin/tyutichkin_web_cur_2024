build:
	docker build -t mospolytech . 

# Запуск без компила
fast-run:
	docker run -p 8083:8083 mospolytech

full-run: build fast-run


clean:
	docker container rm $$(docker ps -aq) -f

