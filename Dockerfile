# ----------------------------------------
FROM node:16-alpine as frontend-builder
ARG BUILD_COUNTRY=""

LABEL Author="Syhily <syhily@gmail.com>"
LABEL Thanks="Rex <talebook@foxmail.com>"
LABEL Thanks="oldiy <oldiy2018@gmail.com>"

WORKDIR /app
RUN if [ "x${BUILD_COUNTRY}" = "xCN" ]; then \
    echo "using repo mirrors for ${BUILD_COUNTRY}"; \
    npm config set registry https://mirrors.tencent.com/npm/; \
    fi

COPY ["app/package.json", "app/package-lock.json*", "/app/"]
RUN npm install

COPY app/ /app/
RUN npm run build

# ----------------------------------------
FROM goreleaser/goreleaser as backend-builder

WORKDIR /build
ADD . /build
RUN goreleaser build --snapshot --rm-dist

# ----------------------------------------
FROM talebook/calibre-docker as server
ARG BUILD_COUNTRY=""

# Set mirrors in china
RUN if [ "x${BUILD_COUNTRY}" = "xCN" ]; then \
    echo "using repo mirrors for ${BUILD_COUNTRY}"; \
    sed 's@deb.debian.org/debian@mirrors.aliyun.com/debian@' -i /etc/apt/sources.list; \
    fi

# install envsubst
RUN apt-get update && apt-get install -y gettext

# intall nodejs for nuxtjs server side render
RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -
RUN apt-get install -y nodejs

# ----------------------------------------
# Production environment.
FROM server as production

COPY docker/nginx.conf /etc/nginx/conf.d/talebook.conf
COPY --from=frontend-builder /app/.nuxt/ /talebook/app/.nuxt/
COPY --from=frontend-builder /app/node_modules/ /talebook/app/node_modules/
COPY --from=frontend-builder /app/src/static/ /talebook/app/dist/
COPY --from=backend-builder /build/dist/talebook_linux_amd64_v1/talebook /talebook/talebook

EXPOSE 7000

VOLUME ["/talebook/repository"]

# TODO Finish this docker ENTRYPOINT
