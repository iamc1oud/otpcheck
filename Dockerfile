FROM ubuntu:20.04
RUN apt update && apt install -y ca-certificates
WORKDIR /app
COPY . .
COPY config.sample.toml config.toml
# COPY static/smtp.tpl ./static/smtp.tpl
CMD ["./otpcheck", "--config", "./config.toml"]