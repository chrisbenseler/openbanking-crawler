# Open Banking Crawler

## Running in docker

```ssh
docker build -t openbankingcrawler .
docker run -p 3000:3000 -e DBHOST=host.docker.internal -t openbankingcrawler:latest
```

```ssh
docker run -e DBHOST=host.docker.internal -e MODE=local -t openbankingcrawler:latest
docker run -e DBHOST=host.docker.internal -e MODE=report -e TYPE=PERSONAL_CC -v /tmp/outputs:/go/src/openbankingcrawler/outputs -t openbankingcrawler:latest
```
