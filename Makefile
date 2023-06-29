currentBranch = $(shell git symbolic-ref HEAD | sed -e 's,.*/\(.*\),\1,')
ENVOY:=envoy-donech
JAEGER:=jaeger-donech


.PHONY: gen
gen: lint
	rm -rf gen/go/donech/*
	docker run --rm -v "$(shell pwd):/work" uber/prototool prototool all

.PHONY: lint
lint:
	docker run --rm -v "$(shell pwd):/work" uber/prototool prototool lint

.PHONY: install
install:
	@echo "当前git分支 $(currentBranch)"
	rm -rf .subsplit
	git subsplit init .
	git subsplit publish  --no-tags --heads=$(currentBranch) gen/go:git@github.com:donech/proto-go.git
	rm -rf .subsplit

.PHONY: depend
depend: envoy jaeger
	@echo "start depends success"

foundEnvoy := $(shell docker ps -af "name=$(ENVOY)" -q | grep -q . && echo Found || echo Not Found)

.PHONY: envoy
ifeq ($(foundEnvoy), Found)
envoy: stop-envoy start-envoy
	echo "stop and start envoy success"
else
envoy: start-envoy
	echo "start envoy success"
endif

.PHONY: stop-envoy
stop-envoy:
	docker stop $(ENVOY)
	docker rm $(ENVOY)

.PHONY: start-envoy
start-envoy:
	docker run -itd --name $(ENVOY) \
    	  -p 51051:51051 \
    	  -p 10000:10000 \
          -v "$(shell pwd)/gen/descriptor/donech/gate/v1/v1.pb:/data/gate.pb:ro" \
          -v "$(shell pwd)/envoy/envoy-config-dev.yml:/etc/envoy/envoy.yaml:ro" \
          envoyproxy/envoy:v1.26-latest
	docker logs $(ENVOY)

.PHONY: in-envory
in-envoy:
	docker exec -it $(ENVOY) /bin/bash

.PHONY: valid-envoy
valid-envoy:
	docker run --rm \
              -v "$(shell pwd)/gen/descriptor/donech/gate/v1/v1.pb:/data/gate.pb:ro" \
              -v "$(shell pwd)/envoy/envoy-config-dev.yml:/etc/envoy/envoy.yaml:ro" \
              envoyproxy/envoy:v1.26-latest \
              --mode validate -c /etc/envoy/envoy.yaml



foundJaeger := $(shell docker ps -f "name=$(JAEGER)" -q | grep -q . && echo Found || echo Not Found)

.PHONY: jaeger
ifeq ($(foundJaeger), Found)
jaeger: stop-jaeger start-jaeger
	@echo "stop and start jaeger success"
else
jaeger: start-jaeger
	@echo "start jaeger success"
endif
.PHONY: start-jaeger
start-jaeger:
	@echo "jaeger starting"
	docker run --rm -d --name=$(JAEGER) \
				  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
				  -p 9411:9411\
				  -p 16686:16686\
                  jaegertracing/all-in-one
	@echo "jaeger starting"

.PHONY: stop-jaeger
stop-jaeger:
	docker stop $(JAEGER)
