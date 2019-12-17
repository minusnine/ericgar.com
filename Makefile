test:
	hugo && dev_appserver.py --enable_host_checking false --host 0.0.0.0 --port 1313 app.yaml
deploy:
	hugo && firebase deploy
serve:
	hugo server -D --bind 0.0.0.0 -b ssh.ericgar.com:1313
