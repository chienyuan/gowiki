## go wiki


go build gowiki.go
gowiki

docker build -t chienyuan/gowiki .
docker run  -p 8080:8080 chienyuan/gowiki
docker push chienyuan/gowiki
kubectl apply -f gowiki.yaml
