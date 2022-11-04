FROM amazon/aws-lambda-nodejs:14

WORKDIR ${LAMBDA_TASK_ROOT}

RUN npm install -g pnpm
COPY --from=build-image /app/dist ${LAMBDA_TASK_ROOT}/dist
COPY --from=build-image /app/package.json ./
COPY --from=build-image /app/pnpm-lock.yaml ./

RUN npm install -g pnpm
RUN pnpm install

CMD [ "app.handler" ]