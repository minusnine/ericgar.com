deploy:
	hugo --gc && hugo && firebase deploy

serve:
	hugo --gc && hugo server -D --bind 0.0.0.0 -b ssh.ericgar.com:1313
