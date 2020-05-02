# go wiki
Simple go example for docker and kubernets
Refer to <https://golang.org/doc/articles/wiki/>

```bash
#!/bin/bash
go build gowiki.go
gowiki
docker build -t chienyuan/gowiki .
docker run  -p 8080:8080 chienyuan/gowiki
docker push chienyuan/gowiki
kubectl delete -f gowiki.yaml
kubectl apply -f gowiki.yaml
```
