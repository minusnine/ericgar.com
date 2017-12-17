deploy:
	hugo && goapp deploy
test:
	hugo && goapp serve -host 0.0.0.0 -port 1313 
serve:
	hugo server -D --bind 0.0.0.0 -b ssh.ericgar.com:1313
