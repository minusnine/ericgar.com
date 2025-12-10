test:
	hugo --gc && hugo server --buildDrafts --bind 0.0.0.0 --port 8080 -b http://100.115.92.200:8080/

deploy:
	hugo --gc && hugo && firebase deploy

testcodespace:
	hugo --gc && hugo server --buildDrafts --appendPort=false --baseURL https://$(CODESPACE_NAME)-1313.$(GITHUB_CODESPACES_PORT_FORWARDING_DOMAIN)
