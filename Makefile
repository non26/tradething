# DOCKER_IMAGE_NAME=ttl-img
# DOCKER_IMAGE_TAG=v1
# PROFILE=root-user
# REGION=ap-northeast-1
# CONTAINER_SERVICE_NAME=ttl-container-service-1
# CONTAINER_IMAGE_LABEL=ttl-img-label
# LOCAL_IMAGE_NAME=ttl-img
# LOCAL_IMAGE_TAG=v1
# INSTANCE_NAME=ttl-nginx-1

# DOCKER_IMAGE_NAME=ttl-img
# DOCKER_IMAGE_TAG=v2
# PROFILE=root-user
# REGION=ap-northeast-1
# CONTAINER_SERVICE_NAME=ttl-container-service-1
# CONTAINER_IMAGE_LABEL=ttl-img-label-v2
# LOCAL_IMAGE_NAME=ttl-img
# LOCAL_IMAGE_TAG=v2
# INSTANCE_NAME=ttl-nginx-1


# create-image:
# 	docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} .

# images:
# 	docker images

# get-container:
# 	aws lightsail get-container-services --profile ${PROFILE}

# push2-registry:
# 	aws lightsail push-container-image --region ${REGION} \
# 	--service-name ${CONTAINER_SERVICE_NAME} \
# 	--label ${CONTAINER_IMAGE_LABEL} \
# 	--image ${LOCAL_IMAGE_NAME}:${LOCAL_IMAGE_TAG} \
# 	--profile ${PROFILE}

# # echo `aws lightsail get-container-images --service-name ttl-container-service-1 --profile root-user | jq -r '.containerImages[0].image'`

# export image := `aws lightsail get-container-images --service-name ttl-container-service-1 --profile root-user | jq -r '.containerImages[0].image'`

# deploy:
# 	aws lightsail create-container-service-deployment \
# 		--profile ${PROFILE} \
# 		--service-name ${CONTAINER_SERVICE_NAME} \
# 		--containers '{"'$(CONTAINER_IMAGE_LABEL)'":{"image":"'$(image)'","environment":{"HOST":"","PORT":"8080","LOG_ENV":"production"},"ports":{"8080":"HTTP"}}}' \
# 		--public-endpoint '{"containerName":"'$(CONTAINER_IMAGE_LABEL)'","containerPort":8080,"healthCheck":{"path":"/"}}'


go-curl:
	go run curl/main.go

local:
	go run cmd/app/echo/main.go

go-zip:
	bash generate-function.sh