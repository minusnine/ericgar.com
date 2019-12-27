deploy:
	hugo && firebase deploy

serve:
	hugo server -D --bind 0.0.0.0 -b ssh.ericgar.com:1313
