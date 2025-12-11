test:
	hugo server --gc --cleanDestinationDir --buildDrafts --bind 0.0.0.0 --port 8080 -b http://100.115.92.200:8080/

deploy:
	hugo --gc --cleanDestinationDir && firebase deploy

testcodespace:
	hugo server  --gc --cleanDestinationDir --buildDrafts --appendPort=false --baseURL https://$(CODESPACE_NAME)-1313.$(GITHUB_CODESPACES_PORT_FORWARDING_DOMAIN)
