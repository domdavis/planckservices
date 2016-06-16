# Planckservice

Microservices are all about breaking down applications into small, self
contained chunks. Typically these are deployed in clusters and communicate via
RESTful APIs. Current thinking is that, for applications and deployments,
smaller is most definitely better.

In physics, the Planck length is the smallest measurement of length with any
real meaning. In the world of microservices the planckservice is the smallest
possible service with any real meaning. By combining planckservices in a
properly configured swarm you can achieve anything that it is possible to
achieve on any digital computer, but in a truly fault tolerant and distributed
way.

## Endpoints

```
/
```

Randomly returns `0` or `1`.

```
/n
```

Where `n = 0|1`. Returns `n`.

```
/n/m
```

Where `n = 0|1` and `m = 0|1`. Returns `n NAND m`.

## Creating a docker image

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o planckservice .
docker build -t planckservice .
docker run --rm -d -p 8000:8000 planckservice
```
