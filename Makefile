deploy:
	hugo && goapp deploy
test:
	hugo && goapp serve -host 0.0.0.0 -port 1313
