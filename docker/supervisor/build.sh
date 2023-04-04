docker build -t jialin-supervisor-image:1.0 .

docker run -it --name js --rm  -v `pwd`/data/:/data/  --network host jialin-supervisor-image

