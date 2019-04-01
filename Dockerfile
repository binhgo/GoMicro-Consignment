# consignment-service/Dockerfile

# We use the official golang image, which contains all the
# correct build tools and libraries. Notice `as builder`,
# this gives this container a name that we can reference later on.
FROM golang:latest as builder
# ARG security: https://bit.ly/2oY3pCn
ARG SSH_PRIVATE_KEY
WORKDIR /builder/
ADD . /builder/
RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
	&& git config --global url."git@bitbucket.org:".insteadOf https://bitbucket.org/ \
	&& git config --global url."git@github.com:".insteadOf https://github.com/ \
	&& ssh-keyscan bitbucket.org >> ~/.ssh/known_hosts \
	&& ssh-keyscan github.com >> ~/.ssh/known_hosts

RUN go get -u github.com/binhgo/GoMicro-Consignment
