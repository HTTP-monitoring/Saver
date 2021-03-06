package config

const Default = `
db:
  host: 127.0.0.1
  port: "5431"
  user: postgres
  dbname: monitor
  password: postgres
  sslmode: disable
redis:
  host: 127.0.0.1
  port: "6378"
  threshold: 4
nats:
  host: nats://localhost:4221
  topic: save
  queue: saver
`
