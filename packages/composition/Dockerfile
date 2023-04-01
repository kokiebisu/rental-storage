FROM node:18 AS build-image

WORKDIR /app

COPY package.json ./
COPY pnpm-lock.yaml ./
COPY tsconfig.docker.json ./

RUN npm install -g pnpm typescript
RUN pnpm install
COPY ./ ./

RUN pnpm build:docker

FROM amazon/aws-lambda-nodejs:16

WORKDIR ${LAMBDA_TASK_ROOT}

RUN npm install -g pnpm
COPY --from=build-image /app/dist ${LAMBDA_TASK_ROOT}/dist
COPY --from=build-image /app/package.json ./
COPY --from=build-image /app/pnpm-lock.yaml ./

RUN pnpm install --prod

CMD [ "app.handler" ]