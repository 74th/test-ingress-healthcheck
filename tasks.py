from invoke import task

@task
def build(c):
    with c.cd("cmd/server"):
        c.run("CGO_ENABLED=0 go build -o server server.go")
    c.run("docker-compose build")
    c.run("docker-compose push")

@task
def app1(c):
    c.run("kubectl apply -f manifest/app1")

@task
def endpoints(c):
    c.run("gcloud endpoints services deploy endpoints/app3.yaml")
