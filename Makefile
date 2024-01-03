test:
	hugo --gc && hugo server -D --bind 0.0.0.0 --port 8080 -b http://100.115.92.200:8080/

deploy:
	hugo --gc && hugo && firebase deploy

